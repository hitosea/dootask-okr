<template >
    <n-drawer v-model:show="show" :on-after-enter="showDrawer" :on-after-leave="closeDrawer" :mask-closable="false"
        :z-index="nextZIndex()" class="okr" style="--n-body-padding:16px 20px 24px 24px;max-width: 998px;width: 90%;"
        :trap-focus="false">
        <n-drawer-content :title="$t('已归档') + ' OKR'" closable>
            <div class="flex flex-col absolute top-[16px] bottom-[24px] left-[24px] right-[20px] overflow-hidden">
                <div class=" flex-auto overflow-hidden">
                  <ArchiveMain @close="(e, id) => { emit('close', e, id)}"></ArchiveMain>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import ArchiveMain from './ArchiveMain.vue';
import { nextZIndex } from "@dootask/tools"


const show = ref(false)

const emit = defineEmits(['close', 'upData'])


// 关闭Drawer
const closeDrawer = () => {
}

// 显示
const showDrawer = () => {

}

</script>
<style scoped lang="less">

:deep(.n-drawer-body-content-wrapper) {
    @apply relative;
}

:deep(.n-drawer-header__close) {
    @apply absolute -left-36;

    &:focus {
        @apply bg-none;
    }
    i {
        @apply text-[#fff];
    }
}

:deep(.n-base-close:not(.n-base-close--disabled):focus::before) {
    @apply bg-transparent;
}
</style>
