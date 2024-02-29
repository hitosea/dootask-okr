<template >
    <div class="page-okr" ref="pageOkrRef">
        <div class="okr-title">
            <div class="flex items-center">
                <div class="okr-nav-back" @click="handleReturn"><i class="okrfont">&#xe676;</i></div>
                <h2 :class="searchShow ? 'title-active' : ''">{{ $t(pageTitle) }}</h2>
                <div :class="searchShow ? 'title-active' : ''" class="okr-app-refresh" v-if="!loadIng" @click="reLoadList"><i class="okrfont">&#xe6ae;</i></div>
            </div>
            <div class="okr-right z-[2]">
                <n-button class="search-button" @mouseover="() => { searchShow = true }" @mouseout="() => { searchShow = false }"
                    :class="searchShow || searchObject ? 'search-active' : ''" circle>
                    <span class="search-button-span border-[rgba(142,142,143,0.5)] h-[16px] leading-4" v-show="searchShow || searchObject">{{ inputName }}</span>
                    <n-input
                        ref="searchInputRef"
                        v-show="searchShow || searchObject"
                        v-model:value="searchObject"
                        class="border-none"
                        clearable
                        :placeholder="$t('请输入目标 (O)')"/>
                    <i v-if="APP_BASE_APPLICATION" class="ivu-icon ivu-icon-ios-search"></i>
                    <i v-else class="okrfont">&#xe6f8;</i>
                </n-button>
                <n-button class="add-button" type="tertiary" @click="handleAdd" circle>
                    <n-spin size="small" :show="btnLoading>0">
                        <i v-if="APP_BASE_APPLICATION" class="ivu-icon ivu-icon-md-add"></i>
                        <i v-else class="okrfont">&#xe6f2;</i>
                    </n-spin>
                </n-button>
                <n-button class="more-button" type="tertiary" @click="moreButtonPopoverShow = true" circle>
                    <n-popover class="okr-more-button-popover"
                    :show="moreButtonPopoverShow"
                    @clickoutside="moreButtonPopoverShow = false"
                    placement="bottom"
                    :z-index="modalTransferIndex()"
                    trigger="click"
                    raw
                    :show-arrow="true">
                        <template #trigger>
                            <i v-if="APP_BASE_APPLICATION" class="ivu-icon ivu-icon-ios-more font-bold"></i>
                            <i v-else class="okrfont">&#xe6f2;</i>
                        </template>
                        <div class="flex flex-col">
                            <p v-if="proxy.$globalStore.electron && !isSingle" @click="[openNewWin(), moreButtonPopoverShow=false]"> {{ $t('新窗口打开') }}</p>
                            <p @click="[handleArchiveShow(), moreButtonPopoverShow=false]"> {{ $t('已归档OKR') }}</p>
                            <p v-if="isAdmin || isDepartmentOwner" @click="[handleDeleteShow(), moreButtonPopoverShow=false]"> {{ $t('离职/删除人员OKR') }}</p>
                            <p v-if="isAdmin" @click="[handleSettingShow(), moreButtonPopoverShow=false]"> {{ $t('设置') }}</p>
                        </div>
                    </n-popover>
                </n-button>
            </div>
        </div>
        <div class="okr-tabs">
            <n-tabs type="line" :value="tabsName" animated :on-update:value="changeTabs">
                <n-tab-pane :tab="$t('我创建的')" name="created">
                    <div class="okr-scrollbar">
                        <Icreated ref="ICreatedRef" :searchObject="searchObject" @edit="handleEdit" @add="handleAdd"/>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('我参与的')" name="katılım">
                    <div class="okr-scrollbar">
                        <OkrParticipant ref="OkrParticipantRef" :searchObject="searchObject" @edit="handleEdit"/>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('部门OKR')" name="dept">
                    <div class="okr-scrollbar">
                        <OkrDepartment ref="OkrDepartmentRef" :searchObject="searchObject" @edit="handleEdit"/>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('我关注的')" name="attention">
                    <div class="okr-scrollbar">
                        <OkrFollow ref="OkrFollowRef" :searchObject="searchObject" @edit="handleEdit"/>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('OKR复盘')" name="review">
                    <div class="okr-scrollbar">
                        <OkrReplay ref="OkrReplayRef" :searchObject="searchObject" @edit="handleEdit"/>
                    </div>
                </n-tab-pane>
            </n-tabs>
        </div>
    </div>

    <!-- 添加Okr -->
    <AddOkrsDrawer v-model:show="addShow" :edit="edit" :editData="editData" @close="handleClose"></AddOkrsDrawer>

    <!-- 已归档 -->
    <ArchiveDrawer v-model:show="archiveShow" @close="handleClose"></ArchiveDrawer>

    <!-- 离职/删除人员OKR -->
    <DeleteDrawer v-model:show="deleteShow" @close="handleClose"></DeleteDrawer>

    <!-- 设置 -->
    <SettingDrawer v-model:show="settingShow" @close="()=>{ settingShow = false }"></SettingDrawer>

    <!-- 强提示 -->
    <TipsModal :show="showModal" :content="tipsContent" @close="() => { showModal = false }"></TipsModal>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import AddOkrsDrawer from './components/AddOkrsDrawer.vue'
import ArchiveDrawer from './components/ArchiveDrawer.vue'
import DeleteDrawer from './components/DeleteDrawer.vue'
import SettingDrawer from './components/SettingDrawer.vue'
import Icreated from '@/views/manage/Icreated.vue'
import OkrReplay from '@/views/manage/OkrReplay.vue'
import OkrFollow from '@/views/manage/OkrFollow.vue'
import OkrParticipant from '@/views/manage/OkrParticipant.vue'
import OkrDepartment from './manage/OkrDepartment.vue';
import { useRouter, useRoute } from 'vue-router'
import TipsModal from '@/views/components/TipsModal.vue';
import { getUserInfo } from '@/api/modules/user'
import { UserStore } from '@/store/user'

const { proxy } = getCurrentInstance();
const isAdmin = UserStore().isAdmin()
const isDepartmentOwner = UserStore().isDepartmentOwner()
const APP_BASE_APPLICATION = computed(() => window.__MICRO_APP_BASE_APPLICATION__ ? 1 : 0)
const router = useRouter()
const isSingle = proxy.$globalStore.isSingle()
const pageTitle = ref("OKR管理")
const loadIng = ref(false)
const route = useRoute()
const searchInputRef = ref(null)
const pageOkrRef = ref(null)
const ICreatedRef = ref(null)
const OkrParticipantRef = ref(null)
const OkrDepartmentRef = ref(null)
const OkrFollowRef = ref(null)
const OkrReplayRef = ref(null)

const addShow = ref(false)
const archiveShow = ref(false)
const deleteShow = ref(false)
const settingShow = ref(false)
const moreButtonPopoverShow = ref(false)
const edit = ref(false)
const searchObject = ref('')
const searchShow = ref(false)
const tabsName = ref('created')

const showModal = ref(false)
const tipsContent = ref('')
const btnLoading = ref(0)

let editData = {}


watch(route,(newValue)=>{
    nextTick(()=>{
        if(newValue.query.active == undefined && OkrFollowRef.value != null){
            OkrFollowRef.value.getList('search')
        }
    })
},{immediate:true})

watch(searchShow,(newValue)=>{
    nextTick(()=>{
        if (newValue) {
            setTimeout(()=>{
                searchInputRef.value?.focus()
            },100)
        }
    })
},{immediate:true})


if (route.query.active == undefined) {
    router.replace({
        path: route.path,
        query: { active: tabsName.value }
    })
} else {
    tabsName.value = route.query.active + ''
}

const inputName = computed(()=>{
    if(tabsName.value == 'created'){
        return $t('我创建的')
    }
    if(tabsName.value == 'katılım'){
        return $t('我参与的')
    }
    if(tabsName.value == 'dept'){
        return $t('部门OKR')
    }
    if(tabsName.value == 'attention'){
        return $t('我关注的')
    }
    if(tabsName.value == 'review'){
        return $t('OKR复盘')
    }
})

const changeTabs = (e) => {
    searchObject.value = ''
    searchShow.value = false
    tabsName.value = e
    router.replace({
        path: route.path,
        query: { active: e }
    })
}

const reLoadList = () => {
    loadIng.value = true;
    //重新获取列表
   if (tabsName.value == 'created' ) {
        ICreatedRef.value.resetGetList('search')
    }
    if (tabsName.value == 'katılım') {
        OkrParticipantRef.value.resetGetList('search')
    }
    if (tabsName.value == 'dept') {
        OkrDepartmentRef.value.resetGetList('search')
    }
    if (tabsName.value == 'attention') {
        OkrFollowRef.value.resetGetList('search')
    }
    if (tabsName.value == 'review') {
        OkrReplayRef.value.resetGetList('search')
    }
    setTimeout(()=>{
        loadIng.value = false;
    },300)
}

//编辑
const handleEdit = (data) => {
    addShow.value = true
    edit.value = true
    editData = data
}

//添加OKR
const handleAdd = () => {
    btnLoading.value = 1
    getUserInfo().then(({ data }) => {
            if (data.identity[0] != 'admin' && data.department && data.department.length == 0) {
                tipsContent.value = $t('您当前未加入任何部门，不能发起！')
                showModal.value = true
                return
            }
            else{
                addShow.value = proxy.$openChildPage('/addOkr')
            }
    })
    .catch()
    .finally(()=>{
        setTimeout(()=>{
            btnLoading.value = 0
        },300)
    })
}

const handleClose = (e, id) => {

    //重新获取列表
    if (tabsName.value == 'created' && e == 1) {
        ICreatedRef.value.resetGetList('search')
    }
    if (tabsName.value == 'katılım' && e == 1) {
        OkrParticipantRef.value.resetGetList('search')
    }
    if (tabsName.value == 'dept' && e == 1) {
        OkrDepartmentRef.value.resetGetList('search')
    }
    if (tabsName.value == 'attention' && e == 1) {
        OkrFollowRef.value.resetGetList('search')
    }

    //更新单条数据
    if (tabsName.value == 'created' && e == 2) {
        ICreatedRef.value.upData(id)
    }
    if (tabsName.value == 'katılım' && e == 2) {
        OkrParticipantRef.value.upData(id)
    }
    if (tabsName.value == 'dept' && e == 2) {
        OkrDepartmentRef.value.upData(id)
    }
    if (tabsName.value == 'attention' && e == 2) {
        OkrFollowRef.value.upData(id)
    }

    edit.value = false
    editData = {}
    addShow.value = false
}

const handleReturn = () => {
    router.go(-1)
}

const modalTransferIndex = () => {
    return window.modalTransferIndex = window.modalTransferIndex + 1
}

// 新窗口打开
const openNewWin = () => {
    proxy.$globalStore.electron.sendMessage('windowRouter', {
        name: `okr`,
        path: `single/apps/okr/list?tab=${tabsName.value}`,
        force: false,
        config: {
            title: $t(pageTitle.value),
            titleFixed: true,
            parent: null,
            width: Math.min(window.screen.availWidth, pageOkrRef.value.clientWidth + 72),
            height: Math.min(window.screen.availHeight, pageOkrRef.value.clientHeight + 36),
            minWidth: 600,
            minHeight: 450,
        }
    });
}

// 已归档
const handleArchiveShow = () => {
    archiveShow.value = proxy.$openChildPage('/archive')
}

// 已离职删除人员
const handleDeleteShow = () => {
    deleteShow.value = proxy.$openChildPage('/deletePersonnel')
}

// 设置
const handleSettingShow = () => {
    settingShow.value = proxy.$openChildPage('/setting')
}
</script>

<style lang="less" scoped>
.page-okr {
    @apply absolute top-0 bottom-0 left-0 right-0 flex flex-col bg-page-bg  px-16 py-20 md:px-20;

    .okr-title {
        @apply h-42 flex justify-between items-center relative mt-12 mb-14;

        .icon-return {
            @apply block md:hidden mr-16 text-20 z-[2];
        }

        h2 {
            @apply text-title-color text-28 font-semibold;
        }

        .title-active {
            @apply hidden md:block;
        }

        .okr-right {
            @apply flex items-center gap-4 ;

            .add-button,
            .more-button,
            .search-button {
                @apply bg-[#f2f3f5] w-36 h-36 rounded-full flex items-center justify-center cursor-pointer;

                i {
                    @apply text-20 text-emoji-users-color;
                }
                text-align: left;
            }

            .search-button {
                @apply flex items-center;

                i,
                span {
                    @apply flex-initial;
                }

                .search-button-span {
                    @apply text-14  text-emoji-users-color pr-8 border-solid border-0 border-r;
                }

                :deep(.n-input) {
                    @apply flex-1 bg-transparent border-0;

                    .n-input__border,
                    .n-input__state-border,
                    .n-input__border:focus,
                    .n-input__state-border:focus {
                        @apply border-0 shadow-none;
                    }
                }
            }

            .search-active {
                @apply w-auto flex-1 md:w-320 px-14;

                i {
                    @apply pl-8;
                }
            }
        }
    }

    .okr-tabs {
        @apply flex-1 relative;

        :deep(.n-tabs) {
            @apply absolute left-0 right-0 top-0 bottom-0;

            .n-tabs-pane-wrapper {
                @apply flex-1 relative;

                .n-tab-pane {
                    @apply max-h-full absolute left-0 right-0 top-0 bottom-0;

                    .okr-scrollbar {
                        @apply absolute left-0 right-0 top-0 bottom-0;
                    }
                }
            }
        }
    }
}

//
body.window-portrait {
    .page-okr {
        .okr-title {
            margin: 4px 0 14px -4px;
        }
    }
}
</style>
