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
.main {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .welcome {
        margin-bottom: 12px;
    }

    .buttons {
        display: flex;
        align-items: center;
        margin-top: 12px;

        > button + button {
            margin-left: 12px;
        }
    }
}
</style>
