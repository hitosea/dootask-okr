<template >
    <n-drawer v-model:show="show" :width="728" :on-after-leave="closeDrawer">
        <n-drawer-content :title="$t('添加OKR')" closable>
            <div class="flex flex-col absolute left-[24px] right-[24px] top-[16px] bottom-[16px]">
                <n-scrollbar>
                    <div class="add-main-box">

                        <n-form ref="formRef" :model="formValue" :rules="rules" size="medium" label-placement="left"
                            label-width="auto" require-mark-placement="left">

                            <n-form-item class="w-full" :label="$t('目标（O）')" path="object">
                                <n-input v-model:value="formValue.object" :placeholder="$t('请输入目标')" />
                            </n-form-item>

                            <n-form-item :label="$t('类型')" path="type">
                                <n-radio-group v-model:value="formValue.type" name="radiogroup1">
                                    <n-space>
                                        <n-radio value="0">{{ $t('承诺型') }}</n-radio>
                                        <n-radio value="1">{{ $t('挑战性') }}</n-radio>
                                    </n-space>
                                </n-radio-group>
                            </n-form-item>

                            <n-form-item :label="$t('优先级')" path="priority">
                                <n-radio-group v-model:value="formValue.priority" name="radiogroup2">
                                    <n-space>
                                        <template v-if="formValue.type == '0'">
                                            <n-radio value="0">
                                                <span class="span span-1">P0</span>
                                            </n-radio>
                                            <n-radio value="1">
                                                <span class="span span-2">P1</span>
                                            </n-radio>
                                        </template>
                                        <template v-else>
                                            <n-radio value="2">
                                                <span class="span span-3">P2</span>
                                            </n-radio>
                                        </template>
                                    </n-space>
                                </n-radio-group>
                            </n-form-item>

                            <n-form-item :label="$t('归属')" path="ownership">
                                <n-radio-group v-model:value="formValue.ownership" name="radiogroup3">
                                    <n-space>
                                        <n-radio value="0">{{ $t('部门') }}</n-radio>
                                        <n-radio value="1">{{ $t('个人') }}</n-radio>
                                    </n-space>
                                </n-radio-group>
                            </n-form-item>

                            <n-form-item :label="$t('可见范围')">
                                <n-select v-model:value="formValue.visibilityRange" :placeholder="$t('请选择可见范围')"
                                    :options="generalOptions" />
                            </n-form-item>

                            <n-form-item :label="$t('周期')">
                                <n-date-picker class="w-full" v-model:value="formValue.time" type="daterange" clearable
                                    size="medium" />
                            </n-form-item>

                            <n-form-item :label="$t('对齐目标')">
                                <div @click="handleGoal"
                                    class="w-full h-[30px] bg-[#F4F5F7] border-[1px] border-[#F4F5F7] border-solid rounded cursor-pointer pl-12 pr-8 flex items-center">
                                    <p class="text-[rgba(194,194,194,1)] text-14">{{ $t('请选择对齐目标') }}</p>
                                    <i class="taskfont text-[rgba(194,194,194,1)] ml-auto">&#xe72b;</i>
                                </div>
                            </n-form-item>

                            <n-form-item :label="$t('关联项目')">
                                <n-select v-model:value="formValue.item" :placeholder="$t('请选择项目')"
                                    :options="itemOptions" />
                            </n-form-item>
                        </n-form>
                        <div>
                            <div class="flex items-center justify-between">
                                <h3 class="text-14 text-text-li font-medium">{{ $t('关键KR') }}</h3>
                                <div class="flex items-center cursor-pointer" @click="handleAddKr">
                                    <i class="taskfont text-14 text-primary-color">&#xe731;</i>
                                    <p class="text-14 text-primary-color ml-4">{{ $t('添加KR') }}</p>
                                </div>
                            </div>

                            <div v-for="(item, index) in formKRValue"
                                class="border-[1px] border-solid border-[#F2F2F2] mt-16 rounded">
                                <div
                                    class="flex items-center justify-between px-[12px] py-[13px] bg-[#FAFAFA] border-0 border-b-[1px] border-solid border-[#F2F2F2]">
                                    <h3 class="text-14 text-text-li font-medium">KR{{ index + 1 }}</h3>
                                    <div class="flex items-center cursor-pointer" @click="handleRemoveKr(index)">
                                        <i class="taskfont text-14 text-text-tips">&#xe787;</i>
                                    </div>
                                </div>

                                <div class="p-24 pb-2">
                                    <n-form ref="formRefs" :model="formKRValue[index]" size="medium" label-placement="left"
                                        label-width="auto" require-mark-placement="left">
                                        <n-grid :cols="4" :x-gap="16">
                                            <n-form-item-gi :span="4" class="w-full" :label="$t('KR')">
                                                <n-input v-model:value="item.KR" :placeholder="$t('请输入')" />
                                            </n-form-item-gi>

                                            <n-form-item-gi :span="4" :label="$t('周期')" :rule="timeRule">
                                                <n-date-picker class="w-full" v-model:value="item.time" type="daterange"
                                                    clearable size="medium" />
                                            </n-form-item-gi>

                                            <n-form-item-gi :span="2" :label="$t('参与人')">
                                                <n-select v-model:value="item.member" multiple :options="options"
                                                    :render-label="renderLabel" :render-tag="renderMultipleSelectTag"
                                                    filterable />
                                            </n-form-item-gi>

                                            <n-form-item-gi :span="2" :label="$t('信心')">
                                                <n-input v-model:value="item.confidence" :placeholder="$t('0-100')" />
                                            </n-form-item-gi>
                                        </n-grid>
                                    </n-form>
                                </div>
                            </div>
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
        <SelectAlignment :value="selectAlignmentShow" @close="() => { selectAlignmentShow = false }"></SelectAlignment>
    </n-drawer>

</template>
<script setup lang="ts">
import { watch } from "vue";
import {
    NAvatar,
    NText,
    NTag,
    SelectRenderTag,
    SelectRenderLabel
} from 'naive-ui'
import SelectAlignment from '@/views/components/SelectAlignment.vue'

const emit = defineEmits(['close'])

const show = ref(false)
const loadIng = ref(false)
const selectAlignmentShow = ref(false)

const formValue = reactive({
    object: "",
    type: "0",
    priority: "0",
    ownership: "0",
    visibilityRange: null,
    time: null,
    goal: null,
    item: null,
})

const formKRValue = reactive([
    {
        KR: null,
        time: null,
        confidence: null,
        member: null,
    },

])

const generalOptions = ref([
    {
        label: $t('全公司'),
        value: 0,
    },
    {
        label: $t('仅相关成员'),
        value: 1,
    },
    {
        label: $t('仅部门成员'),
        value: 2,
    },
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

const options = reactive([
    {
        label: '111',
        value: '07akioni'
    },
    {
        label: '222',
        value: '08akioni'
    },
    {
        label: '333',
        value: '09akioni'
    }
])

const renderLabel: SelectRenderLabel = (option) => {
    return h(
        'div',
        {
            style: {
                display: 'flex',
                alignItems: 'center'
            }
        },
        [
            h(NAvatar, {
                src: 'https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg',
                round: true,
                size: 'small'
            }),
            h(
                'div',
                {
                    style: {
                        marginLeft: '12px',
                        padding: '4px 0'
                    }
                },
                [
                    h('div', null, [option.label as string]),
                    h(
                        NText,
                        { depth: 3, tag: 'div' },
                    )
                ]
            )
        ]
    )
}

const renderMultipleSelectTag: SelectRenderTag = ({
    option,
    handleClose
}) => {
    return h(
        NTag,
        {
            style: {
                padding: '0 6px 0 4px'
            },
            round: true,
            closable: true,
            onClose: (e) => {
                e.stopPropagation()
                handleClose()
            }
        },
        {
            default: () =>
                h(
                    'div',
                    {
                        style: {
                            display: 'flex',
                            alignItems: 'center'
                        }
                    },
                    [
                        h(NAvatar, {
                            src: 'https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg',
                            round: true,
                            size: 22,
                            style: {
                                marginRight: '4px'
                            }
                        }),
                        option.label as string
                    ]
                )
        }
    )
}


watch(() => formValue.type, (newValue) => {
    if (newValue) {
        formValue.type == '0' ? formValue.priority = '0' : formValue.priority = '2'
    }
})

const rules = {
    object: {
        required: true,
        message: $t('请输入目标'),
        trigger: ['input']
    },
    type: {
        required: true,
        trigger: 'change',
        message: $t('请选择类型')
    },
    priority: {
        required: true,
        trigger: 'change',
        message: $t('请选择优先级')
    },
    ownership: {
        required: true,
        trigger: 'change',
        message: $t('请选择归属')
    },
}
const timeRule = {
    required: true,
    trigger: 'change',
    validator() {

    },
}

const handleSubmit = () => {
    emit('close')
}
const closeDrawer = () => {

}

const handleAddKr = () => {
    formKRValue.push(
        {
            KR: null,
            time: null,
            confidence: null,
            member: null,
        },
    )
}

const handleRemoveKr = (index) => {
    formKRValue.splice(index, 1)
}

const handleGoal = () => {
    selectAlignmentShow.value = true
}

</script>
<style lang="less" >
.n-drawer-body-content-wrapper {
    @apply relative;
}

.add-main-box {
    @apply flex-auto shrink-0;
}

.button-box {
    @apply flex gap-2 mt-32 flex-initial;
}

.n-drawer-header__close {
    @apply absolute -left-36;

    i {
        @apply text-[#fff];
    }
}

.span {
    @apply text-14 text-white px-6 py-4 rounded-full flex items-center leading-3 shrink-0;
}

.span-1 {
    @apply bg-[#FF7070];
}

.span-2 {
    @apply bg-[#FC984B];
}

.span-3 {
    @apply bg-[#72A1F7];
}
</style>