<template>
    <div class="flex flex-col h-full">
        <div class="flex justify-between flex-col 2xl:flex-row">
            <div class="flex-[2] hidden md:flex items-center mb-16 mt-24">
                <div v-if="userInfo == 'admin'" class=" mb-2 mr-8 whitespace-nowrap">
                    {{ $t('部门') }}
                </div>
                <n-select v-if="userInfo == 'admin'" v-model:value="departmentsvalue" :options="departments"
                    class="w-[33%] h-[36px] mr-24" :placeholder="$t('全部')" />
                <div class=" mb-2 mr-8 whitespace-nowrap">
                    {{ $t('负责人') }}
                </div>
                <n-select v-model:value="principalvalue" :options="principal" class="w-[33%] h-[36px] "
                    :placeholder="$t('全部')" />

                <div class=" mb-2 mr-8 ml-24 whitespace-nowrap">
                    {{ $t('时间') }}
                </div>
                <n-date-picker class="w-[33%] h-[36px] " v-model:value="daterange" value-format="yyyy.MM.dd HH:mm:ss"
                    type="daterange" clearable size="medium" />

                <n-button :loading="isloading" type="primary" class=" mb-4 ml-24 rounded px-16" @click="handleClick()">
                    <template #icon>
                        <i class="taskfont" v-if="!(isloading)">&#xe72a;</i>
                    </template>
                    {{ $t('搜索') }}
                </n-button>
            </div>
            <div
                class="flex-1 flex items-center mt-16 md:mt-0 mb-16 justify-between md:justify-start 2xl:justify-end 2xl:mb-0">
                <div class=" hidden md:block mb-4 mr-12 whitespace-nowrap 2xl:ml-24">
                    {{ $t('类型') }}
                </div>
                <n-button-group size="medium" class="mb-4 hidden md:flex">
                    <n-button @click="handleClick2(null)" class=" rounded px-16 bg-white whitespace-nowrap">
                        {{ $t('全部') }}
                    </n-button>
                    <div class="w-[4px]"></div>
                    <n-button @click="handleClick2('1')" class=" rounded px-16 bg-white whitespace-nowrap">
                        {{ $t('承诺型') }}
                    </n-button>
                    <n-button @click="handleClick2('2')" class=" rounded px-16 bg-white whitespace-nowrap">
                        {{ $t('挑战型') }}
                    </n-button>
                </n-button-group>

                <n-checkbox v-model:checked="completednotrated" class="md:ml-36 rounded whitespace-nowrap mb-2 "
                    @click="getList('search')">
                    <span class="text-text-tips">{{ $t('已完成未评分') }}</span>
                </n-checkbox>
                <div @click="active = true" class="flex md:hidden text-text-tips text-14">
                    <i class="taskfont">&#xe700;</i>
                    {{ $t('筛选') }}
                </div>

            </div>
        </div>
        <div class="relative h-full">
            <div class="absolute top-0 left-0 bottom-0 right-0">
                <n-scrollbar :on-scroll="onScroll">
                    <div class="okr-department-main">
                        <OkrLoading v-if="loadIng"></OkrLoading>
                        <OkrItems v-if="list.length != 0 && !loadIng" @upData="upData" @edit="handleEdit" :list="list">
                        </OkrItems>
                        <OkrNotDatas v-if="!loadIng && list.length == 0" :msg="$t('暂无OKR')" :types="searchObject">
                        </OkrNotDatas>
                        <OkrLoading v-if="onscrolloading" position="onscroll"></OkrLoading>
                    </div>
                </n-scrollbar>
            </div>
        </div>
        <n-drawer v-model:show="active" default-height="60vh" placement="bottom" resizable>

            <n-drawer-content>
                <template #header>
                    <div class="flex w-full items-center justify-between ">
                        {{ $t('筛选') }}
                        <i class="taskfont text-text-tips" @click="active = false">&#xe6e5;</i>
                    </div>
                </template>
                <div class="flex flex-col h-full">
                    <div v-if="userInfo == 'admin'" class="whitespace-nowrap">
                        {{ $t('部门') }}
                    </div>
                    <n-select v-if="userInfo == 'admin'" v-model:value="departmentsvalue" :options="departments"
                        class=" h-[36px] mr-24" :placeholder="$t('全部')" />

                    <div class=" whitespace-nowrap" :class="userInfo == 'admin' ? 'mt-16' : ''">
                        {{ $t('负责人') }}
                    </div>
                    <n-select v-model:value="principalvalue" :options="principal" class="h-[36px] "
                        :placeholder="$t('全部')" />

                    <div class="mt-16 whitespace-nowrap">
                        {{ $t('类型') }}
                    </div>
                    <n-select v-model:value="types" :options="typeList" class="h-[36px] "
                        :placeholder="$t('全部')" />


                    <div class="mt-16 whitespace-nowrap">
                        {{ $t('时间') }}
                    </div>
                    <n-date-picker class=" h-[36px] " v-model:value="daterange" value-format="yyyy.MM.dd HH:mm:ss"
                        type="daterange" clearable size="medium" />

                    <div class="flex justify-between gap-4 mt-auto">
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

const active = ref(false)
const loadIng = ref(false)
const isloading = ref(false)
const onscrolloading = ref(false)
const userInfo = UserStore().info.identity[0]
const page = ref(1)
const last_page = ref(99999)
const list = ref([])
const departmentsvalue = ref(null)
const principalvalue = ref(null)
const types = ref(null);
const daterange = ref(null)
const completednotrated = ref(false)
const searchTime = ref(null)

const emit = defineEmits(['edit'])

const departments = ref([
    { label: $t('全部'), value: null },
])
const principal = ref([
    { label: $t('全部'), value: null },
])
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
    searchTime.value = setInterval(() => {
        page.value = 1
        getList('search')
        clearInterval(searchTime.value)
    }, 300)
}, { deep: true })

const init = () => {
    const sendata = {
        page: page.value,
        page_size: 10,
    }
    getUserList(sendata).then(({ data }) => {
        data.data.map(item => {
            principal.value.push({
                label: item.nickname,
                value: item.userid,
            })
        })
    })

    getDepartmentList().then(({ data }) => {
        data.data.map(item => {
            departments.value.push({
                label: item.name,
                value: item.id,
            })
        })
    })

}

const getList = (type) => {
    let serstatic = type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic) {
        const sendata = {
            completed: completednotrated.value ? 1 : 0,
            department_id: departmentsvalue.value,
            end_at: daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[1] / 1000) : null,
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
            start_at: daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[0] / 1000) : null,
            type: types.value == "0" ? null : types.value,
            userid: principalvalue.value,
        }
        if (serstatic) {
            loadIng.value = true
        } else if (type == 'onscrollsearch') {
            onscrolloading.value = true
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
    getList('search');
}

//重置
const handleReset = ()=>{
    departmentsvalue.value = null
    principalvalue.value = null
    types.value = null
    daterange.value = null
    getList('');
}

//类型
const handleClick2 = (type) => {
    types.value = type
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
        if (!loadIng.value) {
            page.value++
            getList('onscrollsearch')
        }
    }
}


onMounted(() => {
    init()
    getList('search')
})
defineExpose({
    upData,
    getList
})
</script>
<style lang="less" scoped>
.okr-department-main {
    @apply flex flex-col gap-6 pt-0;
}

:deep(.n-drawer-header__main) {
    @apply flex-1;
}
</style>
