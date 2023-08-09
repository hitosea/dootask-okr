<template >
    <div class="bg-[#FAFAFA] min-h-full flex">
        <div class="nav-top  h-[44px] bg-[#FAFAFA] z-[5]">
            <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
            <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('OKR详情') }}</h2>
            <n-popover placement="bottom-end" trigger="click">
                <template #trigger>
                    <n-icon class="text-text-tips z-[2]" size="20" :component="EllipsisHorizontalSharp" />
                </template>
                <div class="flex flex-col">
                    <p class="py-8" @click="handleEdit"> {{ $t('编辑') }}</p>
                    <p class="py-8" @click="handleFollowOkr"> {{ is_follow ?   $t('取消收藏') : $t('添加收藏')}}</p>
                    <p class="py-8" @click="handleCancel"> {{ cancel == 0 ? $t('取消目标') : $t('重启目标') }}</p>

                </div>
            </n-popover>
        </div>
        <div class="pt-[58px] pb-16 pl-16 pr-16 flex-1 min-h-full">
            <OkrDetailsMain ref="OkrDetailsMainRef" :show="true" :id="id" @isFollow="(e) => { is_follow = e }"
                @canceled="(e) => { cancel = e }"></OkrDetailsMain>
        </div>
    </div>
</template>
<script lang="ts" setup>
import OkrDetailsMain from '@/views/components/OkrDetailsMain.vue';
import { EllipsisHorizontalSharp } from '@vicons/ionicons5'
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()

const id = ref(null)
const OkrDetailsMainRef = ref(null)
const is_follow = ref(false)
const cancel = ref(0)



if (route.query.data != undefined) {
    id.value = Number(route.query.data)
}


const handleReturn = () => {
    router.back()
}

const handleEdit = () => {
    OkrDetailsMainRef.value.handleEdit()
}

const handleCancel = () => {
    OkrDetailsMainRef.value.handleCancel()
}

const handleFollowOkr = () => {
    OkrDetailsMainRef.value.handleFollowOkr()
}



</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-16;

    .icon-return {
        @apply text-20 text-text-tips;
    }
}
</style>