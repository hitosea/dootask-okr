<template>
    <n-scrollbar :on-scroll="onScroll">
        <div class="okr-replay-main">
            <div>
                <OkrNotDatas v-if="items.length == 0 && !loadIng && !onscrolloading" :msg="$t('暂无复盘')" :types="searchObject !=''"></OkrNotDatas>
                <OkrLoading v-if="loadIng"></OkrLoading>
            </div>
            <div v-if="items.length != 0" class="replay">
                <div v-for="(item, index) in items" :key="index"
                    :class="{ 'replay-item': true, 'replay-item-active': item.isActive }" @click="openMultiple">
                    <div class="replay-item-head">
                        <div>
                            <span class="replay-item-okr-level scale-[0.8333]" :class="pStatus(item.priority)">{{
                                item.priority
                            }}</span>
                            <span class="text-[14px] m-[5px] text-title-color"><b>{{ item.replayName }}</b></span>
                            <span class=" text-text-li text-12">{{ $t("的目标复盘") }}</span>
                        </div>
                        <div class="cursor-pointer hidden md:block" @click="() => (item.isActive = !item.isActive)">
                            <span class="mr-[10px]">{{ item.isActive == true ? $t("收起") : $t("展开") }}</span>
                            <i class="replay-item-head-icon taskfont">&#xe705;</i>
                        </div>
                    </div>
                    <div class="flex">
                        <div class="replay-item-okr cursor-pointer" @click.stop="openOkrDetail(item.id)">
                            <div class="replay-item-okr-icon w-[25px] h-[15px]">O</div>
                            <div class="text-[#515A6E] text-14">{{ item.okrName }}</div>
                        </div>
                    </div>
                    <div class="replay-item-body" v-if="item.isActive">
                        <OkrReplayDetail :okrReplayList="item"></OkrReplayDetail>
                    </div>
                </div>
            </div>
            <OkrLoading v-if="onscrolloading" position='onscroll'></OkrLoading>
            <!-- OKR详情 -->
            <OkrDetailsModal ref="RefOkrDetails" :id="detailId" :show="okrDetailsShow" @close="() => {
                    okrDetailsShow = false
                }
                "></OkrDetailsModal>
        </div>
    </n-scrollbar>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import OkrReplayDetail from "@/views/components/OkrReplayDetails.vue"
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import * as http from "../../api/modules/replay"
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import OkrLoading from '../components/OkrLoading.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const addMultipleShow = ref(false)
const items = ref([])
const loadIng = ref(false)
const onscrolloading = ref(false)
const page = ref(1)
const last_page = ref(99999)

const okrDetailsShow = ref(false)
const detailId = ref(0)

const searchTime = ref(null)
const props = defineProps({
    searchObject: {
        type: String,
        default: "",
    }
})

watch(() => props.searchObject, (newValue) => {
    clearInterval(searchTime.value)
    searchTime.value = setInterval(() => {
        page.value = 1
        getData('search')
        clearInterval(searchTime.value)
    }, 300)
}, { deep: true })

//封装复盘列表
const loadResplayList = (data) => {
    items.value = data.map((itemData) => returnReplayItem(itemData))
}

//提取复盘信息
const returnReplayItem = (replay) => {
    return {
        id: replay.okr_id,
        isActive: false,
        replayName: replay.okr_alias.join(","),
        okrName: replay.okr_title,
        priority: replay.okr_priority,
        krList: replay.kr_history ? replay.kr_history : [],
        okrProgress: replay.okr_progress,
        review: replay.review,
        problem:replay.problem
    }
}

// 获取数据
const getData = (type) => {
    let serstatic = type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic) {
        // 获取复盘列表
        const data = {
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
        }
        if (serstatic) {
            loadIng.value = true
        } else if (type == 'onscrollsearch') {
            onscrolloading.value = true
        }
        http.getReplayList(data).then(({ data }) => {
            onscrolloading.value = false
            loadIng.value = false
            if (serstatic) {
                data.data ? loadResplayList(data.data) : []
            } else {
                ; (data.data || []).map((item) => {
                    items.value.push(returnReplayItem(item))
                })
            }
            last_page.value = data.last_page
        })
    }
}

const onScroll = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!onscrolloading.value && !loadIng.value) {
            if (last_page.value > page.value) {
                page.value++
                getData("onscrollsearch")
            }
        }
    }
}

// 状态
const pStatus = (p) => {
    return p == "P0" ? "span-1" : p == "P1" ? "span-2" : "span-3"
}

// 打开复盘
const openMultiple = () => {
    addMultipleShow.value = true
}

//查看okr详情
const openOkrDetail = (id) => {

    if (window.innerWidth < 768) {
        router.push({
            path: '/okrDetails',
            query: { data: id },
        })
    }
    else {
        detailId.value = id
         okrDetailsShow.value = true
    }
}

onMounted(() => {
    getData("search")
})

</script>

<style lang="less" scoped>
.okr-replay-main {
    @apply pt-24 flex flex-col gap-6;
}

.replay {
    width: 100%;
    height: 100%;
    @apply flex flex-col gap-3 md:gap-5;

    &-item {
        @apply bg-white rounded-lg h-auto p-24;

        &-head {
            @apply flex items-center justify-between text-[#87d068];

            &-icon {
                display: inline-block;
                transform: rotate(0deg);
                transition: transform 0.5s;
            }
        }

        &-okr {
            background: #f4f5f7;
            padding: 6px 16px;
            @apply mt-10 h-auto flex rounded-lg items-center;

            &-icon {
                @apply text-center text-10 leading-4 rounded-lg bg-[#87D068] text-[#87d068] mr-10;
                background: rgba(135, 208, 104, 0.2);
            }

            &-level {
                @apply bg-[#FF7070] text-10 text-white px-6 py-2 rounded-full scale-[0.8333] origin-center;
            }
        }

        &-body {
            @apply mt-14;
        }
    }

    &-item-active &-item-head {
        &-icon {
            transform: rotate(-180deg);
        }
    }
}

.span-1 {
    @apply bg-[#FF7070];
}

.span-2 {
    @apply bg-[#FC984B];
}

.span-3 {
    @apply bg-[#72A1F7];
}
</style>
