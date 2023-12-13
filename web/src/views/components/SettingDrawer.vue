<template >
    <n-drawer v-model:show="show" :on-after-enter="showDrawer" :on-after-leave="closeDrawer" :mask-closable="false"
        :z-index="13" class="okr" style="--n-body-padding:16px 20px 26px 26px;max-width: 500px;width: 90%;"
        :trap-focus="false">
        <n-drawer-content :title="$t('设置')" closable>
            <div class="flex flex-col absolute top-[16px] bottom-[24px] left-[26px] right-[26px] overflow-hidden">
                <div class="flex-auto overflow-hidden">
                    <SettingMain ref="settingMainRef" @loadIng="(e)=>(loadIng = e)"></SettingMain>
                </div>
                <div class="button-box flex-initial">
                    <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                        {{ $t('提交') }}
                    </n-button>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import SettingMain from './SettingMain.vue';

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
</script>

<style lang="less" scoped>
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
