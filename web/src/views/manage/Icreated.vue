<template >
    <n-scrollbar :on-scroll="onScroll" ref="scrollbarRef">
        <div class="i-created-main">
            <PersonalStatistics ref="PersonalStatisticsRef" v-if="props.searchObject == ''"></PersonalStatistics>
            <div>
                <OkrLoading v-if="loadIng"></OkrLoading>
                <OkrItems :list="list" @upData="upData" @edit="handleEdit" @getList="resetGetList"
                    v-if="list.length != 0 && !loadIng"></OkrItems>
                <OkrNotDatas v-if="!loadIng && !onscrolloading && list.length == 0" :types="props.searchObject != ''">
                    <template v-slot:content v-if="props.searchObject == '' && !loadIng">
                        <div class="mt-5"><s></s>
                            <div>
                                <n-button type="primary" ghost @click="handleAdd" :loading="btnLoading">
                                    <i class="okrfont mr-5">&#xe731;</i>
                                    {{ $t('添加') }} Objective
                                </n-button>
                            </div>
                        </div>
                    </template>
                </OkrNotDatas>
                <OkrLoading v-if="onscrolloading" position='onscroll'></OkrLoading>
            </div>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import PersonalStatistics from '@/views/components/PersonalStatistics.vue'
import OkrItems from '@/views/components/OkrItems.vue'
import { getMyList, getOkrDetail } from '@/api/modules/okrList'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from '../components/OkrLoading.vue'
import utils from '@/utils/utils'

const loadIng = ref(false)
const onscrolloading = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])
const scrollbarRef = ref(null)
const PersonalStatisticsRef = ref(null)

const searchTime = ref(null)
const props = defineProps({
    btnLoading: {
        type: Boolean,
        default: false,
    },
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

const emit = defineEmits(['edit', 'add'])

const resetGetList = () => {
    page.value = 1
    getList('search')
}

const getList = (type) => {
    let serstatic = type == 'search' ? true : false
    if (last_page.value >= page.value || serstatic  || type == 'updata') {
        const data = {
            objective: props.searchObject,
            page: type == 'updata' ? 1 : page.value,
            page_size: type == 'updata' ?  page.value * 20 : 20,
        }
        if (serstatic) {
            loadIng.value = true
        } else if (type == 'onscrollsearch') {
            onscrolloading.value = true
        }
        getMyList(data).then(({ data }) => {
            onscrolloading.value = false
            loadIng.value = false
            if (serstatic  || type == 'updata') {
                list.value = data.data || []
            }
            else {
                (data.data || []).map(item => {
                    list.value.push(item)
                })
            }
            if (props.searchObject == '') {
                PersonalStatisticsRef.value?.getData()
            }
            last_page.value = data.last_page
        })
    }
}

//编辑
const handleAdd = () => {
    emit('add')
}
//编辑
const handleEdit = (data) => {
    emit('edit', data)
}

//更新数据
const upData = (id) => {
    list.value.map((item, index) => {
        if (item.id == id) {
            getOkrDetail({id}).then(({ data }) => {
                list.value[index] = data
                if (props.searchObject == '') {
                    PersonalStatisticsRef.value?.getData()
                }
                list.value = utils.listSort(list.value)
            })
        }
    })
}

const onScroll = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight >= e.target.scrollHeight) {
        // 重新请求数据
        if (!onscrolloading.value && !loadIng.value) {
            if (last_page.value > page.value) {
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
.i-created-main {
    @apply relative pt-16 md:pt-24 flex flex-col gap-3 md:gap-4;
}
</style>
