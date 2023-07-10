<template>
    <n-layout class="root-layout">
        <router-view />
    </n-layout>
</template>

<script lang="ts" setup>
import { onMounted, watch } from 'vue'
import { useLoadingBar } from 'naive-ui'
import { loadingBarApiRef } from "../routes";
import { UserStore } from "../store/user";
import { GlobalStore } from '@/store';

const userStore = UserStore()
const loadingBar = useLoadingBar()
const globalStore = GlobalStore()

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
<style lang="less">
.root-layout {
    @apply absolute w-full bottom-0 top-0 right-0 left-0 min-h-full;

    .child-view {
        @apply absolute w-full bottom-0 top-0 right-0 left-0 min-h-full;
    }
}
</style>