<template >
    <n-scrollbar :on-scroll="onScroll" ref="scrollbarRef">
        <div class="okr-participant-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <OkrItems @upData="upData" @edit="handleEdit" @getList="resetGetList" v-if="list.length != 0 && !loadIng" :list="list"></OkrItems>
            <OkrNotDatas v-if="!loadIng && !onscrolloading && list.length == 0" :loadIng="loadIng" :types="searchObject !=''"></OkrNotDatas>
            <OkrLoading v-if="onscrolloading" position='onscroll'></OkrLoading>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getParticipantList  } from '@/api/modules/participant'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from '../components/OkrLoading.vue'

const emit = defineEmits(['edit'])

const loadIng = ref(false)
const onscrolloading = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])
const scrollbarRef = ref(null)

const searchTime = ref(null)
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

const resetGetList = ()=>{
    page.value = 1
    getList('search')
}

const getList = (type) => {
    let serstatic =  type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic || type == 'upData') {
        const data = {
            objective: props.searchObject,
            page: type == 'upData' ? 1 : page.value,
            page_size: type == 'upData' ?  page.value * 20 : 20,
        }
        if (serstatic){
            loadIng.value = true
        }else if ( type == 'onscrollsearch' ){
            onscrolloading.value = true
        }
        getParticipantList(data).then(({ data }) => {
            loadIng.value = false
            if ( serstatic || type == 'upData' ) {
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
const upData = (id) => {
    getList('upData')
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
.okr-participant-main {
    @apply pt-16 md:pt-24 flex flex-col gap-6;
}
</style>
