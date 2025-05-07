<template>
    <n-layout class="root-layout" v-show="show">
        <router-view />

        <!-- 新增复盘 -->
        <AddMultipleDrawer
            v-model:show="addMultipleShow"
            :data="addMultipleData"
            :multipleId="multipleId"
            @close="handleCloseMultiple"
        />
    </n-layout>
</template>

<script lang="ts" setup>
import { onMounted } from "vue"
import { GlobalStore } from "@/store"
import AddMultipleDrawer from "@/views/components/AddMultipleDrawer.vue"
import { getAppData } from "@/utils/app"
import { initAppData } from "@/microapp"

const show = ref(getAppData('props.type') !== "details")
const globalStore = GlobalStore()
const cleanupAppData = ref(null)
const { addMultipleShow, multipleId, addMultipleData } = globalStore.multipleSetup()

// 关闭复盘
const handleCloseMultiple = () => {
    globalStore.$patch((state) => {
        state.addMultipleData = null
        state.multipleId = 0
        state.addMultipleShow = false
    })
}

onMounted(() => {
    cleanupAppData.value = initAppData()
})

onUnmounted(() => {
    if (cleanupAppData.value) cleanupAppData.value()
})
</script>
<style lang="less" scoped>
.root-layout {
    @apply absolute w-full bottom-0 top-0 right-0 left-0 min-h-full;
}
</style>
