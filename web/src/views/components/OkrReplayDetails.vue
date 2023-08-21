<template>
    <n-data-table :columns="columns" :data="data" :single-line="false" size="small" />
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
const props = defineProps(["okrReplayList"])

// 回顾
let review = ref("")
let problem = ref("")

// 加载表格
const loadTable = (krReplay) => {
    review = krReplay.review || $t('无')
    problem = krReplay.problem || $t('无')
    let count = 1
    data.value = krReplay.krList.map((itemData, index) => {
        const additionalProps =
            index === 0
                ? { targetTitle: krReplay.okrName, length: krReplay.krList.length, completeness: krReplay.okrProgress }
                : {}
        return {
            id: itemData.id, //目标id
            num: count++, //kr位置
            krList: itemData.kr_history, //kr列表
            krTitle: itemData.title, //标题
            krCompleteness: itemData.progress, //完成度
            score: itemData.score, //评分
            evaluate: itemData.comment == 1 ? $t("做得好的") : $t("可提升的"), //评价
            ...additionalProps,
        }
    })
}

nextTick(() => {
    loadTable(props.okrReplayList)
})

const data = ref([])

//表格定义
const columns = ref<DataTableColumn[]>([
    {
        key: "targetTitle",
        minWidth:'100px',
        rowSpan: (rowData) => rowData.length,
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
        render(row) {
            return h(
                "span",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
                },
                { default: () => row.targetTitle },
            )
        },
    },
    {
        key: "completeness",
        minWidth:'100px',
        rowSpan: (rowData) => rowData.length,
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
                { default: () => row.completeness + "%" },
            )
        },
    },
    {
        key: "krTitle",
        minWidth:'100px',
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
        render(row) {
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
                            return "KR" + row.num
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
                            return row.krTitle
                        },
                    },
                ),
            ])
        },
    },
    {
        key: "krCompleteness",
        minWidth:'100px',
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
                { default: () => row.krCompleteness + "%" },
            )
        },
    },
    {
        key: "score",
        minWidth:'100px',
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
                { default: () => row.score },
            )
        },
    },
    {
        key: "evaluate",
        minWidth:'100px',
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
                { default: () => $t("评价") },
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
                { default: () => row.evaluate },
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
