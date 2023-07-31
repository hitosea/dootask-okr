<template >
    <n-scrollbar :on-scroll="onScroll">
        <div class="okr-follow-main">
            <OkrLoading v-if="loadIng"></OkrLoading>
            <OkrItems @upData="upData" @edit="handleEdit" v-else-if="list.length != 0" :list="list"></OkrItems>
            <OkrNotDatas v-else></OkrNotDatas>
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import OkrItems from '@/views/components/OkrItems.vue'
import { getFollowList } from '@/api/modules/follow'
import { getOkrDetail } from '@/api/modules/okrList'
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import OkrLoading from "@/views/components/OkrLoading.vue"

const emit = defineEmits(['edit'])

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
        if (!loadIng.value) {
            page.value++
            getList('')
        }
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
