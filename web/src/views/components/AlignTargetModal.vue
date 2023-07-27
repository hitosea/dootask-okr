<template >
    <n-modal v-model:show="props.value" transform-origin="center">
        <n-card class="w-[640px]" :title="$t('对齐目标')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <AlignTarget :value="props.value" :id="props.id"></AlignTarget>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import { Close } from "@vicons/ionicons5"
import AlignTarget from "@/views/components/AlignTarget.vue";

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
    id: {
        type: Number,
        default: 0,
    },
})


const emit = defineEmits(['close'])
const handleClose = () => {
    emit('close')
}

</script>

<style lang="less" >

</style>

