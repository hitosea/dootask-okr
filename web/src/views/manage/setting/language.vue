<template>
    <div class="setting-item">
        <n-form label-placement="left" label-width="auto">
            <n-form-item :label="$t('选择语言')" path="language">
                <n-select v-model:value="language" :options="options" @update:value="update" />
            </n-form-item>
        </n-form>
        <div class="setting-footer">
            <n-button @click="handleLanguageUpdate" type="primary">{{ $t("提交") }}</n-button>
            <n-button @click="handleLanguageReset">{{ $t("重置") }}</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { GlobalStore } from "@/store"
import { ref } from "vue"
import utils from "@/utils/utils"
const globalStore = GlobalStore()
const language = ref('')

const options = ref([
    { label: $t("English"), value: "en", },
    { label: $t("中文"), value: "zh", },
])
const update = (e: any) => {
    language.value = e
}

language.value = utils.cloneJSON(globalStore.language)




const handleLanguageUpdate = () => {
    globalStore.setLanguage(language.value)
}

const handleLanguageReset = () => {
    language.value = utils.cloneJSON(globalStore.language)
}
</script>

<style lang="less"></style>
