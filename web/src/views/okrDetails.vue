<template >
    <div class="bg-[#FAFAFA] min-h-full flex relative" :style="{ 'z-index': modalTransferIndex }">
        <div class="nav-top  h-[52px] bg-[#FAFAFA] z-[5]">
            <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
            <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('OKR详情') }}</h2>
            <n-popover placement="bottom-end" trigger="click">
                <template #trigger>
                    <i class="taskfont text-22 mr-4 z-[2]">&#xe6e9;</i>
                </template>
                <div class="flex flex-col">
                    <p class="py-8" @click="handleEdit" v-if="userInfo.userid == userid && detail.canceled == '0' && detail.completed == '0'"> {{ $t('编辑') }}</p>
                    <p class="py-8" @click="handleFollowOkr"> {{ is_follow ? $t('取消收藏') : $t('添加收藏') }}</p>
                    <p class="py-8" @click="handleCancel" v-if="userInfo.userid == userid"> {{ cancel == 0 ? $t('取消目标') :
                        $t('重启目标') }}</p>
                </div>
            </n-popover>
        </div>
        <div class="pt-[52px] pb-16 flex-1 min-h-full">
            <OkrDetailsMain ref="OkrDetailsMainRef" :show="true" :id="id" @isFollow="(e) => { is_follow = e }"
                @canceled="(e) => { cancel = e }" @openDetail="openDetail" @getDetail="getDetail"></OkrDetailsMain>
        </div>
    </div>
</template>
<script lang="ts" setup>
import OkrDetailsMain from '@/views/components/OkrDetailsMain.vue';
import { useRoute, useRouter } from 'vue-router';
import { UserStore } from '@/store/user'

const userInfo = UserStore().info
const route = useRoute()
const router = useRouter()

const id = ref(null)
const userid = ref(null)
const detail = ref<any>({})
const OkrDetailsMainRef = ref(null)
const is_follow = ref(false)
const cancel = ref(0)
const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1

if (route.query.data != undefined) {
    id.value = Number(route.query.data)
    userid.value = Number(route.query.userid)
}


const openDetail = (id, userid) => {
    router.push({
        path: '/okrDetails',
        query: {
            data: id,
            userid: userid,
        },
    })
}

const getDetail = (item) => {
    detail.value = item
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
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
</style>
