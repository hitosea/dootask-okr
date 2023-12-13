<template>
    <!-- :theme="theme" -->
    <n-config-provider :locale="locale" :date-locale="dateLocale" :theme-overrides="theme === null ? lightThemeOverrides : darkThemeOverrides">
        <n-loading-bar-provider>
            <n-message-provider>
                <n-notification-provider>
                    <n-dialog-provider>
                        <Result v-if="resultCode > 0" />
                        <layout v-else />
                    </n-dialog-provider>
                </n-notification-provider>
            </n-message-provider>
        </n-loading-bar-provider>
    </n-config-provider>

</template>

<script lang="ts" setup>
import { GlobalStore } from "../store"
import { zhCN, dateZhCN, enUS, dateEnUS } from "naive-ui"
import Layout from "./layout.vue"
import result from "../utils/result"
import { computed } from "vue"
import { NLocale } from "naive-ui/es/locales/common/enUS"
import { NDateLocale } from "naive-ui/es/locales"
import themeOverrides from "../utils/naive.config"
const lightThemeOverrides = themeOverrides.lightThemeOverrides
const darkThemeOverrides = themeOverrides.darkThemeOverrides
const resultCode = result.code()
const globalStore = GlobalStore()
const { themeName, theme } = globalStore.appSetup()

if(!document.getElementById('okr-okrfont-style')){
    const styleTag = document.createElement('style');
    styleTag.setAttribute('id', 'okr-okrfont-style');
    styleTag.textContent = !import.meta.env.DEV ? `
        @font-face {
            font-family: 'okrfont';  /* Project id 2583385 */
            src: url('${globalStore.baseUrl}/apps/okr/assets/styles/okrfont/iconfont.woff2') format('woff2'),
            url('${globalStore.baseUrl}/apps/okr/assets/styles/okrfont/iconfont.woff') format('woff'),
            url('${globalStore.baseUrl}/apps/okr/assets/styles/okrfont/iconfont.ttf') format('truetype');
        }
        .okrfont {
            font-family: "okrfont", "serif" !important;
            font-style: normal;
            -webkit-font-smoothing: antialiased;
            -moz-osx-font-smoothing: grayscale;
        }
    ` : `
        @font-face {
            font-family: 'okrfont';  /* Project id 2583385 */
            src: url('http://127.0.0.1:5566/apps/okr/assets/styles/okrfont/iconfont.woff2') format('woff2'),
            url('http://127.0.0.1:5566/apps/okr/assets/styles/okrfont/iconfont.woff') format('woff'),
            url('http://127.0.0.1:5566/apps/okr/assets/styles/okrfont/iconfont.ttf') format('truetype');
        }
        .okrfont {
            font-family: "okrfont", "serif" !important;
            font-style: normal;
            -webkit-font-smoothing: antialiased;
            -moz-osx-font-smoothing: grayscale;
        }
    `;
    document.head.appendChild(styleTag);
}

const locale = computed((): NLocale => {
    if (globalStore.language == "en") return enUS
    return zhCN
})

const dateLocale = computed((): NDateLocale => {
    if (globalStore.language == "en") return dateEnUS
    return dateZhCN
})

watch(themeName,(newVal,oldVal)=>{
    document.querySelector("body").classList.remove('okr-theme-'+oldVal)
    document.querySelector("body").classList.add('okr-theme-'+newVal)
},{immediate:true})

</script>

