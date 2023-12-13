<template >
    <div class="delete-box nav-top  h-[52px] bg-[#FAFAFA] z-[5]" :style="{ 'z-index': modalTransferIndex + 1 }">
        <i @click="handleReturn" class="okrfont icon-return z-[2]">&#xe704;</i>
        <h2 class="absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('设置')}}</h2>
        <i @click="handleSubmit" class="okrfont text-primary-color mr-4 z-[3] text-20">&#xe684;</i>
    </div>
    <div class="relative pt-[76px] pb-16 pl-16 pr-16 h-full" :style="{ 'z-index': modalTransferIndex }">
        <SettingMain ref="settingMainRef" @loadIng="(e)=>(loadIng = e)"></SettingMain>
    </div>
</template>
<script lang="ts" setup>
import SettingMain from '@/views/components/SettingMain.vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1
const loadIng = ref(false)
const settingMainRef = ref(null)

const handleReturn = () => {
    router.go(-1)
}

const handleSubmit = () => {
    settingMainRef.value.handleSubmit()
}

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
</style>
