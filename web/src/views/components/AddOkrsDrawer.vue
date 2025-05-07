<template >
    <n-drawer
        v-model:show="show"
        :on-after-enter="showDrawer"
        :on-after-leave="closeDrawer"
        :mask-closable="false"
        :z-index="nextZIndex()"
        class="okr"
        style="--n-body-padding:16px 20px 24px 34px;max-width: 600px;width: 90%;"
        :trap-focus="false">
        <n-drawer-content :title="(props.edit ? $t('编辑') : $t('添加')) + ' Objective'" closable>
            <div class="flex flex-col absolute top-[16px] bottom-[24px] left-[34px] right-[20px] overflow-hidden">
                <div class=" flex-auto overflow-hidden">
                    <AddOkrsMain
                        ref="AddOkrsRef"
                        :edit="props.edit"
                        :editData="props.editData"
                        @close="(e, id) => { emit('close', e, id) }"
                        @loadIng="(e) => { loadIng = e }"/>
                </div>
                <div class="button-box flex-initial">
                    <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                        {{ $t('提交') }}
                    </n-button>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import AddOkrsMain from '@/views/components/AddOkrsMain.vue';
import { nextZIndex } from "dootask-tools"

const loadIng = ref(false)
const show = ref(false)
const AddOkrsRef = ref(null)

const emit = defineEmits(['close', 'upData'])

const props = defineProps({
    edit: {
        type: Boolean,
        default: false,
    },
    editData: {
        type: Object,
        default: {},
    },
})

//提交
const handleSubmit = () => {
    AddOkrsRef.value.handleSubmit()
}


// 关闭Drawer
const closeDrawer = () => {
}

// 显示
const showDrawer = () => {
    AddOkrsRef.value.showDrawer()
}

</script>
<style lang="less" scoped>
:deep(.n-drawer-body-content-wrapper) {
    @apply relative;
}

.add-main-box {
    @apply flex-auto shrink-0;
}

.button-box {
    @apply flex gap-2 mt-24 flex-initial;
}

:deep(.n-drawer-header__close) {
    @apply absolute -left-36;

    &:focus {
        @apply bg-none;
    }

    i {
        @apply text-[#fff];
    }
}

:deep(.n-base-close:not(.n-base-close--disabled):focus::before) {
    @apply bg-transparent;
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

.okr-user-selects {
    @apply w-full bg-none border-none !important;
}
</style>
