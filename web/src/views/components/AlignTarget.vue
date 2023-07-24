<template >
    <n-modal v-model:show="props.value" transform-origin="center">
        <n-card class="w-[640px]" :title="$t('对齐目标')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <div class="align-target">
                <div class="a-t-list">
                    <div class="a-t-tab flex-[0] text-[--n-color-target]">
                        <span class="a-t-tabs w-[22px] h-[14px] bg-[rgba(135,208,104,0.2)] scale-[0.8333]">O</span>
                        <span class="a-t-tab-b">财务部</span>
                    </div>
                    <h3 class="a-t-title">本季度实现1000万元的收入目标本季度实现1000万万元的收入目标</h3>
                    <n-tooltip trigger="hover">
                        <template #trigger>
                            <i class="taskfont cursor-pointer text-text-tips ml-auto"> &#xe680;</i>
                        </template>
                        {{ $t('取消对齐') }}
                    </n-tooltip>
                </div>
                <div class="a-t-list">
                    <div class="a-t-tab flex-[0] text-[--n-color-target] mt-auto mb-[4px]">
                        <span class="a-t-tabs w-[22px] h-[14px] bg-[rgba(135,208,104,0.2)] scale-[0.8333]">KR</span>
                        <span class="a-t-tab-b">财务部</span>
                        <div class="absolute border-solid border-0 border-l border-t border-text-tips h-[calc(100%-2px)] w-[calc(100%-10px)] top-[-14px] left-[10px] rounded-tl-lg">
                             <div class="w-[4px] h-1 bg-text-tips rotate-45 absolute right-0 top-[-3px]"></div>
                             <div class="w-[4px] h-1 bg-text-tips -rotate-45 absolute right-0 top-[1px]"></div>
                        </div>
                    </div>
                    <div>
                        <h4 class="a-t-title-s ">完善团队制度，为了更好的服务客户</h4>
                        <h3 class="a-t-title line-through opacity-25">本季度实现1000万元的收入目标本季度实现1000万万元的收入目标</h3>
                    </div>
                    <n-tooltip trigger="hover">
                        <template #trigger>
                            <i class="taskfont cursor-pointer text-text-tips ml-auto"> &#xe680;</i>
                        </template>
                        {{ $t('取消对齐') }}
                    </n-tooltip>
                </div>
            </div>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import { Close } from "@vicons/ionicons5"

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
})

const emit = defineEmits(['close'])


const handleClose = () => {
    emit('close')
}
</script>

<style lang="less" >
.align-target {

    @apply flex flex-col gap-6;

    .a-t-list {
        @apply flex items-center;

        .a-t-tab {
            @apply flex items-center flex-initial flex-shrink-0 relative;

            .a-t-tabs {
                @apply rounded-full flex items-center justify-center text-12 origin-center leading-3 flex-shrink-0;
            }

            .a-t-tab-b {
                @apply bg-[#F3F0FF] ml-4 text-12 text-[#4D3EF5] py-2 px-4 flex items-center justify-center leading-3 rounded flex-shrink-0 ;
            }
        }

        .a-t-title {
            @apply text-title-color text-14 truncate flex-auto ml-4;
        }

        .a-t-title-s {
            @apply text-text-tips text-12 truncate flex-auto ml-4 font-normal;
        }
    }
}
</style>

