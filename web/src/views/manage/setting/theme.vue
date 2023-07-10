<template>
    <div class="setting-item">
        <n-form label-placement="left" label-width="auto">
            <n-form-item :label="$t('选择主题')" path="theme">
                <n-select v-model:value="theme" :options="options" @update:value="update" />
            </n-form-item>
        </n-form>
        <div class="setting-footer">
            <n-button @click="handleThemeUpdate" type="primary">{{ $t("提交") }}</n-button>
            <n-button @click="handleThemeReset">{{ $t("重置") }}</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { GlobalStore } from "@/store"
import { ref } from "vue"
import utils from "@/utils/utils"
const globalStore = GlobalStore()
const { themeName } = globalStore.appSetup()
const theme = ref('')

const options = ref([
            { label: $t("跟随系统"), value: "flow" },
            { label: $t("明亮"), value: "light" },
            { label: $t("暗黑"), value: "dark" },
])

// 获取主题
theme.value = utils.cloneJSON(themeName.value)

//下拉改变
const update = (e: any) => {
    theme.value = e
}

//提交
const handleThemeUpdate = () => {
    globalStore.setThemeName(theme.value)
}

//重置
const handleThemeReset = () => {
    theme.value = utils.cloneJSON(themeName.value)
}
</script>

<style lang="less"></style>
