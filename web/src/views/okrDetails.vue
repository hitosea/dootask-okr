<template >
    <div class="nav-top  h-[44px] bg-[#FAFAFA] z-[5]">
        <i @click="handleReturn" class="taskfont icon-return z-[2]">&#xe704;</i>
        <h2 class=" absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('OKR详情') }}</h2>
        <n-icon class="text-text-tips z-[2]" size="20" :component="EllipsisHorizontalSharp" />
    </div>
    <div class="pt-[58px] pb-16 pl-16 pr-16">
        <OkrDetailsMain ref="OkrDetailsMainRef" :show="true" :id="id"
            ></OkrDetailsMain>
    </div>
</template>
<script lang="ts" setup>
import OkrDetailsMain from '@/views/components/OkrDetailsMain.vue';
import { EllipsisHorizontalSharp } from '@vicons/ionicons5'
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()

const id = ref(null)
const AddOkrsRef = ref(null)

if (route.query.data != undefined) {
    id.value = Number(route.query.data)
}


const handleReturn = () => {
    router.back()
}
nextTick(()=>{
    AddOkrsRef.value.showDrawer()
})

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-16;

    .icon-return {
        @apply text-20 text-text-tips;
    }
}
</style>