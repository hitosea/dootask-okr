<template>
    <div v-if="show">
        <h1 class="text-center mt-[12%]">{{ $t('欢迎使用okr子系统') }}</h1>
        <div class="text-center mt-[20px]">
            <router-link to="/apps/list" class="mr-[10px]">{{ $t('前往Okr列表') }}</router-link> |
            <router-link to="/apps/analysis" class="ml-[10px]">{{ $t('前往Okr汇总') }}</router-link>
        </div>
    </div>
</template>

<script lang="ts" setup>
const show = ref(false)

setTimeout(()=>{
    show.value = true
},1000);
</script>
