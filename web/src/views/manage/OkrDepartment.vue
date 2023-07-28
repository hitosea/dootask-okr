<template >
    <n-scrollbar>
        <div class="okr-department-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <div v-else>
                <div class=" flex items-center mb-16 mt-0">

                    <div class=" mb-2 mr-8">
                        {{$t('部门')}}
                    </div>
                    <n-select 
                        v-model:value="departmentsvalue" 
                        :options="departments" 
                        class="custom-select" 
                        placeholder="全部"/>

                    <div class=" mb-2 mr-8 ml-24">
                        {{$t('负责人')}}
                    </div>
                    <n-select 
                        v-model:value="principalvalue" 
                        :options="principal" 
                        class="custom-select" 
                        placeholder="全部"/>

                    <div class=" mb-2 mr-8 ml-24">
                        {{$t('类型')}}
                    </div>
                    <n-select 
                        v-model:value="typevalue" 
                        :options="types" 
                        class="custom-select" 
                        placeholder="全部"/>

                    <div class=" mb-2 mr-8 ml-24">
                        {{$t('时间')}}
                    </div>
                    <n-date-picker class="custom-select" v-model:value="daterange"
                    value-format="yyyy.MM.dd HH:mm:ss" type="daterange" clearable size="medium" placeholder="请选择时间"/>
  
                    <n-button type="primary" class=" mb-4 ml-24 rounded" @click="getList('search')">
                        <i class="taskfont mr-5">&#xe72a;</i>
                        {{ $t('搜索') }}  
                    </n-button>
                    <n-checkbox v-model:checked="completednotrated" class="ml-auto rounded" @click="getList('search')">
                        {{ $t('已完成未评分') }}    
                    </n-checkbox>             
                </div>
                <OkrItems :list="list" v-if="list.length != 0"></OkrItems>
                <OkrNotDatas v-else :msg="$t('暂无OKR')">
                    <template v-slot:content>
                        <div class="mt-5">
                            <div class="mb-10"></div>     
                            <div>
                                <n-button type="primary" ghost>
                                    <i class="taskfont mr-5">&#xe731;</i>
                                    {{ $t('创建OKR') }}
                                </n-button>
                            </div>    
                        </div>
                    </template>
                </OkrNotDatas>
            </div>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getDepartmentOkrList } from '@/api/modules/department' 
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"
import { getDepartmentList } from '@/api/modules/department'
import { getUserList } from '@/api/modules/created'
import utils from '@/utils/utils'

const loadIng = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])
const departmentsvalue = ref(null)
const principalvalue = ref(null)
const typevalue = ref(null)
const departments = ref([
    { label: "全部", value: null },
])
const principal = ref([
    { label: "全部", value: null },
])
const types = ref([
    { label: "全部", value: null },
    { label: "承诺型", value: "1" },
    { label: "挑战型", value: "2" },
]);
const daterange = ref(null)
const completednotrated = ref(false)

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
    
    if (last_page.value >= page.value || type == 'search') {
        const sendata = {
            completed: completednotrated.value ? 1 : 0,
            department_id : departmentsvalue.value,
            end_at : daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[1] / 1000) : null,
            objective : null,
            page: page.value,
            page_size: 10,
            start_at : daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[0] / 1000) : null,
            type: typevalue.value,
            userid: principalvalue.value,
        }   

        loadIng.value = true
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
        })
        loadIng.value = false
    }
}

onMounted(() => {
    init()
    getList('')
})

</script>
<style lang="less" scoped>
.okr-department-main {
    @apply py-24 flex flex-col gap-6;
}

.custom-select {
    width: 160px;
    height: 36px;
}


</style>