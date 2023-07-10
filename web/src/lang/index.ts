import { createI18n } from "vue-i18n"
import zh from "./modules/zh"
import en from "./modules/en"
import auto_zh from "./auto/zh.json"
import auto_en from "./auto/en.json"

const i18n = createI18n({
    legacy: false,
    locale: localStorage.getItem("lang") || "zh",
    fallbackLocale: "zh",
    globalInjection: true,
    messages: {
        zh: { ...zh, ...auto_zh },
        en: { ...en, ...auto_en },
    },
    warnHtmlMessage: false,
})

export default i18n

export const I18nGlobal = i18n.global
