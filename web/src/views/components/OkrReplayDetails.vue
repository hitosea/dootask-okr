<template>
    <n-data-table :columns="columns" :data="tableData" :single-line="false" :hover="false" size="small" />
    <div class="replay-details">
        <div class="replay-details-title">{{ $t("回顾") }}</div>
        <h3 class="mb-8 text-text-li text-20 font-medium flex justify-between items-center ">{{ $t('价值与收获') }}</h3>
        <p class="text-text-li text-14" v-html="review"></p>
        <h3 class="mt-16 mb-8 text-text-li text-20 font-medium flex justify-between items-center ">{{ $t('问题与不足') }}</h3>
        <p class="text-text-li text-14" v-html="problem"></p>
    </div>
</template>

<script lang="ts" setup>
import { DataTableColumn } from "naive-ui"
const props = defineProps({
    okrReplayList: {
        type: Object,
        default: {},
    }
}
)
const tableData = ref([])

watch(() => props.okrReplayList, (newValue) => {
    if (newValue) {
        tableData.value.push({
            O: newValue.okr_title,
            lenght: newValue.key_results.length,
            Ocomplete: newValue.okr_progress + '%',
            KR: newValue.key_results[0].title,
            KRcomplete: newValue.key_results[0].progress + '%',
            KRMark: newValue.key_results[0].score,
            evaluate: null,
        })
        for (let index = 1; index < newValue.key_results.length; index++) {
            tableData.value.push({
                KR: newValue.key_results[index].title,
                KRcomplete: newValue.key_results[index].progress + '%',
                KRMark: newValue.key_results[index].score,
                evaluate: null,
            })
        }
    }

}, { immediate: true })

// 回顾
let review = ref("")
let problem = ref("")

//表格定义
const columns = ref<DataTableColumn[]>([
    {
        key: "O",
        minWidth: 200,
        rowSpan: (rowData) => rowData.lenght,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                    size: 24,
                },
                { default: () => $t("目标（O）") },
            )
        },
        render(rowData) {
            let arr = []
            arr.push(
                h('span', {
                    class: 'line-clamp-2'
                }, rowData.O)
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
        key: "Ocomplete",
        minWidth: 100,
        rowSpan: (rowData) => rowData.lenght,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                    size: 24,
                },
                { default: () => $t("O完成度") },
            )
        },
        render(row) {
            return h(
                "span",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                },
                { default: () => row.Ocomplete  },
            )
        },
    },
    {
        key: "KR",
        minWidth: 300,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                    size: 24,
                },
                { default: () => $t("关键KR") },
            )
        },
        render(row, index) {
            return h("div", [
                h(
                    "span",
                    {
                        style: {
                            color: "#8BCF70",
                            fontSize: "12px",
                        },
                        strong: true,
                        size: "small",
                    },
                    {
                        default: () => {
                            return "KR" + (index + 1)
                        },
                    },
                ),
                h(
                    "span",
                    {
                        style: {
                            marginLeft: "10px",
                            color: "#515A6E",
                            fontSize: "14px",
                        },
                        strong: true,
                        size: "small",
                    },
                    {
                        default: () => {
                            return row.KR
                        },
                    },
                ),
            ])
        },
    },
    {
        key: "KRcomplete",
        minWidth: 100,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                    size: 24,
                },
                { default: () => $t("KR完成度") },
            )
        },
        render(row) {
            return h(
                "span",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                },
                { default: () => row.KRcomplete },
            )
        },
    },
    {
        key: "KRMark",
        minWidth: 100,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                    size: 24,
                },
                { default: () => $t("KR评分") },
            )
        },
        render(row) {
            return h(
                "span",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                },
                { default: () => row.KRMark },
            )
        },
    },
])
</script>

<style lang="less" scoped>
.replay-details {
    &-title {
        @apply mt-10 py-12 text-18 text-[#515A6E] font-medium;
    }
}
</style>
