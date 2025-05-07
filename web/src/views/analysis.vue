<template>
    <div class="page-okr-analysis" ref="pageOkrAnalysisRef">
        <div class="h-full flex flex-col">
            <div class="page-title">
                <div class="flex items-center">
                    <div v-if="isPortrait" class="okr-nav-back" @click="handleReturn"><i class="okrfont">&#xe676;</i></div>
                    <h2>{{ pageTitle }}</h2>
                    <div class="okr-app-refresh" v-if="!loadIng" @click="getData"><i class="okrfont">&#xe6ae;</i></div>
                </div>
                <n-tooltip v-if="isMainElectron" trigger="hover">
                    <template #trigger>
                        <div class="open-new-win" type="tertiary" @click="openNewWin">
                            <i class="okrfont open">&#xe776;</i>
                        </div>
                    </template>
                    <p class="max-w-[300px]"> {{ $t('新窗口打开') }}</p>
               </n-tooltip>
            </div>
            <n-tabs v-if="(departments.length) > 1" class="ml-[2px] mt-[-10px]" type="line" animated :value="tabsValue" :on-update:value="(e)=>{ tabsValue = e }">
                <n-tab-pane v-for="item in departments" :name="item.id" :tab="item.id == 0 ? $t(item.name) : item.name"></n-tab-pane>
            </n-tabs>
            <div v-if="!deptLoadIng" class="flex overflow-hidden">
                <n-scrollbar>
                    <n-grid x-gap="24" cols="1 600:2 800:3">

                        <!-- 整体平均完成度 -->
                        <n-gi class="bg-white mb-12 md:mb-20">
                            <n-spin :show="loadIng" size="small">
                                <div class="list-body">
                                    <div class="echarts-pie">
                                        <div class="text-16 font-medium">{{ $t('整体平均完成度') }}</div>
                                        <div class="pie">
                                            <div id="degreeOfCompletion"></div>
                                        </div>
                                        <div class="legend text-center">
                                            <span>
                                                <span class="legend-name">{{ $t('O的数量') }}: </span>
                                                <span class=" font-medium text-center text-12">{{ analyzeDatas.completes.total }}</span>
                                            </span>
                                            <span>
                                                <span class="legend-name">{{ $t('已完成O') }}:</span>
                                                <span class="font-medium text-12">{{ analyzeDatas.completes.complete }}</span>
                                            </span>
                                        </div>
                                    </div>
                                    <n-divider class="py-8" />
                                    <div class="list-progress">
                                        <div class="text-16 font-medium">
                                            {{ tabsValue == 0 ? $t('各部门平均完成度') : $t('各人员平均完成度') }}
                                        </div>
                                        <div class="text-14 text-center py-50"
                                            v-if="analyzeDatas.deptCompletes?.length == 0">{{ $t('暂无数据') }}</div>
                                        <div class="mt-20" v-else v-for="item in analyzeDatas.deptCompletes">
                                            <p class="progress-name max-w-[calc(100%-50px)] line-clamp-1">
                                                {{ item.department_name }}</p>
                                            <n-progress type="line" color="#8BCF70"
                                                :percentage="calculatingProgress(item.complete, item.total)">
                                                <span class="text-[#8BCF70] w-[50px] block text-right">{{calculatingProgress(item.complete, item.total) }}%</span>
                                            </n-progress>
                                        </div>
                                    </div>
                                </div>
                            </n-spin>
                        </n-gi>

                        <!-- OKR评分分布 -->
                        <n-gi class="bg-white mb-12 md:mb-20">
                            <n-spin :show="loadIng" size="small">
                                <div class="list-body">
                                    <div class="echarts-pie">
                                        <div class="text-16 font-medium">{{ $t('评分分布') }}</div>
                                        <div class="pie">
                                            <div id="scoreDistribution"></div>
                                        </div>
                                        <div class="legend flex justify-center flex-wrap" >
                                            <span class="flex justify-center items-center float-left mb-10">
                                                <p class="dot"></p>
                                                <span class="legend-name">{{ $t('未评分') }}: </span>
                                                <span class="font-medium text-12">{{ analyzeDatas.score.unscored }}</span>
                                            </span>
                                            <span class="flex justify-center items-center float-left mb-10">
                                                <p class="dot ff"></p>
                                                <span class="legend-name">{{ $t('0-3分') }}: </span>
                                                <span class="font-medium text-12">{{ analyzeDatas.score.zero_to_three}}</span>
                                            </span>
                                            <span class="flex justify-center items-center float-left mb-10">
                                                <p class="dot fc"></p>
                                                <span class="legend-name">{{ $t('3-7分') }}: </span>
                                                <span class="font-medium text-12">{{ analyzeDatas.score.three_to_seven }}</span>
                                            </span>
                                            <span class="flex justify-center items-center float-left mb-10">
                                                <p class="dot bc"></p>
                                                <span class="legend-name">{{ $t('7-10分') }}: </span>
                                                <span class="font-medium text-12">{{ analyzeDatas.score.seven_to_ten }}</span>
                                            </span>
                                        </div>
                                    </div>
                                    <n-divider class="py-8" />
                                    <div class="list-progress">
                                        <div class="text-16 font-medium">
                                            {{ tabsValue == 0 ? $t('各部门评分分布') : $t('各人员评分分布') }}
                                        </div>
                                        <div class="text-14 text-center py-50" v-if="analyzeDatas.deptScores?.length == 0">
                                            {{ $t('暂无数据') }}</div>
                                        <div class="mt-20" v-else v-for="item in analyzeDatas.deptScores">
                                            <p class="text-text-li max-w-[calc(100%-50px)] line-clamp-1">
                                                {{ item.department_name }}</p>
                                            <div class="custom-progres">
                                                <div class="progres h-[16px] leading-4">
                                                    <p class="bg-[#FF7070]" v-if="item.zero_to_three"
                                                        :style="{ width: calculatingProgress(item.zero_to_three, item.total) + '%' }">
                                                        {{ item.zero_to_three }}</p>
                                                    <p class="bg-[#FC984B]" v-if="item.three_to_seven"
                                                        :style="{ width: calculatingProgress(item.three_to_seven, item.total) + '%' }">
                                                        {{ item.three_to_seven }}</p>
                                                    <p class="bg-[#8BCF70]" v-if="item.seven_to_ten"
                                                        :style="{ width: calculatingProgress(item.seven_to_ten, item.total) + '%' }">
                                                        {{ item.seven_to_ten }}</p>
                                                    <p class="bg-[#E0E1E4]" v-if="item.unscored"
                                                        :style="{ width: calculatingProgress(item.unscored, item.total) + '%' }">
                                                        {{ item.unscored }}</p>
                                                </div>
                                                <div class="collect">{{ item.total }}{{ $t('个') }}</div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </n-spin>
                        </n-gi>

                        <!-- OKR评分率 -->
                        <n-gi class="bg-white mb-12 md:mb-20">
                            <n-spin :show="loadIng" size="small">
                                <div class="list-body">
                                    <div class="echarts-pie">
                                        <div class="text-16 flex font-medium">{{ $t('评分率') }}
                                            <n-tooltip trigger="hover" :width="widthWindow < 768 ? 200 : 300">
                                                <template #trigger>
                                                    <img class="ml-8 w-15" :src="utils.apiUrl(tipsSvgfrom)" />
                                                </template>
                                                <p class="max-w-[300px]">
                                                    {{ $t('已完成评分的 OKR 所占比例，一个 OKR 里负责人与上级都完成评分，才能计为完成评分的 OKR') }}</p>
                                            </n-tooltip>
                                        </div>
                                        <div class="pie">
                                            <div id="scoreRatingRate"></div>
                                        </div>
                                        <div class="legend text-center">
                                            <span>
                                                <p class="dot"></p>
                                                <span class="legend-name">{{ $t('未完成') }}: </span>
                                                <span class="font-medium">{{ analyzeDatas.scoreRate.total - analyzeDatas.scoreRate.complete }}</span>
                                            </span>
                                            <span>
                                                <p class="dot bc"></p>
                                                <span class="legend-name">{{ $t('已完成') }}:</span>
                                                <span class="font-medium">{{ analyzeDatas.scoreRate.complete }}</span>
                                            </span>
                                        </div>
                                    </div>

                                    <n-divider class="py-8" />
                                    <div class="list-progress">
                                        <div class="text-16 font-medium flex">
                                            {{ tabsValue == 0 ? $t('部门评分占比') : $t('个人评分占比') }}
                                            <n-tooltip trigger="hover">
                                                <template #trigger>
                                                    <img class="ml-8 w-15" :src="utils.apiUrl(tipsSvgfrom)" />
                                                </template>
                                                {{ tabsValue == 0 ? $t('各个部门完成 OKR 评分的所占比例') : $t('各个人员完成 OKR 评分的所占比例') }}
                                            </n-tooltip>
                                        </div>
                                        <div class="text-14 text-center py-50"
                                            v-if="analyzeDatas.deptScoreProportion?.length == 0">{{ $t('暂无数据') }}</div>
                                        <div class="mt-20" v-else v-for="item in analyzeDatas.deptScoreProportion">
                                            <p class="text-text-li max-w-[calc(100%-50px)] line-clamp-1">
                                                {{ item.department_name }}</p>
                                            <div class="custom-progres">
                                                <div class="progres h-[16px] leading-4">
                                                    <p class="bg-[#8BCF70]" v-if="item.already_reviewed"
                                                        :style="{ width: calculatingProgress(item.already_reviewed, item.total) + '%' }">
                                                        {{ item.already_reviewed }}</p>
                                                    <p class="bg-[#E0E1E4]" v-if="item.unscored"
                                                        :style="{ width: calculatingProgress(item.unscored, item.total) + '%' }">
                                                        {{ item.unscored }}</p>
                                                </div>
                                                <div class="collect">{{ item.total }}{{ $t('个') }}</div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </n-spin>
                        </n-gi>

                    </n-grid>
                </n-scrollbar>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import * as echarts from 'echarts';
import * as http from "@/api/modules/analysis";
import utils from '@/utils/utils';
import tipsSvgfrom from '@/assets/images/icon/tips.svg';
import { getAppData, handleBackApp } from "@/utils/app"

const pageTitle = 'OKR ' + $t('结果分析')
const deptLoadIng = ref(false)
const loadIng = ref(false)
const tabsValue = ref<any>(null)
const departments = ref<any>([])
const pageOkrAnalysisRef = ref(null)

const isMainElectron = computed(() => getAppData('initialData.isMainElectron') ? 1 : 0)
const isPortrait = computed(() => getAppData('instance.store.state.windowPortrait') ? 1 : 0)

const handleReturn = () => {
    handleBackApp()
}

// 总数据
const analyzeDatas = ref({
    completes: {
        total: 0,
        complete: 0
    },
    deptCompletes: [],
    score: {
        total: 0,
        unscored: 0,
        zero_to_three: 0,
        three_to_seven: 0,
        seven_to_ten: 0
    },
    deptScores: [],
    scoreRate: {
        total: 0,
        complete: 0
    },
    deptScoreProportion: [],
})

// 计算进度
const getDepartment = () => {
    deptLoadIng.value = true;
    http.getAnalyzeDepartment().then(res=>{
        departments.value = res.data?.sort((a, b) => a.id - b.id);
        tabsValue.value = departments.value[0].id;
    }).finally(()=> {
        deptLoadIng.value = false;
    })
}

// 计算进度
const calculatingProgress = (complete: number, total: number) => {
    return parseFloat((complete / (total || 1) * 100).toFixed(2))
}

//计算屏幕宽度
const widthWindow = computed(() => {
    return window.innerWidth
})

// 整体平均完成度
const completeEcharts = ref(null);
const loadComplete = () => {
    if(!document.getElementById('degreeOfCompletion')){
        return;
    }
    let data = analyzeDatas.value.completes;
    completeEcharts.value = completeEcharts.value || echarts.init(document.getElementById('degreeOfCompletion'));
    completeEcharts.value.setOption({
        series: [{
            type: 'pie',
            radius: ['65%', '90%'],
            avoidLabelOverlap: false,
            label: {
                position: 'center',
                formatter: () => {
                    return '{p|' + calculatingProgress(data.complete, data.total) + '}' + '{span|' + '%' + '}';
                },
                rich: {
                    p: {
                        fontSize: 22,
                        fontWeight: 'bold',
                    },
                    span: {
                        color: '#515A6E',
                        fontSize: 16,
                        fontWeight: '400'
                    }
                },
            },
            color: ['#8BCF70', '#bcbfca40'],
            data: [
                { value: data.complete, name: $t('已完成') },
                { value: data.total < 1 ? 1 : data.total - data.complete, name: $t('未完成') },
            ]
        }]
    });
}

// OKR评分分布
const scoreDistributeEcharts = ref(null);
const loadScoreDistribute = () => {
    if(!document.getElementById('scoreDistribution')){
        return;
    }
    let data = analyzeDatas.value.score
    scoreDistributeEcharts.value = scoreDistributeEcharts.value || echarts.init(document.getElementById('scoreDistribution'));
    scoreDistributeEcharts.value.setOption({
        series: [{
            type: 'pie',
            radius: ['65%', '90%'],
            avoidLabelOverlap: false,
            label: {
                position: 'center',
                formatter: () => {
                    return '{p|' + data.total + '}' + '\n {span|OKR ' + $t('数量') + '}';
                },
                rich: {
                    p: {
                        fontSize: 22,
                        fontWeight: 'bold',
                        textAlign:'center',
                    },
                    span: {
                        color: '#515A6E',
                        fontSize: 12,
                        opacity: 0.5,
                        textAlign:'center',
                    }
                },
                fontSize: 20,
                fontWeight: 'bold'
            },
            color: ['#8BCF70', '#FC984B', '#FF7070', '#bcbfca40'],
            data: [
                { value: data.seven_to_ten, name: '7-10分' },
                { value: data.three_to_seven, name: '3-7分' },
                { value: data.zero_to_three, name: '0-3分' },
                { value: data.total < 1 ? 1 : data.unscored, name: '未评分' },
            ]
        }]
    });
}

// 整体平均完成度
const scoreScoreRate = ref(null);
const loadScoreRate = () => {
    if(!document.getElementById('scoreRatingRate')){
        return;
    }
    let data = analyzeDatas.value.scoreRate
    scoreScoreRate.value = scoreScoreRate.value || echarts.init(document.getElementById('scoreRatingRate'));
    scoreScoreRate.value.setOption({
        series: [{
            type: 'pie',
            radius: ['65%', '90%'],
            avoidLabelOverlap: false,
            label: {
                position: 'center',
                formatter: () => {
                    return '{p|' + calculatingProgress(data.complete, data.total) + '}' + '{span|' + '%' + '}';
                },
                rich: {
                    p: {
                        fontSize: 22,
                        fontWeight: 'bold',
                    },
                    span: {
                        color: '#515A6E',
                        fontSize: 16,
                        fontWeight: '400'
                    }
                },
            },
            color: ['#8BCF70', '#bcbfca40'],
            data: [
                { value: data.complete, name: $t('已完成') },
                { value: data.total < 1 ? 1 : data.total - data.complete, name: $t('未完成') },
            ]
        }]
    });
}

// 获取数据
const getData = () => {
    loadIng.value = true;
    const params = {
        department: tabsValue.value
    }
    // 整体平均完成度
    http.getAnalyzeComplete(params).then(({ data }) => {
        analyzeDatas.value.completes = data
        loadComplete()
    })
    http.getAnalyzeDeptComplete(params).then(({ data }) => {
        analyzeDatas.value.deptCompletes = data
    })
    // OKR评分分布
    http.getAnalyzeScore(params).then(({ data }) => {
        analyzeDatas.value.score = data
        loadScoreDistribute()
    })
    http.getAnalyzeDeptScore(params).then(({ data }) => {
        analyzeDatas.value.deptScores = data
    })
    // OKR人员评分率
    http.getAnalyzeScoreSate(params).then(({ data }) => {
        analyzeDatas.value.scoreRate = data
        loadScoreRate()
    })
    http.getAnalyzeDeptScoreProportion(params).then(({ data }) => {
        analyzeDatas.value.deptScoreProportion = data
    })
    setTimeout(() => {
        loadIng.value = false;
    }, 300)
}

// 新窗口打开
const openNewWin = () => {
    const param = {
        name: `okr`,
        path: `/single/apps/okr/analysis`,
        force: false,
        config: {
            title: pageTitle,
            titleFixed: true,
            parent: null,
            width: Math.min(window.screen.availWidth, pageOkrAnalysisRef.value.clientWidth + 72),
            height: Math.min(window.screen.availHeight, pageOkrAnalysisRef.value.clientHeight + 36),
            minWidth: 600,
            minHeight: 450,
        }
    }
    //
    getAppData('openChildWindow')?.(param)
}


//
watch(tabsValue, (value) => {
    loadComplete()
    loadScoreDistribute();
    loadScoreRate();
    getData()
})

// 加载
nextTick(() => {
    getDepartment()
})
</script>

<style lang="less" scoped>
.page-okr-analysis {
    @apply p-20 md:px-24 h-full w-full bg-page-bg box-border;

    .page-title {
        @apply h-40 flex mb-16 md:mb-24 text-title-color font-medium justify-between;

        .open-new-win .okrfont{
            font-size: 24px;
            cursor: pointer;
        }

        h2 {
            @apply text-28;
        }
    }
    .list-body {
        @apply p-24;

        .echarts-pie {
            .pie {
                height: 150px;
                width: 150px;
                @apply m-auto mt-20 mb-16;

                >div {
                    @apply h-full w-full;
                }
            }

            .legend {
                height: 28px;

                >div {
                    @apply whitespace-nowrap;
                }

                >span,
                >div>span {
                    @apply mr-20;
                }

                span.legend-name {
                    @apply text-text-li opacity-50 text-12 mr-5;
                }

                .dot {
                    @apply bg-[#bcbfca40] mr-3 inline-block w-10 h-10 rounded-full;

                    &.ff {
                        @apply bg-[#FF7070];
                    }

                    &.fc {
                        @apply bg-[#FC984B];
                    }

                    &.bc {
                        @apply bg-[#8BCF70];
                    }
                }
            }
        }

        .list-progress {
            .progress-name {
                @apply text-text-li;
            }

            .custom-progres {
                @apply flex text-center mt-5;

                .progres {
                    @apply flex flex-1 text-12 text-right rounded overflow-hidden bg-[#bcbfca40] text-[#FFFFFF];

                    >p {
                        @apply whitespace-nowrap px-5;
                    }
                }

                .collect {
                    @apply w-50 text-[#8F8F8E] text-right;
                }
            }
        }
    }
}

.okr-theme-dark {
    .progres {
        @apply text-[#505050] !important;
    }
}

.nav-top {
    @apply flex md:hidden items-center justify-between absolute top-0 left-0 right-0 px-16;

    .icon-return {
        @apply text-20 text-text-tips;
    }
}
</style>
