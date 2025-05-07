<template>
    <div class=" hidden md:block">
        <n-data-table  :columns="columns" :data="tableData" :single-line="false" :hover="false"
            style="--n-merged-td-color-hover:#ffffff" size="large"/>
    </div>
    <div class="flex flex-col md:gap-[16px] gap-[12px] mt-16 md:mt-[24px]" v-if="props.okrReplayList.replays">
        <div class="bg-[#F4F5F7] p-[16px] rounded-lg cursor-pointer" v-for="(item, index) in props.okrReplayList.replays"
            @click="openActive(index)">
            <div class="flex items-center justify-between">
                <h3 class="text-text-li text-16">{{ item.okr_title }} ({{ utils.GoDateHMS(item.created_at, '/') }})</h3>
                <img class="ml-8 w-15" :src="index == active ? utils.apiUrl(packUp) : utils.apiUrl(unfold)" />
            </div>
            <template v-if="index == active">
                <div class="bg-[#fff] border-[#F2F2F2] border-[1px] rounded-lg overflow-hidden mt-12 md:mt-24">
                    <div class="flex items-center  px-[16px] text-text-li text-14"
                        :class="(indexs + 1) < (item.kr_history || []).length ? 'border-solid border-0 border-b-[1px] border-[#F2F2F2]' : ''"
                        v-for="(items, indexs) in item.kr_history">

                        <div class="flex-[7] py-[14px] hidden md:block">
                            <span class="text-primary-color text-12 mr-4">KR{{ indexs + 1 }}</span> {{ items.title }}
                        </div>
                        <div class="flex-[3] py-[14px] md:pl-16 border-solid border-0  md:border-l-[1px] border-[#F2F2F2]">
                            <span class="inline-block md:hidden text-primary-color text-12 mr-4">KR{{ indexs + 1 }}</span>
                            <span class="inline-block md:hidden font-medium">{{ $t('评价：') }}</span>
                            {{ items.evaluate == 1 ? $t('做得好的') : $t('可提升的') }}
                        </div>
                    </div>
                </div>
                <div class="replay-details">
                    <h3 class="md:mb-12 mb-6 text-text-li md:text-20 text-18 font-medium flex justify-between items-center ">{{ $t('价值与收获') }}
                    </h3>
                    <p class="text-text-li text-14" v-html="item.review"></p>
                    <h3 class="md:mt-24 mt-20 mb-12 text-text-li md:text-20 text-18 font-medium flex justify-between items-center ">{{
                        $t('问题与不足')
                    }}
                    </h3>
                    <p class="text-text-li text-14" v-html="item.problem"></p>
                </div>
            </template>
        </div>
    </div>
    <div class="flex mt-[24px] justify-center flex-col items-center" v-else>
        <img class="w-[120px]" :src="utils.apiUrl(notDataSvg)" />
        <p class="text-[14px] text-text-tips mt-[16px]">{{ $t('暂无数据') }}</p>
        <n-button class="mt-16 px-16" type="primary" ghost @click="handleAddMultiple">
            <i class="okrfont mr-5">&#xe731;</i>
            {{ $t('添加复盘') }}
        </n-button>
    </div>
    <!-- 强提示 -->
    <TipsModal :show="showModal" :content="tipsContent" @close="() => { showModal = false }"></TipsModal>
</template>

<script lang="ts" setup>
import { DataTableColumn } from "naive-ui"
import utils from '@/utils/utils';
import packUp from '@/assets/images/icon/packUp.svg';
import unfold from '@/assets/images/icon/unfold.svg';
import notDataSvg from "@/assets/images/icon/notData.svg";
import { UserStore } from '@/store/user'
import { GlobalStore } from '@/store';

import TipsModal from '@/views/components/TipsModal.vue';

const userInfo = UserStore().info
const globalStore = GlobalStore()

const tableData = ref([])
const active = ref(-1)

const tipsContent = ref('')
const showModal = ref(false)

const props = defineProps({
    okrReplayList: {
        type: Object,
        default: {},
    }
}
)

watch(() => props.okrReplayList, (newValue) => {
    if (newValue) {
        tableData.value.push({
            O: newValue.okr_title,
            lenght: (newValue.key_results || []).length,
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
                    style: {
                        color: "#515A6E",
                        fontSize: "14px",
                    },
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
    if (active.value == index) {
        active.value = -1
    } else {
        active.value = index
    }

}

//添加复盘
const handleAddMultiple = () => {
    if (userInfo.userid != props.okrReplayList.key_results[0].userid) {
        tipsContent.value = $t('仅限负责人操作')
        showModal.value = true
        return
    }
    const data = {
        id:null,
        title:null,
        progress:null,
        key_results:[],
    }
    data.id = props.okrReplayList.okr_id
    data.title = props.okrReplayList.okr_title
    data.progress = props.okrReplayList.okr_progress
    data.key_results = props.okrReplayList.key_results

    globalStore.$patch((state) => {
        state.addMultipleData = data
        state.addMultipleShow = true
    })
}
</script>

<style lang="less" scoped>
.replay-details {
    @apply mt-12 md:mt-24;
}

</style>
