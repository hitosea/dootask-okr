<template >
    <n-scrollbar>
        <div class="okr-follow-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <OkrItems v-else-if="list.length != 0" :list="list" ></OkrItems>
            <OkrNotDatas v-else></OkrNotDatas>
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
            loadIng.value = false
            if (type == 'search') {
                list.value = data.data || []
            }else {
                (data.data || []).map(item => {
                    list.value.push(item)
                })
            }
            last_page.value = data.last_page
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