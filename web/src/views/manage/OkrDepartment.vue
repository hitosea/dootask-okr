<template >
    <div class=" flex items-center my-16 nowrap">
        <div class=" mb-2 mr-8  custom-div">
            {{$t('部门')}}
        </div>
        <n-select
            v-model:value="departmentsvalue"
            :options="departments"
            class="custom-select"
            :placeholder="$t('全部')"/>

        <div class=" mb-2 mr-8 ml-24 custom-div">
            {{$t('负责人')}}
        </div>
        <n-select
            v-model:value="principalvalue"
            :options="principal"
            class="custom-select"
            :placeholder="$t('全部')"/>

        <div class=" mb-2 mr-8 ml-24 custom-div">
            {{$t('类型')}}
        </div>
        <n-select
            v-model:value="typevalue"
            :options="types"
            class="custom-select"
            :placeholder="$t('全部')"/>

        <div class=" mb-2 mr-8 ml-24 custom-div">
            {{$t('时间')}}
        </div>
        <n-date-picker class="custom-select" v-model:value="daterange"
        value-format="yyyy.MM.dd HH:mm:ss" type="daterange" clearable size="medium" :placeholder="$t('请选择时间')"/>

        <n-button :loading="isLoading" type="primary" class=" mb-4 ml-24 rounded" @click="handleClick()">
            <i class="taskfont mr-5" v-if="!(isLoading)">&#xe72a;</i>
            {{ $t('搜索') }}
        </n-button>

        <n-checkbox v-model:checked="completednotrated" class="ml-auto rounded custom-div" @click="getList('search')">
            {{ $t('已完成未评分') }}
        </n-checkbox>
    </div>
    <n-scrollbar :on-scroll="onScroll">
        <div class="okr-department-main">
<<<<<<< Updated upstream
            <OkrLoading v-if="loadIng"></OkrLoading>
            <div class=" flex items-center mb-16 mt-0 nowrap">

                <div class=" mb-2 mr-8  custom-div">
                    {{$t('部门')}}
                </div>
                <n-select
                    v-model:value="departmentsvalue"
                    :options="departments"
                    class="custom-select"
                    :placeholder="$t('全部')"/>

                <div class=" mb-2 mr-8 ml-24 custom-div">
                    {{$t('负责人')}}
                </div>
                <n-select
                    v-model:value="principalvalue"
                    :options="principal"
                    class="custom-select"
                    :placeholder="$t('全部')"/>

                <div class=" mb-2 mr-8 ml-24 custom-div">
                    {{$t('类型')}}
                </div>
                <n-select
                    v-model:value="typevalue"
                    :options="types"
                    class="custom-select"
                    :placeholder="$t('全部')"/>

                <div class=" mb-2 mr-8 ml-24 custom-div">
                    {{$t('时间')}}
                </div>
                <n-date-picker class="custom-select" v-model:value="daterange"
                value-format="yyyy.MM.dd HH:mm:ss" type="daterange" clearable size="medium" :placeholder="$t('请选择时间')"/>

                <n-button type="primary" class=" mb-4 ml-24 rounded" @click="getList('search')">
                    <i class="taskfont mr-5">&#xe72a;</i>
                    {{ $t('搜索') }}
                </n-button>
                <n-checkbox v-model:checked="completednotrated" class="ml-auto rounded custom-div ml-10" @click="getList('search')">
                    {{ $t('已完成未评分') }}
                </n-checkbox>
            </div>
            <OkrItems @upData="upData" @edit="handleEdit" :list="list" v-if="list.length != 0"></OkrItems>
=======
            <OkrItems v-if="list.length != 0" @upData="upData" @edit="handleEdit" :list="list"></OkrItems>
>>>>>>> Stashed changes
            <OkrNotDatas v-else :msg="$t('暂无OKR')"></OkrNotDatas>
        </div>
    </n-scrollbar>
    <div class="mask" v-if="loadIng">
        <OkrLoading></OkrLoading>
    </div>
    <div class="maskbottom" v-if="loadIngBottom">
        <OkrLoading></OkrLoading>
    </div>
</template>

<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getDepartmentOkrList } from '@/api/modules/department'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"
import { getDepartmentList } from '@/api/modules/department'
import { getOkrDetail } from '@/api/modules/okrList'
import { getUserList } from '@/api/modules/created'
import utils from '@/utils/utils'

function handleClick() {
    isLoading.value = true
    getList('search');
}

const loadIng = ref(false)
const loadIngBottom = ref(false)
const isLoading = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])
const departmentsvalue = ref(null)
const principalvalue = ref(null)
const typevalue = ref(null)
const departments = ref([
    { label: $t('全部'), value: null },
])
const principal = ref([
    { label: $t('全部'), value: null },
])
const types = ref([
    { label: $t('全部'), value: null },
    { label: $t('承诺型'), value: "1" },
    { label: $t('挑战型'), value: "2" },
]);
const daterange = ref(null)
const completednotrated = ref(false)

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
            objective : null,
            page: page.value,
            page_size: 10,
            start_at : daterange.value ? utils.formatDate('Y-m-d 00:00:00', daterange.value[0] / 1000) : null,
            type: typevalue.value,
            userid: principalvalue.value,
        }
        if (type != 'search'){
            loadIngBottom.value = true
        }else{
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
            loadIngBottom.value = false         

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
    getList('search')
})

</script>
<style lang="less" scoped>
.okr-department-main {
    @apply py-24 flex flex-col gap-6 pt-0;
}

.custom-select {
    width: 160px;
    height: 36px;
}
.custom-div{
    white-space: nowrap;

}

.mask {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 70vh;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999; 
  }
  .maskbottom {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 90vh;
    display: flex;
    justify-content: center;
    align-items: flex-end;
    z-index: 9999; 
  }

</style>
