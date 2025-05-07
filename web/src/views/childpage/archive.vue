<template >
    <div class="delete-box nav-top h-[52px] bg-[#FAFAFA] z-[5]" :style="{ 'z-index': modalZIndex + 1 }">
        <i @click="handleReturn" class="okrfont icon-return z-[2]">&#xe676;</i>
        <h2 class="absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('已归档') }} OKR</h2>
    </div>
    <div class="relative pt-[76px] pb-16 pl-16 pr-16 h-full" :style="{ 'z-index': modalZIndex }">
        <ArchiveMain></ArchiveMain>
    </div>
</template>
<script lang="ts" setup>
import ArchiveMain from '@/views/components/ArchiveMain.vue';
import { useRouter } from 'vue-router';
import { nextZIndex } from "dootask-tools"


const router = useRouter()
const modalZIndex = nextZIndex()


const handleReturn = () => {
    router.go(-1)
}
</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between absolute top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
</style>
