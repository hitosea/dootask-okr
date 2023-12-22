<template >
    <div class="delete-box">
        <div class="delete-box-search">
            <div class="flex flex-1 gap-[16px]">
                <div class="delete-box-search-input">
                    <p class="delete-box-search-title text-text-li shrink-0">{{ $t('目标名称') }}</p>
                    <n-input v-model:value="tableKeyWord" :placeholder="$t('目标名称')" clearable />
                </div>
                <div class="flex items-center overflow-hidden" >
                    <p class="delete-box-search-title text-text-li mr-8 ">{{ $t('负责人') }}</p>
                    <n-select v-model:value="principalValue" :options="principalOptions" :on-blur="getPrincipalList" :on-search="getPrincipalList" :class="isAdmin ? '' :'max-w-[225px] ' " class="flex-1 overflow-hidden"
                        filterable :placeholder="$t('全部')" clearable>
                        <template #action>
                            <div v-if="principalLastPage > principalPage" quaternary
                                class=" h-full w-full whitespace-nowrap text-center" @click.stop="getPrincipalList('more')">
                                {{ $t('更多...') }}
                            </div>
                            <div v-else quaternary class=" h-full w-full whitespace-nowrap text-center">
                                {{ $t('已经到底了') }}
                            </div>
                        </template>
                    </n-select>
                </div>
                <n-button class="search-button" :loading="tableLoadIng > 0" type="primary" size="small" @click="getList('search')">
                    <template #icon>
                        <i v-if="APP_BASE_APPLICATION" class="ivu-icon ivu-icon-ios-search"></i>
                    <i v-else class="okrfont">&#xe6f8;</i>
                    </template>
                    {{ $t('搜索') }}
                </n-button>
            </div>
            <div class="delete-box-search-btn-row">
                <n-button class="search-button" :loading="tableLoadIng > 0" type="primary" size="small" @click="getList('search')">
                    <template #icon>
                        <i v-if="APP_BASE_APPLICATION" class="ivu-icon ivu-icon-ios-search"></i>
                    <i v-else class="okrfont">&#xe6f8;</i>
                    </template>
                    {{ $t('搜索') }}
                </n-button>
                <n-button
                    :strong="tableCheckedRowKeys.length == 0"
                    :secondary="tableCheckedRowKeys.length == 0"
                    :type="tableCheckedRowKeys.length == 0 ? 'tertiary' : 'primary'"
                    :disabled="tableCheckedRowKeys.length == 0"
                    class="float-right"
                    size="small"
                    @click="handleAssign"
                >
                    {{ $t('批量分配') }}
                </n-button>
            </div>
        </div>
        <div class="mt-16 flex flex-col flex-1 overflow-hidden">
            <div class="flex-1 overflow-auto">
                <n-data-table
                    :columns="tableColumns"
                    :data="tableData"
                    striped
                    :bordered="false"
                    :loading="tableLoadIng > 0"
                    :row-key="(row) => row.id"
                    @update:checked-row-keys="handleCheck"
                />
            </div>
            <div class="pagination mt-auto flex justify-center shrink-0 flex-0">
                <div  class=" hidden md:flex">
                    <n-pagination
                   
                    v-model:page="tablePage"
                    :default-page-size="tablePageSize"
                    :page-sizes="[10, 20, 30, 50, 100]"
                    :page-count="tableLastPage"
                    size="medium"
                    show-quick-jumper
                    show-size-picker
                    :on-update:page="onPage"
                    :on-update:page-size="onPageSize"
                >
                <template #prefix> {{ $t('共') }} {{ tableTotal }} {{ $t('条') }}
                    </template>
                    <template #suffix>{{ $t('页') }}</template>
            </n-pagination>
                </div>
                <div class="md:hidden flex">
                    <n-pagination class="pagination-web " simple  v-model:page="tablePage" v-model:page-size="tablePageSize" :page-count="tableLastPage" :on-update:page="onPage"/>
                </div>
            </div>
        </div>
    </div>

    <!-- OKR详情 -->
    <OkrDetailsModal :id="okrDetailsId" :show="okrDetailsShow" @getList="()=>{  getList('init') }" @close="() => { okrDetailsShow = false }"/>

    <!-- 警告框 -->
    <WarningPopup v-model:show="popupShow" :loading="deleteLoadIng>0" :title="popupTitle" :content="popupContent" @close="popupShow = false" @submit="handleDelete"></WarningPopup>

    <!-- 分配框 -->
    <n-modal v-model:show="assignShow" >
        <n-card class="w-[480px] max-w-[90%]" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <i class="okrfont cursor-pointer" @click="assignShow=false">&#xe6e5;</i>
            </template>
            <template #header>
                <h3 class="flex items-center text-16 font-medium text-title-color">
                    {{ $t('分配OKR') }}
                </h3>
            </template>
            <n-select v-model:value="assignUserId" :options="assignOptions" :on-blur="getAssignList" :on-search="getAssignList" class="flex-1 overflow-hidden"
                filterable :placeholder="$t('请选择用户')" clearable>
                <template #action>
                    <div v-if="assignLastPage > assignPage" quaternary
                        class=" h-full w-full whitespace-nowrap text-center" @click.stop="getAssignList('more')">
                        {{ $t('更多...') }}
                    </div>
                    <div v-else quaternary class=" h-full w-full whitespace-nowrap text-center">
                        {{ $t('已经到底了') }}
                    </div>
                </template>
            </n-select>
            <template #footer>
                <div class="flex justify-end gap-3">
                    <n-button secondary strong @click="assignShow = false">{{ $t('取消') }}</n-button>
                    <n-button type="primary" :loading="assignLoadIng > 0" @click="submitAssign">{{ $t('确定') }}</n-button>
                </div>
            </template>
        </n-card>
    </n-modal>

</template>
<script setup lang="ts">
import { DataTableColumn } from 'naive-ui';
import { useMessage } from "@/utils/messageAll"
import { getOwnerList ,getUserList} from '@/api/modules/created'
import { getLeaveList, okrDelete, okrAssign } from '@/api/modules/okrList'
import { UserStore } from '@/store/user'
import { NButton, DataTableRowKey } from 'naive-ui'
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import  WarningPopup from './WarningPopup.vue';
import { ResultDialog } from "@/api"
import { getUserInfo } from '@/api/modules/user';
const { proxy } = getCurrentInstance();
const APP_BASE_APPLICATION = computed(() => window.__MICRO_APP_BASE_APPLICATION__ ? 1 : 0)

const emit = defineEmits(['close'])

const isAdmin = UserStore().isAdmin()
const okrAdminOwner = ref(false)
const message = useMessage()


const popupShow  = ref(false)
const popupTitle  = ref('')
const popupContent  = ref('')
const popupDataId  = ref(0)

const principalValue = ref(null)
const principalOptions = ref([])
const principalPage = ref(1)
const principalLastPage = ref(99999)

const okrDetailsId = ref(0)
const okrDetailsShow = ref(false)
const deleteLoadIng = ref(0)

const assignId  = ref(0)
const assignUserId  = ref(null)
const assignShow  = ref(false)
const assignLoadIng  = ref(0)
const assignOptions = ref([])
const assignPage = ref(1)
const assignLastPage = ref(99999)

const tableLoadIng = ref(0)
const tableKeyWord = ref('')
const tableData = ref<any>([])
const tableTotal = ref(0);
const tablePage = ref(1);
const tablePageSize = ref(20);
const tableLastPage = ref(1)
const tableCheckedRowKeys = ref<DataTableRowKey[]>([])
const tableColumns = ref<DataTableColumn[]>([
    {
        type: 'selection',
        minWidth: 60,
    },
    {
        title: $t('目标名称'),
        key: 'title',
        minWidth: 200,
        render(rowData:any) {
            return h("div",{class: 'flex items-center'}, [
                    h("span",{class: 'okr-objective-num'}, rowData.objective_num || ('O' + rowData.id) ),
                    h("span",{class: 'line-clamp-1'}, rowData.title || "" ),
                ])
        }
    },
    {
        title: $t('负责人'),
        key: 'user',
        minWidth: 120,
        render(rowData:any) {
            return rowData.user?.nickname || ""
        }
    },
    {
        title: $t('当前O进度'),
        minWidth: 100,
        key: 'progress',
        render(rowData:any) {
            return rowData.progress + '%';
        }
    },
    {
        title: $t('操作'),
        minWidth: 160,
        width: 160,
        key: "action",
        render(rowData:any) {
            return [h(
                NButton,{
                    quaternary: true,
                    size: 'small',
                    type: 'primary',
                    style:{
                        height:'auto',
                    },
                    onClick: _ => {
                        okrDetailsId.value = rowData.id
                        okrDetailsShow.value = proxy.$openChildPage('/okrDetails',{ id: rowData.id })
                    }
                },
                { default: () => $t('查看') }
            ), h(
                NButton,{
                    quaternary: true,
                    size: 'small',
                    type: 'primary',
                    style:{
                        height:'auto',
                    },
                    onClick:  _ => handleAssign(rowData)
                },
                { default: () => $t('分配') }
            ), isAdmin ? h(
                NButton,{
                    quaternary: true,
                    size: 'small',
                    type: 'error',
                    style:{
                        height:'auto',
                    },
                    onClick: _ => {
                        popupShow.value = true
                        popupTitle.value = $t('删除')
                        popupDataId.value = rowData.id
                        popupContent.value = $t('确定要删除吗？', {title: rowData.title})
                    }
                },
                { default: () => $t('删除') }
            ) : '']
        }
    }
])

//负责人
const getPrincipalList = (type:any) => {
    let keyWord = '';
    if (type == 'init' || type === false) {
        principalPage.value = 1
        principalOptions.value = ([
            { label: $t('全部'), value: null }
        ])
    } else if(type == 'more') {
        principalPage.value++
    } else if(typeof type != 'object'){
        keyWord = type + ''
    }
    const sendata = {
        status: '2',
        page: principalPage.value,
        page_size: 20,
        keyword: keyWord,
    }
    getOwnerList(sendata).then(({ data }) => {
        principalOptions.value = data.data?.map(item=>{
            return {
                label: item.nickname,
                value: item.userid,
            }
        }) || []
        
        principalLastPage.value = data.last_page
    })
}

//用户列表
const getAssignList = (type:any) => {
    let keyWord = '';
    if (type == 'init' || type === false) {
        assignPage.value = 1
        assignOptions.value = ([
            { label: $t('全部'), value: null }
        ])
    } else if(type == 'more') {
        assignPage.value++
    } else if(typeof type != 'object'){
        keyWord = type + ''
    }
    const sendata = {
        status: '2',
        page: assignPage.value,
        page_size: 20,
        keyword: keyWord,
    }
    getUserList(sendata).then(({ data }) => {
        assignOptions.value = data.data?.map(item=>{
            return {
                label: item.nickname,
                value: item.userid,
            }
        }) || []
        
        assignLastPage.value = data.last_page
    })
}

// 获取列表
const getList = (type:string) => {
    if(type == 'search'){
        tablePage.value = 1
    }
    tableLoadIng.value = 1;
    getLeaveList({
        objective: tableKeyWord.value,
        page: tablePage.value,
        page_size: tablePageSize.value,
        userid: principalValue.value
    }).then(({ data }) => {
        tableLoadIng.value = 0
        data.data ? tableData.value = data.data : tableData.value = []
        tableTotal.value = data.count
        tableLastPage.value = data.last_page
    })
}

// 分页
const onPage = (page) => {
    tablePage.value = page
    getList('init')
}
const onPageSize = (pageSize) => {
    tablePage.value = 1
    tablePageSize.value = pageSize
    getList('init')
}

// 选中事件
const handleCheck = (rowKeys) => {
    tableCheckedRowKeys.value = rowKeys
}


// 分配
const handleAssign = (rowData) => {
    if(rowData && rowData.id){
        assignId.value = rowData.id
    }else{
        assignId.value = 0
    }
    getAssignList('init')
    assignUserId.value = null
    assignShow.value = true
}
const submitAssign = () => {
    if(!assignUserId.value){
        message.error($t('请选择用户'))
        return;
    }
    assignLoadIng.value = 1
    okrAssign({
        okr_ids: [assignId.value,...tableCheckedRowKeys.value].filter(h=>h),
        userid: assignUserId.value,
    }).then(({ data }) => {
        message.success($t('分配成功'))
        emit('close', 1)
        getList('init')
    }).finally(() => {
        assignLoadIng.value = 0
        assignShow.value = false
        assignUserId.value = null
        tableCheckedRowKeys.value = []
    })
}

// 删除
const handleDelete = () => {
    deleteLoadIng.value = 1
    okrDelete({id:popupDataId.value }).then(({ msg }) => {
        message.success($t('删除成功'))
        getList('init')
    })
    .catch(ResultDialog)
    .finally(() => {
        deleteLoadIng.value = 0
        popupShow.value = false
    })
}
//获取能否能搜索部门
const handleGetUserInfo = () => {
    getUserInfo().then(({ data }) => {
        okrAdminOwner.value = data.okr_admin_owner
    })
}
//
onMounted(() => {
    getPrincipalList('init')
    getList('search')
    handleGetUserInfo()
})
</script>

<style lang="less" scoped>
.delete-box {
    @apply flex flex-col h-full;

    .delete-box-search {
        @apply flex gap-16 items-center;

        .delete-box-search-input {
            @apply flex items-center gap-2;
        }
        .delete-box-search-btn-row{
            .search-button{
                @apply hidden;
            }
        }
    }

    .pagination{
        @apply pt-16;
    }
}

body.window-portrait {
    .delete-box {
        .delete-box-search{
            @apply block;
            >div{
                @apply block;
            }
            .delete-box-search-title{
                width: 60px;
            }
            .delete-box-search-input{
                @apply mb-16;
            }
            .delete-box-search-btn-row{
                @apply flex mt-16 justify-between;
                .search-button{
                    @apply flex flex;
                }
            }
            .search-button{
                @apply hidden;
            }
        }

    }
}
</style>
<style>
.okr .delete-box .n-checkbox .n-checkbox-box{
    border-radius: 4px !important;
}
</style>
