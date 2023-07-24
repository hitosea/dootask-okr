<template >
    <n-modal v-model:show="show" transform-origin="center">
        <n-card class="w-[480px]" :title="$t('信心指数')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <n-form ref="formRef" :model="formValue" size="medium" label-placement="left" label-width="auto">
                    <n-form-item >
                        <n-input-number class="w-full" v-model:value="formValue.complete" :show-button="false">
                         
                        </n-input-number>
                    </n-form-item>

                </n-form>
            </div>

            <template #footer>
                <div class="button-box flex justify-end mt-0">
                    <div class="flex items-center gap-4">
                        <n-button :loading="loadIng" type="default" @click="handleClose">
                            {{ $t('取消') }}
                        </n-button>
                        <n-button :loading="loadIng" type="primary" @click="">
                            {{ $t('确定') }}
                        </n-button>
                    </div>

                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { Close } from "@vicons/ionicons5"
const show = ref(false)
const loadIng = ref(false)
const formValue = ref({
    complete: 0,
    status: 0,
})

const emit = defineEmits(['close'])

const handleClose = () => {
    emit('close')
}
</script>
<style lang="less" scoped>
:deep(.n-card__content){
    @apply pb-0 !important;
}
</style>