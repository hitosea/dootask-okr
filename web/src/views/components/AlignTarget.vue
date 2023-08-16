<template >
    <div class="align-target">
        <div class="a-t-list" v-for="item in dataList" v-if="dataList">
            <div class="a-t-tab flex-[0] text-[--n-color-target] mt-auto mb-[4px]">
                <span class="a-t-tabs w-[22px] h-[14px] bg-[rgba(135,208,104,0.2)] scale-[0.8333]">{{ item.prefix
                }}</span>
                <n-tooltip trigger="hover" v-if="item.alias[0]">
                    <template #trigger>
                        <span v-if="item.alias[0]" class="a-t-tab-b "
                            :class="item.prefix == 'O' ? 'text-[#4D3EF5] bg-[#F3F0FF]' : 'text-[#0066FF] bg-[#EDF4FF]'">{{
                                item.alias[0] }} {{ item.alias.length > 1 ? "+1" : '' }}</span>
                    </template>
                    {{ item.alias.join(',') }}
                </n-tooltip>

                <div v-if="item.align_objective"
                    class="absolute border-solid border-0 border-l border-t border-text-tips h-[calc(100%-2px)] w-[calc(100%-10px)] top-[-14px] left-[10px] rounded-tl-lg">
                    <div class="w-[4px] h-1 bg-text-tips rotate-45 absolute right-0 top-[-3px]"></div>
                    <div class="w-[4px] h-1 bg-text-tips -rotate-45 absolute right-0 top-[1px]"></div>
                </div>
            </div>

            <h3 class="a-t-title cursor-pointer" v-if="!item.align_objective" @click="handleDetail(item.id, item.userid)">{{
                item.title }}</h3>
            <div class="flex-1 overflow-hidden" v-else>
                <h4 class="a-t-title-s max-w-[90%]">{{ item.align_objective }}</h4>
                <h3 class="a-t-title max-w-[90%] cursor-pointer" @click="handleDetail(item.parent_id, item.userid)"
                    :class="item.deleted_at == null ? '' : 'line-through opacity-25'">{{
                        item.title }}</h3>
            </div>
            <div v-if="props.progressShow" class="flex ml-auto min-w-[55px] items-center cursor-pointer"
                :class="cancelShow ? 'mr-24' : ''">
                <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle" :show-indicator="false"
                    :offset-degree="180" :stroke-width="15" color="var(--primary-color)" status="success"
                    :percentage="item.progress" />
                <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
            </div>
            <n-tooltip trigger="hover" v-if="props.cancelShow">
                <template #trigger>
                    <i v-if="props.cancelShow" class="taskfont cursor-pointer text-text-tips ml-auto hidden md:block"
                        @click="alignCancel(item.id)">
                        &#xe680;</i>
                </template>
                {{ $t('取消对齐') }}
            </n-tooltip>
        </div>

        <div v-else class="flex flex-initial items-center justify-center py-[60px]">
            <div >
                <img class="w-80" src="@/assets/images/icon/notData.svg" />
                <p class="mt-10">{{ $t('暂无数据') }}</p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useMessage, useDialog } from "naive-ui"
import { getAlignDetail, getAlignCancel } from '@/api/modules/okrList'

const loadIng = ref(false)
const message = useMessage()
const dialog = useDialog()
const dataList = ref<any>([])

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
})


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

const alignCancel = (itemID) => {
    dialog.info({
        title: $t('提示'),
        content: $t('你确定要取消对齐吗？'),
        positiveText: $t('确定'),
        negativeText: $t('取消'),
        onPositiveClick: () => {
            const upData = {
                okr_id: props.id,
                align_okr_id: itemID,
            }
            loadIng.value = true
            getAlignCancel(upData)
                .then(({ msg }) => {
                    message.success($t('修改成功'))
                    getList();
                    emit('unalign', props.id)
                })
                .catch(({ msg }) => {
                    message.error(msg)
                })
                .finally(() => {
                    loadIng.value = false
                })
        },
        onNegativeClick: () => {

        }
    })

}

const handleDetail = (id, userid) => {
    emit('openDetail', id, userid)
}

watch(() => props.value, (newValue) => {
    if (newValue) {
        getList()
    }
}, { immediate: true })

defineExpose({
    getList
})

</script>

<style lang="less" scoped>
.align-target {

    @apply flex flex-col gap-6 w-full;

    .a-t-list {
        @apply flex items-center overflow-hidden;

        .a-t-tab {
            @apply flex items-center flex-initial flex-shrink-0 relative;

            .a-t-tabs {
                @apply rounded-full flex items-center justify-center text-12 origin-center leading-3 flex-shrink-0;
            }

            .a-t-tab-b {
                @apply ml-4 text-12 py-2 px-4 flex items-center justify-center leading-3 rounded flex-shrink-0;
            }
        }

        .a-t-title {
            @apply text-title-color text-14 flex-auto ml-4 line-clamp-1;
        }

        .a-t-title-s {
            @apply text-text-tips text-12 flex-auto ml-4 font-normal line-clamp-1;
        }
    }
}</style>
