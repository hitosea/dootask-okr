<template>
    <div class="flex flex-col h-full">
        <div class="flex justify-between flex-col 2xl:flex-row" >
            <div class="flex-[2] flex items-center mb-16 mt-24 ">
                <div v-if="userInfo == 'admin'" class=" mb-2 mr-8 whitespace-nowrap">
                    {{$t('部门')}}
                </div>
                <n-select
                    v-if="userInfo == 'admin'"
                    v-model:value="departmentsvalue"
                    :options="departments"
                    class="w-[33%] h-[36px] mr-24"
                    :placeholder="$t('全部')"/>      
                <div class=" mb-2 mr-8 whitespace-nowrap">
                    {{$t('负责人')}}
                </div>
                <n-select
                    v-model:value="principalvalue"
                    :options="principal"
                    class="w-[33%] h-[36px] "
                    :placeholder="$t('全部')"/>
        
                <div class=" mb-2 mr-8 ml-24 whitespace-nowrap">
                    {{$t('时间')}}
                </div>
                <n-date-picker class="w-[33%] h-[36px] " v-model:value="daterange"
                value-format="yyyy.MM.dd HH:mm:ss" type="daterange" clearable size="medium"/>
        
                <n-button :loading="isLoading" type="primary" class=" mb-4 ml-24 rounded px-16" @click="handleClick()">
                    <template #icon>
                        <i class="taskfont" v-if="!(isLoading)">&#xe72a;</i>
                    </template>
                    {{ $t('搜索') }}
                </n-button>
            </div>
            <div class="flex-1 flex items-center mb-16 2xl:justify-end 2xl:mb-0">
                <div class=" mb-4 mr-12 whitespace-nowrap 2xl:ml-24">
                    {{$t('类型')}}
                </div>
                    <n-button-group size="medium" class="mb-4">
                        <n-button  @click="handleClick2('0')" class=" rounded px-16 bg-white whitespace-nowrap">
                            {{$t('全部')}}
                        </n-button>
                        <div class="w-[4px]"></div>
                        <n-button  @click="handleClick2('1')" class=" rounded px-16 bg-white whitespace-nowrap">
                            {{$t('承诺型')}}
                        </n-button>
                        <n-button  @click="handleClick2('2')" class=" rounded px-16 bg-white whitespace-nowrap">
                            {{$t('挑战型')}}
                        </n-button>
                    </n-button-group>
    
        
                <n-checkbox v-model:checked="completednotrated" class="ml-36 rounded whitespace-nowrap mb-2 " @click="getList('search')">
                    {{ $t('已完成未评分') }}
                </n-checkbox>
            </div>
        </div>
        <div class="relative h-full" >
            <div class="absolute top-0 left-0 bottom-0 right-0">
                <n-scrollbar :on-scroll="onScroll">
                    <div class="okr-department-main">
                        <n-spin size="small" :show="loadIng" :class="loadIng && list.length == 0 ? 'mt-[10%]':''">
                        <OkrItems v-if="list.length != 0" @upData="upData" @edit="handleEdit" :list="list"></OkrItems>
                        </n-spin>
                        <OkrNotDatas v-if="!loadIng && list.length == 0" :msg="$t('暂无OKR')" :types="searchObject"></OkrNotDatas>
                    </div>
                </n-scrollbar>
            </div>
        </div>
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

function handleClick() {
    isLoading.value = true
    getList('search');
}

function handleClick2(type) {
    types.value = type
    getList('search');
}

const loadIng = ref(false)
const isLoading = ref(false)
const userInfo = UserStore().info.identity[0]
const page = ref(1)
const last_page = ref(99999)
const list = ref([])
const departmentsvalue = ref(null)
const principalvalue = ref(null)
const departments = ref([
    { label: $t('全部'), value: null },
])
const principal = ref([
    { label: $t('全部'), value: null },
])
const types  = ref('0');
const daterange = ref(null)
const completednotrated = ref(false)

const searchTime = ref(null)
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

const emit = defineEmits(['edit'])

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

    getDepartmentList().then(({data}) => {
        data.data.map(item =>{
            departments.value.push({
                label : item.name,
                value : item.id,
            })
        })
    })

}

const getList = (type) => {
    if (type == 'search'){
        page.value = 1
    }
    if (last_page.value >= page.value || type == 'search') {
        const sendata = {
            completed: completednotrated.value ? 1 : 0,
            department_id : departmentsvalue.value,
            end_at : daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[1] / 1000) : null,
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
            start_at : daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[0] / 1000) : null,
            type: types.value=="0" ? null: types.value,  
            userid: principalvalue.value,
        }
        if (type != 'search'){
            loadIng.value = true
        }
        getDepartmentOkrList(sendata).then(({ data }) => {
            if (type == 'search') {
                data.data ?  list.value = data.data : list.value = []
            }
            else {
                if (data.data) {
                    data.data.map(item => {
                        list.value.push(item)
                    })
                }
            }
            last_page.value = data.last_page
            loadIng.value = false
            isLoading.value = false
        })
    }
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
            getList('')
        }
    }
}


onMounted(() => {
    init()
    getList('')
})
defineExpose({
    upData,
    getList
})
</script>
<style lang="less" scoped>
.okr-department-main {
    @apply pt-24 flex flex-col gap-6 pt-0;
}

</style>
