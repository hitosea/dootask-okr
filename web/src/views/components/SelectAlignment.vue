<template >
    <n-modal v-model:show="props.value" transform-origin="center" :on-after-leave="handleClear" :mask-closable="false">
        <n-card
            class="max-w-[640px] w-full h-[calc(100vh-52px)] md:h-auto top-[52px] md:top-0 rounded-none rounded-t-xl md:rounded-xl"
            :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header>
                <h3 class="text-16 md:text-18 text-title-color ">{{ $t('选择对齐目标') }}</h3>
            </template>
            <template #header-extra>
                <i class="taskfont text-16 cursor-pointer text-[#A7ABB5]" @click="handleClose">&#xe6e5;</i>
            </template>
            <div class="flex flex-col flex-1 overflow-hidden">
                <n-input v-model:value="searchName" placeholder="搜索" :autofocus="false" :on-update:value="searchInput">
                    <template #prefix>
                        <i class="taskfont"> &#xe6f8;</i>
                    </template>
                </n-input>

                <n-checkbox-group  v-model:value="cities" class="flex mt-[16px] overflow-hidden" v-if="okrList.length > 0">
                    <div class="flex flex-col flex-auto ">
                        <n-scrollbar class="max-h-full md:max-h-[300px]" :on-scroll="onScroll">
                            <div class="align-okr" v-for="(item, index) in okrList">
                                <div class="object-field ">
                                    <n-checkbox :value="item.id" />
                                    <i class="taskfont" :class="[item.key_results == null ? 'text-[#ededed] cursor-not-allowed' : 'text-text-tips cursor-pointer',
                                        openList.indexOf(index) != -1 ? 'active' : ''
                                        ]
                                        " @click="handleOpen(index, item.key_results)">&#xe745;</i>
                                    <span class="span " :class="pStatus(item.priority)">{{ item.priority
                                    }}</span>
                                    <h3 class="text-14 text-title-color font-normal line-clamp-1 ml-4 pr-16">
                                        {{ item.title }}
                                    </h3>
                                </div>
                                <div class="flex-1 w-full" v-if="openList.indexOf(index) != -1">
                                    <div v-for="(items, indexs) in item.key_results"
                                        class="flex items-center hover:bg-[#F4F5F7] pl-[43px] py-[9px] rounded">
                                        <n-checkbox :value="items.id" />
                                        <span class="ml-12 mr-4 text-12 leading-3 text-text-tips">KR{{ indexs + 1 }}</span>
                                        <p class=" flex-1 text-14 leading-4 text-title-color line-clamp-1 pr-16">
                                            {{ items.title }}
                                        </p>
                                    </div>
                                </div>
                            </div>


                            <!-- <div class="flex justify-center">
                                <n-spin size="small" :show="loadIng" ></n-spin>
                            </div> -->
                        </n-scrollbar>
                    </div>
                </n-checkbox-group>
                <div v-else class="flex flex-initial items-center justify-center py-[60px]">
                    <div v-if="searchName !='' && okrList.length ==0">
                        <img class="w-80" src="@/assets/images/icon/notSearch.svg" />
                        <p class="mt-10 text-[#515A6E] opacity-50 text-center">{{ $t('没有找到匹配的结果') }}</p>     
                    </div>
                    <div v-if="searchName =='' && okrList.length ==0">
                        <img class="w-80" src="@/assets/images/icon/notData.svg" />
                        <p class="mt-10 text-[#515A6E] opacity-50 text-center">{{ $t('暂无数据') }}</p>     
                    </div>
                </div>
            </div>
            <template #footer>
                <div class="button-box">
                    <p>{{ $t('已选：') }}
                        <span class=" text-primary-color" v-if="cities">{{ cities.length }}</span>
                        <span v-else>0</span>
                        {{ $t('项') }}
                    </p>
                    <div class="flex items-center gap-4">
                        <n-button class="hidden md:flex" :loading="loadIng" strong secondary @click="handleClose">
                            {{ $t('取消') }}
                        </n-button>
                        <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                            {{ $t('确定') }}
                        </n-button>
                    </div>

                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { getAlignList } from '@/api/modules/created'
import utils from "@/utils/utils";

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
    editData: {
        type: Object,
        default: null,
    },
})


const emit = defineEmits(['close', 'submit'])

const loadIng = ref(false)
const cities = ref<any>([])
const searchName = ref('')
const openList = ref([])
const okrList = ref<any>([])
const page = ref(1)
const last_page = ref(999999)

watch(() => props.value, (newValue) => {
    if (newValue) {
        cities.value = utils.cloneJSON(props.editData)
        getList('')
    }
})

const onScroll = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!loadIng.value) {
            page.value++
            getList('')
        }
    }
}

const searchInput = (e) => {
    searchName.value = e
    page.value = 1
    getList('search');
}

const handleClear = () => {
    cities.value = []
    openList.value = []
    okrList.value = []
    last_page.value = 999999
    page.value = 1
    searchName.value = ''
}

const handleClose = () => {
    emit('close')
}

const handleSubmit = () => {
    emit('submit', cities.value)
    emit('close')
}

const getList = (type) => {
    if (last_page.value >= page.value || type == 'search') {
        const data = {
            objective: searchName.value,
            page: page.value,
            page_size: 20,
        }
        loadIng.value = true
        getAlignList(data).then(({ data }) => {
            if (type == 'search') {
                okrList.value = data.data
            }
            else {
                data.data.map(item => {
                    okrList.value.push(item)
                })
            }
            last_page.value = data.last_page
            loadIng.value = false
            if (cities.value != null) {
                okrList.value.map((item, index) => {
                    item.key_results?.map(items => {
                        if (cities.value.indexOf(items.id) != -1) {
                            openList.value.push(index)
                        }
                    })
                })
            }
        })
    }
}

const handleOpen = (index, keyResults) => {
    if (keyResults == null) return;
    if (openList.value.indexOf(index) == -1) {
        openList.value.push(index)
    } else {
        openList.value.splice(openList.value.indexOf(index), 1)
    }
}


const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

</script>
<style lang="less" scoped>
.align-okr {
    @apply flex flex-col items-start shrink-0;

    .object-field {
        @apply flex flex-initial items-center shrink-0 py-9;

        i {
            @apply text-12 ml-12 transition-all;
        }

        i.active {
            @apply rotate-90;
        }

        .span {
            @apply text-12 md:text-14 ml-4 font-medium text-white px-6 py-2 rounded-full origin-center flex items-center leading-3 shrink-0;
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

    }
}

.button-box {
    @apply flex items-center justify-between mt-0 ;
}
:deep(.n-card__content){
    @apply flex-auto flex overflow-hidden;
}
:deep(.n-card__footer){
    @apply flex-initial pt-12;
    border-top: solid 1px #F2F2F2;
}
</style>
