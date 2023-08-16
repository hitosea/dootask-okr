<template >
    <n-scrollbar class="h-full">
        <div class="flex flex-col shrink-0 h-full">
            <h3 class="text-text-li text-18 font-medium mb-16  hidden md:block">OKR</h3>
            <n-data-table class="hidden md:block" :columns="columns" :data="tableData" :single-line="false" :hover="false"
                size="small" style="--n-td-color-hover-modal:#ffffff" />
            <div class="flex flex-col md:hidden" v-if="props.data">
                <h3 class=" text-text-li text-14 font-normal flex justify-between items-center">{{ $t('目标（O）') }} <span
                        class="text-primary-color">{{ props.data.progress }}%</span></h3>
                <div
                    class="border-solid border-[1px] border-[#F2F2F2] rounded mt-12 p-16 text-15 text-title-color font-medium">
                    {{ props.data.title }}
                </div>
                <h3 class="mt-16 text-text-li text-14 font-medium flex justify-between items-center ">{{ $t('关键KR') }}</h3>
                <div v-for="(item, index) in tableData"
                    class="border-solid border-[1px] border-[#F2F2F2] rounded mt-12 p-16 ">
                    <h3 class="text-15 text-title-color font-medium"><span class="mr-4 text-12 text-text-tips">KR{{ index +
                        1 }}</span>{{ item.KR }}</h3>
                    <div class="mt-12 flex items-center">
                        <p class="text-12 text-text-li flex-1 border-solid border-0 border-r-[1px] border-[#F2F2F2]">{{
                            $t('KR完成度:') }}<span class=" text-primary-color">{{ item.KRcomplete }}</span></p>
                        <p
                            class="text-12 text-text-li flex-1 text-center border-solid border-0 border-r-[1px] border-[#F2F2F2]">
                            {{ $t('KR评分:') }}<span class=" text-primary-color">{{ item.KRMark }}</span></p>
                        <div class="flex-1 flex justify-end">
                            <n-select v-if="props.multipleId == 0" class="w-[90%]" :placeholder="$t('请选择评价')"
                                :options="itemOptions" :disabled="props.multipleId > 0"
                                v-model:value="tableData[index].evaluate"></n-select>
                            <p v-else class="text-12 text-text-li flex-1 text-center "><span
                                    class="">{{ tableData[index].evaluate == 1 ? $t('做得好的') : $t('可提升的') }}</span></p>
                        </div>
                    </div>
                </div>
            </div>
            <h3 class="text-text-li text-14 md:text-18 font-medium mt-24">{{ $t('回顾') }}</h3>
            <div class="flex-auto flex flex-col shrink-0 min-h-[250px]">
                <!-- <TEditor v-if="props.multipleId == 0" v-model:value="review" :readOnly="false"></TEditor>
                <div v-else v-html="review"></div> -->
                <h3 class="mt-12 text-text-li text-18 md:text-20 font-medium flex justify-between items-center ">{{ $t('价值与收获') }}</h3>
                <n-input v-if="props.multipleId == 0" class="mt-8" :rows="8" v-model:value="review" type="textarea"
                    maxlength="255" show-count :placeholder="$t('我们从过程中学到了什么新东西')" />
                <p class="mt-12" v-else v-html="review"></p>
                <h3 class="mt-12 text-text-li text-18 md:text-20 font-medium flex justify-between items-center ">{{ $t('问题与不足') }}</h3>
                <n-input v-if="props.multipleId == 0" class="mt-8" :rows="8" v-model:value="problem" type="textarea"
                    maxlength="255" show-count :placeholder="$t('请描述出现的某个问题并针对该问题展开分析')" />
                <p class="mt-12" v-else v-html="problem"></p>
            </div>
        </div>
    </n-scrollbar>
</template>
<script setup lang="ts">
import { NSelect, DataTableColumn } from 'naive-ui';
import { replayCreate, replayDetail } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"

const message = useMessage()
const loadIng = ref(false)
const review = ref(``)
const problem = ref(``)

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
        minWidth: 200,
        rowSpan: (rowData) => (rowData.lenght),
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
        title: $t('O完成度'),
        key: 'Ocomplete',
        minWidth: 100,
        rowSpan: (rowData) => (rowData.lenght),
    },
    {
        title: $t('关键KR'),
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
        title: $t('KR完成度'),
        minWidth: 100,
        key: 'KRcomplete',
    },
    {
        title: $t('KR评分'),
        minWidth: 100,
        key: 'KRMark',
    },
    {
        title: $t('评价'),
        key: 'evaluate',
        minWidth: 150,
        render(rowData) {
            if (props.multipleId == 0) {
                return h(NSelect, {
                    placeholder: $t('请选择评价'),
                    value: rowData.evaluate,
                    options: itemOptions.value,
                    disabled: props.multipleId > 0,
                    onUpdateValue: (e: any) => {
                        rowData.evaluate = e
                    },
                })
            }else{
                return h('p', rowData.evaluate == 1 ? $t('做得好的') : $t('可提升的'))
            }
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


const emit = defineEmits(['close', 'loadIng'])

const closeDrawer = () => {
    tableData.value = []
    review.value = ``
    problem.value = ``
}

const handleSubmit = () => {
    const upData = {
        okr_id: props.data.id,
        review: review.value,
        problem: problem.value,
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
            message.success($t('添加成功'))
            emit('close')
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })

}

watch(() => props.data, (newValue) => {
    if (newValue && props.multipleId == 0) {
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
            .then(({ data }) => {
                tableData.value.push({
                    O: data.okr_title,
                    lenght: data.kr_history.length,
                    Ocomplete: data.okr_progress + '%',
                    KR: data.kr_history[0].title,
                    KRcomplete: data.kr_history[0].progress + '%',
                    KRMark: data.kr_history[0].score,
                    evaluate: data.kr_history[0].comment,
                })
                if (data.kr_history.length > 1) {
                    for (let index = 1; index < data.kr_history.length; index++) {
                        tableData.value.push({
                            KR: data.kr_history[index].title,
                            KRcomplete: data.kr_history[index].progress + '%',
                            KRMark: data.kr_history[index].score,
                            evaluate: data.kr_history[index].comment,
                        })
                    }
                }
                review.value = data.review || $t('无')
                problem.value = data.problem || $t('无')
            })
            .catch(({ msg }) => {
                message.error(msg)
            })
            .finally(() => {
                loadIng.value = false
            })
    }
}, { immediate: true })

watch(() => loadIng.value, (newValue) => {
    emit('loadIng', newValue)
}, { immediate: true })

defineExpose({
    closeDrawer,
    handleSubmit
})

</script>
<style lang="less" scoped></style>
