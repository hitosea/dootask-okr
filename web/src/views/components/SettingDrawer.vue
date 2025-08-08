<template >
    <n-drawer 
        v-model:show="show" 
        :on-after-enter="showDrawer" 
        :on-after-leave="closeDrawer" 
        :z-index="nextZIndex()"
        class="okr" 
        style="--n-body-padding:16px 20px 26px 26px;max-width: 500px;width: 90%;"
        :trap-focus="false">
        <n-drawer-content :title="$t('设置')" closable>
            <div class="flex flex-col absolute top-[16px] bottom-[24px] left-[26px] right-[26px] overflow-hidden">
                <div class="flex-auto overflow-hidden">
                    <SettingMain ref="settingMainRef" @loadIng="(e)=>(loadIng = e)" @close="()=>{ emit('close') }"></SettingMain>
                </div>
                <div class="button-box flex-initial gap-[8px] flex">
                    <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                        {{ $t('提交') }}
                    </n-button>
                    <n-button :loading="loadIng" strong secondary @click="handleReset">
                        {{ $t('重置') }}
                    </n-button>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import SettingMain from './SettingMain.vue';
import { nextZIndex } from "@dootask/tools"

const loadIng = ref(false)
const show = ref(false)
const emit = defineEmits(['close', 'upData'])
const settingMainRef = ref(null)

// 关闭Drawer
const closeDrawer = () => {
    settingMainRef.value.closeDrawer()
}

// 显示
const showDrawer = () => {
    settingMainRef.value.showDrawer()
}

//提交
const handleSubmit = () => {
    settingMainRef.value.handleSubmit()
}
//提交
const handleReset = () => {
    settingMainRef.value.handleSubmit('get')
}


</script>

<style lang="less" scoped>
:deep(.n-drawer-body-content-wrapper) {
    @apply relative;
}
</style>
