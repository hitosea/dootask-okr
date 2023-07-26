<template >
    <n-scrollbar>
        <div class="i-created-main">
            <PersonalStatistics></PersonalStatistics>
            <okrItem :list="list" v-if="list.length != 0"></okrItem>
        </div>
        <okrNotData v-if="list.length == 0">
            <template v-slot:content>
                <div class="mt-5">
                    <div class="mb-10">{{$t('暂无OKR')}}</div>     
                    <div>
                        <n-button type="primary" ghost>
                            <i class="taskfont mr-5">&#xe731;</i>
                            {{ $t('创建OKR') }}
                        </n-button>
                    </div>    

                </div>
            </template>
        </okrNotData>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import PersonalStatistics from '@/views/components/PersonalStatistics.vue'
import okrItem from '@/views/components/okrItem.vue'
import { getMyList } from '@/api/modules/orkList'
import okrNotData from "@/views/components/okrNotData.vue"

const loadIng = ref(false)
const page = ref(1)
const last_page = ref(99999)
const list = ref([])


const getList = (type) => {
    if (last_page.value >= page.value || type == 'search') {
        const data = {
            page: page.value,
            page_size: 10,
        }
        loadIng.value = true
        getMyList(data).then(({ data }) => {
            console.log(data);
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

onMounted(() => {
    getList('')
})
</script>
<style lang="less" scoped>
.i-created-main {
    @apply py-24 flex flex-col gap-6;
}
</style>