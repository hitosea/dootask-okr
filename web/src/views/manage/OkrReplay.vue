<template>
    <div class="okr-replay-main">

        <div class="flex flex-wrap flex-col 2xl:flex-row 2xl:items-center flex-[0_0_auto]">
            <div class="hidden md:flex items-center mb-16  min-w-full">
                <div class="flex-1 flex items-center overflow-hidden" v-if="(userInfo == 'admin' || okrAdminOwner)">
                    <div class="mb-2 mr-8 text-text-li whitespace-nowrap font-medium">
                        {{ $t('部门') }}
                    </div>
                    <n-select v-model:value="departmentsvalue" :options="departments" clearable
                        class="mr-16 flex-1 overflow-hidden" :placeholder="$t('全部')" />
                </div>

                <div class=" flex items-center overflow-hidden"
                    :class="(userInfo == 'admin' || okrAdminOwner) ? 'flex-1' : ''">
                    <div class="text-text-li mr-8 whitespace-nowrap font-medium">
                        {{ $t('负责人') }}
                    </div>
                    <n-select v-model:value="principalvalue" :options="principal" :on-search="getUser"
                        :class="(userInfo == 'admin' || okrAdminOwner) ? '' : 'max-w-[225px] '"
                        class="flex-1 overflow-hidden" filterable :placeholder="$t('全部')" clearable>
                        <template #action>
                            <div v-if="principallast_page > principalpage" quaternary
                                class=" h-full w-full whitespace-nowrap text-center" @click.stop="principalClick('')">
                                {{ $t('更多...') }}
                            </div>
                            <div v-else quaternary class=" h-full w-full whitespace-nowrap text-center">
                                {{ $t('已经到底了') }}
                            </div>
                        </template>
                    </n-select>
                </div>

                <div class="flex items-center" :class="(userInfo == 'admin' || okrAdminOwner) ? 'flex-1' : ''">
                    <div class="mr-8 ml-16 whitespace-nowrap text-text-li font-medium">{{ $t('时间') }}</div>
                    <div v-if="showDatePickers" :class="(userInfo == 'admin' || okrAdminOwner) ? '' : 'max-w-[225px] '"
                        class="okr-date-picker-waps ">
                        <DatePickers />
                    </div>
                    <n-date-picker v-else :class="(userInfo == 'admin' || okrAdminOwner) ? '' : 'max-w-[225px] '"
                        v-model:value="daterange" value-format="yyyy.MM.dd HH:mm:ss" type="daterange" clearable
                        size="medium" />
                </div>

                <n-button :loading="isloading" type="primary" size="small" class="ml-16 rounded px-16"
                    @click="handleClick()">
                    <template #icon>
                        <i class="okrfont" v-if="!(isloading)">&#xe72a;</i>
                    </template>
                    {{ $t('搜索') }}
                </n-button>
            </div>
            <div class="flex items-center mb-12 justify-between md:justify-start">
                <n-checkbox v-model:checked="completednotrated" class="rounded whitespace-nowrap mb-2 "
                    @click="handleClick()">
                    <span class="text-text-li">{{ $t('已评分未复盘') }}</span>
                </n-checkbox>
                <div @click="active = true" class="flex items-center md:hidden text-14"
                    :class="searchActive ? 'text-primary-color' : 'text-text-tips'">
                    <i class="okrfont text-18 mr-4">&#xe700;</i>
                    {{ $t('筛选') }}
                </div>

            </div>
        </div>
        <div>
            <OkrNotDatas v-if="items.length == 0 && !loadIng && !onscrolloading" :loadIng="loadIng" :msg="$t('暂无数据')"
                :types="searchObject != ''"></OkrNotDatas>
            <OkrLoading v-if="loadIng"></OkrLoading>
        </div>
        <n-scrollbar :on-scroll="onScroll">
            <div v-if="items.length != 0" class="replay flex-[1_1_auto] ">
                <div v-for="(item, index) in items" :key="index" :class="{ 'replay-item': true, 'replay-item-active': item.isActive }">
                    <div class="replay-item-head">
                        <div>
                            <span class="replay-item-okr-level py-[0.5px]" :class="pStatus(item.okr_priority)">{{
                                item.okr_priority
                            }}</span>
                            <span class="text-[14px] m-[5px] text-title-color font-medium"><b class="font-medium">{{
                                (item.okr_alias || []).join(",") }}</b></span>
                            <span class=" text-text-li text-12">{{ $t("的目标复盘") }}</span>
                        </div>
                        <div class="cursor-pointer hidden md:block" @click="() => (item.isActive = !item.isActive)">
                            <span class="mr-[10px]">{{ item.isActive == true ? $t("收起") : $t("展开") }}</span>
                            <i class="replay-item-head-icon okrfont">&#xe705;</i>
                        </div>
                    </div>
                    <div class="flex">
                        <div class="replay-item-okr cursor-pointer px-[16px] py-[9.5px] bg-[#f4f5f7]"
                            @click.stop="openOkrDetail(item.okr_id)">
                            <div class="replay-item-okr-icon w-[25px] h-[16px] shrink-0">{{ item.objective_num || ('0' +
                                item.okr_id) }}</div>
                            <div class="text-[#515A6E] text-14 line-clamp-1">{{ item.okr_title }}</div>
                        </div>
                    </div>
                    <div class="replay-item-body" v-if="item.isActive">
                        <OkrReplayDetail :okrReplayList="item"></OkrReplayDetail>
                    </div>
                </div>
            </div>
        </n-scrollbar>
        <n-drawer v-model:show="active" default-height="370px" placement="bottom" resizable>

            <n-drawer-content class="screen-d">
                <template #header>
                    <div class="flex w-full items-center justify-between text-text-li text-16 md:text-18">
                        {{ $t('筛选') }}
                        <i class="okrfont text-[#999]" @click="active = false">&#xe6e5;</i>
                    </div>
                </template>
                <div class="flex flex-col h-full">
                    <div v-if="(userInfo == 'admin' || okrAdminOwner)" class="whitespace-nowrap text-text-li mb-4">
                        {{ $t('部门') }}
                    </div>
                    <n-select v-if="(userInfo == 'admin' || okrAdminOwner)" v-model:value="departmentsvalue"
                        :options="departments" clearable class="mr-24" :placeholder="$t('全部')" />

                    <div class=" whitespace-nowrap text-text-li mb-4"
                        :class="(userInfo == 'admin' || okrAdminOwner) ? 'mt-16' : ''">
                        {{ $t('负责人') }}
                    </div>
                    <n-select v-model:value="principalvalue" :options="principal" :on-search="getUser" class="" filterable
                        :placeholder="$t('全部')" clearable>
                        <template #action>
                            <div v-if="principallast_page > principalpage" quaternary
                                class=" h-full w-full whitespace-nowrap text-center" @click.stop="principalClick('')">
                                {{ $t('更多...') }}
                            </div>
                            <div v-else quaternary class=" h-full w-full whitespace-nowrap text-center">
                                {{ $t('已经到底了') }}
                            </div>
                        </template>
                    </n-select>

                    <div class="mt-16 whitespace-nowrap mb-4 text-text-li">
                        {{ $t('时间') }}
                    </div>
                    <div v-if="showDatePickers" class="okr-date-picker-waps">
                        <DatePickers />
                    </div>
                    <n-date-picker v-else class="h-[36px]" v-model:value="daterange" value-format="yyyy.MM.dd HH:mm:ss"
                        type="daterange" clearable size="medium" />


                    <div
                        class="flex justify-between gap-4 mt-auto border-solid pt-12 border-0  border-t-[1px] border-[#F2F2F2]">
                        <n-button :loading="isloading" strong secondary type="tertiary" class="flex-1" @click="handleReset">
                            {{ $t('重置') }}
                        </n-button>
                        <n-button :loading="isloading" type="primary" class="flex-1" @click="handleClick()">
                            {{ $t('搜索') }}
                        </n-button>
                    </div>

                </div>
            </n-drawer-content>
        </n-drawer>

        <OkrLoading v-if="onscrolloading" position='onscroll'></OkrLoading>
        <!-- OKR详情 -->
        <OkrDetailsModal ref="RefOkrDetails" :id="detailId" :show="okrDetailsShow" @openDetail="openOkrDetail" @close="() => {
            okrDetailsShow = false
        }
            "></OkrDetailsModal>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import OkrReplayDetail from "@/views/components/OkrReplayDetails.vue"
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import { getReplayList } from '@/api/modules/replay'
import { getDepartmentList } from '@/api/modules/department'
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import OkrLoading from '../components/OkrLoading.vue'
import { GlobalStore } from '@/store';
import { UserStore } from '@/store/user'
import { getUserList } from '@/api/modules/created'
import { getUserInfo } from '@/api/modules/user'
import utils from "@/utils/utils"
import { isMicroApp, getAppData } from "@/utils/app"

const datePickerApps = ref([])
const globalStore = GlobalStore()
const items = ref([])
const loadIng = ref(false)
const isloading = ref(false)
const onscrolloading = ref(false)
const page = ref(1)
const last_page = ref(99999)

const active = ref(false)

const keyWord = ref('')

const departmentsvalue = ref(null)
const departments = ref([
    { label: $t('全部'), value: null },
])

const principal = ref([])
const principalvalue = ref(null)
const principallast_page = ref(99999)
const principalpage = ref(1)

const userInfo = UserStore().info.identity[0]
const okrAdminOwner = ref(false)
const showDatePickers = computed(() => isMicroApp() ? 1 : 0)

const daterange = ref<[number, number]>([0, 0])

const completednotrated = ref(false)

const okrDetailsShow = ref(false)
const detailId = ref(0)

const searchTime = ref(null)
const props = defineProps({
    searchObject: {
        type: String,
        default: "",
    }
})

const searchActive = computed(() => {
    return departmentsvalue.value != null || principalvalue.value != null || daterange.value[0] != 0 || daterange.value[1] != 0
})


watch(() => props.searchObject, (newValue) => {
    clearInterval(searchTime.value)
    loadIng.value = true
    searchTime.value = setInterval(() => {
        page.value = 1
        getData('search')
        clearInterval(searchTime.value)
    }, 300)
}, { deep: true })

watch(() => globalStore.addMultipleChange, (newValue) => {
    if (newValue) {
        getData('search')
    }
}, { deep: true })

watch(() => active.value, (newValue) => {
    if (newValue) {
        loadDatePickers()
    } else {
        unmountDatePickerApps()
    }
})

// 获取数据
const getData = (type) => {
    let serstatic = type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic) {
        // 获取复盘列表
        const data = {
            select_replayed: completednotrated.value ? 1 : 0,//完成未完成
            department_id: departmentsvalue.value,//部门
            start_at: daterange.value[0] ? utils.TimeHandle(daterange.value[0],2) : '',//开始时间
            end_at: daterange.value[1] ? utils.TimeHandle(daterange.value[1]) : '',//结束时间
            userid: principalvalue.value,//用户
            objective: props.searchObject,//关键词
            page: page.value,
            page_size: 20,
        }
        if (serstatic) {
            loadIng.value = true
        } else if (type == 'onscrollsearch') {
            onscrolloading.value = true
        }
        getReplayList(data).then(({ data }) => {
            onscrolloading.value = false
            loadIng.value = false
            isloading.value = false
            if (serstatic) {
                data.data ?
                    items.value = data.data
                    : items.value = []
            } else {
                (data.data || []).map((item) => {
                    item.isActive = false
                    items.value.push(item)
                })
            }
            last_page.value = data.last_page
        }).catch(({ msg }) => {
            onscrolloading.value = false
            loadIng.value = false
            isloading.value = false
        })
        .finally(() => {
            onscrolloading.value = false
            loadIng.value = false
            isloading.value = false
        })
    }
}

const init = () => {
    getDepartmentList().then(({ data }) => {
        data.data.map(item => {
            departments.value.push({
                label: item.name,
                value: item.id,
            })
        })
    })
}



//获取能否能搜索部门
const handleGetUserInfo = () => {
    getUserInfo().then(({ data }) => {
        okrAdminOwner.value = data.okr_admin_owner
        principalClick('init')
    })
}

const onScroll = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!onscrolloading.value && !loadIng.value) {
            if (last_page.value > page.value) {
                page.value++
                getData("onscrollsearch")
            }
        }
    }
}

// 状态
const pStatus = (p) => {
    return p == "P0" ? "span-1" : p == "P1" ? "span-2" : "span-3"
}

//查看okr详情
const openOkrDetail = (id) => {
    okrDetailsShow.value = true
    detailId.value = id
}

//重新获取
const resetGetList = () => {
    page.value = 1
    getData('search')
}

//负责人
const principalClick = (type) => {
    if (type == 'init') {
        principalpage.value = 1
        principal.value = ([
            { label: $t('全部'), value: null }
        ])
    } else {
        principalpage.value++
    }
    const sendata = {
        dept_only: (!(userInfo == 'admin' || okrAdminOwner.value)),
        page: principalpage.value,
        page_size: 20,
        keyword: keyWord.value,
    }
    getUserList(sendata).then(({ data }) => {
        if (data.data) {
            data.data.map(item => {
                principal.value.push({
                    label: item.nickname,
                    value: item.userid,
                })
            })
        }
        principallast_page.value = data.last_page
    })
}

// 获取用户
const getUser = (keyword) => {
    if (keyword == '') {
        principalpage.value = 1
    }
    keyWord.value = keyword
    const sendata = {
        dept_only: (!(userInfo == 'admin' || okrAdminOwner.value)),
        page: principalpage.value,
        page_size: 20,
        keyword: keyword,
    }

    getUserList(sendata).then(({ data }) => {
        if (keyword == '') {
            principal.value = ([
                { label: $t('全部'), value: null }
            ])
            data.data.map(item => {
                principal.value.push({
                    label: item.nickname,
                    value: item.userid,
                })
            })
        }
        else {
            principal.value = ([])
            if (data.data) {
                data.data.map(item => {
                    principal.value.push({
                        label: item.nickname,
                        value: item.userid,
                    })
                })
            }
        }
        principallast_page.value = data.last_page
    })

}


//搜索
const handleClick = () => {
    isloading.value = true
    page.value = 1
    getData('search');
}

//重置
const handleReset = () => {
    departmentsvalue.value = null
    principalvalue.value = null
    daterange.value = [0, 0]
    active.value = false
    page.value = 1
    getData('search');
}

// 加载时间组件
const loadDatePickers = () => {
    nextTick(() => {
        if (!isMicroApp()) return false;
        unmountDatePickerApps()
        document.querySelectorAll('datepickers').forEach(e => {
            const instance = getAppData('instance')
            const app = new instance.Vue({
                el: document.querySelector('DatePickers'),
                store: instance.store,
                render: (h: any) => {
                    return h(instance.components?.DatePicker, {
                        class: "okr-app-date-pickers",
                        props: {
                            editable: false,
                            placeholder: $t("请选择时间"),
                            format: "yyyy-MM-dd",
                            type: "daterange",
                            placement: "bottom-end",
                            confirm: true,
                            options: {shortcuts: getAppData('methods.extraCallA')?.('timeOptionShortcuts')}
                        },
                        on: {
                            "on-change": (e: any) => {
                                daterange.value = e
                            }
                        }
                    })
                }
            });
            datePickerApps.value.push(app);
        })
    })
}
// 卸载时间组件
const unmountDatePickerApps = () => {
    if (datePickerApps.value) {
        datePickerApps.value.forEach(app => {
            app.$el.replaceWith(document.createElement("DatePickers"));
            app.$destroy()
        })
        datePickerApps.value = [];
    }
}

// 重载
window.addEventListener('resize', function () {
    loadDatePickers()
})

// 卸载
window.addEventListener('apps-unmount', function () {
    unmountDatePickerApps()
})

onBeforeUnmount(() => {
    unmountDatePickerApps()
})

onMounted(() => {
    init()
    getData('search')
    loadDatePickers()
    handleGetUserInfo()
})
defineExpose({
    resetGetList,
})
</script>

<style lang="less" scoped>
.okr-replay-main {
    @apply pt-16 md:pt-24 flex flex-col h-full;
}

.replay {
    width: 100%;
    height: 100%;
    @apply flex flex-col gap-3 md:gap-4;

    &-item {
        @apply bg-white rounded-lg h-auto p-16 md:p-24;

        &-head {
            @apply flex items-center justify-between text-[#87d068];

            &-icon {
                display: inline-block;
                transform: rotate(0deg);
                transition: transform 0.5s;
            }
        }

        &-okr {
            @apply mt-10 h-auto flex rounded-lg items-center;

            &-icon {
                @apply text-center text-10 leading-4 rounded-lg bg-[#87D068] text-[#87d068] mr-10;
                background: rgba(135, 208, 104, 0.2);
            }

            &-level {
                @apply bg-[#FF7070] text-12 font-medium text-white px-6 rounded-full;
            }
        }

        &-body {
            @apply mt-14;
        }
    }

    &-item-active &-item-head {
        &-icon {
            transform: rotate(-180deg);
        }
    }
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

:deep(.n-drawer-header__main) {
    @apply flex-1;
}

:deep(.n-checkbox-box__border) {
    @apply rounded !important;
}

:deep(.n-checkbox .n-checkbox-box) {
    @apply rounded !important;
}
</style>
