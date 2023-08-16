<template>
    <n-layout class="root-layout">
        <router-view />
    <!-- 新增复盘 -->
    <AddMultipleDrawer v-model:show="addMultipleShow" :data="addMultipleData" :multipleId="multipleId" @close="handleCloseMultiple"></AddMultipleDrawer>
    </n-layout>
</template>

<script lang="ts" setup>
import { onMounted, watch } from 'vue'
import { useLoadingBar } from 'naive-ui'
import { loadingBarApiRef } from "../routes";
import { UserStore } from "../store/user";
import { GlobalStore } from '@/store';
import AddMultipleDrawer from '@/views/components/AddMultipleDrawer.vue';

const userStore = UserStore()
const loadingBar = useLoadingBar()
const globalStore = GlobalStore()
const { addMultipleShow, multipleId,addMultipleData} = globalStore.multipleSetup()


//关闭复盘
const handleCloseMultiple = () => {
    globalStore.$patch((state) => {
        state.addMultipleData = null
        state.multipleId = 0
        state.addMultipleShow = false
    })
}

watch(
    () => userStore.info,
    () => { },
    { immediate: true }
)


onMounted(() => {
    loadingBarApiRef.value = loadingBar
    loadingBar.finish()
    window.addEventListener('scroll', windowScrollListener);

})

const windowScrollListener = () => {
    globalStore.$patch((state) => {
        state.windowScrollY = window.scrollY
    })
}

const otherEvents = () => {
    // 非客户端监听窗口激活
    const hiddenProperty = 'hidden' in document ? 'hidden' : 'webkitHidden' in document ? 'webkitHidden' : 'mozHidden' in document ? 'mozHidden' : '';
    const visibilityChangeEvent = hiddenProperty.replace(/hidden/i, 'visibilitychange');
    document.addEventListener(visibilityChangeEvent, () => {
        globalStore.$patch((state) => {
            state.windowActive = !document[hiddenProperty]
        })
    });
}


otherEvents()
</script>
<style lang="less" scoped>
.root-layout {
    @apply absolute w-full bottom-0 top-0 right-0 left-0 min-h-full;


}
</style>
