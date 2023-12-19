<template >
    <div class="archive-box">
        <div class="archive-box-search">
            <div class="archive-search-input">
                <p class=" text-text-li shrink-0">{{ $t('目标名称') }}</p><n-input v-model:value="objective" :loading="loadIng"
                    :placeholder="$t('目标名称')" clearable />
            </div>
            <n-button :loading="loadIng" type="primary" size="small" @click="onSearch">
                <template #icon>
                    <i class="okrfont">&#xe72a;</i>
                </template>
                {{ $t('搜索') }}
            </n-button>
        </div>
        <div class="mt-16 flex flex-col flex-1">
            <div>
                <n-data-table :columns="columns" :data="tableData" :single-line="false" :hover="false" :loading="loadIng"
                    style="--n-td-color-hover-modal:#ffffff" />
            </div>
            <div class="mt-auto md:flex hidden justify-center shrink-0">
                <n-pagination v-model:page="tablePage" v-model:page-size="tablePageSize" :page-count="tableTotal"
                    size="medium" :page-sizes="[10, 20, 30, 50, 100]" show-quick-jumper show-size-picker
                    :on-update:page="onPage" :on-update:page-size="onPageSize" />
            </div>
            <div class="flex md:hidden justify-center shrink-0 mt-auto">
                <n-pagination class="pagination-web" simple v-model:page="tablePage" :page-count="tableTotal" />
            </div>
        </div>

        <WarningPopup v-model:show="showModal" :title="OTitle" :content="OContent" @submit="handleSubmit"
            @close="showModal = false">
        </WarningPopup>

        <!-- OKR详情 -->
        <OkrDetailsModal :id="okrDetailsId" :show="okrDetailsShow" @getlist="handleGetOkrArchive"
            @close="() => { okrDetailsShow = false }" />

    </div>
</template>
<script setup lang="ts">
import { DataTableColumn } from 'naive-ui';
import WarningPopup from '@/views/components/WarningPopup.vue';
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import { getOkrArchive, okrArchiveRestore, okrDelete } from '@/api/modules/okrList'
import utils from '@/utils/utils';
import { useMessage } from "@/utils/messageAll"

const message = useMessage()

const { proxy } = getCurrentInstance();

const loadIng = ref(false)
const objective = ref('')

const tablePage = ref(1);
const tablePageSize = ref(20);
const tableTotal = ref(0);
const tableData = ref<any>([])

const okrDetailsId = ref(0)
const okrDetailsShow = ref(false)

const showModal = ref(false);

const OTitle = ref('');
const OContent = ref('');
const OId = ref(0)
const OType = ref(0)

const emit = defineEmits(['close'])

const columns = ref<DataTableColumn[]>([
    {
        title: $t('目标名称'),
        key: 'title',
        minWidth: 200,
        render(rowData) {
            let arr = []
            arr.push(
                h("span", { class: 'okr-objective-num' }, rowData.objective_num || ('O' + rowData.id)),
            )
            arr.push(
                h('span', {
                    class: 'line-clamp-1'
                }, rowData.title)
            )
            return h('div', {
                class: " flex items-center"
            }, arr)
        }
    },
    {
        title: $t('完成时间'),
        key: 'end_at',
        minWidth: 200,
        render(rowData) {
            return h('div', {
            }, utils.GoDateHMS(rowData.end_at || 0))
        }
    },
    {
        title: $t('归档时间'),
        minWidth: 200,
        key: 'archive_at',
        render(rowData) {
            return h('div', {
            }, utils.GoDateHMS(rowData.archive_at || 0))
        }
    },
    {
        title: $t('归档人员'),
        minWidth: 100,
        key: 'nickname',
        render(rowData) {
            return h('div', {
            }, rowData.archive_user.nickname)
        }
    },
    {
        title: $t('操作'),
        minWidth: 160,
        width: 160,
        key: "",
        render(rowData) {
            let arr = []
            arr.push(
                h('span', {
                    class: "text-primary-color cursor-pointer",
                    onClick: () => {
                        okrDetailsId.value = rowData.id
                        okrDetailsShow.value = proxy.$openChildPage('/okrDetails', { id: rowData.id })
                    }
                }, $t('查看'))
            )
            arr.push(
                h('span', {
                    class: "text-primary-color cursor-pointer",
                    onClick: () => {
                        OTitle.value = $t('还原')
                        OContent.value = $t('确定要还原') + `【${rowData.title}】？`
                        OId.value = rowData.id
                        OType.value = 1
                        showModal.value = true
                    }
                }, $t('还原'))
            )
            arr.push(
                h('span', {
                    class: "text-[#ED4014] cursor-pointer",
                    onClick: () => {
                        OTitle.value = $t('删除')
                        OContent.value = $t('确定要删除') + `【${rowData.title}】？`
                        OId.value = rowData.id
                        OType.value = 2
                        showModal.value = true
                    }
                }, $t('删除'))
            )
            return h('div', {
                style: {
                    display: 'flex',
                    gap: '12px',
                    alignItems: 'items-center',
                },
            }, arr)
        }
    }
])

// 归档列表
const handleGetOkrArchive = () => {

    loadIng.value = true
    getOkrArchive({
        objective: objective.value,
        page: tablePage.value,
        page_size: tablePageSize.value,
    }).then(({ data }) => {
        tableData.value = (data.data || [])
        tableTotal.value = data.last_page
    })
        .catch(({ msg }) => {

        })
        .finally(() => {
            loadIng.value = false
        })

}

const handleSubmit = () => {
    if (OType.value == 1) {
        handleRestore()
    }
    if (OType.value == 2) {
        handleDelete()
    }
}


//还原归档目标
const handleRestore = () => {
    loadIng.value = true
    okrArchiveRestore({
        id: OId.value,
    }).then(({ data }) => {
        message.success($t('操作成功'))
        handleGetOkrArchive();
        emit('close', 1)
    })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
            showModal.value = false
        })
}

//删除归档目标
const handleDelete = () => {
    loadIng.value = true
    okrDelete({
        id: OId.value,
    }).then(({ data }) => {
        message.success($t('操作成功'))
        handleGetOkrArchive();
    })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
            showModal.value = false
        })
}


const onSearch = () => {
    tablePage.value = 1
    handleGetOkrArchive();
}

// 分页
const onPage = (page) => {
    tablePage.value = page
    handleGetOkrArchive()
}
const onPageSize = (pageSize) => {
    tablePage.value = 1
    tablePageSize.value = pageSize
    handleGetOkrArchive()
}

onMounted(() => {
    handleGetOkrArchive();
})


</script>
<style scoped lang="less">
.archive-box {
    @apply flex flex-col h-full;
}

.archive-box-search {
    @apply flex gap-4 items-center md:justify-start justify-between;
}

.archive-search-input {
    @apply flex items-center gap-2;
}
</style >
<style >
.pagination-web .n-pagination .n-pagination-item.n-pagination-item--disabled.n-pagination-item--active,
.n-pagination .n-pagination-item.n-pagination-item--disabled.n-pagination-item--button {
    background: none;
    border: none;
}

.pagination-web .n-pagination-quick-jumper {
    width: 30px;
}
.pagination-web .n-pagination-quick-jumper .n-input{
    background: none;
}
</style>