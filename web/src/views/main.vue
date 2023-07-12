<template>
    <div class="main">
        <div class="welcome">
            {{ $t("main.welcome") }}
        </div>
        <n-date-picker panel type="date" />

        <div class="buttons">
            <n-button tertiary @click="handleThemeUpdate">
                {{ themeLabel }}
            </n-button>

            <n-button tertiary @click="handleLanguageUpdate">
                {{ languageLabel }}
            </n-button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { GlobalStore } from "../store"
import { computed } from "vue"
const globalStore = GlobalStore()

const themeLabel = computed(
    () =>
        ({
            dark: "浅色",
            light: "深色",
        }[globalStore.themeName]),
)
const handleThemeUpdate = () => {
    if (globalStore.themeName === "dark") {
        globalStore.setThemeName("light")
    } else {
        globalStore.setThemeName("dark")
    }
}

const languageLabel = computed(
    () =>
        ({
            zh: "English",
            en: "中文",
        }[globalStore.language]),
)

const handleLanguageUpdate = () => {
    if (globalStore.language === "en") {
        globalStore.setLanguage("zh")
    } else {
        globalStore.setLanguage("en")
    }
}
</script>

<style lang="less" scoped>

</style>
