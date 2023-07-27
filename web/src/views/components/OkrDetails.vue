<template >
    <n-modal v-model:show="props.show" transform-origin="center" @after-enter="showDrawer" @after-leave="closeDrawer"
        :z-index="1">
        <n-card class="w-[1240px] relative" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <div class="flex">
                <div class="flex-1 flex flex-col pb-[64px]">
                    <div
                        class="flex h-[28px] items-center justify-between pb-[15px] border-solid border-0 border-b-[1px] border-[#F2F3F5] cursor-pointer">
                        <div class="flex items-center gap-4">
                            <n-popover v-if="detialData.completed == '0'" placement="bottom" trigger="click">
                                <template #trigger>
                                    <div v-if="detialData.completed == '0'"
                                        class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid "
                                        :class="detialData.completed == '0' || detialData.canceled == '0' ? 'border-[#A8ACB6]' : 'border-primary-color bg-primary-color'">
                                        <n-icon v-if="detialData.completed == '1'"
                                            :class="detialData.completed == '0' || detialData.canceled == '0' ? '' : 'text-primary-color'"
                                            size="14" :component="CheckmarkSharp" />
                                    </div>
                                </template>
                                <span class="cursor-pointer" @click="() => { console.log(123123) }">
                                    {{ detialData.canceled == '0' ? $t('取消目标') : $t('重启目标') }}
                                </span>
                            </n-popover>
                            <div v-else
                                class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid "
                                :class="detialData.completed == '0' || detialData.canceled == '0' ? 'border-[#A8ACB6]' : 'border-primary-color bg-primary-color'">
                                <n-icon v-if="detialData.completed == '1'"
                                    :class="detialData.completed == '0' || detialData.canceled == '0' ? '' : 'text-primary-color'"
                                    size="14" :component="CheckmarkSharp" />
                            </div>

                            <n-tag v-if="detialData.canceled == '1'">{{ $t('已取消') }} </n-tag>
                            <n-tag v-if="detialData.completed == '1'" type="success">{{ $t('已完成') }}</n-tag>
                            <p class="flex items-center text-text-li opacity-50 text-14">
                                {{ detialData.type == "1" ? $t('承诺型') : $t('挑战性') }} /
                                {{ detialData.ascription == "2" ? $t('个人') : $t('部门') }}
                            </p>
                        </div>
                        <div class="flex items-center gap-6">
                            <i class="taskfont icon-title text-[#A7ACB6]">&#xe779;</i>
                            <i class="taskfont icon-title text-[#A7ACB6]">&#xe679;</i>
                        </div>
                    </div>
                    <h3 class=" text-title-color mt-[28px] text-24 font-normal line-clamp-1">{{ detialData.title }}</h3>


                    <div class="mt-24 flex flex-col gap-4">
                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i class="taskfont icon-item text-[#A7ABB5]">&#xe6e4;</i>
                                <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('负责人') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14" v-if="detialData.alias">{{  detialData.alias[0] }}</p>
                        </div>


                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i class="taskfont icon-item text-[#A7ABB5]">&#xe6e8;</i>
                                <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('起止时间') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14">{{ utils.GoDate(detialData.start_at || 0) }} ~
                                {{ utils.GoDate(detialData.end_at || 0) }}</p>
                        </div>

                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i class="taskfont icon-item text-[#A7ABB5]">&#xe6ec;</i>
                                <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('优先级') }}</span>
                            </p>
                            <span class="span" :class="pStatus(detialData.priority)">{{ detialData.priority }}</span>
                        </div>

                    </div>

                    <h4 class="text-text-li text-16 font-medium mt-36">KR</h4>

                    <div class="flex flex-col mt-20 gap-6">
                        <div class="flex flex-col" v-for="(item, index) in detialData.key_results">
                            <div class="flex items-center">
                                <span class=" text-primary-color text-12 leading-5 mr-8">KR{{ index + 1 }}</span>
                                <h4 class=" text-title-color text-14 font-normal line-clamp-1">{{ item.title }}</h4>
                            </div>
                            <div class="flex items-center justify-between mt-8">
                                <div class="flex items-center">
                                    <div class="flex items-center gap-2">
                                        <n-avatar round :size="20"
                                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                        <n-avatar round :size="20"
                                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                    </div>
                                    <i class="taskfont icon-item ml-24 text-[#A7ABB5]">&#xe6e8;</i>
                                    <p class="flex-1 text-text-li text-14">{{ utils.GoDate(item.start_at || 0) }} ~
                                        {{ utils.GoDate(item.end_at || 0) }}</p>
                                </div>
                                <div class="flex items-center gap-6">
                                    <div class="flex items-center cursor-pointer"
                                        @click="handleSchedule(item.id, item.progress, item.progress_status)">
                                        <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle"
                                            :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                            color="var(--primary-color)" status="success" :percentage="item.progress" />
                                        <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
                                    </div>

                                    <div v-if="item.confidence == '0'" class="flex items-center cursor-pointer"
                                        @click="handleConfidence">
                                        <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ $t('信心') }}</p>
                                    </div>
                                    <div v-else class="flex items-center cursor-pointer" @click="handleConfidence">
                                        <i class="taskfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ item.confidence }}</p>
                                    </div>

                                    <div v-if="item.score == '0'" class="flex items-center cursor-pointer"
                                        @click="handleMark">
                                        <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ $t('评分') }}</p>
                                    </div>
                                    <div v-else class="flex items-center cursor-pointer" @click="handleMark">
                                        <img class="mr-6 -mt-2" src="@/assets/images/icon/fen.svg" />
                                        <p class="text-text-li opacity-50 text-12">{{ item.score }}{{ $t('分') }}</p>
                                    </div>

                                </div>
                            </div>
                        </div>

                    </div>

                    <div class="border-solid border-0 border-b-[1px] border-[#F2F3F5] my-36"></div>

                    <h4 class="text-text-li text-16 font-medium mb-12">{{ $t('对齐目标') }} <i
                            class="taskfont text-16 cursor-pointer text-[#A7ABB5]">&#xe779;</i></h4>

                    <AlignTarget :value="props.show" :id="props.id"></AlignTarget>

                </div>
                <div v-if="loadIng"
                    class="absolute top-[71px] bottom-[24px] left-0 right-0 z-10 bg-[rgba(255,255,255,0.8)] flex-1 flex justify-center items-center">
                    <n-spin size="small" />
                </div>
                <div
                    class="w-[414px] flex flex-col flex-initial border-solid border-0 border-l-[2px] border-[#F2F3F5] pl-24 ml-24">
                    <div
                        class="flex items-center justify-between border-solid border-0 border-b-[1px] border-[#F2F3F5] pb-[15px] h-[28px]">
                        <ul v-if="!loadIng" class="flex items-center gap-8">
                            <li class="li-nav" :class="navActive == 0 ? 'active' : ''" @click="handleNav(0)">{{ $t('评论') }}
                            </li>
                            <li class="li-nav" :class="navActive == 1 ? 'active' : ''" @click="handleNav(1)">{{ $t('动态') }}
                            </li>
                            <li class="li-nav" :class="navActive == 2 ? 'active' : ''" @click="handleNav(2)">{{ $t('复盘') }}
                            </li>
                        </ul>
                        <ul v-else></ul>
                        <i class="taskfont text-16 cursor-pointer text-[#A7ABB5]" @click="closeModal">&#xe6e5;</i>
                    </div>
                    <div class="flex-auto relative">
                        <div class="absolute top-[24px] bottom-0 left-0 right-0">
                            <div class="text-center" v-show="navActive == 0">
                                <DialogWrappers />
                                <n-spin size="small" />
                            </div>
                            <n-scrollbar v-if="navActive == 1">
                                <div class="flex text-start">
                                    <n-avatar round :size="28" class="mr-8"
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                    <div class="flex flex-col gap-3">
                                        <p class="text-12 leading-3 text-text-li">张三<span class="opacity-60 ml-8">08-06
                                                16:45</span></p>
                                        <h4 class="text-14 leading-[14px] text-title-color font-normal">责任人打分 <span
                                                class=" font-medium">搭建一个团队使用的客户成功系统</span></h4>
                                    </div>
                                </div>
                            </n-scrollbar>
                            <n-scrollbar v-if="navActive == 2">
                                <p class="cursor-pointer"> <i class="taskfont mr-4 text-16 text-text-tips">&#xe6f2;</i><span
                                        class="text-14 text-text-tips">{{ $t('添加复盘') }}</span></p>
                                <div class="flex flex-col gap-3 mt-20">
                                    <div class="flex items-center justify-between">
                                        <h4 class="text-14 text-text-li font-normal">复盘 (2023/08/05)</h4>
                                        <p class="text-12 text-text-li opacity-60">08-04 12:45</p>
                                    </div>
                                    <div class="flex items-center justify-between">
                                        <h4 class="text-14 text-text-li font-normal">复盘 (2023/08/05)</h4>
                                        <p class="text-12 text-text-li opacity-60">08-04 12:45</p>
                                    </div>
                                    <div class="flex items-center justify-between">
                                        <h4 class="text-14 text-text-li font-normal">复盘 (2023/08/05)</h4>
                                        <p class="text-12 text-text-li opacity-60">08-04 12:45</p>
                                    </div>
                                </div>
                            </n-scrollbar>
                        </div>
                    </div>
                </div>
            </div>


        </n-card>

    </n-modal>
</template>
<script setup lang="ts">
import { CheckmarkSharp } from '@vicons/ionicons5'
import { getOkrDetail } from '@/api/modules/okrList'
import AlignTarget from "@/views/components/AlignTarget.vue";
import { ResultDialog } from "@/api"
import utils from '@/utils/utils';

const navActive = ref(0)
const dialogWrappersApp = ref()
const loadIng = ref(false)
const detialData = ref<any>({})

const emit = defineEmits(['close', 'schedule', 'confidence', 'mark'])

const props = defineProps({
    show: {
        type: Boolean,
        default: false,
    },
    id: {
        type: Number,
        default: 0,
    },
})
watch(() => props.show, (newValue) => {
    if (newValue) {
        getDetail()
    }
})

const getDetail = () => {
    const upData = {
        id: props.id,
    }
    loadIng.value = true
    getOkrDetail(upData)
        .then(({ data }) => {
            console.log(data);
            detialData.value = data
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}



const handleNav = (index) => {
    navActive.value = index
}

//更新进度
const handleSchedule = (id, progress, progress_status) => {
    emit('schedule', id, progress, progress_status)
}

const handleConfidence = () => {
    emit('confidence')
}

const handleMark = () => {
    emit('mark')
}

const closeModal = () => {
    emit('close')
}

// 加载聊天组件
const loadDialogWrappers = () => {
    nextTick(() => {
        if (!window.Vues) return false;
        dialogWrappersApp.value && dialogWrappersApp.value.$destroy();
        dialogWrappersApp.value = new window.Vues.Vue({
            el: document.querySelector('DialogWrappers'),
            store: window.Vues.store,
            render: (h: any) => {
                return h(window.Vues?.components?.DialogWrapper, {
                    props: {
                        dialogId: 59
                    }
                }, [h("div", { slot: "head" })])
            }
        });
        window.addEventListener('apps-unmount', function () {
            dialogWrappersApp.value.$destroy()
        })
    })
}

// 显示
const showDrawer = () => {
    loadDialogWrappers()
}

// 关闭
const closeDrawer = () => {
    dialogWrappersApp.value && dialogWrappersApp.value.$destroy();
}

const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

defineExpose({
    getDetail
})
</script>
<style lang="less" >
.icon-title {
    @apply text-16 cursor-pointer;
}

.icon-item {
    @apply text-18 mr-8;
}

.span {
    @apply text-14 text-white px-6 py-4 rounded-full flex items-center leading-3 shrink-0;
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

.align-target {
    @apply flex flex-col gap-6 mt-8;

    .a-t-list {
        @apply flex items-center;

        .a-t-tab {
            @apply flex items-center flex-initial flex-shrink-0 relative;

            .a-t-tabs {
                @apply rounded-full flex items-center justify-center text-12 origin-center leading-3 flex-shrink-0;
            }

            .a-t-tab-b {
                @apply bg-[#F3F0FF] ml-4 text-12 text-[#4D3EF5] py-2 px-4 flex items-center justify-center leading-3 rounded flex-shrink-0;
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

.li-nav {
    @apply list-none cursor-pointer text-text-li opacity-50 text-14;
}

.li-nav.active {
    @apply text-title-color text-16 opacity-100;
}
</style>
