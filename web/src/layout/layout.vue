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
import utils from "@/utils/utils"
import AddMultipleDrawer from "@/views/components/AddMultipleDrawer.vue"
import { getAppData } from "@/utils/app"
import { initAppData } from "@/microapp"

const show = ref(getAppData("initialData.type") !== "details")
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

// 监听按键事件
const handleKeydown = (event: any) => {
    if (event.key === "Escape" || ((event.metaKey || event.ctrlKey) && event.key === "w")) {
        if (utils.closeLastModel()) {
            event.preventDefault() // 阻止默认的浏览器行为
        }
    }
}

onMounted(() => {
    cleanupAppData.value = initAppData()
    window.addEventListener("keydown", handleKeydown)
})

onUnmounted(() => {
    if (cleanupAppData.value) cleanupAppData.value()
    window.removeEventListener("keydown", handleKeydown)
})
</script>
<style lang="less" scoped>
.root-layout {
    @apply absolute w-full bottom-0 top-0 right-0 left-0 min-h-full;
}
</style>
