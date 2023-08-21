<template >
    <n-scrollbar :on-scroll="onScroll">
        <div class="okr-follow-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <OkrItems @upData="upData" @edit="handleEdit" @getList="resetGetList" v-if="list.length != 0 && !loadIng" :list="sortList"></OkrItems>
            <OkrNotDatas  v-if="!loadIng && !onscrolloading && list.length == 0" :types="searchObject !=''"></OkrNotDatas>
            <OkrLoading v-if="onscrolloading" position='onscroll'></OkrLoading>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getFollowList } from '@/api/modules/follow'
import { getOkrDetail } from '@/api/modules/okrList'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from '../components/OkrLoading.vue'


const emit = defineEmits(['edit'])

const loadIng = ref(false)
const onscrolloading = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])

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


const sortList = computed(()=>{
    return list.value.sort((a,b)=>{
        return -1;
    }).sort((a,b)=>{
        if( a.completed > 0 || a.canceled > 0){
            return b.completed - a.completed
        }
        return -1
    })
})

const resetGetList = ()=>{
    page.value = 1
    getList('search')
}

const getList = (type) => {
    let serstatic =  type == 'search' ? true  : false
    if (last_page.value >= page.value || serstatic ) {
        const data = {
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
        }
        if ( serstatic ){
            loadIng.value = true
        }else if ( type == 'onscrollsearch' ){
            onscrolloading.value = true
        }
        loadIng.value = true
        getFollowList(data).then(({ data }) => {
            loadIng.value = false
            if ( serstatic ) {
                list.value = data.data || []
            } else {
                (data.data || []).map(item => {
                    list.value.push(item)
                })
            }
            last_page.value = data.last_page
        })
    }
}

//编辑
const handleEdit = (data) => {
    emit('edit', data)
}

//更新数据
const upData = (id, type) => {
    list.value.map((item, index) => {
        if (item.id == id) {
            const upData = {
                id: id,
            }
            getOkrDetail(upData)
                .then(({ data }) => {
                    list.value[index] = data
                    if (type == 'FollowOkr') {
                        list.value.map((item, index) => {
                            if (item.id == id) {
                                list.value.splice(index, 1)
                            }
                        })
                    }
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
onMounted(() => {
    getList('search')
})
defineExpose({
    upData,
    getList,
    resetGetList
})
</script>
<style lang="less" scoped>
.okr-follow-main {
    @apply pt-16 md:pt-24 flex flex-col gap-6;
}
</style>
