<template >
    <n-scrollbar :on-scroll="onScroll" ref="scrollbarRef">
        <div class="okr-participant-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <OkrItems @upData="upData" @edit="handleEdit" v-if="list.length != 0 && !loadIng" :list="list"></OkrItems>
            <OkrNotDatas v-if="!loadIng && list.length == 0" :types="searchObject"></OkrNotDatas>
            <OkrLoading v-if="onscrolloading" position='onscroll'></OkrLoading>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getParticipantList  } from '@/api/modules/participant'
import {  getOkrDetail } from '@/api/modules/okrList'
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
    searchTime.value = setInterval(() => {
        page.value = 1
        getList('search')
        clearInterval(searchTime.value)
    }, 300)
}, { deep: true })

const getList = (type) => {
    let serstatic =  type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic) {
        const data = {
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
        }
        if (serstatic){
            loadIng.value = true
        }else if ( type == 'onscrollsearch' ){
            onscrolloading.value = true
        }
        getParticipantList(data).then(({ data }) => {
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
        if (!onscrolloading.value) {
            page.value++
            getList('onscrollsearch')
        }
    }
}


onMounted(() => {
    getList('search')
})

defineExpose({
    upData,
    getList
})
</script>
<style lang="less" scoped>
.okr-participant-main {
    @apply pt-24 flex flex-col gap-6;
}
</style>
