<template >
    <n-drawer 
        v-model:show="show" 
        :on-after-enter="showDrawer" 
        :on-after-leave="closeDrawer" 
        :z-index="nextZIndex()" 
        class="okr" 
        style="--n-body-padding:16px 20px 24px 24px;max-width: 998px;width: 90%;"
        :trap-focus="false">
        <n-drawer-content :title="$t('离职/删除人员') + ' OKR'" closable>
            <div class="flex flex-col absolute top-[16px] bottom-[24px] left-[24px] right-[20px] overflow-hidden">
                <div class="flex-auto overflow-hidden">
                    <DeleteMain @close="(e, id) => { emit('close', e, id)}"></DeleteMain>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import DeleteMain from './DeleteMain.vue';
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

<style lang="less" scoped>
:deep(.n-drawer-body-content-wrapper) {
    @apply relative;
}
</style>
