<template >
    <div class="nav-top  h-[52px] bg-[#FAFAFA] z-[5]" :style="{ 'z-index': modalZIndex + 1 }">
        <i @click="handleReturn" class="okrfont icon-return z-[2]">&#xe676;</i>
        <h2 class="absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ okrEdit ? $t('编辑') : $t('添加') }} Objective</h2>
        <i @click="handleSubmit" class="okrfont text-primary-color mr-4 z-[3] text-22">&#xe684;</i>
    </div>
    <div class="relative pt-[76px] pb-32 pl-16 pr-2 bg-[#fff]" :style="{ 'z-index': modalZIndex }">
        <AddOkrsMain ref="AddOkrsRef" :edit="okrEdit" :editData="okrEditData" :labelPlacement="'top'" :labelAlign="'left'"
            @submit="handleReturn" @loadIng="(e) => { loadIng = e }">
        </AddOkrsMain>
    </div>
</template>
<script lang="ts" setup>
import AddOkrsMain from '@/views/components/AddOkrsMain.vue';
import { useRouter } from 'vue-router';
import { GlobalStore } from '@/store';
import { nextModalIndex } from "@/utils/app"

const router = useRouter()
const loadIng = ref(false)

const AddOkrsRef = ref(null)
const modalZIndex = nextModalIndex()
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
      router.go(-1)
}

nextTick(() => {
    AddOkrsRef.value.showDrawer()
})

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between absolute top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
</style>
