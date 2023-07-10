<template>
    <div class="edit-text" @click="handleOnClick">
        <n-input
            v-if="isEdit"
            ref="inputRef"
            v-model:value="inputValue"
            autosize
            :status="status"
            :size="size"
            :placeholder="placeholder"

            @keyup.esc="handleCancel"

            @click.stop="handleUnStatus"
            @updateValue="handleUnStatus"
            @focus="handleUnStatus"

            @keyup.enter="handleUpdate"
            @blur="handleUpdate">
            <template #suffix>
                <n-spin v-if="isLoad" class="edit-suffix-icon" :size="12" />
                <n-icon v-else class="edit-suffix-icon" :component="CheckmarkCircleSharp" @click="handleUpdate"/>
            </template>
        </n-input>
        <slot v-else></slot>
    </div>
</template>

<style scoped lang="less">
.edit-text {
    display: inline-block;
    .edit-suffix-icon {
        cursor: pointer;
        margin-left: 6px;
        &.n-icon {
            font-size: 14px;
            padding: 1px;
        }
    }
}
</style>
<script lang="ts" setup>
import {defineProps, nextTick, ref} from 'vue';
import { CheckmarkCircleSharp } from '@vicons/ionicons5'

const props = defineProps({
    value: {
        type: String,
        default: '',
    },
    size: {
        default: 'medium',
    },
    placeholder: {
        default: '',
    },
    params: {
        default: () => ({}),
    },
    onEdit: {
        type: Function,
    },
    onUpdate: {
        type: Function,
    },
});

const status = ref()
const isLoad = ref(false)
const isEdit = ref(false)
const inputRef = ref(null)
const inputValue = ref(props.value)
const handleOnClick = () => {
    inputValue.value = props.value
    isEdit.value = true
    if (props.onEdit) {
        props.onEdit()
    }
    nextTick(() => {
        inputRef.value.select()
    })
}
const handleCancel = () => {
    if (!isLoad.value) {
        isEdit.value = false
    }
}
const handleUnStatus = () => {
    status.value = undefined
}
const handleUpdate = (e = null) => {
    if (!isEdit.value) {
        return
    }
    if (inputValue.value == props.value) {
        isEdit.value = false
        return
    }
    if (isLoad.value) {
        return
    }
    isLoad.value = true
    //
    if (!props.onUpdate) {
        isEdit.value = false    // 没有返回
        isLoad.value = false
        return
    }
    const call = props.onUpdate(inputValue.value, props.params, e);
    if (!call) {
        isEdit.value = false    // 返回无内容
        isLoad.value = false
        return
    }
    if (call.then) {
        call.then(_ => {
            isEdit.value = false
        }).catch(_ => {
            status.value = 'error'
        }).finally(() => {
            isLoad.value = false
        })
    } else {
        isEdit.value = false
    }
};
</script>
