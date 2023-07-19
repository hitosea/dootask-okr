<template >
    <n-modal v-model:show="props.value" transform-origin="center">
        <n-card style="width: 80%" :title="$t('添加OKR')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer" size="24" :component="Close" @click="handleClose" />
            </template>
            <div class="add-main-box">
                <n-form ref="formRef" :model="formValue" :rules="rules" size="medium">
                    <n-grid :cols="4" :x-gap="12">
                        <n-form-item-gi  class="w-full" :span="4" :label="$t('目标（O）')" path="object">
                            <n-input v-model:value="formValue.object" :placeholder="$t('名称')" />
                        </n-form-item-gi>

                        <n-form-item-gi :span="2" :label="$t('类型')" path="type">
                            <n-radio-group v-model:value="formValue.type" name="radiogroup1">
                                <n-space>
                                    <n-radio value="0">{{ $t('承诺型') }}</n-radio>
                                    <n-radio value="1">{{ $t('挑战性') }}</n-radio>
                                </n-space>
                            </n-radio-group>
                        </n-form-item-gi>

                        <n-form-item-gi :span="2" :label="$t('优先级')" path="priority">
                            <n-radio-group v-model:value="formValue.priority" name="radiogroup2">
                                <n-space>
                                    <template v-if="formValue.type == '0'">
                                        <n-radio value="0">P0</n-radio>
                                        <n-radio value="1">P1</n-radio>
                                    </template>
                                    <template v-else>
                                        <n-radio value="2">P2</n-radio>
                                    </template>
                                </n-space>
                            </n-radio-group>
                        </n-form-item-gi>

                        <n-form-item-gi :span="4" :label="$t('归属')" path="ownership">
                            <n-radio-group v-model:value="formValue.ownership" name="radiogroup3">
                                <n-space>
                                    <n-radio value="0">{{ $t('部门') }}</n-radio>
                                    <n-radio value="1">{{ $t('个人') }}</n-radio>
                                </n-space>
                            </n-radio-group>
                        </n-form-item-gi>

                        <n-form-item-gi :span="2" :label="$t('可见范围')" path="">
                            <n-select v-model:value="formValue.visibilityRange" placeholder="Select"
                                :options="generalOptions" />
                        </n-form-item-gi>

                        <n-form-item-gi :span="2" :label="$t('周期')" path="">
                            <n-date-picker class="w-full" v-model:value="formValue.time" type="daterange" clearable size="medium"/>
                        </n-form-item-gi>
                        <n-form-item-gi :span="2" :label="$t('对齐目标')" path="">
                            <n-select v-model:value="formValue.item" placeholder="Select" :options="itemOptions" />
                        </n-form-item-gi>
                        <n-form-item-gi :span="2" :label="$t('关联项目')" path="">
                            <n-select v-model:value="formValue.item" placeholder="Select" :options="itemOptions" />
                        </n-form-item-gi>
                    </n-grid>

                </n-form>
            </div>
            <template #footer>
                <div class="button-box">
                    <n-button :loading="loadIng" type="default" @click="handleClose">
                        {{ $t('取消') }}
                    </n-button>
                    <n-button :loading="loadIng" type="primary" @click="">
                        {{ $t('确定') }}
                    </n-button>
                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { Close } from "@vicons/ionicons5"
import { watch } from "vue";
const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
})

const emit = defineEmits(['close'])


const loadIng = ref(false)

const formValue = reactive({
    object: "",
    type: "0",
    priority: "0",
    ownership: "0",
    visibilityRange: "",
    time: null,
    item: "",
})
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

const handleClose = () => {
    emit('close')
}


</script>
<style lang="less" >
.add-main-box {}

.button-box {
    @apply flex justify-end gap-2;
}
</style>