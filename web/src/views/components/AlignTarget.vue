<template>
    <div class="align-target">
        <div class="a-t-list" v-for="item in dataList" v-if="dataList">
            <div class="a-t-tab flex-[0] text-[--n-color-target] mt-auto mb-[4px]">
                <span class="a-t-tabs text-[#8BCF70] w-[24px] h-[16px] bg-[rgba(135,208,104,0.2)]">
                    {{props.active == 0 ? item.prefix : item.align_objective_prefix }}
                </span>
                <n-tooltip trigger="hover" v-if="item?.alias && item?.alias[0]">
                    <template #trigger>
                        <span v-if="item.alias[0]" class="a-t-tab-b " :class="item.prefix == 'O' ? 'text-[#4D3EF5] bg-[#F3F0FF]' : 'text-[#0066FF] bg-[#EDF4FF]'">{{
            item.alias[0] }} {{ item.alias.length > 1 ? '+' + (item.alias.length - 1) : '' }}</span>
                    </template>
                    {{ item.alias.join(',') }}
                </n-tooltip>

                <div v-if="item.align_objective"
                    class="absolute border-solid border-0 border-l border-t border-text-tips h-[calc(100%-2px)] w-[calc(100%-10px)] top-[-14px] left-[10px] rounded-tl-lg">
                    <div class="w-[4px] h-1 bg-text-tips rotate-45 absolute right-0 top-[-3px]"></div>
                    <div class="w-[4px] h-1 bg-text-tips -rotate-45 absolute right-0 top-[1px]"></div>
                </div>
            </div>

            <h3 class="a-t-title cursor-pointer w-[10px] mr-[36px]" :class="item.deleted_at ? 'line-through  opacity-50' : ''" v-if="!item.align_objective"
                @click="handleDetail(item.id, item.userid, item.deleted_at)">{{ item.title }}</h3>
            <div class="flex-1 overflow-hidden" v-else>
                <h4 class="a-t-title-s mr-[36px]">{{ props.active == 0 ? item.align_objective :$t('被对齐')+ item.align_objective_prefix +"：" + item.align_objective }}</h4>
                <h3 class="a-t-title mr-[36px] cursor-pointer" @click="handleDetail(item.parent_id, item.userid, item.deleted_at)"
                    :class="item.deleted_at == null ? '' : 'line-through opacity-50'">{{ item.title }}</h3>
            </div>
            <div v-if="props.progressShow" class="flex ml-auto min-w-[55px] items-center cursor-pointer" :class="cancelShow ? 'md:mr-24' : ''">
                <n-progress class="-mt-7 mr-[6px]" style="width: 15px; " type="circle" :show-indicator="false" :offset-degree="180" :stroke-width="15"
                    :color="colorStatus(item.progress_status)" status="success" :percentage="item.progress" />
                <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
            </div>
            <n-tooltip trigger="hover" v-if="props.cancelShow">
                <template #trigger>
                    <i v-if="props.cancelShow" class="okrfont cursor-pointer text-[#A7ABB5] ml-auto hidden md:block" @click="alignCancel(item.id)">
                        &#xe680;</i>
                </template>
                {{ $t('取消对齐') }}
            </n-tooltip>
        </div>

        <div v-else class="flex flex-initial items-center justify-center">
            <div>
                <img class="w-[60px]" :src="utils.apiUrl(notDataSvg)" />
                <p class="mt-10 text-[#515A6E] opacity-50 text-center">{{ $t('暂无数据') }}</p>
            </div>
        </div>

        <div v-if="loadIng" class="flex justify-center absolute bg-[rgba(255,255,255,0.6)] left-0 right-0 top-0 bottom-0">
            <n-spin size="small" :show="loadIng"></n-spin>
        </div>
        <InfoModal v-model:show="infoShow" :title="infoTitle" :content="infoContent" @submit="handleSubmit" @close="handleCancel"></InfoModal>
    </div>
</template>

<script setup lang="ts">
import { useMessage } from "@/utils/messageAll"
import { getAlignDetail, getAlignCancel, getByAlignDetail } from '@/api/modules/okrList'
import InfoModal from "./InfoModal.vue";
import utils from "@/utils/utils";
import notDataSvg from "@/assets/images/icon/notData.svg";


const loadIng = ref(false)
const message = useMessage()
const dataList = ref<any>([])
const infoShow = ref(false)
const infoTitle = ref('')
const infoContent = ref('')
const infoId = ref(0)

const emit = defineEmits(['unalign', 'openDetail'])

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
    cancelShow: {
        type: Boolean,
        default: false,
    },
    progressShow: {
        type: Boolean,
        default: false,
    },
    id: {
        type: Number,
        default: 0,
    },
    userid: {
        type: Number,
        default: 0,
    },
    active: {
        type: Number,
        default: 0,
    },
})

//获取对齐的目标
const getList = () => {
    const upData = {
        id: props.id,
    }
    loadIng.value = true
    getAlignDetail(upData)
        .then(({ data }) => {
            dataList.value = data
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

//获取被对齐的目标
const getByList = () => {
    const upData = {
        id: props.id,
    }
    loadIng.value = true
    getByAlignDetail(upData)
        .then(({ data }) => {
            dataList.value = data
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

const alignCancel = (itemID) => {
    infoShow.value = true
    infoTitle.value = $t('取消对齐')
    infoContent.value = $t('你确定要取消对齐吗？')
    infoId.value = itemID
}
//确定
const handleSubmit = () => {
    const upData = {
        okr_id: props.id,
        align_okr_id: infoId.value,
    }
    loadIng.value = true
    getAlignCancel(upData)
        .then(({ msg }) => {
            message.success($t('修改成功'))
            handleCancel();
            getList();
            emit('unalign', props.id)
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}


//取消
const handleCancel = () => {
    infoShow.value = false
    infoTitle.value = ''
    infoContent.value = ''
    infoId.value = 0
}

//颜色判断
const colorStatus = (color) => {
    let result = ''
    if (color == 1 || color == 0) {
        result = '#8BCF70'
    }
    if (color == 2) {
        result = '#FFA25A'
    }
    if (color == 3) {
        result = '#FF7070'
    }
    return result
}

const handleDetail = (id, userid, deleted_at) => {
    if (deleted_at) return
    emit('openDetail', id, userid, 'open')
}

watch(() => props.value, (newValue) => {
    if (newValue) {
        getList()
    }
}, { immediate: true })

watch(() => props.active, (newValue, oldValue) => {
    if (newValue != oldValue) {
        if (newValue == 0) {
            getList()
        } else {
            getByList()
        }
    }
})


defineExpose({
    getList
})

</script>

<style lang="less" scoped>
.align-target {

    @apply flex flex-col gap-4 w-full overflow-hidden relative min-h-24;

    .a-t-list {
        @apply flex items-center overflow-hidden;

        .a-t-tab {
            @apply flex items-center flex-initial flex-shrink-0 relative;

            .a-t-tabs {
                @apply rounded-full flex items-center justify-center text-12 font-medium leading-3 flex-shrink-0;
            }

            .a-t-tab-b {
                @apply ml-4 text-12 py-2 px-4 flex items-center justify-center leading-3 rounded flex-shrink-0;
            }
        }

        .a-t-title {
            @apply text-text-li text-left text-14 flex-auto ml-4 truncate font-medium;
        }

        .a-t-title-s {
            @apply text-text-tips text-left text-12 flex-auto ml-4 font-normal line-clamp-1;
        }
    }
}
</style>
