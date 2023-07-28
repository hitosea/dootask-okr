<template>
    <n-scrollbar>
        <div class="py-24">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <OkrNotDatas v-else-if="items.length == 0" :msg="$t('暂无复盘')"></OkrNotDatas>
            <div v-else class="replay">
                <div
                    v-for="(item, index) in items"
                    :key="index"
                    :class="{ 'replay-item': true, 'replay-item-active': item.isActive }"
                    @click="openMultiple"
                >
                    <div class="replay-item-head">
                        <div>
                            <span class="replay-item-okr-level scale-[0.8333]" :class="pStatus(item.priority)">{{ item.priority }}</span>
                            <span class="text-[14px] m-[5px] text-[#333333]" ><b>{{ item.replayName }}</b></span>
                            <span class="text-[#515a6e] text-12">{{ $t('的目标复盘') }}</span>
                        </div>
                        <div class="cursor-pointer" @click="() => (item.isActive = !item.isActive)">
                            <span class="mr-[10px]">{{ item.isActive == true ? $t('收起') : $t('展开') }}</span>
                            <i class="replay-item-head-icon taskfont">&#xe705;</i>
                        </div>
                    </div>
                    <div class="flex">
                        <div class="replay-item-okr">
                            <div class="replay-item-okr-icon w-[25px] h-[15px]">O</div>
                            <div class="text-[#515A6E] text-14">{{ item.okrName }}</div>
                        </div>
                    </div>
                    <div class="replay-item-body" v-if="item.isActive">
                        <OkrReplayDetail :okrReplayList="item"></OkrReplayDetail>
                    </div>
                </div>
            </div>
        </div>
    </n-scrollbar>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import OkrReplayDetail from "@/views/components/OkrReplayDetails.vue"
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"
import * as http from "../../api/modules/replay"

const addMultipleShow = ref(false)
const items = ref([])
const loadIng = ref(false)
const page = ref(1)
const last_page = ref(99999)

//封装复盘列表
const loadResplayList = (data)=>{
    items.value = data.map((itemData) => (returnReplayItem(itemData)));    
}

//提取复盘信息
const returnReplayItem = (replay)=>{
    return {
        id:replay.okr_id,
        isActive: false,
        replayName: replay.okr_alias.join(","),
        okrName: replay.okr_title,
        priority: replay.okr_priority,
        krList: (replay.kr_history?replay.kr_history:[]),
        okrProgress:replay.okr_progress,
        review: replay.review
  }
}

// 获取数据
const getData = (type) => {
    if (last_page.value >= page.value || type == 'search') {
        // 获取复盘列表
        const data = {
            page: page.value,
            page_size: 10,
        }
        loadIng.value = true
        http.getReplayList(data).then(({ data }) => {
            if (type == 'search') {
                data.data?loadResplayList(data.data):[]
            }
            else {
                if (data.data) {
                    data.data.map(item => {
                        items.value.push(returnReplayItem(item))
                    })
                }
            }
            loadIng.value = false
            last_page.value = data.last_page
        }) 
    }    
}

// 状态
const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}


// 打开复盘
const openMultiple = () => {
    addMultipleShow.value = true;
}

onMounted(()=>{
    getData('')
})
</script>
<style lang="less" scope>
.replay {
    width: 100%;
    height: 100%;
    @apply flex flex-col gap-5;
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
