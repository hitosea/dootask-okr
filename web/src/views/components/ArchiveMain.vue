<template >
  <div class="archive-box">
    <div class="archive-box-search">
      <div class="archive-search-input">
        <p class=" text-text-li shrink-0">{{ $t('目标名称') }}</p><n-input v-model="keyWord" :placeholder="$t('目标名称')"
          clearable />
      </div>
      <n-button :loading="loadIng > 0" type="primary" size="small" @click="onSearch">
        <template #icon>
          <i class="okrfont">&#xe72a;</i>
        </template>
        {{ $t('搜索') }}
      </n-button>
    </div>
    <div class="mt-16 flex flex-col flex-1">
      <n-data-table :columns="columns" :data="tableData" :single-line="false" :hover="false"
        style="--n-td-color-hover-modal:#ffffff" />
      <div class="mt-auto flex justify-center">
        <n-pagination v-model:page="page" :page-count="total" size="medium" show-quick-jumper show-size-picker />
      </div>

    </div>
  </div>
</template>
<script setup lang="ts">
import { DataTableColumn } from 'naive-ui';

const loadIng = ref(0)
const keyWord = ref('')
const page = ref(1);
const total = ref(0);

const tableData = ref<any>([])

const columns = ref<DataTableColumn[]>([
  {
    title: $t('目标名称'),
    key: 'O',
    minWidth: 200,
    rowSpan: (rowData) => (rowData.lenght),
    render(rowData) {
      let arr = []
      arr.push(
        h('span', {
          class: 'line-clamp-1'
        }, rowData.O)
      )
      return h('div', {
        style: {
          display: 'flex',
        },
      }, arr)
    }
  },
  {
    title: $t('完成时间'),
    key: 'Ocomplete',
    minWidth: 100,
    rowSpan: (rowData) => (rowData.lenght),
  },
  {
    title: $t('归档时间'),
    minWidth: 200,
    key: 'KR',
    render(rowData, index) {
      let arr = []
      arr.push(
        h('span', {
          class: 'text-primary-color shrink-0 mr-4',
        }, 'KR' + (index + 1))
      )
      arr.push(
        h('span', {
          class: 'line-clamp-2'
        }, rowData.KR)
      )
      return h('div', {
        style: {
          display: 'flex',
          alignItems: 'start',
        },
      }, arr)
    }
  },
  {
    title: $t('归档人员'),
    minWidth: 100,
    key: 'KRcomplete',
  },
  {
    title: $t('操作'),
    minWidth: 150,
    key:"",
    render(rowData) {
      let arr = []
      return h('div', {
        style: {
          display: 'flex',
          alignItems: 'items-center',
        },
      }, arr)
    }
  }
])

const onSearch = () => { }
</script>
<style scoped>
.archive-box {
  @apply flex flex-col h-full;
}

.archive-box-search {
  @apply flex gap-[16px] items-center;
}

.archive-search-input {
  @apply flex items-center gap-[8px];
}
</style>