<template >
    <n-drawer v-model:show="show" style="max-width: 948px;width: 90%;" :on-after-leave="closeDrawer" :z-index="13" :mask-closable="false"
        :trap-focus="false" :on-update-show="()=>{emit('close') }" class="okr">
        <n-drawer-content :title="$t('复盘')" closable>
            <div class="flex flex-col h-full">
                <n-scrollbar class="h-full flex-auto">
                    <div class="flex flex-col shrink-0 h-full">
                        <h3 class="text-text-li text-18 font-normal mb-16">OKR</h3>
                        <n-data-table :columns="columns" :data="tableData" :single-line="false" :hover="false" size="small"
                            style="--n-td-color-hover-modal:#ffffff" />
                        <h3 class="text-text-li text-18 font-normal mb-16 mt-24">{{ $t('回顾') }}</h3>
                        <div class="flex-auto shrink-0 min-h-[250px]">
                            <TEditor v-if="props.multipleId == 0 " v-model:value="editorContent" :readOnly="false"></TEditor>
                            <div v-else v-html="editorContent"></div>
                        </div>
                    </div>
                </n-scrollbar>
                <div class="button-box">
                    <n-button  v-if="props.multipleId == 0 " :loading="loadIng" type="primary" @click="handleSubmit">
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
import { replayCreate, replayDetail } from '@/api/modules/okrList'
import { ResultDialog } from "@/api"
import { useMessage } from "naive-ui"

const message = useMessage()
const show = ref(false)
const loadIng = ref(false)
const editorContent = ref(``)
// const editorContent = ref(`<p><span style="font-size: 24pt;"><strong>价值与收获</strong></span></p> <p>&nbsp;</p> <p><span style="font-size: 24pt;"><span style="font-size: 32px;"><strong>问题与不足</strong></span></span></p> <p>&nbsp;</p>`)

const props = defineProps({
    data: {
        type: Object,
        default: () => { },
    },
    multipleId: {
        type: Number,
        default: 0,
    },
})


const tableData = ref<any>([])

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
        key: 'KR',
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
                placeholder: $t('请选择评价'),
                value: rowData.evaluate,
                options: itemOptions.value,
                disabled: props.multipleId > 0,
                onUpdateValue: (e: any) => {
                    rowData.evaluate = e
                },
            })
        }
    }
])

const itemOptions = ref([
    {
        label: $t('做得好的'),
        value: 1,
    },
    {
        label: $t('可提升的'),
        value: 2,
    },
])


const emit = defineEmits(['close'])

const closeDrawer = () => {
    tableData.value = []
    editorContent.value = ``
    // editorContent.value = `<p><span style="font-size: 24pt;"><strong>价值与收获</strong></span></p> <p>&nbsp;</p> <p><span style="font-size: 24pt;"><span style="font-size: 32px;"><strong>问题与不足</strong></span></span></p> <p>&nbsp;</p>`
}

const handleSubmit = () => {
    const upData = {
        okr_id: props.data.id,
        review: editorContent.value,
        comments: [],
    }
    tableData.value.map((item, index) => {
        upData.comments.push({
            okr_id: props.data.key_results[index].id,
            comment: item.evaluate,
        })
    })

    loadIng.value = true
    replayCreate(upData)
        .then(({ msg }) => {
            message.success(msg)
            emit('close')
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })

}

watch(() => props.data, (newValue) => {
    if (newValue) {
        tableData.value.push({
            O: newValue.title,
            lenght: newValue.key_results.length,
            Ocomplete: newValue.progress + '%',
            KR: newValue.key_results[0].title,
            KRcomplete: newValue.key_results[0].progress + '%',
            KRMark: newValue.key_results[0].kr_score,
            evaluate: null,
        })
        for (let index = 1; index < newValue.key_results.length; index++) {
            tableData.value.push({
                KR: newValue.key_results[index].title,
                KRcomplete: newValue.key_results[index].progress + '%',
                KRMark: newValue.key_results[index].kr_score,
                evaluate: null,
            })
        }
    }
}, { immediate: true })

//查看详情
watch(() => props.multipleId, (newValue) => {
    if (newValue) {
        const upData = {
            id: newValue,
        }
        loadIng.value = true
        replayDetail(upData)
            .then(({ data  }) => {
                tableData.value.push({
                    O: data.okr_title,
                    lenght: data.kr_history.length,
                    Ocomplete: data.okr_progress + '%',
                    KR: data.kr_history[0].title,
                    KRcomplete: data.kr_history[0].progress + '%',
                    KRMark: data.kr_history[0].score,
                    evaluate: data.kr_history[0].comment,
                })
                for (let index = 1; index < data.kr_history.length; index++) {
                    tableData.value.push({
                        KR: data.kr_history[index].title,
                        KRcomplete: data.kr_history[index].progress + '%',
                        KRMark: data.kr_history[index].score,
                        evaluate: data.kr_history[index].comment,
                    })
                }
                editorContent.value = data.review
            })
            .catch(ResultDialog)
            .finally(() => {
                loadIng.value = false
            })

    }
})

</script>
<style lang="less" scoped>
.button-box {
    @apply flex gap-2 mt-32 flex-initial;
}

:deep(.n-scrollbar-content) {
    @apply h-full;
}
:deep(.n-drawer-header__close) {
    @apply absolute -left-36 ;
    &:focus{
        @apply bg-none;
    }
    i {
        @apply text-[#fff];
    }
}

:deep(.n-base-close:not(.n-base-close--disabled):focus::before) {
        @apply bg-transparent;
}
</style>
