<template >
    <div class="page-okr">
        <div class="okr-title">
            <i @click="handleReturn" class="taskfont icon-return">&#xe704;</i>
            <h2 :class="searchShow ? 'title-active' : ''">{{ $t('OKR管理') }}</h2>
            <div class="okr-right">
                <div class="search-button" @mouseover="()=>{searchShow = true }" @mouseout="()=>{searchShow = false }" :class="searchShow || searchObject? 'search-active' : ''">
                    <span class="search-button-span" v-show="searchShow || searchObject" >{{ tabsName }}</span>
                    <n-input  v-show="searchShow || searchObject" class="border-none" clearable v-model:value="searchObject" :placeholder="$t('请输入目标 (O)')" />
                    <i class="taskfont" >&#xe6f8;</i>
                </div>
                <div class="add-button" type="tertiary" @click="handleAdd">
                    <i class="taskfont">&#xe6f2;</i>
                </div>
            </div>
        </div>
        <div class="okr-tabs">
            <n-tabs type="line" :value="tabsName" animated :on-update:value="changeTabs">
                <n-tab-pane :tab="$t('我创建的OKR')" :name="$t('我创建的OKR')">
                    <div class="okr-scrollbar">
                        <Icreated ref="ICreatedRef" :searchObject="searchObject" @edit="handleEdit" @add="handleAdd"></Icreated>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('参与的OKR')" :name="$t('参与的OKR')">
                    <div class="okr-scrollbar">
                        <OkrParticipant ref="OkrParticipantRef" :searchObject="searchObject"  @edit="handleEdit" ></OkrParticipant>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('部门OKR')" :name="$t('部门OKR')">
                    <div class="okr-scrollbar">
                        <OkrDepartment ref="OkrDepartmentRef" :searchObject="searchObject" @edit="handleEdit"></OkrDepartment>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('关注的OKR')" :name="$t('关注的OKR')">
                    <div class="okr-scrollbar">
                        <OkrFollow ref="OkrFollowRef" :searchObject="searchObject" @edit="handleEdit"></OkrFollow>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('OKR复盘')" :name="$t('OKR复盘')">
                    <div class="okr-scrollbar">
                        <OkrReplay ref="OkrReplayRef" :searchObject="searchObject" @edit="handleEdit"></OkrReplay>
                    </div>
                </n-tab-pane>
            </n-tabs>
        </div>
    </div>
    <AddOkrsDrawer v-model:show="addShow" :edit="edit" :editData="editData" @close="handleClose"></AddOkrsDrawer>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import AddOkrsDrawer from './components/AddOkrsDrawer.vue'
import Icreated from '@/views/manage/Icreated.vue'
import OkrReplay from '@/views/manage/OkrReplay.vue'
import OkrFollow from '@/views/manage/OkrFollow.vue'
import OkrParticipant from '@/views/manage/OkrParticipant.vue'
import OkrDepartment from './manage/OkrDepartment.vue';
import { useRouter } from 'vue-router'

const router = useRouter()
const ICreatedRef = ref(null)
const OkrParticipantRef = ref(null)
const OkrDepartmentRef = ref(null)
const OkrFollowRef = ref(null)
const OkrReplayRef = ref(null)

const addShow = ref(false)
const edit = ref(false)
let editData = {}
const searchObject = ref('')
const searchShow = ref(false)
const tabsName = ref($t('我创建的OKR'))


const changeTabs = (e)=>{
    searchObject.value = ''
    searchShow.value = false
    tabsName.value = e
}

//编辑
const handleEdit = (data) => {
    addShow.value = true
    edit.value = true
    editData = data
}

//添加OKR
const handleAdd = () => {
    if (window.innerWidth < 768) {
        router.push('/addOkr')
    }
    else{
        addShow.value = true
    }
}

const handleClose = (e,id) => {
    //重新获取列表
    if(tabsName.value == $t('我创建的OKR') && e == 1){
        ICreatedRef.value.getList('search')     
    }
    else if(tabsName.value == $t('参与的OKR') && e == 1){
        OkrParticipantRef.value.getList('search')
    }
    else if(tabsName.value == $t('部门OKR') && e == 1){
        OkrDepartmentRef.value.getList('search')
    }
    else if(tabsName.value == $t('关注的OKR') && e == 1){
        OkrFollowRef.value.getList('search')
    }

    //更新单条数据
    if(tabsName.value == $t('我创建的OKR') && e == 2){
        ICreatedRef.value.upData(id)      
    }
    else if(tabsName.value == $t('参与的OKR') && e == 2){
        OkrParticipantRef.value.upData(id)
    }
    else if(tabsName.value == $t('部门OKR') && e == 2){
        OkrDepartmentRef.value.upData(id)
    }
    else if(tabsName.value == $t('关注的OKR') && e == 2){
        OkrFollowRef.value.upData(id)
    }

    edit.value = false
    editData = {}
    addShow.value = false
}

const handleReturn = () => {
    router.back()
}

</script>

<style lang="less" scoped>
.page-okr {
    @apply absolute top-0 bottom-0 left-0 right-0 flex flex-col p-16 md:p-24 bg-page-bg;

    .okr-title {
        @apply flex justify-between items-center mb-8 md:mb-14 relative;

        .icon-return{
            @apply block md:hidden mr-16 text-20 z-[2];
        }

        h2 {
            @apply text-title-color text-17 md:text-28 font-normal absolute left-0 right-0 text-center md:relative;
        }
        .title-active{
            @apply hidden md:block;
        }

        .okr-right{
            @apply flex items-center gap-4 md:gap-6 z-[2];
            .add-button,
            .search-button {
                @apply bg-bg-manage-menu w-36 h-36 md:w-12 md:h-12 rounded-full flex items-center justify-center cursor-pointer;

                i {
                    @apply text-15 md:text-20 text-emoji-users-color;
                }
            }
            .search-button{
                @apply flex items-center;
                i,span{
                    @apply flex-initial;
                }

                .search-button-span{
                    @apply text-14 text-emoji-users-color pr-8 border-solid border-0 border-r border-text-tips;
                }
                :deep(.n-input){
                    @apply flex-1 bg-transparent border-0;
                    .n-input__border,.n-input__state-border,.n-input__border:focus,.n-input__state-border:focus{
                        @apply  border-0 shadow-none;
                    }
                }
            }
            .search-active{
                @apply w-auto flex-1 md:w-320 px-14;
                i{
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
}</style>
