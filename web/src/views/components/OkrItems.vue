<template >
    <div class="okr-item-main">
        <div class="okr-item-box" @click="handleOpenDetail(item.id,item.userid)" v-for="(item) in props.list">

            <div class="okr-item-progress hidden md:block" :class="item.completed == '1' || item.canceled == '1' ? 'opacity-60' : ''">
                <n-progress  
                    :color="item.canceled == '1' ? '#A7ABB5' : '#87D068'"
                    indicator-text-color="#87D068" type="circle" :percentage="item.progress" :offset-degree="180"
                    :stroke-width="8">
                    <p v-if="item.canceled == '0'" class="text-primary-color text-14">{{ item.progress }}<span class="text-12">%</span></p>
                    <p v-else class="text-[#A7ABB5] text-12 scale-[0.8333] origin-center break-keep">{{ $t('已取消') }}</p>
                </n-progress> 
            </div>
            

            <div class="okr-list" :class="item.completed == '1' || item.canceled == '1' ? 'opacity-60' : ''">
                <div class="okr-title">
                    <div class="okr-title-l">
                        <h3 class="leading-[1.4]" :class="item.completed == '1' || item.canceled == '1' ? 'line-through' : ''"> <span class="h-[16px] " :class="pStatus(item.priority)">{{ item.priority }}</span>{{ item.title }}</h3>
                    </div>
                    <div class="okr-title-r">
                        <i class="okrfont okr-title-star text-[16px]" v-if="item.is_follow && item.completed == '0' && item.canceled == '0'"
                            @click.stop="handleFollowOkr(item.id)">&#xe683;</i>
                        <i class="okrfont md:pr-16 text-[#8F8F8E] text-[16px]" v-if="!item.is_follow && item.completed == '0' && item.canceled == '0'" @click.stop="handleFollowOkr(item.id)">&#xe679;</i>
                        <i class="okrfont okr-title-icon text-[16px]">&#xe671;</i>
                        <p>{{ item.kr_finish_count }}/{{ item.kr_count }}</p>
                        <div  v-if="item.completed == '1'"
                                class="flex md:hidden items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid cursor-pointer border-primary-color bg-primary-color"
                                >
                                <n-icon class="text-white" size="14" :component="CheckmarkSharp" />
                        </div>
                    </div>
                </div>
                <div class="okr-time">
                    <i class="okrfont text-[12px]">&#xe6e4;</i>
                    <p>{{ item.alias.join(',') }}</p>
                    <div class="w-1 bg-[#F2F3F5] mx-12 h-[12px]"></div>
                    <template v-if="item.canceled == '0' && item.completed =='0'">
                        <i class="okrfont text-[12px]" :class="isOverdue(item.end_at) ?'text-[#ED4014]' : ''"> &#xe6e8;</i>
                        <p :class="isOverdue(item.end_at) ?'text-[#ED4014]' : ''">{{ expiresFormat(item.end_at) }}</p>
                    </template>
                    <template v-if="item.canceled == '1' || item.completed =='1'">
                        <i class="okrfont text-[12px]" > &#xe6e8;</i>
                        <p >{{utils.GoDate(item.start_at) + "~" + utils.GoDate(item.end_at) }}</p>
                    </template>
                </div>

                <div class="okr-time-web">
                    <n-avatar round :size="24" :src="item.user_avatar" />
                    <div class="flex items-center text-12">
                        <i class="okrfont text-14 mr-6">&#xe671;</i>
                        <p class="text-12 mr-16">{{ item.kr_finish_count }}/{{ item.kr_count }}</p>
                        <n-progress class="-mt-7 mr-[6px]" style="width: 14px; " type="circle" :show-indicator="false"
                            :offset-degree="180" :stroke-width="15" color="var(--primary-color)" status="success"
                            :percentage="item.progress" />
                        {{ item.progress }}%
                    </div>

                </div>

                <div class="kr-list">
                    <template v-for="(childItem, index) in item.key_results">
                        <div class="kr-list-item" v-if="index < 3">
                            <span class="bg-[rgba(135,208,104,0.2);] scale-[0.8333]">KR{{ index + 1 }}</span>
                            <p class="max-w-[70%]">{{ childItem.title }}</p>
                            <div class="kr-list-schedule w-[60px] text-text-li">
                                <n-progress class="-mt-7 mr-[6px]" style="width: 15px; " type="circle"
                                    :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                    :color="colorStatus(childItem.progress_status)" status="success" :percentage="childItem.progress" />
                                {{ childItem.progress }}%
                            </div>
                        </div>
                    </template>
                    <div v-if="item.key_results && item.key_results.length > 3" class=" text-12 text-primary-color ml-[2px]">KR(+{{ item.key_results.length - 3 }})</div>
                </div>
                <div class="align-target" v-if="item.align_count > 0">
                    <div class=" cursor-pointer" @click.stop="handleTarget(1, item)">{{ $t('对齐目标') }}({{ item.align_count
                    }}）</div>
                </div>
                <div class="align-target" v-else>
                    <div class=" cursor-pointer" @click.stop="handleTarget(2, item)">
                        + {{ $t('向上对齐') }}
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!-- 查看对齐OKR -->
    <AlignTargetModal :value="alignTargetShow" :eidtItem="eidtItem" @close="() => { alignTargetShow = false }"
        @upData="(id) => { emit('upData', id) }" @openSelectAlignment="(item) => { handleTarget(2, item) }" @openDetail="handleOpenDetail"></AlignTargetModal>

    <!-- 选择对齐OKR -->
    <SelectAlignment :value="selectAlignmentShow" :editData="alignObjective" @close="() => { selectAlignmentShow = false }"
        @submit="submitSelectAlignment"></SelectAlignment>

    <!-- OKR详情 -->
    <OkrDetailsModal ref="RefOkrDetails" :id="eidtId" :show="okrDetailsShow" @close="() => { okrDetailsShow = false }"
        @edit="handleEdit" @upData="(id) => { emit('upData', id) }" @getList="()=>{ emit('getList') }" @openDetail="handleOpenDetail"></OkrDetailsModal>
</template>
<script setup lang="ts">
import AlignTargetModal from '@/views/components/AlignTargetModal.vue';
import SelectAlignment from '@/views/components/SelectAlignment.vue'
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import { CheckmarkSharp } from '@vicons/ionicons5'
import utils from '@/utils/utils';
import webTs from '@/utils/web'
import { alignUpdate, okrFollow } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"
import { useRouter } from 'vue-router'
import { UserStore } from '@/store/user'

const userInfo = UserStore().info
const router = useRouter()
const alignTargetShow = ref(false)
const selectAlignmentShow = ref(false)
const alignObjective = ref([])
const okrDetailsShow = ref(false)

const nowInterval = ref<any>(null)
const nowTime = ref(0)
const loadIng = ref(false)
const message = useMessage()
const eidtId = ref(0)
const userId = ref(0)
const eidtItem = ref({})

const RefOkrDetails = ref(null)

const props = defineProps({
    list: {
        type: undefined,
        default: () => []
    }
});

const emit = defineEmits(['upData', 'edit','getList'])

//对齐
const handleTarget = (e, item) => {
    eidtItem.value = item
    eidtId.value = item.id
    userId.value = item.userid
    if (e == 1) {
        alignTargetShow.value = true
    } else {
        if (userInfo.userid != item.userid) return message.error($t('仅负责人可操作'))
        alignObjective.value = item.align_objective
        selectAlignmentShow.value = true
    }

}

const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

//打开详情
const handleOpenDetail = (id,userid) => {
    if (window.innerWidth < 768) {
        router.push({
            path: '/okrDetails',
            query: {
                data: id,
                userid:userid,
             },
        })
    }
    else {
        eidtId.value = id
        okrDetailsShow.value = true
    }
}


//关注
const handleFollowOkr = (id) => {
    const upData = {
        id: id,
    }
    loadIng.value = true
    okrFollow(upData)
        .then(({ msg }) => {
            message.success($t('操作成功'))
            emit('upData', id, 'FollowOkr')
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

//提交关注
const submitSelectAlignment = (e) => {
    const upData = {
        align_objective: e.join(','),
        id: eidtId.value,
    }
    loadIng.value = true
    alignUpdate(upData)
        .then(({ msg }) => {
            message.success($t('操作成功'))
            emit('upData', eidtId.value)
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

//编辑
const handleEdit = (data) => {
    emit('edit', data)
    okrDetailsShow.value = false
}

//时间
const expiresFormat = (date) => {
    const Dates = new Date(date);
    const timestamp = Dates.getTime();
    return webTs.countDownFormat(timestamp, nowTime.value)
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

const isOverdue = (end_at) => {
    let time = utils.GoDateHMS(end_at)
    return Number(utils.Date(time, true)) < nowTime.value;
}

onMounted(() => {
    nowInterval.value = setInterval(() => {
        nowTime.value = utils.Time();
    }, 1000);
})

onUnmounted(()=>{
    clearInterval(nowInterval.value);
})

</script>
<style lang="less" scoped>
.okr-item-main {
    @apply flex flex-col gap-3 md:gap-4;

    .okr-item-box {
        @apply p-16 md:px-24 md:py-32 bg-white rounded-lg flex gap-4 cursor-pointer;

        .okr-list {
            @apply flex flex-col flex-1 overflow-hidden;

            .okr-title {
                @apply flex justify-between items-start md:items-center;

                .okr-title-l {
                    @apply md:max-w-75p w-full;

                    span {
                        @apply  text-12 mr-4 font-medium text-white px-6 py-0 rounded-full origin-center leading-3 mt-3 md:mt-0;
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
                        @apply text-text-li text-14  font-normal md:font-medium line-clamp-3 md:line-clamp-1 overflow-hidden;
                    }
                }

                .okr-title-r {
                    @apply flex items-center ml-24;

                    .okr-title-star {
                        @apply text-[#FFD023] md:pr-16;
                    }

                    .okr-title-icon {
                        @apply text-text-tips mr-4 hidden md:block;
                    }

                    p {
                        @apply text-text-tips hidden md:block;
                    }
                }
            }

            .okr-time {
                @apply hidden md:flex items-center text-12 mt-4 text-text-tips;

                i {
                    @apply mr-4;
                }
            }

            .okr-time-web {
                @apply flex md:hidden items-center mt-12 text-text-tips justify-between;
            }

            .kr-list {
                @apply mt-16 hidden md:flex flex-col gap-4;

                .kr-list-item {
                    @apply flex items-center;

                    span {
                        @apply text-primary-color text-10 origin-center py-2 px-6 rounded-full flex items-center leading-3;
                    }

                    p {
                        @apply text-text-li text-14 ml-4 truncate;
                    }

                    .kr-list-schedule {
                        @apply ml-auto flex items-center justify-start;
                    }
                }
            }

            .align-target {
                @apply items-start mt-12 text-text-tips text-12 hidden md:flex;
            }
        }
    }

    .okr-item-progress{
        width: 52px;
        .n-progress--circle{
            width: 100% !important;
        }
    }
}
</style>
