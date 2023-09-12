<template >
    <n-modal v-model:show="show" transform-origin="center" :mask-closable="false">
        <n-card class="w-[480px] max-w-[90%]" :bordered="false" size="huge" role="dialog" aria-modal="true" >
            <template #header>
                <h3 class="text-16 md:text-18 text-title-color ">{{ $t('信心指数') }}</h3>
            </template>
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <n-form ref="formRef" :model="formValue" size="medium" label-placement="left" label-width="auto">
                    <n-form-item >
                        <n-input-number class="w-full" :max="100" :min="0" :precision="0" placeholder="0-100" v-model:value="formValue.confidence" :show-button="false">

                        </n-input-number>
                    </n-form-item>

                </n-form>
            </div>

            <template #footer>
                <div class="button-box flex justify-end mt-0">
                    <div class="flex flex-1 md:flex-initial items-center gap-4">
                        <n-button class="flex-1 md:flex-initial" :loading="loadIng" strong secondary @click="handleClose">
                            {{ $t('取消') }}
                        </n-button>
                        <n-button class="flex-1 md:flex-initial" :loading="loadIng" type="primary" @click="handleSubmit">
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
import { confidenceUpdate } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"

const message = useMessage()
const show = ref(false)
const loadIng = ref(false)
const formValue = ref({
    confidence: 0,
})

const props = defineProps({
    id: {
        type: Number,
        default: 0,
    },
    confidence: {
        type: Number,
        default: 0,
    },
})

watch(() => props.id, (newValue) => {
    if (newValue) {
        formValue.value.confidence = props.confidence || null
    }
}, { immediate: true })

const emit = defineEmits(['close'])



const handleSubmit = () => {
    const upData = {
        id: props.id,
        confidence: formValue.value.confidence,
    }
    loadIng.value = true
    confidenceUpdate(upData)
        .then(({ msg }) => {
            message.success($t('操作成功'))
        })
        .catch(ResultDialog)
        .finally(() => {
            emit('close',1,props.id,formValue.value.confidence)
            loadIng.value = false
        })
}

const handleClose = () => {
    emit('close',2)
}
</script>
<style lang="less" scoped>
:deep(.n-card__content){
    @apply pb-0 !important;
}
</style>
