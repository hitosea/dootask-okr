<template >
    <div class="okr-item-main">
        <div class="okr-item-box" @click="handleOpenDetail(item.id)" v-for="(item) in props.list">
            <n-progress :class="item.progress == 100 ? 'opacity-60' : ''" style="width: 52px;" :color="item.canceled == '1' ? '#A7ABB5' :'var(--primary-color)' " indicator-text-color="var(--primary-color)"
                type="circle" :percentage="item.progress" :offset-degree="180" :stroke-width="8">
                <p v-if="item.canceled == '0'" class="text-primary-color text-14">{{ item.progress }}<span class="text-12">%</span></p>
                <p v-else class="text-[#A7ABB5] text-12 scale-[0.8333] origin-center break-keep">{{ $t('已取消') }}</p>
            </n-progress>
            <div class="okr-list" :class="item.progress == 100 ? 'opacity-60' : ''">
                <div class="okr-title">
                    <div class="okr-title-l">
                        <span class="scale-[0.8333]" :class="pStatus(item.priority)">{{ item.priority }}</span>
                        <h3 :class="item.progress == 100 ? 'line-through' : ''">{{ item.title }}</h3>
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
         @edit="handleEdit" @upData="(id)=>{emit('upData',id)}"></OkrDetails>

</template>
<script setup lang="ts">
import AlignTargetModal from '@/views/components/AlignTargetModal.vue';
import SelectAlignment from '@/views/components/SelectAlignment.vue'
import OkrDetails from './OkrDetails.vue';



import utils from '@/utils/utils';
import webTs from '@/utils/web'
import { alignUpdate, okrFollow } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"


const alignTargetShow = ref(false)
const selectAlignmentShow = ref(false)
const okrDetailsShow = ref(false)

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





//编辑
const handleEdit = (data) => {
    emit('edit',data)
    okrDetailsShow.value = false
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
<style lang="less" scoped>
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
