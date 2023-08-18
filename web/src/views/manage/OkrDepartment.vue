<template>
    <div class="flex flex-col h-full">
        <div class="flex justify-between flex-col 2xl:flex-row 2xl:items-center md:mb-16 md:mt-24">
            <div class="flex-[2] hidden md:flex items-center mb-16 2xl:mb-0">
                <div v-if="userInfo == 'admin'" class="mb-2 mr-8 text-text-li whitespace-nowrap">
                    {{ $t('部门') }}
                </div>
                <n-select v-if="userInfo == 'admin'" v-model:value="departmentsvalue" :options="departments" clearable
                    class="w-[33%] mr-16" :placeholder="$t('全部')" />
                <div class="text-text-li mr-8 whitespace-nowrap">
                    {{ $t('负责人') }}
                </div>
                <n-select v-model:value="principalvalue" :options="principal" :on-search="getUser" class="w-[33%]" filterable :placeholder="$t('全部')" clearable>
                    <template #action>
                        <div v-if="principallast_page > principalpage" quaternary class=" h-full w-full whitespace-nowrap text-center" @click.stop="principalClick('')">
                            {{ $t('更多...') }}
                        </div>
                        <div v-else quaternary class=" h-full w-full whitespace-nowrap text-center">
                            {{ $t('已经到底了') }}
                        </div>
                    </template>
                </n-select>

                <div class="text-text-li mr-8 ml-16 whitespace-nowrap">{{ $t('时间') }}</div>
                <div v-if="showDatePickers" class="okr-date-picker-waps" style="width: 30%;">
                    <DatePickers />
                </div>
                <n-date-picker v-else class="w-[33%]" v-model:value="daterange" value-format="yyyy.MM.dd HH:mm:ss"
                    type="daterange" clearable size="medium" />

                <n-button :loading="isloading" type="primary" size="small" class="ml-24 rounded px-16" @click="handleClick()">
                    <template #icon>
                        <i class="taskfont" v-if="!(isloading)">&#xe72a;</i>
                    </template>
                    {{ $t('搜索') }}
                </n-button>
            </div>
            <div
                class="flex-1 flex items-center mt-12 md:mt-0 mb-12 justify-between md:justify-start 2xl:justify-end 2xl:mb-0 ">
                <div class=" hidden md:block mr-12 whitespace-nowrap 2xl:ml-24 ">
                    {{ $t('类型') }}
                </div>
                <div class="hidden md:flex">
                    <div @click="handleClick2(null)" :class="types == null ? 'bg-[rgba(139,207,112,0.05)] border-[#8BCF70] text-primary-color':'bg-white border-[#F2F2F2]' " class="rounded px-12 py-6 border-solid border-[1px] cursor-pointer">
                        {{ $t('全部') }}
                    </div>
                    <div class="flex ml-4">
                        <div @click="handleClick2(1)" :class="types == '1' ? 'bg-[rgba(139,207,112,0.05)] border-[#8BCF70] text-primary-color z-[2]':'bg-white border-[#F2F2F2]' " class="rounded-l px-12 py-6 border-solid border-[1px] border-[#F2F2F2] bg-white cursor-pointer">
                            {{ $t('承诺型') }}
                        </div>
                        <div @click="handleClick2(2)" :class="types == '2' ? 'bg-[rgba(139,207,112,0.05)] border-[#8BCF70] text-primary-color':'bg-white border-[#F2F2F2]' " class="-ml-1 rounded-r px-12 py-6 border-solid border-[1px] border-[#F2F2F2] bg-white cursor-pointer">
                            {{ $t('挑战型') }}
                        </div>
                    </div>
                </div>

                <n-checkbox v-model:checked="completednotrated" class="md:ml-16 rounded whitespace-nowrap mb-2 "
                    @click="handleClick()">
                    <span class="text-text-li">{{ $t('已完成未评分') }}</span>
                </n-checkbox>
                <div @click="active = true" class="flex items-center md:hidden text-14" :class=" searchActive ? 'text-primary-color' : 'text-text-tips'">
                    <i class="taskfont text-18 mr-4" >&#xe700;</i>
                    {{ $t('筛选') }}
                </div>

            </div>
        </div>
        <div class="relative h-full">
            <div class="absolute top-0 left-0 bottom-0 right-0">
                <n-scrollbar :on-scroll="onScroll">
                    <div class="okr-department-main">
                        <OkrLoading v-if="loadIng"></OkrLoading>
                        <OkrItems v-if="list.length != 0 && !loadIng" @upData="upData" @edit="handleEdit" @getList="resetGetList" :list="list">
                        </OkrItems>
                        <OkrNotDatas v-if="!loadIng && !onscrolloading && list.length == 0" :loadIng="loadIng" :msg="$t('暂无OKR')" :types="searchObject !='' || loadingstatus ">
                        </OkrNotDatas>
                        <OkrLoading v-if="onscrolloading" position="onscroll"></OkrLoading>
                    </div>
                </n-scrollbar>
            </div>
        </div>

        <n-drawer v-model:show="active" default-height="60vh" placement="bottom" resizable>

            <n-drawer-content>
                <template #header>
                    <div class="flex w-full items-center justify-between text-16 md:text-18">
                        {{ $t('筛选') }}
                        <i class="taskfont text-text-tips" @click="active = false">&#xe6e5;</i>
                    </div>
                </template>
                <div class="flex flex-col h-full">
                    <div v-if="userInfo == 'admin'" class="whitespace-nowrap mb-4">
                        {{ $t('部门') }}
                    </div>
                    <n-select v-if="userInfo == 'admin'" v-model:value="departmentsvalue" :options="departments" clearable
                        class=" h-[36px] mr-24" :placeholder="$t('全部')" />

                    <div class=" whitespace-nowrap mb-4" :class="userInfo == 'admin' ? 'mt-16' : ''">
                        {{ $t('负责人') }}
                    </div>
                    <n-select v-model:value="principalvalue" :options="principal" :on-search="getUser" class="h-[36px]" filterable :placeholder="$t('全部')" clearable>
                        <template #action>
                            <div v-if="principallast_page > principalpage" quaternary class=" h-full w-full whitespace-nowrap text-center" @click.stop="principalClick('')">
                                {{ $t('更多...') }}
                            </div>
                            <div v-else quaternary class=" h-full w-full whitespace-nowrap text-center">
                                {{ $t('已经到底了') }}
                            </div>
                        </template>
                    </n-select>

                    <div class="mt-16 whitespace-nowrap mb-4">
                        {{ $t('类型') }}
                    </div>
                    <n-select v-model:value="types" :options="typeList" class="h-[36px] " clearable
                        :placeholder="$t('全部')" />


                    <div class="mt-16 whitespace-nowrap mb-4">
                        {{ $t('时间') }}
                    </div>
                    <div v-if="showDatePickers" class="okr-date-picker-waps"><DatePickers /></div>
                    <n-date-picker v-else class="h-[36px]" v-model:value="daterange" value-format="yyyy.MM.dd HH:mm:ss"
                        type="daterange" clearable size="medium" />


                    <div class="flex justify-between gap-4 mt-auto border-solid pt-12 border-0  border-t-[1px] border-[#F2F2F2]">
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
    </div>
</template>

<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getDepartmentOkrList } from '@/api/modules/department'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import { getDepartmentList } from '@/api/modules/department'
import { getOkrDetail } from '@/api/modules/okrList'
import { getUserList } from '@/api/modules/created'
import utils from '@/utils/utils'
import { UserStore } from '@/store/user'
import OkrLoading from '../components/OkrLoading.vue'

const datePickerApps = ref([])
const active = ref(false)
const loadIng = ref(false)
const isloading = ref(false)
const onscrolloading = ref(false)
const loadingstatus = ref(false)
const userInfo = UserStore().info.identity[0]
const page = ref(1)
const principalpage = ref(1)
const last_page = ref(99999)
const principallast_page = ref(99999)
const list = ref([])
const departmentsvalue = ref(null)
const principalvalue = ref(null)
const types = ref(null);
const daterange = ref([])
const completednotrated = ref(false)
const searchTime = ref(null)
const keyWord = ref('')

const emit = defineEmits(['edit'])

const showDatePickers = computed(() => window.__MICRO_APP_BASE_APPLICATION__ ? 1 : 0)

const departments = ref([
    { label: $t('全部'), value: null },
])
const principal = ref([])
const typeList = ref([
    { label: $t('全部'), value: null },
    { label: $t('承诺型'), value: 1 },
    { label: $t('挑战型'), value: 2 },
])

const props = defineProps({
    searchObject: {
        type: String,
        default: "",
    }
})

watch(() => props.searchObject, (newValue) => {
    clearInterval(searchTime.value)
    loadIng.value = true
    searchTime.value = setInterval(() => {
        page.value = 1
        getList('search')
        clearInterval(searchTime.value)
    }, 300)
}, { deep: true })

watch(() => active.value, (newValue) => {
    if(newValue){
        loadDatePickers()
    }else{
        unmountDatePickerApps()
    }
})


const searchActive = computed(()=>{
    return departmentsvalue.value != null || principalvalue.value  != null || types.value  != null ||  daterange.value.length  != 0
})

const getUser = (keyword) => {
    if ( keyword == '' ) {
        principalpage.value = 1
    }
    keyWord.value = keyword
    const sendata = {
        dept_only: userInfo == 'admin' ? false : true,
        page: principalpage.value,
        page_size: 20,
        keyword: keyword,
    }

    getUserList(sendata).then(({ data }) => {
        if ( keyword == '' ) {
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

const init = () => {
    principalClick('init')
    getDepartmentList().then(({ data }) => {
        data.data.map(item => {
            departments.value.push({
                label: item.name,
                value: item.id,
            })
        })
    })
}

const resetGetList = ()=>{
    page.value = 1
    getList('search')
}


const getList = (type) => {
    let serstatic = type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic) {
        if (serstatic) {
            loadIng.value = true
        } else if (type == 'onscrollsearch') {
            onscrolloading.value = true
        }
        const sendata = {
            completed: completednotrated.value ? 1 : 0,
            department_id: departmentsvalue.value,
            end_at: daterange.value[1] ? utils.TimeHandle(daterange.value[1],2) : '',
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
            start_at: daterange.value[0] ? utils.TimeHandle(daterange.value[0]) : '',
            type: types.value == "0" ? null : types.value,
            userid: principalvalue.value,
        }
        getDepartmentOkrList(sendata).then(({ data }) => {
            loadIng.value = false
            isloading.value = false
            onscrolloading.value = false
            if (serstatic) {
                data.data ? list.value = data.data : list.value = []
            }
            else {
                if (data.data) {
                    data.data.map(item => {
                        list.value.push(item)
                    })
                }
            }
            last_page.value = data.last_page
        })
    }
}

//搜索
const handleClick = () => {
    isloading.value = true
    page.value = 1
    loadingstatus.value=true
    active.value = false
    getList('search');
}

//类型
const handleClick2 = (type) => {
    types.value = type
    page.value=1
    loadingstatus.value=true
    getList('search');
}

//负责人
const principalClick = (type) => {
    if (type == 'init'){
        principalpage.value = 1
        principal.value = ([
            { label: $t('全部'), value: null }
        ])
    }else{
        principalpage.value++
    }
    const sendata = {
        dept_only: userInfo == 'admin' ? false : true,
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

//重置
const handleReset = ()=>{
    departmentsvalue.value = null
    principalvalue.value = null
    types.value = null
    daterange.value = null
    loadingstatus.value=false
    active.value = false
    page.value=1
    getList('search');
}

//编辑
const handleEdit = (data) => {
    emit('edit', data)
}

//更新数据
const upData = (id) => {
    list.value.map((item, index) => {
        if (item.id == id) {
            const upData = {
                id: id,
            }
            getOkrDetail(upData)
                .then(({ data }) => {
                    list.value[index] = data
                })
                .catch()
                .finally(() => {
                })
        }
    })
}


const onScroll = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!onscrolloading.value && !loadIng.value) {
            if(last_page.value > page.value){
                page.value++
                getList('onscrollsearch')
            }
        }
    }
}

// 加载时间组件
const loadDatePickers = () => {
    nextTick(() => {
        if (!window.Vues || !showDatePickers) return false;
        unmountDatePickerApps()
        document.querySelectorAll('datepickers').forEach(e => {
            let app = new window.Vues.Vue({
                el: document.querySelector('DatePickers'),
                store: window.Vues.store,
                render: (h: any) => {
                    return h(window.Vues?.components?.DatePicker, {
                        class: "okr-app-date-pickers",
                        props: {
                            editable: false,
                            placeholder: $t("选择时间范围"),
                            format: "yyyy-MM-dd",
                            type:"daterange",
                            placement: "bottom-end",
                            confirm: true,
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
    if(datePickerApps.value){
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
    loadDatePickers()
    getList('search')
})

defineExpose({
    upData,
    getList,
    resetGetList
})
</script>
<style lang="less" scoped>
.okr-department-main {
    @apply flex flex-col gap-6 pt-0;
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
:deep(.n-checkbox__label) {
    @apply pl-4 !important;
}
:deep(.n-drawer-body-content-wrapper) {
    @apply px-16 !important;
}
</style>
