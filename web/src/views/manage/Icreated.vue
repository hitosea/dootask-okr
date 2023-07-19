<template >
    <n-scrollbar>
        <div class="i-created-main">
            <PersonalStatistics></PersonalStatistics>
            <okrItem></okrItem>
            
        </div>
    </n-scrollbar>
</template>
<script lang="ts" setup>
import PersonalStatistics from '@/views/components/PersonalStatistics.vue'
import okrItem from '@/views/components/okrItem.vue'

</script>
<style lang="less" scoped>
.i-created-main{
    @apply py-24 flex flex-col gap-6;
}
</style>