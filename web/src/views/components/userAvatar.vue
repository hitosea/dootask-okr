<template >
    <div>
        <n-avatar round :size="28" class="mr-8 shrink-0"
            :src="props.userUrl" />
    </div>
</template>
<script setup lang="ts">

const props = defineProps({
    userUrl: {
        type: String,
        default: '',
    },
    disable:{
        type: String,
        default: '',
    }
})
</script>
<style lang="less" scoped>
</style>