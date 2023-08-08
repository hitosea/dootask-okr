<template >
    <div class="flex flex-col h-full">
        <div class="nav-top  h-[44px] bg-[#FAFAFA] z-[5]">
            <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
            <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ route.query.id == undefined ? $t('添加复盘') :$t('复盘详情') }}</h2>
            <n-icon v-if="route.query.id == undefined" @click="handleSubmit" class="text-primary-color z-[3]" size="20" :component="Checkmark" />
        </div>
        <div class="pt-[68px] pb-16 pl-16 pr-16 flex flex-1">
            <AddMultipleMain ref="AddMultipleMainRef" :data="addMultipleData" :multipleId="multipleId"
                @loadIng="(e) => { loadIng = e }" @close="handleReturn"></AddMultipleMain>
        </div>
    </div>
</template>
<script setup lang="ts">
import AddMultipleMain from './components/AddMultipleMain.vue';
import { Checkmark } from '@vicons/ionicons5'
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()
const loadIng = ref(false)
const multipleId = ref(0)
const addMultipleData = ref({})
const AddMultipleMainRef = ref(null)


if (route.query.data != undefined) {
    addMultipleData.value = JSON.parse(route.query.data as string)
}

if (route.query.id !=undefined){
    multipleId.value = Number(route.query.id)
}

//提交
const handleSubmit = () => {
    if (loadIng.value) return;
    AddMultipleMainRef.value.handleSubmit()
}

const handleReturn = () => {
    router.back()
}

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-16;

    .icon-return {
        @apply text-20 text-text-tips;
    }
}
:deep(.n-scrollbar-content) {
    @apply h-full;
}
</style>