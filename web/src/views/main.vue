<template>
    <h1 class="text-center mt-[12%]">欢迎使用okr子系统</h1>
    <div class="text-center mt-[20px]">
        <router-link to="/list" class="mr-[10px]">前往Okr列表</router-link> |
        <router-link to="/analysis" class="ml-[10px]">前往Okr汇总</router-link>
    </div>
    <OkrDetailsModal ref="RefOkrDetails" :id="globalStore.okrDetail.id" :show="okrDetailsShow" @close="close">
    </OkrDetailsModal>
</template>

<script lang="ts" setup>
import { watch } from 'vue';
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import { GlobalStore } from "@/store"
import { useRouter } from 'vue-router'

const router = useRouter()
const okrDetailsShow = ref(false)
const globalStore = GlobalStore()

// 监听打开
watch(() => globalStore.okrDetail, (newValue) => {
    if (newValue.show) {
        if (window.innerWidth < 768) {
            router.push({
                path: '/okrDetails',
                query: { data: globalStore.okrDetail.id },
            })
        }
        else {
            okrDetailsShow.value = false
            nextTick(() => {
                okrDetailsShow.value = true;
            })
        }
    }
}, { immediate: true })

// 关闭
const close = () => {
    okrDetailsShow.value = false
    globalStore.okrDetail.id = 0
    globalStore.okrDetail.show = false
}
</script>
