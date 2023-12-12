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
      <div>
        <n-data-table :columns="columns" :data="tableData" :single-line="false" :hover="false"
        style="--n-td-color-hover-modal:#ffffff" />
      </div>
      <div class="mt-auto flex justify-center shrink-0">
        <n-pagination v-model:page="page" :page-count="total" size="medium" show-quick-jumper show-size-picker />
      </div>
    </div>

   <WarningPopup v-model:show="showModal" :title="OTitle" :content="OContent" @close=" showModal = false "></WarningPopup>
  </div>
</template>
<script setup lang="ts">
import { DataTableColumn } from 'naive-ui';
import  WarningPopup from './WarningPopup.vue';

const loadIng = ref(0)
const keyWord = ref('')
const page = ref(1);
const total = ref(0);
const showModal = ref(false);
const OTitle = ref('');
const OContent = ref('');

const tableData = ref<any>([
  // {
  //   O:"打卡的话旷达科技啊快进到撒寄快递会尽快的哈时间跨度哈看的h",
  //   Ocomplete:"21323",
  // }
])

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
      arr.push(
        h('span',{
          class:"text-primary-color cursor-pointer",
        },$t('查看'))
      )
      arr.push(
        h('span',{
          class:"text-primary-color cursor-pointer",
          onClick: () => {
            OTitle.value = $t('还原')
            OContent.value = $t('确定要还原')+`【${rowData.O}】？`
             showModal.value = true
            }
        },$t('还原'))
      )
      arr.push(
        h('span',{
          class:"text-[#ED4014] cursor-pointer",
        },$t('删除'))
      )
      return h('div', {
        style: {
          display: 'flex',
          gap:'12px',
          alignItems: 'items-center',
        },
      }, arr)
    }
  }
])

const onSearch = () => { }
</script>
<style scoped lang="less">
.archive-box {
  @apply flex flex-col h-full;
}

.archive-box-search {
  @apply flex gap-4 items-center;
}

.archive-search-input {
  @apply flex items-center gap-2;
}
</style>