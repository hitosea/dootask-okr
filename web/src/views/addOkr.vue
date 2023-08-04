<template >
    <div class="nav-top  h-[44px] bg-[#FAFAFA] z-[5]">
        <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
        <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('OKR管理') }}</h2>
        <n-icon @click="handleSubmit" class="text-primary-color z-[2]" size="20" :component="Checkmark" />
    </div>
    <div class="pt-[68px] pb-16 pl-16 pr-2">
        <AddOkrs ref="AddOkrsRef" :edit="edit" :editData="editData" :labelPlacement="'top'" :labelAlign="'left'" @close="handleReturn" @loadIng="(e) => { loadIng = e }">
        </AddOkrs>
    </div>
</template>
<script lang="ts" setup>
import AddOkrs from '@/views/components/AddOkrs.vue';
import { Checkmark } from '@vicons/ionicons5'
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()
const loadIng = ref(false)
const edit = ref(false)
const editData = ref({})
const AddOkrsRef = ref(null)

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
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-16;

    .icon-return {
        @apply text-20 text-text-tips;
    }
}
</style>