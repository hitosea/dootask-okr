<template >
    <div class="nav-top  h-[52px] bg-[#FAFAFA] z-[5]" :style="{'z-index':modalTransferIndex}">
        <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
        <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('OKR管理') }}</h2>
        <n-icon @click="handleSubmit" class="text-primary-color mr-4 z-[3]" size="20" :component="Checkmark" />
    </div>
    <div class="pt-[76px] pb-16 pl-16 pr-2">
        <AddOkrsMain ref="AddOkrsRef" :edit="edit" :editData="editData" :labelPlacement="'top'" :labelAlign="'left'" @close="handleReturn" @loadIng="(e) => { loadIng = e }">
        </AddOkrsMain>
    </div>
</template>
<script lang="ts" setup>
import AddOkrsMain from '@/views/components/AddOkrsMain.vue';
import { Checkmark } from '@vicons/ionicons5'
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()
const loadIng = ref(false)
const edit = ref(false)
const editData = ref({})
const AddOkrsRef = ref(null)
const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1

if (route.query.data != undefined) {
    edit.value = true
    editData.value = JSON.parse(route.query.data as string)
}


//提交
const handleSubmit = () => {
    AddOkrsRef.value.handleSubmit()
}

const handleReturn = () => {
    router.back()
}
nextTick(()=>{
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
