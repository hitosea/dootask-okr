<template>
    <div v-if="props.position=='onscroll'" class="h-full w-full relative">
        <n-spin size="small" class="w-full mt-[24px]">
        </n-spin>
    </div>
    <div v-else class="h-full w-full relative">
        <n-spin size="small"  class=" absolute left-1/2 translate-x-[-14px] top-24 z-[9999]">
        </n-spin>
    </div>
</template>

<script lang="ts" setup>
    const props = defineProps({
        position: {
            type: String,
            default: ()=> '',
        }
    });
</script>