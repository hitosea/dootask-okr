<template >
    <n-scrollbar :on-scroll="onScroll">
        <div class="i-created-main">
            <PersonalStatistics></PersonalStatistics>
            <div>
                <OkrLoading v-if="loadIng"></OkrLoading>
                <OkrItems :list="list" @upData="upData" @edit="handleEdit" v-else-if="list.length != 0"></OkrItems>
                <OkrNotDatas v-else>
                    <template v-slot:content>
                        <div class="mt-5">
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
import PersonalStatistics from '@/views/components/PersonalStatistics.vue'
import OkrItems from '@/views/components/OkrItems.vue'
import { getMyList, getOkrDetail } from '@/api/modules/okrList'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"

const loadIng = ref(false)
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

const emit = defineEmits(['edit'])

const getList = (type) => {
    if (last_page.value >= page.value || type == 'search') {
        const data = {
            objective: props.searchObject,
            page: page.value,
            page_size: 10,
        }
        loadIng.value = true
        getMyList(data).then(({ data }) => {
            loadIng.value = false
            if (type == 'search') {
                list.value = data.data || []
            }
            else {
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
        if (!loadIng.value) {
            page.value++
            getList('')
        }
    }
}

onMounted(() => {
    getList('')
})
defineExpose({
    upData
})
</script>
<style lang="less" scoped>
.i-created-main {
    @apply relative py-24 flex flex-col gap-6;
}
</style>
