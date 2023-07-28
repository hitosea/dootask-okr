<template >
    <n-scrollbar>
        <div class="okr-participant-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <div v-else>
                <OkrItems :list="list" v-if="list.length != 0"></OkrItems>
                <OkrNotDatas v-else></OkrNotDatas>
            </div>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getParticipantList } from '@/api/modules/participant'
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
        getParticipantList(data).then(({ data }) => {
            if (type == 'search') {
                data.data ? list.value = data.data : []
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
.okr-participant-main {
    @apply py-24 flex flex-col gap-6;
}
</style>