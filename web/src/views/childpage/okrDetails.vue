<template >
    <div ref="pageOkrDetailRef" :class="isSingle ? ['bg-white'] : ['bg-[#FAFAFA]']" class="page-okr-details  min-h-full flex relative" :style="{ 'z-index': modalTransferIndex }">
        <div v-if="!isSingle" class="nav-top  h-[52px] bg-[#FAFAFA] z-[5]">
            <i @click="handleReturn" class="okrfont icon-return z-[2]">&#xe676;</i>
            <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">OKR {{ $t('详情') }}</h2>
            <n-popover placement="bottom-end" :show="showPopover" :z-index="modalTransferIndex" @clickoutside="showPopover = false">
                <template #trigger>
                    <i @click="showPopover = !showPopover" class="okrfont text-22 mr-4 z-[2]">&#xe6e9;</i>
                </template>
                <div class="flex flex-col">
                    <p class="py-8" @click="handleEdit" v-if="userInfo.userid == detail.userid && detail.canceled == '0' && detail.completed == '0'"> {{ $t('编辑') }}</p>
                    <p class="py-8" @click="handleFollowOkr"> {{ is_follow ? $t('取消关注') : $t('关注目标') }}</p>
                    <p class="py-8" @click="handleCancel" v-if="userInfo.userid == detail.userid && detail.completed == '0'"> {{ cancel == 0 ? $t('取消目标') : $t('重启目标') }}</p>
                    <p class="py-8" @click="handleWarningShow(2)" > {{  $t('归档') }}</p>
                    <p class="py-8" @click="handleWarningShow(1)" > {{ $t('删除') }}</p>
                </div>
            </n-popover>
        </div>
        <div v-if="detailsShow" :class="isSingle ? ['py-16','pl-24','h-[100vh]'] : ['pt-[52px]', 'pb-16', 'min-h-full']" class="flex-1 overflow-hidden">
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
const isSingle = computed(() => document.querySelector('.electron-single-micro-apps') && route.name == 'okrDetails' ? 1 : 0  )

const id = ref(null)
const userid = ref(null)
const detail = ref<any>({})
const detailsShow = ref(false)
const pageOkrDetailRef = ref(null)
const OkrDetailsMainRef = ref(null)
const is_follow = ref(false)
const showPopover = ref(false)
const cancel = ref(0)
const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1

if (route.query.id != undefined || route.query.data != undefined) {
    id.value = Number(route.query.id || route.query.data)
    userid.value = Number(route.query.userid)
}

setTimeout(() => {
    detailsShow.value = true;
}, 0);

// 监听打开
const openDetail = (changeId, userId) => {
    router.replace({
        path: route.path,
        query: {
            id: changeId,
            userid: userId,
        },
    })
    id.value = changeId
    userid.value = userId
}

const getDetail = (item) => {
    detail.value = item
}
const handleReturn = () => {
    router.go(-1)
}

const handleEdit = () => {
    OkrDetailsMainRef.value.handleEdit()
    showPopover.value = false
}

const handleCancel = () => {
    OkrDetailsMainRef.value.handleCancel()
    showPopover.value = false
}

const handleFollowOkr = () => {
    OkrDetailsMainRef.value.handleFollowOkr()
    showPopover.value = false
}

const handleWarningShow = (e) => {
    OkrDetailsMainRef.value.handleWarningShow(e)
    showPopover.value = false
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
