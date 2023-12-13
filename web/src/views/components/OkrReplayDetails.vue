<template>
    <n-data-table :columns="columns" :data="tableData" :single-line="false" :hover="false"
        style="--n-merged-td-color-hover:#ffffff" />
    <div class="flex flex-col gap-[16px] mt-[24px]">
        <div class="bg-[#F4F5F7] p-[16px] rounded-lg cursor-pointer" v-for="(item, index) in props.okrReplayList.replays"
            @click="openActive(index)">
            <div class="flex items-center justify-between">
                <h3 class="text-text-li text-16">{{ item.okr_title }}</h3>
                <img class="ml-8 w-15" :src="index == active ? utils.apiUrl(packUp) : utils.apiUrl(unfold)" />
            </div>
            <template v-if="index == active">
                <div class="bg-[#fff] border-[#F2F2F2] border-[1px] rounded-lg overflow-hidden mt-24">
                    <div class="flex items-center py-[14px] px-[16px] text-text-li text-14"
                        :class="(indexs + 1) < item.kr_history.length ? 'border-solid border-0 border-b-[1px] border-[#F2F2F2]' : ''"
                        v-for="(items, indexs) in item.kr_history">

                        <div class="flex-[7]">
                            <span class="text-primary-color text-12 mr-4">KR{{ indexs + 1 }}</span> {{ items.title }}
                        </div>
                        <div class="flex-[3]">

                        </div>
                    </div>
                </div>
                <div class="replay-details">
                    <h3 class="mb-12 text-text-li text-20 font-medium flex justify-between items-center ">{{ $t('价值与收获') }}
                    </h3>
                    <p class="text-text-li text-14" v-html="item.review"></p>
                    <h3 class="mt-24 mb-12 text-text-li text-20 font-medium flex justify-between items-center ">{{
                        $t('问题与不足')
                    }}
                    </h3>
                    <p class="text-text-li text-14" v-html="item.problem"></p>
                </div>
            </template>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { DataTableColumn } from "naive-ui"
import utils from '@/utils/utils';
import packUp from '@/assets/images/icon/packUp.svg';
import unfold from '@/assets/images/icon/unfold.svg';


const props = defineProps({
    okrReplayList: {
        type: Object,
        default: {},
    }
}
)
const tableData = ref([])
const active = ref(-1)


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
                { default: () => row.Ocomplete },
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

const openActive = (index) => {
    if(active.value == index){
        active.value = -1
    }else{
        active.value = index
    }

}

</script>

<style lang="less" scoped>.replay-details {
    @apply mt-24;
}</style>
