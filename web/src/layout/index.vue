<template>
    <!-- :theme="theme" -->
    <n-config-provider  :locale="locale" :date-locale="dateLocale"
        :theme-overrides="theme === null ? lightThemeOverrides : darkThemeOverrides">
        <n-loading-bar-provider>
            <n-message-provider>
                <n-notification-provider>
                    <n-dialog-provider>
                        <Result v-if="resultCode > 0" />
                        <layout v-else v-wechat-title="title"/>
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
import { useRoute } from 'vue-router'
const route = useRoute()
const lightThemeOverrides = themeOverrides.lightThemeOverrides
const darkThemeOverrides = themeOverrides.darkThemeOverrides
const resultCode = result.code()
const globalStore = GlobalStore()
const { themeName, theme } = globalStore.appSetup()

const title = computed(()=>{
    return route.meta.title
})


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
<style lang="less" scoped>

</style>
