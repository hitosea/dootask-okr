<template>
    <n-data-table :columns="columns" :data="data" :single-line="false" />
    <div class="replay">
        <div class="replay-title">{{$t('回顾')}}</div>
        <div>{{ review }}</div>
    </div>
</template>

<script lang="ts" setup>
import {DataTableColumn } from 'naive-ui';
const props = defineProps(['okrReplayList']);

// 回顾
let review = ref('')

// 加载表格
const loadTable = (krReplay)=> {
    review = krReplay.review
    data.value = krReplay.krList.map((itemData, index) => {
    const additionalProps = index === 0 ? { targetTitle: krReplay.okrName,length: krReplay.krList.length,completeness: krReplay.okrProgress} : {};
    return {
      id: itemData.id,
      krList: itemData.kr_history,
      krTitle: itemData.title,
      krCompleteness: itemData.progress,
      score: itemData.score,
      evaluate: itemData.comment,
      ...additionalProps,
    };
  });
}

nextTick(()=>{
    loadTable(props.okrReplayList)
    
})

const data = ref([])

//表格定义
const columns = ref<DataTableColumn[]>([
    {
        key: "targetTitle",
        rowSpan: (rowData) => rowData.length,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                        fontWeight: "bold",
                    },
                    size: 24,
                },
                { default: () => $t('目标（O）') },
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
        rowSpan: (rowData) => rowData.length,
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                        fontWeight: "bold",
                    },
                    size: 24,
                },
                { default: () => $t('O完成度') },
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
                { default: () => row.completeness+'%' },
            )
        },
    },
    {
        key: "krTitle",
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                        fontWeight: "bold",
                    },
                    size: 24,
                },
                { default: () => $t('关键KR') },
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
                            return "KR" + row.id
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
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                        fontWeight: "bold",
                    },
                    size: 24,
                },
                { default: () => $t('KR完成度') },
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
                { default: () => row.krCompleteness+'%' },
            )
        },
    },
    {
        key: "score",
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                        fontWeight: "bold",
                    },
                    size: 24,
                },
                { default: () => $t('KR评分') },
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
        title() {
            return h(
                "p",
                {
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                        fontWeight: "bold",
                    },
                    size: 24,
                },
                { default: () => "评价" },
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

<style lang="less" scope>
.replay {
    &-title {
        @apply text-18 text-[#515A6E];
    }
    &-head {
        @apply mt-10 text-24 text-[#333333];
    }
    &-text {
        @apply mt-10 leading-7 text-14 text-[#333333];
    }
}
</style>
