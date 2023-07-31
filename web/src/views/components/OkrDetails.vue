<template >
    <n-modal v-model:show="props.show" transform-origin="center" @after-enter="showDrawer" :mask-closable="false"
        @after-leave="closeDrawer" :z-index="13">
        <n-card class="w-[90%] max-w-[1200px] " :bordered="false" size="huge" role="dialog" aria-modal="true">
            <div class="flex max-h-[640px]">
                <div class="flex-1 flex flex-col  relative">
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
                                <span class="cursor-pointer" @click="handleCancel">
                                    {{ detialData.canceled == '0' ? $t('取消目标') : $t('重启目标') }}
                                </span>
                            </n-popover>
                            <div v-else
                                class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid "
                                :class="detialData.completed == '0' ? 'border-[#A8ACB6]' : 'border-primary-color bg-primary-color'">
                                <n-icon v-if="detialData.completed == '1'"
                                    :class="detialData.completed == '0' ? 'text-[#A8ACB6]' : ' text-white'" size="14"
                                    :component="CheckmarkSharp" />
                            </div>

                            <n-tag v-if="detialData.canceled == '1'">{{ $t('已取消') }} </n-tag>
                            <n-tag v-if="detialData.completed == '1'" type="success">{{ $t('已完成') }}</n-tag>
                            <p class="flex items-center text-text-li opacity-50 text-14">
                                {{ detialData.type == "1" ? $t('承诺型') : $t('挑战性') }} /
                                {{ detialData.ascription == "2" ? $t('个人') : $t('部门') }}
                            </p>
                        </div>
                        <div class="flex items-center gap-6">
                            <i class="taskfont icon-title text-[#A7ACB6]" @click="handleEdit">&#xe779;</i>

                            <i class="taskfont text-[#FFD023]" v-if="detialData.is_follow"
                                @click.stop="handleFollowOkr(detialData.id)">&#xe683;</i>
                            <i class="taskfont text-[#A7ACB6]" v-else
                                @click.stop="handleFollowOkr(detialData.id)">&#xe679;</i>
                        </div>
                    </div>
                    <n-scrollbar >
                    <n-spin :show="loadIng">
                        <h3 class=" text-title-color mt-[28px] text-24 font-normal line-clamp-1 ">{{ detialData.title }}</h3>


                        <div class="mt-24 flex flex-col gap-4">
                            <div class="flex items-center">
                                <p class="flex items-center w-[115px]">
                                    <i class="taskfont icon-item text-[#A7ABB5]">&#xe6e4;</i>
                                    <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('负责人') }}</span>
                                </p>
                                <p class="flex-1 text-text-li text-14" v-if="detialData.alias">{{ detialData.alias[0] }}</p>
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
                                            <div v-if="showUserSelect">
                                                        <UserSelects :formkey="index" />
                                                    </div>
                                            <n-avatar v-if="!showUserSelect" round :size="20"
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
                                            @click="handleConfidence(item.id, item.confidence)">
                                            <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ $t('信心') }}</p>
                                        </div>
                                        <div v-else class="flex items-center cursor-pointer"
                                            @click="handleConfidence(item.id, item.confidence)">
                                            <i class="taskfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ item.confidence }}</p>
                                        </div>

                                        <div v-if="item.kr_score == '0'" class="flex items-center cursor-pointer"
                                            @click="handleMark(item.id, item.score)">
                                            <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ $t('评分') }}</p>
                                        </div>
                                        <div v-else class="flex items-center cursor-pointer">
                                            <img class="mr-6 -mt-2" src="@/assets/images/icon/fen.svg" />
                                            <p class="text-text-li opacity-50 text-12">{{ item.kr_score }}{{ $t('分') }}</p>
                                        </div>

                                    </div>
                                </div>
                            </div>

                        </div>

                        <div class="border-solid border-0 border-b-[1px] border-[#F2F3F5] my-36"></div>

                        <h4 class="text-text-li text-16 font-medium mb-12">{{ $t('对齐目标') }} <i
                                class="taskfont text-16 cursor-pointer text-[#A7ABB5]">&#xe779;</i></h4>

                        <AlignTarget class="pb-[28px]" :value="props.show" :id="props.id"></AlignTarget>
                    </n-spin>
                </n-scrollbar>
                </div>

                <div
                    class="w-[414px] relative flex flex-col flex-initial border-solid border-0 border-l-[2px] border-[#F2F3F5] pl-24 ml-24">
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
                                <n-spin size="small" class="absolute top-0 bottom-0 left-0 right-0"> </n-spin>
                                <DialogWrappers />
                            </div>
                            <n-scrollbar v-if="navActive == 1">
                                <div class="flex text-start mb-[24px]" v-for="item in logList">
                                    <n-avatar round :size="28" class="mr-8" :src="item.user_avatar" />
                                    <div class="flex flex-col gap-3">
                                        <p class="text-12 leading-3 text-text-li">{{ item.user_nickname }}<span
                                                class="opacity-60 ml-8">{{ utils.GoDateHMS(item.created_at) }}</span></p>
                                        <h4 class="text-14 leading-[14px] text-title-color font-normal"> <span
                                                class=" font-medium">{{ item.content }}</span></h4>
                                    </div>
                                </div>
                            </n-scrollbar>
                            <n-scrollbar v-if="navActive == 2">
                                <p class="cursor-pointer" @click="handleAddMultiple"> <i
                                        class="taskfont mr-4 text-16 text-text-tips">&#xe6f2;</i><span
                                        class="text-14 text-text-tips">{{ $t('添加复盘') }}</span></p>
                                <div class="flex flex-col gap-3 mt-20" v-if="replayList.length">
                                    <div class="flex items-center justify-between cursor-pointer" v-for="item in replayList" @click="handleView(item.id)">
                                        <h4 class="text-14 text-text-li font-normal">{{ $t('复盘') }} ({{
                                            utils.GoDate(item.created_at) }})</h4>
                                        <p class="text-12 text-text-li opacity-60">{{ utils.GoDateHMS(item.created_at) }}
                                        </p>
                                    </div>
                                </div>
                                <p v-else class="text-12 mt-20 text-text-tips text-center">{{ $t('暂无复盘') }}</p>
                            </n-scrollbar>
                        </div>
                    </div>
                    <div v-if="loadIngR"
                        class="absolute top-[71px] bottom-[24px] left-0 right-0 z-10 bg-[rgba(255,255,255,0.8)] flex-1 flex justify-center items-center">
                        <n-spin size="small" />
                    </div>
                </div>
            </div>


        </n-card>

    </n-modal>
</template>
<script setup lang="ts">
import { CheckmarkSharp } from '@vicons/ionicons5'
import { getOkrDetail, okrFollow, getLogList, getReplayList ,okrCancel} from '@/api/modules/okrList'
import AlignTarget from "@/views/components/AlignTarget.vue";
import { ResultDialog } from "@/api"
import utils from '@/utils/utils';
import { useMessage } from "naive-ui"

const userSelectApps = ref([]);
const navActive = ref(0)
const dialogWrappersApp = ref()
const loadIng = ref(false)
const loadIngR = ref(false)
const detialData = ref<any>({})
const message = useMessage()

const logListPage = ref(1)
const logList = ref<any>([])

const replayListPage = ref(1)
const replayList = ref([])
const showUserSelect = computed(() => window.Vues?.components?.UserSelect ? 1 : 0)

const emit = defineEmits(['close', 'schedule', 'confidence', 'mark', 'edit', 'addMultiple','checkMultiple','upData'])

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

watch(() => detialData.value.dialog_id, (newValue) => {
    if (newValue) {
        loadUserSelects();
        loadDialogWrappers()
    }
})

const getDetail = (type) => {
    const upData = {
        id: props.id,
    }
    if(type=='first') loadIng.value = true
    getOkrDetail(upData).then(({ data }) => {
        detialData.value = data
    })
    .catch(ResultDialog)
    .finally(() => {
        if(type=='first') loadIng.value = false
    })
}

const handleFollowOkr = (id) => {
    const upData = {
        id: id,
    }
    loadIng.value = true
    okrFollow(upData)
        .then(({ msg }) => {
            message.success(msg)
            emit('upData',id)
            getDetail('')
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}


const handleNav = (index) => {
    navActive.value = index
    if (navActive.value == 1) {
        logListPage.value = 1
        logList.value = []
        handleGetLogList()
    }
    if (navActive.value == 2) {
        replayListPage.value = 1
        replayList.value = []
        handleGetReplayList()
    }
}

const handleGetLogList = () => {
    const upData = {
        id: detialData.value.id,
        page: logListPage.value,
        page_size: 10,
    }
    loadIngR.value = true
    getLogList(upData)
        .then(({ data }) => {
            data.data.map(item => {
                logList.value.push(item)
            })
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIngR.value = false
        })
}
const handleGetReplayList = () => {
    const upData = {
        id: detialData.value.id,
        page: replayListPage.value,
        page_size: 10,
    }
    loadIngR.value = true
    getReplayList(upData)
        .then(({ data }) => {
            data.data.map(item => {
                replayList.value.push(item)
            })

        })
        .catch(ResultDialog)
        .finally(() => {
            loadIngR.value = false
        })
}

const handleEdit = () => {
    emit('edit', utils.cloneJSON(detialData.value))
}

//更新进度
const handleSchedule = (id, progress, progress_status) => {
    emit('schedule', id, progress, progress_status)
}

const handleConfidence = (id, confidence) => {
    emit('confidence', id, confidence)
}

const handleMark = (id, score) => {
    emit('mark', id, score)
}

const handleAddMultiple = () => {
    emit('addMultiple', utils.cloneJSON(detialData.value))
}
const handleView = (id) => {
    emit('checkMultiple',id)
}

const closeModal = () => {
    emit('close')
}

const handleCancel = () => {

    const upData = {
        id: detialData.value.id,
    }
    loadIng.value = true
    okrCancel(upData)
        .then(({ msg }) => {
            message.success(msg)
            emit('upData',detialData.value.id)
            getDetail('')
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
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
                        dialogId: detialData.value.dialog_id
                    }
                }, [h("div", { slot: "head" })])
            }
        });
    })
}

// 加载选择用户组件
const loadUserSelects = () => {
    nextTick(() => {
        if (!window.Vues) return false;
        document.querySelectorAll('UserSelects').forEach(e => {
            let item = detialData.value.key_results[e.getAttribute('formkey')];
            let app = new window.Vues.Vue({
                el: e,
                store: window.Vues.store,
                render: (h: any) =>{
                    return h(window.Vues?.components?.UserSelect, {
                        class: "okr-user-selects",
                        props: {
                            value: (item.participant).split(',').map(h=> Number(h) )  ,
                            title: $t('选择参与人'),
                            border: false,
                            avatarSize: 20,
                            addIcon: false,
                            disable:true
                        },
                        on: {
                            "on-show-change": (show: any, values: any) => {
                                if (!show) {
                                    item.participant = values.join(',');
                                }
                            }
                        }
                    })
                },
            });
            userSelectApps.value.push(app)
        })
    })
}

// 卸载
window.addEventListener('apps-unmount', function () {
    dialogWrappersApp.value && dialogWrappersApp.value.$destroy()
    userSelectApps.value.forEach(app => app.$destroy())
})

// 显示
const showDrawer = () => {
    getDetail('first')
    if(detialData.value.dialog_id>0){
        loadUserSelects();
        loadDialogWrappers()
    }
}

// 关闭
const closeDrawer = () => {
    dialogWrappersApp.value && dialogWrappersApp.value.$destroy();
    navActive.value = 0
    detialData.value = {};
    userSelectApps.value.forEach(app => app.$destroy())
}

const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

defineExpose({
    getDetail
})
</script>
<style lang="less" scoped>
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
