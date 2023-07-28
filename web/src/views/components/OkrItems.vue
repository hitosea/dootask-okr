<template >
    <div class="okr-item-main">
        <div class="okr-item-box" @click="handleOpenDetail(item.id)" v-for="(item) in props.list">
            <n-progress style="width: 52px;" color="var(--primary-color)" indicator-text-color="var(--primary-color)"
                type="circle" :percentage="item.progress" :offset-degree="180" :stroke-width="8">
                <p class="text-primary-color text-14">{{ item.progress }}<span class="text-12">%</span></p>
            </n-progress>
            <div class="okr-list">
                <div class="okr-title">
                    <div class="okr-title-l">
                        <span class="scale-[0.8333]" :class="pStatus(item.priority)">{{ item.priority }}</span>
                        <h3>{{ item.title }}</h3>
                    </div>
                    <div class="okr-title-r">
                        <i class="taskfont okr-title-star" v-if="item.is_follow"
                            @click.stop="handleFollowOkr(item.id)">&#xe683;</i>
                        <i class="taskfont pr-16" v-else @click.stop="handleFollowOkr(item.id)">&#xe679;</i>
                        <i class="taskfont okr-title-icon">&#xe671;</i>
                        <p>{{ item.kr_finish_count }}/{{ item.kr_count }}</p>
                    </div>
                </div>
                <div class="okr-time">
                    <i class="taskfont"> &#xe61a;</i>
                    <p>{{ item.alias.join(',') }}</p>
                    <div class="w-1 bg-[#F2F3F5] mx-12 h-full"></div>
                    <i class="taskfont"> &#xe71d;</i>
                    <p>{{ expiresFormat(item.end_at) }}</p>
                </div>
                <div class="kr-list">
                    <div class="kr-list-item" v-for="(childItem, index) in item.key_results">
                        <span class="bg-[rgba(135,208,104,0.2);] scale-[0.8333]">KR{{ index + 1 }}</span>
                        <p>{{ childItem.title }}</p>
                        <div class="kr-list-schedule">
                            <n-progress class="-mt-6 mr-[6px]" style="width: 15px; " type="circle" :show-indicator="false"
                                :offset-degree="180" :stroke-width="15" color="var(--primary-color)" status="success"
                                :percentage="childItem.progress" />
                            {{ childItem.progress }}%
                        </div>
                    </div>

                </div>
                <div class="align-target" v-if="item.align_count > 0">
                    <div class=" cursor-pointer" @click.stop="handleTarget(1, item.id)">{{ $t('对齐目标') }}({{ item.align_count
                    }}）</div>
                </div>
                <div class="align-target" v-else>
                    <div class=" cursor-pointer" @click.stop="handleTarget(2, item.id)">
                        {{ $t('向上对齐') }}
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!-- 查看对齐OKR -->
    <AlignTargetModal :value="alignTargetShow" :id="eidtId" @close="() => { alignTargetShow = false }"></AlignTargetModal>

    <!-- 选择对齐OKR -->
    <SelectAlignment :value="selectAlignmentShow" @close="() => { selectAlignmentShow = false }"
        @submit="submitSelectAlignment"></SelectAlignment>

    <!-- OKR详情 -->
    <OkrDetails ref="RefOkrDetails" :id="eidtId" :show="okrDetailsShow" @close="() => { okrDetailsShow = false }"
        @schedule="handleOpenSchedule" @confidence="handleConfidence" @mark="handleMark" @edit="handleEdit" @addMultiple="handleAddMultiple" @checkMultiple="handleCheckMultiple"></OkrDetails>

    <!-- 更新进度   -->
    <DegreeOfCompletion v-model:show="degreeOfCompletionShow" :id="degreeOfCompletionId"
        :progress="degreeOfCompletionProgress" :progress_status="degreeOfCompletionProgressStatus"
        @close="handleCloseDedree">
    </DegreeOfCompletion>

    <!-- 更新信心 -->
    <Confidences :id="confidencesId" :confidence="confidence" v-model:show="confidenceShow" @close="handleCloseConfidenes">
    </Confidences>

    <!-- 更新评分 -->
    <MarkVue v-model:show="markShow" :id="markId" :score="score" @close="handleCloseMarks"></MarkVue>

    <!-- 新增复盘 -->
    <AddMultiple v-model:show="addMultipleShow" :data="addMultipleData" :multipleId="multipleId" @close="handleCloseMultiple"></AddMultiple>
</template>
<script setup lang="ts">
import AlignTargetModal from '@/views/components/AlignTargetModal.vue';
import SelectAlignment from '@/views/components/SelectAlignment.vue'
import OkrDetails from './OkrDetails.vue';
import DegreeOfCompletion from '@/views/components/DegreeOfCompletion.vue'
import Confidences from '@/views/components/Confidences.vue';
import MarkVue from '@/views/components/Marks.vue';
import AddMultiple from '@/views/components/AddMultiple.vue';
import utils from '@/utils/utils';
import webTs from '@/utils/web'
import { alignUpdate, okrFollow } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"


const alignTargetShow = ref(false)
const selectAlignmentShow = ref(false)
const okrDetailsShow = ref(false)


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

const addMultipleShow = ref(false)
const multipleId = ref(0)
const addMultipleData = ref(null)


const nowInterval = ref<any>(null)
const nowTime = ref(0)
const loadIng = ref(false)
const message = useMessage()
const eidtId = ref(0)

const RefOkrDetails = ref(null)

const props = defineProps({
    list: {
        type: undefined,
        default:() => []
    }
});

const emit = defineEmits(['upData','edit'])

const handleTarget = (e, id) => {
    eidtId.value = id
    if (e == 1) {
        alignTargetShow.value = true
    } else {
        selectAlignmentShow.value = true
    }

}

const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

//打开详情
const handleOpenDetail = (id) => {
    eidtId.value = id
    okrDetailsShow.value = true
}

//打开进度
const handleOpenSchedule = (id, progress, progress_status) => {
    degreeOfCompletionId.value = id
    degreeOfCompletionProgress.value = progress
    degreeOfCompletionProgressStatus.value = progress_status
    degreeOfCompletionShow.value = true
}

//关闭进度
const handleCloseDedree = (type) => {
    if (type == 1) {
        RefOkrDetails.value.getDetail()
        props.list.map((item,index)=>{
            item.key_results.map(CItem=>{
                if (CItem.id == degreeOfCompletionId.value) { 
                    emit('upData',props.list[index].id)
                }
            })
        })
    }
    degreeOfCompletionId.value = 0
    degreeOfCompletionProgress.value = 0
    degreeOfCompletionProgressStatus.value = 0
    degreeOfCompletionShow.value = false
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
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

const submitSelectAlignment = (e) => {
    const upData = {
        align_objective: e.join(','),
        id: eidtId.value,
    }
    loadIng.value = true
    alignUpdate(upData)
        .then(({ msg }) => {
            message.success(msg)
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

//打开信心
const handleConfidence = (id, confidences) => {
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
        RefOkrDetails.value.getDetail()
        emit('upData')
    }
}

//打开评分
const handleMark = (id, scores) => {
    markId.value = id
    score.value = scores
    markShow.value = true
}
//关闭评分
const handleCloseMarks = (type) => {
    markId.value = 0
    score.value = 0
    markShow.value = false
    if (type == 1) {
        RefOkrDetails.value.getDetail()
        emit('upData')
    }
}

//编辑
const handleEdit = (data) => {
    emit('edit',data)
    okrDetailsShow.value = false
}

//添加复盘
const handleAddMultiple = (data) => {
    console.log(data);
    
    okrDetailsShow.value = false
    addMultipleData.value = data
    addMultipleShow.value = true
}

//查看复盘
const handleCheckMultiple = (id) => {
    console.log(id);
    okrDetailsShow.value = false
    multipleId.value = id
    addMultipleShow.value = true
}

//关闭复盘
const handleCloseMultiple = () => {
    addMultipleData.value = null
    multipleId.value = 0
    addMultipleShow.value = false
}



const expiresFormat = (date) => {
    const Dates = new Date(date);
    const timestamp = Dates.getTime();
    return webTs.countDownFormat(timestamp, nowTime.value)
}

onMounted(() => {
    nowInterval.value = setInterval(() => {
        nowTime.value = utils.Time();
    }, 1000);
})

</script>
<style lang="less">
.okr-item-main {
    @apply flex flex-col gap-6;

    .okr-item-box {
        @apply px-24 py-32 bg-white rounded-lg flex gap-4 cursor-pointer;

        .okr-list {
            @apply flex flex-col flex-1;

            .okr-title {
                @apply flex justify-between items-center;

                .okr-title-l {
                    @apply flex items-center gap-2;

                    span {
                        @apply text-10 text-white px-6 py-2 rounded-full origin-center flex items-center leading-3;
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

                    h3 {
                        @apply text-title-color text-14 font-normal;
                    }
                }

                .okr-title-r {
                    @apply flex items-center;

                    .okr-title-star {
                        @apply text-[#FFD023] pr-16;
                    }

                    .okr-title-icon {
                        @apply text-text-tips mr-4;
                    }

                    p {
                        @apply text-text-tips;
                    }
                }
            }

            .okr-time {
                @apply flex items-center mt-12 text-text-tips;

                i {
                    @apply mr-4;
                }
            }

            .kr-list {
                @apply mt-16 flex flex-col gap-5;

                .kr-list-item {
                    @apply flex items-center;

                    span {
                        @apply text-primary-color text-10 origin-center py-2 px-6 rounded-full flex items-center leading-3;
                    }

                    p {
                        @apply text-title-color text-14 ml-16 truncate;
                    }

                    .kr-list-schedule {
                        @apply ml-auto flex items-center;
                    }
                }
            }

            .align-target {
                @apply flex items-start mt-20 text-text-tips text-12;
            }
        }
    }
}
</style>
