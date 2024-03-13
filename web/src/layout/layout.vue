<template>
    <n-layout class="root-layout">
        <router-view />
        <!-- 新增复盘 -->
        <AddMultipleDrawer v-model:show="addMultipleShow" :data="addMultipleData" :multipleId="multipleId" :superiorUser="superiorUser" @close="handleCloseMultiple"></AddMultipleDrawer>
    </n-layout>
</template>

<script lang="ts" setup>
import { onMounted, watch } from 'vue'
import { useLoadingBar } from 'naive-ui'
import { loadingBarApiRef } from "../routes";
import { UserStore } from "../store/user";
import { GlobalStore } from '@/store';
import utils from '@/utils/utils';
import AddMultipleDrawer from '@/views/components/AddMultipleDrawer.vue';

const userStore = UserStore()
const loadingBar = useLoadingBar()
const globalStore = GlobalStore()
const { addMultipleShow, multipleId,addMultipleData ,superiorUser} = globalStore.multipleSetup()

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
    window.addEventListener('keydown', handleKeydown);
})

onUnmounted(() => {
    window.removeEventListener('scroll', windowScrollListener);
    window.removeEventListener('keydown', handleKeydown);
})

const windowScrollListener = () => {
    globalStore.$patch((state) => {
        state.windowScrollY = window.scrollY
    })
}

const handleKeydown = (event) => {
    if (event.key === 'Escape' || ((event.metaKey || event.ctrlKey) && event.key === "w") ) {
        if(utils.closeLastModel()){
            event.preventDefault(); // 阻止默认的浏览器行为
        }
    }
}

</script>
<style lang="less" scoped>
.root-layout {
    @apply absolute w-full bottom-0 top-0 right-0 left-0 min-h-full;


}
</style>
