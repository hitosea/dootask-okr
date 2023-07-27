<template >
    <n-scrollbar>
        <div class="i-created-main">
            <PersonalStatistics></PersonalStatistics>
            <div>
                <OkrItems :list="list" @upData="" @edit="handleEdit" v-if="list.length != 0"></OkrItems>
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
            <OkrLoading v-if="loadIng"></OkrLoading>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import PersonalStatistics from '@/views/components/PersonalStatistics.vue'
import OkrItems from '@/views/components/OkrItems.vue'
import { getMyList } from '@/api/modules/okrList'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"

const loadIng = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])

const emit = defineEmits(['edit'])

const getList = (type) => {
    if (last_page.value >= page.value || type == 'search') {
        const data = {
            page: page.value,
            page_size: 10,
        }
        loadIng.value = true
        getMyList(data).then(({ data }) => {
            if (type == 'search') {
                list.value = data.data
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
        })
    }
}

//编辑
const handleEdit = (data) => {
    emit('edit',data)
}

onMounted(() => {
    getList('')
})


</script>
<style lang="less" scoped>
.i-created-main {
    @apply relative py-24 flex flex-col gap-6;
}
</style>
