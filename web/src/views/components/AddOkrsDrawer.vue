<template >
    <n-drawer v-model:show="show" @after-enter="showDrawer" :on-after-leave="closeDrawer" :mask-closable="false"
        :z-index="13" class="okr" style="--n-body-padding:16px 20px 16px 34px;max-width: 600px;width: 90%;">
        <n-drawer-content :title="props.edit ? $t('编辑OKR') : $t('添加OKR')" closable>
            <div class="flex flex-col h-full">
                <AddOkrs ref="AddOkrsRef" :edit="props.edit" :editData="props.editData"
                    @close="(e, id) => { emit('close', e, id) }" @loadIng="(e) => { loadIng = e }"></AddOkrs>
                <div class="button-box">
                    <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                        {{ $t('提交') }}
                    </n-button>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import AddOkrs from '@/views/components/AddOkrs.vue';

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
    AddOkrsRef.value.closeDrawer()
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
    @apply flex gap-2 mt-32 flex-initial;
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
