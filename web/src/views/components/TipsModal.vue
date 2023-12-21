<template >
    <n-modal v-model:show="showModal" :on-after-leave="()=>{ emit('reset') }">
        <n-card class="w-[420px] max-w-[90%]" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>

            </template>
            <template #header>
                <h3 class="flex items-center text-16 font-medium text-title-color"> <n-icon
                        class="mr-12 " :class="props.color" size="28" :component="icon==1 ? CloseCircleSharp : AlertCircleSharp" />{{ $t('温馨提示') }}</h3>
            </template>
            <p class=" text-text-li text-14 pl-[40px]">{{ props.content }}</p>
            <template #footer>
                <div class="flex justify-end">
                    <n-button type="primary" size="small" @click="handleClose">
                        {{ $t('确定') }}
                    </n-button>
                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { CloseCircleSharp ,AlertCircleSharp } from '@vicons/ionicons5'

const showModal = ref(false)

const emit = defineEmits(['close','reset'])

const props = defineProps({
    content: {
        type: String,
        default: '',
    },
    color:{
        type: String,
        default: 'text-[rgb(237,64,20)]',
    },
    icon :{
        type: Number,
        default: 1,
    },
    tipsClose :{
        type: Number,
        default: 0,
    },
})
const handleClose = ()=>{
    if(props.tipsClose == 1){
        emit('close',1)
    }else{
        emit('close',0)
    }
}

</script>
<style lang="less" scoped>
:deep(.n-card-header){
    @apply pb-12;
}
:deep(.n-card__content){
    @apply pb-20;
}
</style>