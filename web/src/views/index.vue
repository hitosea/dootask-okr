<template >
    <div class="page-okr">
        <div class="okr-title">
            <h2>{{ $t('OKR管理') }}</h2>
            <div class="okr-right">
                <div class="search-button" @click="()=>{searchShow = true}" :class="searchShow ? 'search-active' : ''">
                    <span class="search-button-span" v-if="searchShow" >{{ tabsName }}</span>
                    <n-input v-if="searchShow" class="border-none" clearable v-model:value="searchObject" :placeholder="$t('请输入目标 (O)')" />
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
                        <Icreated></Icreated>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('参与的OKR')" :name="$t('参与的OKR')">
                    <div class="okr-scrollbar">
                        
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('部门OKR')" :name="$t('部门OKR')">
                    <div class="okr-scrollbar">
                        
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('关注的OKR')" :name="$t('关注的OKR')">
                    <div class="okr-scrollbar">
                        <OkrFollow></OkrFollow>
                    </div>
                </n-tab-pane>
                <n-tab-pane :tab="$t('OKR复盘')" :name="$t('OKR复盘')">
                    <div class="okr-scrollbar">
                        <OkrReplay></OkrReplay>
                    </div>
                </n-tab-pane>
            </n-tabs>
        </div>
    </div>
    <AddOkr v-model:show="addShow" @close="handleClose"></AddOkr>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import AddOkr from '@/views/components/AddOkrs.vue';
import Icreated from '@/views/manage/Icreated.vue'
import OkrReplay from '@/views/manage/OkrReplay.vue'
import OkrFollow from '@/views/manage/OkrFollow.vue'
const addShow = ref(false)
const searchObject = ref('')
const searchShow = ref(false)
const tabsName = ref('我创建的OKR')

const changeTabs = (e)=>{
    searchObject.value = ''
    searchShow.value = false
    tabsName.value = e
}

const handleAdd = () => {
    addShow.value = true
}

const handleClose = () => {
    addShow.value = false
}

</script>

<style lang="less">
.page-okr {
    @apply absolute top-0 bottom-0 left-0 right-0 flex flex-col p-24 bg-page-bg;

    .okr-title {
        @apply flex justify-between items-center;

        h2 {
            @apply text-title-color text-36 font-normal;
        }
        .okr-right{
            @apply flex items-center gap-6;
            .add-button,
            .search-button {
                @apply bg-bg-manage-menu w-12 h-12 rounded-full flex items-center justify-center cursor-pointer;
    
                i {
                    @apply text-20 text-emoji-users-color;
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
                .n-input{
                    @apply flex-1 bg-transparent border-0;
                    .n-input__border,.n-input__state-border,.n-input__border:focus,.n-input__state-border:focus{
                        @apply  border-0 shadow-none;
                    }
                }
            }
            .search-active{
                @apply w-320 px-16;
                i{
                    @apply pl-8;
                }
            }
        }
    }

    .okr-tabs {
        @apply flex-1 relative;

        .n-tabs {
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