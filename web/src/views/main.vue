<template>
    <OkrDetails ref="RefOkrDetails" 
        :id="globalStore.okrDetail.id" 
        :show="okrDetailsShow" 
        @close="close" 
    ></OkrDetails>
</template>

<script lang="ts" setup>
import {watch} from 'vue';
import OkrDetails from './components/OkrDetails.vue';
import { GlobalStore } from "@/store"

const okrDetailsShow = ref(false)
const globalStore = GlobalStore()

// 监听打开
watch(() => globalStore.okrDetail, (newValue) => {
    if(newValue.show){
        okrDetailsShow.value = false
        nextTick(() => {
            okrDetailsShow.value = true;
        })
    }
}, { immediate: true })

// 关闭
const close = () => {
    okrDetailsShow.value = false
    globalStore.okrDetail.id = 0
    globalStore.okrDetail.show = false
}
</script>
