<template >
    <n-modal v-model:show="showModal">
        <n-card class="w-[420px] max-w-[90%]" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>

            </template>
            <template #header>
                <h3 class="flex items-center text-16 font-medium text-title-color"> <n-icon
                        class="mr-12 text-[rgb(237,64,20)]" size="28" :component="CloseCircleSharp" />{{ $t('温馨提示') }}</h3>
            </template>
            <p class=" text-text-li text-14 pl-[40px]">{{ props.content }}</p>
            <template #footer>
                <div class="flex justify-end">
                    <n-button type="primary" size="small" @click="() => { emit('close') }">
                        {{ $t('确定') }}
                    </n-button>
                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { CloseCircleSharp } from '@vicons/ionicons5'

const showModal = ref(false)

const emit = defineEmits(['close'])

const props = defineProps({
    content: {
        type: String,
        default: '',
    }
})


</script>
<style lang="less" scoped>
:deep(.n-card-header){
    @apply pb-12;
}
:deep(.n-card__content){
    @apply pb-20;
}
</style>