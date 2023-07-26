<template >
    <n-scrollbar>
        <div class="okr-follow-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <div v-else>
                <OkrItems :list="list" v-if="list.length != 0"></OkrItems>
                <OkrNotDatas v-else>
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
                </OkrNotDatas>
            </div>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getFollowList } from '@/api/modules/follow'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"

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
        getFollowList(data).then(({ data }) => {
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
.okr-follow-main {
    @apply py-24 flex flex-col gap-6;
}
</style>