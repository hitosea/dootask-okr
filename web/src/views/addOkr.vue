<template >
    <div class="nav-top  h-[52px] bg-[#FAFAFA] z-[5]" :style="{ 'z-index': modalTransferIndex }">
        <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
        <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ edit ? $t('编辑OKR') :
            $t('添加OKR') }}</h2>
        <i @click="handleSubmit" class="taskfont text-primary-color mr-4 z-[3] text-20">&#xe684;</i>
    </div>
    <div class="pt-[76px] pb-16 pl-16 pr-2">
        <AddOkrsMain ref="AddOkrsRef" :edit="okrEdit" :editData="okrEditData" :labelPlacement="'top'" :labelAlign="'left'"
            @close="handleReturn" @loadIng="(e) => { loadIng = e }">
        </AddOkrsMain>
    </div>
</template>
<script lang="ts" setup>
import AddOkrsMain from '@/views/components/AddOkrsMain.vue';
import { useRouter } from 'vue-router';
import { GlobalStore } from '@/store';

const router = useRouter()
const loadIng = ref(false)
const edit = ref(false)

const AddOkrsRef = ref(null)
const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1
const globalStore = GlobalStore()
const { okrEditData, okrEdit } = globalStore.okrSetup()



//提交
const handleSubmit = () => {
    AddOkrsRef.value.handleSubmit()
}

const handleReturn = () => {
    globalStore.$patch((state) => {
        state.okrEditData = null
        state.okrEdit = false
    })
    router.back()
}
nextTick(() => {
    AddOkrsRef.value.showDrawer()
})

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
</style>
