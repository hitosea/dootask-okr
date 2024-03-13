<template >
    <div class="relative flex flex-col h-full" :style="{'z-index':modalTransferIndex}">
        <div class="nav-top  h-[52px] bg-[#FAFAFA] z-[5]">
            <i @click="handleReturn" class="okrfont icon-return z-[2]">&#xe676;</i>
            <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ multipleId == 0 ? $t('添加复盘') :$t('复盘详情') }}</h2>
            <i v-if="multipleId == 0" @click="handleSubmit" class="okrfont text-primary-color mr-4 z-[3] text-20">&#xe684;</i>
            <i v-if="canComment" @click="handleComment" class="okrfont text-primary-color mr-4 z-[3] text-20">&#xe684;</i>
        </div>
        <div class="pt-[52px] pb-32 pl-16 pr-16 flex flex-1 bg-[#FFF]">
            <AddMultipleMain ref="AddMultipleMainRef" :data="addMultipleData" :multipleId="multipleId" :superiorUser="superiorUser"
                @loadIng="(e) => { loadIng = e }" @canComment="()=>{canComment = true}" @close="handleReturn"></AddMultipleMain>
        </div>
    </div>
</template>

<script setup lang="ts">
import AddMultipleMain from '@/views/components/AddMultipleMain.vue';
import { useRouter } from 'vue-router';
import { GlobalStore } from '@/store';

const router = useRouter()
const loadIng = ref(false)
const canComment = ref(false)
const AddMultipleMainRef = ref(null)
const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1
const globalStore = GlobalStore()
const {  multipleId,addMultipleData ,superiorUser} = globalStore.multipleSetup()

//提交
const handleSubmit = () => {
    if (loadIng.value) return;
    AddMultipleMainRef.value.handleSubmit()
}
const handleComment = () => {
    if (loadIng.value) return;
    AddMultipleMainRef.value.handleComment()
}

const handleReturn = () => {
    router.back()
}

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
:deep(.n-scrollbar-content) {
    @apply h-full;
}
</style>
