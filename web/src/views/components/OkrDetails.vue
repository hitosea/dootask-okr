<template >
        <n-modal v-model:show="show" transform-origin="center" @after-enter="showDrawer"  @after-leave="closeDrawer" :z-index="1">
            <n-card class="w-[1240px]" :bordered="false" size="huge" role="dialog" aria-modal="true">
                <div class="flex">
                    <div class="flex-1 flex flex-col pb-[64px]">
                        <div
                            class="flex items-center justify-between pb-[15px] border-solid border-0 border-b-[1px] border-[#F2F3F5] cursor-pointer">
                            <div class="flex items-center gap-4">
                                <n-popover placement="bottom" trigger="click">
                                    <template #trigger>
                                        <div
                                            class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid border-[#A8ACB6;]">
                                            <n-icon class=" text-primary-color" size="14" :component="CheckmarkSharp" />
                                        </div>
                                    </template>
                                    <span class="cursor-pointer" @click="() => { console.log(123123) }">重启目标</span>
                                </n-popover>
    
                                <n-tag> 已取消 </n-tag>
                                <p class="flex items-center text-text-li opacity-50 text-14">承诺型 / 个人</p>
                            </div>
                            <div class="flex items-center gap-6">
                                <i class="taskfont icon-title text-[#A7ACB6]">&#xe779;</i>
                                <i class="taskfont icon-title text-[#A7ACB6]">&#xe679;</i>
                            </div>
                        </div>
                        <h3 class=" text-title-color mt-[28px] text-24 font-normal line-clamp-1">开创新绩效考核制度，打造企业持续绩效考核模式</h3>
    
    
                        <div class="mt-24 flex flex-col gap-4">
                            <div class="flex items-center">
                                <p class="flex items-center w-[115px]">
                                    <i class="taskfont icon-item text-[#A7ABB5]">&#xe6e4;</i>
                                    <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('负责人') }}</span>
                                </p>
                                <p class="flex-1 text-text-li text-14">张三</p>
                            </div>
    
    
                            <div class="flex items-center">
                                <p class="flex items-center w-[115px]">
                                    <i class="taskfont icon-item text-[#A7ABB5]">&#xe6e8;</i>
                                    <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('起止时间') }}</span>
                                </p>
                                <p class="flex-1 text-text-li text-14">2023/07/01~2023/09/30</p>
                            </div>
    
                            <div class="flex items-center">
                                <p class="flex items-center w-[115px]">
                                    <i class="taskfont icon-item text-[#A7ABB5]">&#xe6ec;</i>
                                    <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('优先级') }}</span>
                                </p>
                                <span class="span span-1">P0</span>
                            </div>
    
                        </div>
    
                        <h4 class="text-text-li text-16 font-medium mt-36">KR</h4>
    
                        <div class="flex flex-col mt-20 gap-6">
                            <div class="flex flex-col">
                                <div class="flex items-center">
                                    <span class=" text-primary-color text-12 mr-8">KR1</span>
                                    <h4 class=" text-title-color text-14 font-normal line-clamp-1">确定一款持续绩效考核理念的考核工具</h4>
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
                                        <p class="flex-1 text-text-li text-14">2023/07/01~2023/09/30</p>
                                    </div>
                                    <div class="flex items-center gap-6">
                                        <div class="flex items-center cursor-pointer" @click="handleSchedule">
                                            <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle"
                                                :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                                color="var(--primary-color)" status="success" :percentage="100" />
                                            <p class="text-text-li opacity-50 text-12">100%</p>
                                        </div>
                                        <div class="flex items-center cursor-pointer" @click="handleConfidence">
                                            <i class="taskfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                            <p class="text-text-li opacity-50 text-12">100</p>
                                        </div>
                                        <div class="flex items-center cursor-pointer" @click="handleMark">
                                            <img class="mr-6 -mt-2" src="@/assets/images/icon/fen.svg" />
                                            <p class="text-text-li opacity-50 text-12">8.5分</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="flex flex-col">
                                <div class="flex items-center">
                                    <span class=" text-primary-color text-12 mr-8">KR1</span>
                                    <h4 class=" text-title-color text-14 font-normal line-clamp-1">确定一款持续绩效考核理念的考核工具</h4>
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
                                        <p class="flex-1 text-text-li text-14">2023/07/01~2023/09/30</p>
                                    </div>
                                    <div class="flex items-center gap-6">
                                        <div class="flex items-center cursor-pointer" @click="handleSchedule">
                                            <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle"
                                                :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                                color="var(--primary-color)" status="success" :percentage="20" />
                                            <p class="text-text-li opacity-50 text-12">20%</p>
                                        </div>
                                        <div class="flex items-center cursor-pointer" @click="handleConfidence">
                                            <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                            <p class="text-text-li opacity-50 text-12">信心</p>
                                        </div>
                                        <div class="flex items-center cursor-pointer" @click="handleMark">
                                            <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                            <p class="text-text-li opacity-50 text-12">评分</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
    
                        <div class="border-solid border-0 border-b-[1px] border-[#F2F3F5] my-36"></div>
    
                        <h4 class="text-text-li text-16 font-medium">{{ $t('对齐目标') }} <i
                                class="taskfont text-16 cursor-pointer text-[#A7ABB5]">&#xe779;</i></h4>
    
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
                                    <div
                                        class="absolute border-solid border-0 border-l border-t border-text-tips h-[calc(100%-2px)] w-[calc(100%-10px)] top-[-14px] left-[10px] rounded-tl-lg">
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
    
                    </div>
                    <div class="w-[414px] flex flex-col flex-initial border-solid border-0 border-l-[2px] border-[#F2F3F5] pl-24 ml-24">
                        <div
                            class="flex items-center justify-between border-solid border-0 border-b-[1px] border-[#F2F3F5] pb-[15px] h-[28px]">
                            <ul class="flex items-center gap-8">
                                <li class="li-nav" :class="navActive == 0 ? 'active' : ''" @click="handleNav(0)">{{ $t('评论') }}
                                </li>
                                <li class="li-nav" :class="navActive == 1 ? 'active' : ''" @click="handleNav(1)">{{ $t('动态') }}
                                </li>
                                <li class="li-nav" :class="navActive == 2 ? 'active' : ''" @click="handleNav(2)">{{ $t('复盘') }}
                                </li>
                            </ul>
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

const show = ref(false)
const navActive = ref(0)
const dialogWrappersApp = ref()


const emit = defineEmits(['close','schedule','confidence','mark'])

const handleNav = (index) => {
    navActive.value = index
}

const handleSchedule = () => {
    emit('schedule')
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
    nextTick(()=>{
        if(!window.Vues) return false;
        dialogWrappersApp.value && dialogWrappersApp.value.$destroy();
        dialogWrappersApp.value = new window.Vues.Vue({
            el: document.querySelector('DialogWrappers'), 
            store: window.Vues.store,
            render: (h:any) => {
                return h( window.Vues?.components?.DialogWrapper ,{
                    props:{
                        dialogId: 59
                    }
                },[ h("div",{slot:"head"}) ])
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