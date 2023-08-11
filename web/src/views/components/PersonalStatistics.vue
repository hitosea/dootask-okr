<template >
    <div class="personal-statistics">
        <div class="p-s-box p-s-one min-w-[350px]">
            <div class="p-s-one-box mr-24 pr-24 border-solid border-0 border-r border-[#F2F3F5]">
                <div class="p-s-one-title">
                    <h3>{{ $t('待完成数') }}</h3>
                    <div class="p-title">
                        <span>{{ analyzeDatas.okrCompletedAndUncompleted.uncompleted }}</span>
                        <div class="p-s-one-icon bg-[rgba(114,161,247,0.2)] shrink-0  flex xl:hidden">
                            <i class="taskfont text-[#72A1F7]">&#xe6ff;</i>
                        </div>
                    </div>
                </div>
                <div class="p-s-one-icon bg-[rgba(114,161,247,0.2)] shrink-0  hidden xl:flex">
                    <i class="taskfont text-[#72A1F7]">&#xe6ff;</i>
                </div>
            </div>
            <div class="p-s-one-box">
                <div class="p-s-one-title">
                    <h3>{{ $t('已完成/已取消数') }}</h3>
                    <div class="p-title ">
                        <span>{{ analyzeDatas.okrCompletedAndUncompleted.completed }}</span>
                        <div class=" p-s-one-icon p-s-one-icon-2 bg-[rgba(135,208,104,0.2)] shrink-0 flex xl:hidden">
                            <i class="taskfont text-[#87D068]">&#xe707;</i>
                        </div>
                    </div>
                </div>
                <div class=" p-s-one-icon p-s-one-icon-2 bg-[rgba(135,208,104,0.2)] shrink-0 hidden xl:flex">
                    <i class="taskfont text-[#87D068]">&#xe707;</i>
                </div>
            </div>
        </div>
        <div class="sm:flex gap-3 xl:gap-6">
            <div class="p-s-box p-s-two mb-12 sm:mb-0">
                <div class="percent-box">
                    <h3>{{ $t('整体完成度') }}</h3>
                    <p>{{ $t('这是你的目标整体完成度。') }}</p>
                </div>
                <div class="relative w-[150px] h-[67px]">
                    <div class="h-full w-[150px]" id="okrCompletedAndUncompleted"></div>
                    <h3 class=" absolute text-[20px] text-title-color text-center leading-5 left-0 right-0 bottom-0">
                        {{ calculating(analyzeDatas.okrDegreeOfCompletionAndScore.completion_sum,
                            analyzeDatas.okrDegreeOfCompletionAndScore.total) }}%</h3>
                </div>
            </div>
            <div class="p-s-box p-s-two">
                <div class="percent-box">
                    <h3>{{ $t('整体评分') }}</h3>
                    <p>{{ $t('这是你的目标整体获得评分。') }}</p>
                </div>
                <div class="relative w-[150px] h-[67px]">
                    <div class="h-full w-[150px]" id="okrDegreeOfCompletionAndScore"></div>
                    <h3 class=" absolute text-[20px] text-title-color text-center leading-5 left-0 right-0 bottom-0">
                        {{ calculating(analyzeDatas.okrDegreeOfCompletionAndScore.score_sum,
                            analyzeDatas.okrDegreeOfCompletionAndScore.score_total, 1) }}</h3>
                </div>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import * as echarts from 'echarts';
import * as http from '../../api/modules/icreated';

// 总提数据
const analyzeDatas = ref({
    okrCompletedAndUncompleted: {
        uncompleted: 0,
        completed: 0,
    },
    okrDegreeOfCompletionAndScore: {
        completion_sum: 0,
        score_sum: 0,
        total: 0,
        score_total: 0,
    }
})



// 计算完成度
const calculating = (complete: number, total: number, precision = 0) => {
    return parseFloat((complete / (total || 1)).toFixed(precision))
}

// 获取数据
const getData = () => {
    // 我创建的OKR整体完成度
    http.getNumberofOkrIcreatedPending().then(({ data }) => {
        analyzeDatas.value.okrCompletedAndUncompleted = data
    })
    http.getOkrDegreeOfCompletionAndScore().then(({ data }) => {
        analyzeDatas.value.okrDegreeOfCompletionAndScore = data
        loadComplete()
    })
}

// 我创建的OKR整体完成度
const completeEcharts = ref(null);
const scoreEcharts = ref(null);
const loadComplete = () => {
    let data = analyzeDatas.value.okrDegreeOfCompletionAndScore

    // 整体完成度饼图
    const completeness = data.completion_sum / (data.total || 1)
    completeEcharts.value = completeEcharts.value || echarts.init(document.getElementById('okrCompletedAndUncompleted'));
    completeEcharts.value.setOption({
        tooltip: {
            show: false,
        },
        grid: {
            x: 90,
            y: 0
        },
        series: [
            {
                name: 'Access From',
                type: 'pie',
                radius: ['150%', '200%'],
                center: ['50%', '100%'],
                label: {
                    show: false,
                    position: 'center'
                },
                silent: true,
                startAngle: 180,
                color: ['#8BCF70', 'rgba(188, 191, 202, 0.25)'],
                data: [
                    { value: completeness, name: 'Search Engine' },
                    { value: 100 - completeness, name: 'Direct' },
                    {
                        value: 100,
                        itemStyle: {
                            color: 'none',
                            decal: {
                                symbol: 'none'
                            }
                        },
                        label: {
                            show: false
                        }
                    }
                ]
            }
        ]
    });
    // 整体评分饼图
    const averageScore = data.score_sum / (data.score_total || 1)
    scoreEcharts.value = scoreEcharts.value || echarts.init(document.getElementById('okrDegreeOfCompletionAndScore'));
    scoreEcharts.value.setOption({
        tooltip: {
            show: false,
        },
        grid: {
            x: 90,
            y: 0
        },
        series: [
            {
                name: 'Access From',
                type: 'pie',
                radius: ['150%', '200%'],
                center: ['50%', '100%'],
                label: {
                    show: false,
                    position: 'center'
                },
                silent: true,
                startAngle: 180,
                color: ['#8BCF70', 'rgba(188, 191, 202, 0.25)'],
                data: [
                    { value: averageScore, name: 'Search Engine' },
                    { value: 10 - averageScore, name: 'Direct' },
                    {
                        value: 10,
                        itemStyle: {
                            color: 'none',
                            decal: {
                                symbol: 'none'
                            }
                        },
                        label: {
                            show: false
                        }
                    }
                ]
            }
        ]
    });
}

defineExpose({
    getData
})

</script>
<style lang="less" scoped>
.personal-statistics {
    @apply flex-col xl:flex-row flex gap-3 xl:gap-6;

    .p-s-box {
        @apply flex-1 bg-white rounded-lg p-16 md:px-24 md:py-30 flex;
    }

    .p-s-one {
        @apply items-center;

        .p-s-one-box {
            @apply flex-1 flex justify-between items-center;

            .p-s-one-title {
                @apply flex flex-col flex-1;

                h3 {
                    @apply text-12 text-text-tips font-normal;
                }

                .p-title {
                    @apply text-30 flex flex-1 justify-between;
                }
            }

            .p-s-one-icon {
                @apply w-56 h-56 rounded-full items-center justify-center;

                i {
                    @apply text-30;
                }
            }
        }
    }

    .p-s-two {
        @apply justify-between items-center;

        .percent-box {
            @apply flex flex-col gap-3;

            h3 {
                @apply text-title-color text-18;
            }

            p {
                @apply text-14 text-text-tips;
            }
        }
    }
}
</style>
