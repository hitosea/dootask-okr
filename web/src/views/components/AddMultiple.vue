<template >
    <n-drawer v-model:show="show" :width="948" :on-after-leave="closeDrawer" :z-index="1" :trap-focus="false">
        <n-drawer-content :title="$t('复盘')" closable>
            <div class="flex flex-col absolute left-[24px] right-[24px] top-[16px] bottom-[16px]">
                <n-scrollbar class="h-full flex-auto">
                    <div class="flex flex-col shrink-0 h-full">
                        <h3 class="text-text-li text-18 font-normal mb-16">OKR</h3>
                        <n-data-table :columns="columns" :data="data" :single-line="false" :hover="false" size="small" style="--n-td-color-hover-modal:#ffffff"/>
                        <h3 class="text-text-li text-18 font-normal mb-16 mt-24">{{ $t('回顾') }}</h3>
                        <div class="flex-auto shrink-0 min-h-[250px]">
                            <TEditor></TEditor>
                        </div>
                    </div>
                </n-scrollbar>
                <div class="button-box">
                    <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                        {{ $t('提交') }}
                    </n-button>
                </div>
            </div>

        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import { NSelect, DataTableColumn } from 'naive-ui';
import TEditor from '@/components/TEditor.vue';
const show = ref(false)
const loadIng = ref(false)

const data = ref([
    {
        O: '这个是目标啊',
        lenght: 3,
        Ocomplete: 32,
        crux: '这个是关键KR',
        KR: 'New York No. 1 Lake Park',
        KRcomplete: '60%',
        KRMark: 60,
        evaluate: 2,
    },
    {
        crux: '这个是关键KR',
        KR: 'New York No. 1 Lake Park',
        KRcomplete: '60%',
        KRMark: 80,
        evaluate: 0,
    },
    {
        crux: '这个是关键KR',
        KR: 'New York No. 1 Lake Park',
        KRcomplete: '60%',
        KRMark: 80,
        evaluate: null,
    }
])

const columns = ref<DataTableColumn[]>([
    {
        title: $t('目标（O）'),
        key: 'O',
        rowSpan: (rowData) => (rowData.lenght),
    },
    {
        title: $t('O完成度'),
        key: 'Ocomplete',
        rowSpan: (rowData) => (rowData.lenght),
    },
    {
        title: $t('关键KR'),
        key: 'crux',
    },
    {
        title: $t('KR完成度'),
        key: 'KRcomplete',
    },
    {
        title: $t('KR评分'),
        key: 'KRMark',
    },
    {
        title: $t('评价'),
        key: 'evaluate',
        minWidth: 150,
        render(rowData) {
            return h(NSelect, {
                placeholder: $t('请选择项目'),
                value: rowData.evaluate,
                options: itemOptions.value,
                onUpdateValue: (e: any) => {
                    rowData.evaluate = e
                },
            })
        }
    }
])

const itemOptions = ref([
    {
        label: $t('项目1'),
        value: 0,
    },
    {
        label: $t('项目2'),
        value: 1,
    },
    {
        label: $t('项目3'),
        value: 2,
    },
])


const emit = defineEmits(['close'])

const closeDrawer = () => {

}


const handleSubmit = () => {
    emit('close')
}
</script>
<style lang="less" scoped>
.button-box {
    @apply flex gap-2 mt-32 flex-initial;
}
:deep(.n-scrollbar-content){
    @apply h-full;
}
</style>