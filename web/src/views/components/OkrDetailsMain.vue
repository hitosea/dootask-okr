<template >
    <div class="flex flex-col h-full md:h-auto md:flex-row md:max-h-[640px] md:min-h-[640px]">
        <div class="md:flex-1 flex flex-col relative md:overflow-hidden">
            <div
                class="hidden md:flex min-h-[36px] items-center justify-between pb-[15px] border-solid border-0 border-b-[1px] border-[#F2F3F5] relative md:mr-24">
                <div class="flex items-center gap-4">
                    <n-popover v-if="detialData.completed == '0'" placement="bottom" :show="showPopover" trigger="manual">
                        <template #trigger>
                            <div @click="showPopover = !showPopover" v-if="detialData.completed == '0'"
                                class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid cursor-pointer"
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
                    <i v-if="detialData.canceled == '0' && detialData.completed == '0'"
                        class="taskfont icon-title text-[#A7ACB6]" @click="handleEdit">&#xe779;</i>

                    <i class="taskfont text-[#FFD023] cursor-pointer" v-if="detialData.is_follow"
                        @click.stop="handleFollowOkr">&#xe683;</i>
                    <i class="taskfont text-[#A7ACB6] cursor-pointer" v-else @click.stop="handleFollowOkr">&#xe679;</i>

                    <div v-if="detialData.score  > -1" class="flex items-center cursor-pointer">
                        <img class="mr-8" src="@/assets/images/icon/fen.svg" />
                        <p class=" text-title-color text-12">{{ detialData.score }}{{ $t('分') }}
                        </p>
                    </div>
                </div>
                <img v-if="detialData.completed == '1'" class="absolute right-24 -bottom-[50px]" src="@/assets/images/icon/complete.png" />
            </div>

            <n-scrollbar class="pr-[5px] ">
                <n-spin class="md:mr-24" :show="false">
                    <h3 class=" text-title-color md:mt-[28px]  text-18 md:text-24 leading-6 font-normal md:min-h-[40px]">
                        {{ detialData.title }}
                    </h3>
                    <div class="mt-16 md:mt-24 flex flex-col gap-4">
                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i v-if="detialData.ascription == '2'"
                                    class="taskfont icon-item text-[#A7ABB5]">&#xe6e4;</i>
                                <i v-else class="taskfont icon-item text-[#A7ABB5]">&#xe682;</i>

                                <span class="text-[#515A6E] text-[14px] opacity-50"> {{ detialData.ascription == "2" ?
                                    $t('负责人') : $t('部门') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14">
                                {{ detialData.alias && detialData.alias[0] }}
                            </p>
                        </div>


                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i class="taskfont icon-item text-[#A7ABB5]">&#xe6e8;</i>
                                <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('起止时间') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14">
                                <span v-if="detialData.start_at">{{ utils.GoDate(detialData.start_at || 0) }} ~ {{
                                    utils.GoDate(detialData.end_at || 0) }}</span>
                            </p>
                        </div>

                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i class="taskfont icon-item text-[#A7ABB5]">&#xe6ec;</i>
                                <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('优先级') }}</span>
                            </p>
                            <span class="span" :class="pStatus(detialData.priority)">{{ detialData.priority
                            }}</span>
                        </div>


                        <div class="flex md:hidden items-center" v-if="detialData.score > -1">
                            <p class="flex items-center w-[115px]">
                                <img class="mr-6 " src="@/assets/images/icon/fen.svg" />
                                <span class="text-[#515A6E] text-[14px] opacity-50">{{ $t('评分') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14">
                                {{ detialData.score }}{{ $t('分') }}
                            </p>
                        </div>
                    </div>

                    <h4 class="hidden md:block text-text-li text-16 font-medium mt-36">KR</h4>

                    <div class="hidden md:flex flex-col mt-20 gap-6 min-h-[60px]">
                        <div class="flex flex-col" v-for="(item, index) in detialData.key_results">
                            <div class="flex items-start">
                                <span class=" text-primary-color text-12 leading-5 mr-8 mt-2 shrink-0">KR{{ index + 1 }}</span>
                                <h4 class=" text-title-color text-14 font-normal">{{ item.title }}</h4>
                            </div>
                            <div class="flex items-center justify-between mt-8">
                                <div class="flex items-center mr-24">
                                    <div class="flex items-center gap-2">
                                        <div v-if="showUserSelect">
                                            <UserSelects :formkey="index" />
                                        </div>
                                        <n-avatar v-if="!showUserSelect" round :size="20"
                                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                    </div>
                                    <i class="taskfont icon-item ml-24 text-[#A7ABB5]">&#xe6e8;</i>
                                    <p class="flex-1 text-text-li text-14 min-w-[140px] shrink-0">{{
                                        utils.GoDate(item.start_at || 0) }} ~{{ utils.GoDate(item.end_at || 0) }}
                                    </p>
                                </div>
                                <div class="flex items-center gap-6 shrink-0">
                                    <div class="flex items-center cursor-pointer"
                                        @click="handleSchedule(item.id, item.progress, item.progress_status, item.score)">
                                        <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle"
                                            :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                            color="var(--primary-color)" status="success" :percentage="item.progress" />
                                        <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
                                    </div>
                                    <div v-if="item.confidence == '0'" class="flex items-center cursor-pointer"
                                        @click="handleConfidence(item.id, item.confidence, item.score)">
                                        <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ $t('信心') }}</p>
                                    </div>
                                    <div v-else class="flex items-center cursor-pointer"
                                        @click="handleConfidence(item.id, item.confidence, item.score)">
                                        <i class="taskfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ item.confidence }}</p>
                                    </div>
                                    <div v-if="item.kr_score == '0'" class="flex items-center cursor-pointer"
                                        @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                        <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ $t('评分') }}</p>
                                    </div>
                                    <div v-else class="flex items-center cursor-pointer"
                                        @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                        <img class="mr-6 -mt-2" src="@/assets/images/icon/fen.svg" />
                                        <p class="text-text-li opacity-50 text-12">{{ item.kr_score }}{{ $t('分') }}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="hidden md:block border-solid border-0 border-b-[1px] border-[#F2F3F5] my-36"></div>

                    <h4 class="hidden md:block text-text-li text-16 font-medium mb-12">
                        {{ $t('对齐目标') }}
                        <i v-if="detialData.completed == '0'" class="taskfont text-16 cursor-pointer text-[#A7ABB5]"
                            @click="() => { selectAlignmentShow = true }">&#xe779;</i>
                    </h4>

                    <div class="pb-[28px] w-full overflow-hidden hidden md:block">
                        <AlignTarget ref="AlignTargetRef" :value="props.show" :id="props.id" @unalign="handleUnalign">
                        </AlignTarget>
                    </div>

                </n-spin>
            </n-scrollbar>
        </div>

        <div
            class="md:min-w-[35.8%] relative flex flex-col flex-1 md:flex-initial border-solid border-0 md:border-l-[2px] border-[#F2F3F5] mt-[32px] md:mt-0">
            <div
                class="flex items-center justify-between border-solid border-0 border-b-[1px] border-[#F2F3F5] pb-[5px] md:pb-[15px] md:ml-24 min-h-[36px]">
                <ul class="flex w-full items-center gap-8 justify-between md:justify-start">
                    <li class="block md:hidden li-nav" :class="navActive == 3 ? 'active' : ''" @click="handleNav(3)">KR</li>
                    <li class="block md:hidden li-nav" :class="navActive == 4 ? 'active' : ''" @click="handleNav(4)">{{
                        $t('对齐') }}
                    </li>
                    <li class="li-nav" :class="navActive == 0 ? 'active' : ''" @click="handleNav(0)">{{ $t('评论') }}
                    </li>
                    <li class="li-nav" :class="navActive == 1 ? 'active' : ''" @click="handleNav(1)">{{ $t('动态') }}
                    </li>
                    <li class="li-nav" :class="navActive == 2 ? 'active' : ''" @click="handleNav(2)">{{ $t('复盘') }}
                    </li>
                </ul>
                <i class="taskfont text-16 cursor-pointer text-[#A7ABB5] hidden md:block" @click="closeModal">&#xe6e5;</i>
            </div>
            <div class="flex-auto relative mt-16 md:mt-0">
                <div class="md:absolute md:top-[24px] md:bottom-0 md:left-0 md:right-0">
                    <div class="text-center" v-if="navActive == 3">
                        <div class="flex md:hidden flex-col gap-3 min-h-[60px]">
                            <div class="flex flex-col bg-[#fff] px-16 pt-24 rounded-lg"
                                v-for="(item, index) in detialData.key_results">
                                <div class="flex items-center">
                                    <h4 class=" text-title-color text-15 md:text-14 font-normal line-clamp-1">{{ item.title
                                    }}</h4>
                                </div>
                                <div class="flex flex-col justify-between mt-12">
                                    <div class="flex items-center justify-between">
                                        <div class="flex items-center gap-2">
                                            <div v-if="showUserSelect">
                                                <UserSelects :formkey="index" />
                                            </div>
                                            <n-avatar v-if="!showUserSelect" round :size="20"
                                                src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                        </div>
                                        <div class="flex items-center">
                                            <i class="taskfont text-14 mr-6 text-[#A7ABB5]">&#xe6e8;</i>
                                            <p class="flex-1 text-text-tips text-14  shrink-0">{{ utils.GoDate(item.end_at
                                                || 0) }}
                                            </p>
                                        </div>
                                    </div>
                                    <div
                                        class="flex mt-16 py-8 shrink-0 border-solid border-0 border-t-[1px] border-[#F2F3F5]">
                                        <div class="flex items-center justify-center cursor-pointer flex-1 border-solid border-0 border-r-[1px] border-[#F2F3F5]"
                                            @click="handleSchedule(item.id, item.progress, item.progress_status, item.score)">
                                            <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle"
                                                :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                                color="var(--primary-color)" status="success" :percentage="item.progress" />
                                            <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
                                        </div>
                                        <div v-if="item.confidence == '0'"
                                            class="flex flex-1 items-center justify-center cursor-pointer border-solid border-0 border-r-[1px] border-[#F2F3F5]"
                                            @click="handleConfidence(item.id, item.confidence, item.score)">
                                            <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ $t('信心') }}</p>
                                        </div>
                                        <div v-else
                                            class="flex flex-1 items-center justify-center cursor-pointer border-solid border-0 border-r-[1px] border-[#F2F3F5]"
                                            @click="handleConfidence(item.id, item.confidence, item.score)">
                                            <i class="taskfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ item.confidence }}</p>
                                        </div>
                                        <div v-if="item.kr_score == '0'"
                                            class="flex flex-1 items-center justify-center cursor-pointer"
                                            @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                            <i class="taskfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ $t('评分') }}</p>
                                        </div>
                                        <div v-else class="flex flex-1 items-center justify-center cursor-pointer"
                                            @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                            <img class="mr-6 -mt-2" src="@/assets/images/icon/fen.svg" />
                                            <p class="text-text-li opacity-50 text-12">{{ item.kr_score }}{{ $t('分') }}
                                            </p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="text-center" v-if="navActive == 4">
                        <AlignTarget ref="AlignTargetRef" :value="props.show" :id="props.id" @unalign="handleUnalign">
                        </AlignTarget>
                    </div>
                    <div class="text-center flex-1" v-show="navActive == 0">
                        <!-- <n-spin size="small" class="absolute top-0 bottom-0 left-0 right-0"> </n-spin> -->
                        <span v-if="!showDialogWrapper">{{ $t('子应用无法加载') }}</span>
                        <DialogWrappers v-else />
                    </div>
                    <n-scrollbar v-if="navActive == 1" :on-scroll="onScrollLogList">
                        <div class="flex text-start mb-[24px] md:pl-24 pr-[10px] " v-for="item in logList">
                            <n-avatar round :size="28" class="mr-8 shrink-0" :src="item.user_avatar" />
                            <div class="flex flex-col gap-3">
                                <p class="text-12 leading-3 text-text-li">{{ item.user_nickname }}<span
                                        class="opacity-60 ml-8">{{ utils.GoDateHMS(item.created_at) }}</span></p>
                                <h4 class="text-14 leading-[14px] text-title-color font-normal"> <span
                                        class=" font-normal">{{ item.content }}</span></h4>
                            </div>
                        </div>
                    </n-scrollbar>
                    <n-scrollbar v-if="navActive == 2" :on-scroll="onScrollReplayList">
                        <div class="pl-24 pr-[10px]">
                            <p class="cursor-pointer" :class="detialData.score < 0 ? 'text-text-tips' : 'text-primary-color' " @click="handleAddMultiple"> <i
                                    class="taskfont mr-4 text-16 ">&#xe6f2;</i><span
                                    class="text-14" >{{ $t('添加复盘') }}</span></p>
                            <div class="flex flex-col gap-3 mt-20" v-if="replayList.length">
                                <div class="flex items-center justify-between cursor-pointer" v-for="item in replayList"
                                    @click="handleCheckMultiple(item.id)">
                                    <h4 class="text-14 text-text-li font-normal">{{ $t('复盘') }} ({{
                                        utils.GoDate(item.created_at) }})</h4>
                                    <p class="text-12 text-text-li opacity-60">{{ utils.GoDateHMS(item.created_at)
                                    }}
                                    </p>
                                </div>
                            </div>
                            <p v-else class="text-12 mt-20 text-text-tips text-center">{{ $t('暂无复盘') }}</p>
                        </div>
                    </n-scrollbar>
                </div>
            </div>
            <div v-if="loadIngR"
                class="absolute top-[71px] bottom-[24px] left-0 right-0 z-10 bg-[rgba(255,255,255,0.8)] flex-1 flex justify-center items-center">
                <n-spin size="small" />
            </div>
        </div>
    </div>
    <!-- 对齐目标 -->
    <SelectAlignment :value="selectAlignmentShow" :editData="detialData.align_objective"
        @close="() => { selectAlignmentShow = false }" @submit="submitSelectAlignment"></SelectAlignment>

    <!-- 更新进度   -->
    <DegreeOfCompletion v-model:show="degreeOfCompletionShow" :id="degreeOfCompletionId"
        :progress="degreeOfCompletionProgress" :progress_status="degreeOfCompletionProgressStatus"
        @close="handleCloseDedree">
    </DegreeOfCompletion>

    <!-- 更新信心 -->
    <Confidences :id="confidencesId" :confidence="confidence" v-model:show="confidenceShow" @close="handleCloseConfidenes">
    </Confidences>

    <!-- 更新评分 -->
    <MarkVue v-model:show="markShow" :id="markId" :score="score" :superior_score="superiorScore" @close="handleCloseMarks">
    </MarkVue>
</template>
<script setup lang="ts">
import { CheckmarkSharp } from '@vicons/ionicons5'
import { getOkrDetail, okrFollow, getLogList, getReplayList, okrCancel, alignUpdate, participantUpdate } from '@/api/modules/okrList'
import AlignTarget from "@/views/components/AlignTarget.vue";
import { ResultDialog } from "@/api"
import utils from '@/utils/utils';
import { useMessage } from "naive-ui"
import SelectAlignment from '@/views/components/SelectAlignment.vue'
import DegreeOfCompletion from '@/views/components/DegreeOfCompletion.vue'
import Confidences from '@/views/components/Confidences.vue';
import MarkVue from '@/views/components/Marks.vue';
import { GlobalStore } from '@/store';
import { useRouter } from 'vue-router'

const router = useRouter()
const globalStore = GlobalStore()
const userSelectApps = ref([]);
const navActive = ref(0)
const dialogWrappersApp = ref()
const loadIng = ref(false)
const loadIngR = ref(false)
const showPopover = ref(false)
const detialData = ref<any>({})
const message = useMessage()

const logListPage = ref(1)
const logListLastPage = ref(99999)
const logList = ref<any>([])

const replayListPage = ref(1)
const replayListLastPage = ref(99999)
const replayList = ref([])
const showDialogWrapper = computed(() => window.Vues?.components?.DialogWrapper ? 1 : 0)
const showUserSelect = computed(() => window.Vues?.components?.UserSelect ? 1 : 0)

const selectAlignmentShow = ref(false)
const AlignTargetRef = ref(null)

const degreeOfCompletionShow = ref(false)
const degreeOfCompletionId = ref(0)
const degreeOfCompletionProgress = ref(0)
const degreeOfCompletionProgressStatus = ref(0)

const confidenceShow = ref(false)
const confidencesId = ref(0)
const confidence = ref(0)

const markShow = ref(false)
const markId = ref(0)
const score = ref(0)
const superiorScore = ref(0)


const emit = defineEmits(['close', 'edit', 'upData', 'isFollow', 'canceled'])

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


// 明细
const getDetail = (type) => {
    if (type == 'first') loadIng.value = true
    getOkrDetail({
        id: props.id,
    }).then(({ data }) => {
        detialData.value = data
        emit('isFollow', detialData.value.is_follow)
        emit('canceled', detialData.value.canceled)
    })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            if (type == 'first') loadIng.value = false
        })
}

// 关注
const handleFollowOkr = () => {
    loadIng.value = true
    okrFollow({
        id: detialData.value.id,
    }).then(({ msg }) => {
        message.success($t('操作成功'))
        emit('upData', detialData.value.id)
        getDetail('')
    })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

//
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

// 日志下一页
const onScrollLogList = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!loadIng.value) {
            logListPage.value++
            handleGetLogList()
        }
    }
}

// 日志列表
const handleGetLogList = () => {
    if (logListLastPage.value >= logListPage.value) {
        loadIngR.value = true
        getLogList({
            id: detialData.value.id,
            page: logListPage.value,
            page_size: 10,
        }).then(({ data }) => {
            data.data.map(item => {
                if (item.content.includes('创建OKR')) {
                    item.content = $t('创建OKR') + item.content.replace('创建OKR', '')
                }
                if (item.content.includes('修改O目标标题')) {
                    item.content = $t('修改O目标标题') + item.content.replace('修改O目标标题', '')
                }
                if (item.content.includes('修改O目标周期')) {
                    item.content = $t('修改O目标周期') + item.content.replace('修改O目标周期', '')
                }
                if (item.content.includes('修改O目标状态')) {
                    const regex = /\[([^[\]]+)]/;
                    const match = item.content.match(regex);
                    let result = []
                    if (match) {
                        result = match[1].split('=>');
                    }
                    item.content = $t('修改O目标状态') + `[${$t(result[0])} => ${$t(result[1])}]`
                }
                if (item.content.includes('修改对齐目标')) {
                    item.content = $t('修改对齐目标') + item.content.replace('修改对齐目标', '')
                }
                if (item.content.includes('修改KR标题')) {
                    item.content = $t('修改KR标题') + item.content.replace('修改KR标题', '')
                }
                if (item.content.includes('修改KR周期')) {
                    item.content = $t('修改KR周期') + item.content.replace('修改KR周期', '')
                }
                if (item.content.includes('修改KR参与人')) {
                    item.content = $t('修改KR参与人') + item.content.replace('修改KR参与人', '')
                }
                if (item.content.includes('修改KR进度')) {
                    item.content = $t('修改KR进度') + item.content.replace('修改KR进度', '')
                }
                if (item.content.includes('修改KR状态')) {
                    const regex = /\[([^[\]]+)]/;
                    const match = item.content.match(regex);
                    let result = []
                    if (match) {
                        result = match[1].split('=>');
                    }
                    item.content = $t('修改KR状态') + item.content.replace('修改KR状态', '').replace(match[0], '') + `[${$t(result[0])} => ${$t(result[1])}]`
                }
                if (item.content.includes('修改KR信心指数')) {
                    item.content = $t('修改KR信心指数') + item.content.replace('修改KR信心指数', '')
                }
                if (item.content.includes('责任人打分')) {
                    item.content = $t('责任人打分') + item.content.replace('责任人打分', '')
                }
                if (item.content.includes('上级打分')) {
                    item.content = $t('上级打分') + item.content.replace('上级打分', '')
                }
                logList.value.push(item)
            })


            logListLastPage.value = data.last_page
        })
            .catch(({ msg }) => {
                message.error(msg)
            })
            .finally(() => {
                loadIngR.value = false
            })
    }
}

//复盘下一页
const onScrollReplayList = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!loadIng.value) {
            replayListPage.value++
            handleGetLogList()
        }
    }
}

// 复盘列表
const handleGetReplayList = () => {
    if (replayListLastPage.value >= replayListPage.value) {
        loadIngR.value = true
        getReplayList({
            id: detialData.value.id,
            page: replayListPage.value,
            page_size: 10,
        }).then(({ data }) => {
            data.data.map(item => {
                replayList.value.push(item)
            })
            replayListLastPage.value = data.last_page
        })
            .catch(({ msg }) => {
                message.error(msg)
            })
            .finally(() => {
                loadIngR.value = false
            })
    }
}

//编辑
const handleEdit = () => {
    if (window.innerWidth < 768) {
        router.push({
            path: '/addOkr',
            query: { data: JSON.stringify(detialData.value) },
        })
    }
    else {
        emit('edit', utils.cloneJSON(detialData.value))
    }
}

const closeModal = () => {
    emit('close')
}



//打开进度
const handleSchedule = (id, progress, progress_status, score) => {
    if (detialData.value.canceled == '1') return message.error($t('取消目标后不允许更改'))
    if (score > -1) return message.error($t('KR已评分无法操作'))
    degreeOfCompletionId.value = id
    degreeOfCompletionProgress.value = progress
    degreeOfCompletionProgressStatus.value = progress_status
    degreeOfCompletionShow.value = true
}

//关闭进度
const handleCloseDedree = (type) => {
    if (type == 1) {
        getDetail('')
        emit('upData', detialData.value.id)
    }
    degreeOfCompletionId.value = 0
    degreeOfCompletionProgress.value = 0
    degreeOfCompletionProgressStatus.value = 0
    degreeOfCompletionShow.value = false
}

//打开信心
const handleConfidence = (id, confidences, score) => {
    if (detialData.value.canceled == '1') return message.error($t('取消目标后不允许更改'))
    if (score > -1) return message.error($t('KR已评分无法操作'))
    confidencesId.value = id
    confidence.value = confidences
    confidenceShow.value = true
}
//关闭信心
const handleCloseConfidenes = (type) => {
    confidencesId.value = 0
    confidence.value = 0
    confidenceShow.value = false
    if (type == 1) {
        getDetail('')
    }
}

//打开评分
const handleMark = (id, scores, superior_score, progress) => {
    if (detialData.value.canceled == '1') return message.error($t('取消目标后不允许更改'))
    if (progress < 100) return message.error($t('KR进度尚未达到100%'))
    markId.value = id
    score.value = scores
    superiorScore.value = superior_score
    markShow.value = true
}
//关闭评分
const handleCloseMarks = (type) => {
    markId.value = 0
    score.value = 0
    superiorScore.value = 0
    markShow.value = false
    if (type == 1) {
        getDetail('')
    }
}

// 取消O
const handleCancel = () => {
    loadIng.value = true
    okrCancel({
        id: detialData.value.id,
    }).then(({ msg }) => {
        message.success($t('修改成功'))
        showPopover.value = false
        emit('upData', detialData.value.id)
        getDetail('')
    })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

// 取消
const handleUnalign = () => {
    emit('upData', detialData.value.id)
    getDetail('')
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

//添加复盘
const handleAddMultiple = () => {
    if (detialData.value.score < 0) return  message.error($t('KR评分未完成'))
    if (window.innerWidth < 768) {
        router.push({
            path: '/addMultiple',
            query: { data: JSON.stringify(detialData.value) },
        })
    }
    else {
        closeModal()
        globalStore.$patch((state) => {
            state.addMultipleData = detialData.value
            state.addMultipleShow = true
        })
    }
}


//查看复盘
const handleCheckMultiple = (id) => {
    if (window.innerWidth < 768) {
        router.push({
            path: '/addMultiple',
            query: {
                data: JSON.stringify(detialData.value),
                id: id || 0,
            },
        })
    }
    else {
        closeModal()
        globalStore.$patch((state) => {
            state.multipleId = id
            state.addMultipleShow = true
        })
    }
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
                render: (h: any) => {
                    return h(window.Vues?.components?.UserSelect, {
                        class: "okr-user-selects",
                        props: {
                            value: (item.participant).split(',').map(h => Number(h)),
                            title: $t('选择参与人'),
                            border: false,
                            avatarSize: 20,
                            addIcon: false,
                            disable: false
                        },
                        on: {
                            "on-show-change": (show: any, values: any) => {
                                if (!show) {
                                    if (item.participant != values.join(',')) {
                                        item.participant = values.join(',');
                                        participantChange(item, e.getAttribute('formkey'))
                                    }
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



//更新参与人
const participantChange = (item, index) => {
    if (detialData.value.key_results[index].score > -1 || detialData.value.key_results[index].superior_score > -1) {
        return
    }
    const upData = {
        participant: item.participant,
        id: item.id,
    }
    loadIng.value = true
    participantUpdate(upData)
        .then(({ msg }) => {
            message.success(msg)
            getDetail('')
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

//添加对齐目标
const submitSelectAlignment = (e) => {
    const upData = {
        align_objective: e.join(','),
        id: detialData.value.id,
    }
    loadIng.value = true
    alignUpdate(upData)
        .then(({ msg }) => {
            message.success($t('修改成功'))
            emit('upData', detialData.value.id)
            getDetail('')
            AlignTargetRef.value.getList()
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

// 卸载
window.addEventListener('apps-unmount', function () {
    dialogWrappersApp.value && dialogWrappersApp.value.$destroy()
    userSelectApps.value.forEach(app => app.$destroy())
})


// 关闭
const closeDrawer = () => {
    dialogWrappersApp.value && dialogWrappersApp.value.$destroy();
    navActive.value = 0
    detialData.value = {};
    loadIng.value = true;
    userSelectApps.value.forEach(app => app.$destroy())
}

const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

watch(() => props.show, (newValue) => {
    if (newValue) {
        getDetail('first')
    }
}, { immediate: true })

watch(() => detialData.value.dialog_id, (newValue) => {
    if (newValue) {
        loadUserSelects();
        loadDialogWrappers()
    }
}, { immediate: true })

nextTick(() => {
    if (window.innerWidth < 768) {
        navActive.value = 3
    }
})

defineExpose({
    getDetail,
    closeDrawer,
    handleEdit,
    handleCancel,
    handleFollowOkr
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


.li-nav {
    @apply list-none cursor-pointer text-text-li opacity-50 text-14 px-4 relative;
}

.li-nav.active {
    @apply text-primary-color md:text-title-color md:text-16 opacity-100 relative;
}

.li-nav.active::before {
    @apply w-full bg-primary-color absolute left-0 right-0 -bottom-14;
    height: 2px;
    content: ' ';
}</style>
