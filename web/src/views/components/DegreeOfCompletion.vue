<template >
    <n-modal v-model:show="show" transform-origin="center" :mask-closable="false">
        <n-card class="w-[480px] max-w-[90%]"  :bordered="false" size="huge" role="dialog"
            aria-modal="true">
            <template #header>
                <h3 class="text-16 md:text-18 text-title-color ">{{ $t('更新 KR 完成度') }}</h3>
            </template>
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <n-form ref="formRef" :model="formValue" size="medium" label-placement="left" label-width="auto">
                    <n-form-item :label="$t('完成度')">
                        <n-input-number v-model:value="formValue.progress" :max="100" :min="0" :precision="0"
                            placeholder="0" :show-button="false">
                            <template #suffix> %</template>
                        </n-input-number>
                    </n-form-item>
                    <n-form-item :label="$t('状态')" class="">
                        <div class="flex items-center  gap-6">
                            <div class="flex items-center gap-2 cursor-pointer " @click="()=>{ formValue.progress_status = 1 }">
                                <div class="border-[1px] border-solid border-[#8BCF70] rounded-full w-[16px] h-[16px] relative">
                                    <div v-if="formValue.progress_status == 1" class="absolute bg-[#8BCF70] w-[10px] h-[10px] rounded-full top-2 left-2"></div>
                                </div>
                                <span>{{ $t('正常') }}</span>
                            </div>
                            <div class="flex items-center gap-2 cursor-pointer" @click="()=>{ formValue.progress_status = 2 }">
                                <div class="border-[1px] border-solid border-[#FFA25A] rounded-full w-[16px] h-[16px] relative">
                                    <div v-if="formValue.progress_status == 2" class="absolute bg-[#FFA25A] w-[10px] h-[10px] rounded-full top-2 left-2"></div>
                                </div>
                                <span>{{ $t('有风险') }}</span>
                            </div>
                            <div class="flex items-center gap-2 cursor-pointer" @click="()=>{ formValue.progress_status = 3 }">
                                <div class="border-[1px] border-solid border-[#FF7070] rounded-full w-[16px] h-[16px] relative">
                                    <div v-if="formValue.progress_status == 3" class="absolute bg-[#FF7070] w-[10px] h-[10px] rounded-full top-2 left-2"></div>
                                </div>
                                <span>{{ $t('已延期') }}</span>
                            </div>
                        </div>
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
import { updateProgress } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"

const message = useMessage()
const show = ref(false)
const loadIng = ref(false)

const formValue = ref({
    progress: 0,
    progress_status: 0,
})

const props = defineProps({
    id: {
        type: Number,
        default: 0,
    },
    progress: {
        type: Number,
        default: 0,
    },
    progress_status: {
        type: Number,
        default: 0,
    },
})

watch(() => props.id, (newValue) => {
    if (newValue) {
        formValue.value.progress = props.progress || null
        formValue.value.progress_status = props.progress_status
    }
}, { immediate: true })



const emit = defineEmits(['close'])
//提交
const handleSubmit = () => {
    const upData = {
        id: props.id,
        progress: formValue.value.progress,
        status: formValue.value.progress_status,
    }
    loadIng.value = true
    updateProgress(upData)
        .then(({ msg }) => {
            message.success($t('操作成功'))
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            emit('close', 1)
            loadIng.value = false
        })
}
//关闭
const handleClose = () => {
    emit('close', 2)
}
</script>
<style lang="less" scoped>
:deep(.n-card__content) {
    @apply pb-0 !important;
}
</style>
