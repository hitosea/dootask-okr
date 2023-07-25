<template >
    <n-modal v-model:show="props.value" transform-origin="center" :on-after-leave="handleClear">
        <n-card class="w-[640px]" :title="$t('选择对齐目标')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <n-input v-model:value="searchName" placeholder="搜索" :on-update:value="searchInput">
                    <template #prefix>
                        <i class="taskfont"> &#xe6f8;</i>
                    </template>
                </n-input>

                <n-checkbox-group v-model:value="cities" class="mt-[16px]">
                    <div class="flex flex-col ">
                        <n-scrollbar class="max-h-[300px]" :on-scroll="onScroll">
                            <div class="align-okr" v-for="(item, index) in okrList">
                                <div class="object-field ">
                                    <n-checkbox :value="item.id" />
                                    <i class="taskfont" @click="handleOpen(index, $event.target)">&#xe745;</i>
                                    <span class="span scale-[0.8333]" :class="pStatus(item.priority)">{{ item.priority
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
                                        <p class=" flex-1 text-14 leading-3 text-title-color line-clamp-1 pr-16">
                                            {{ items.title }}
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div class="flex justify-center">
                                <n-spin size="small" v-if="loadIng" />
                            </div>
                        </n-scrollbar>
                    </div>
                </n-checkbox-group>
            </div>
            <template #footer>
                <div class="button-box">
                    <p>已选：{{ cities.length }} 项</p>
                    <div class="flex items-center gap-4">
                        <n-button :loading="loadIng" type="default" @click="handleClose">
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
import { Close } from "@vicons/ionicons5"
import { getAlignList } from '@/api/modules/created'

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
})

watch(() => props.value, (newValue) => {
    if (newValue) {
        getList('')
    }
})


const emit = defineEmits(['close', 'submit'])

const loadIng = ref(false)
const cities = ref([])
const searchName = ref('')
const openList = ref([])
const okrList = ref<any>([])
const page = ref(1)
const last_page = ref(999999)


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
    handleClear()
    emit('close')
}

const handleSubmit = () => {
    emit('submit', cities.value)
    handleClear()
    emit('close')
}



const getList = (type) => {
    if (last_page.value >= page.value || type == 'search') {
        const data = {
            objective: searchName.value,
            page: page.value,
            page_size: 10,
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
        })
    }
}

const handleOpen = (index, event) => {
    if (openList.value.indexOf(index) == -1) {
        openList.value.push(index)
        event.classList.add('active')
    } else {
        openList.value.splice(openList.value.indexOf(index), 1)
        event.classList.remove('active')
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
            @apply text-12 text-text-tips ml-12 cursor-pointer transition-all;
        }

        i.active {
            @apply rotate-90;
        }

        .span {
            @apply text-10 text-white px-6 py-2 rounded-full origin-center flex items-center leading-3 shrink-0;
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
    @apply flex items-center justify-between mt-0;
}
</style>