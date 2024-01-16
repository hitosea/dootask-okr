import { createI18n } from "vue-i18n"
import auto_zh from "./auto/zh.json"
import auto_en from "./auto/en.json"
import auto_fr from "./auto/fr.json"
import auto_id from "./auto/id.json"
import auto_ja from "./auto/ja.json"
import auto_ko from "./auto/ko.json"
import auto_zhHant from "./auto/zh-hant.json"

const i18n = createI18n({
    legacy: false,
    locale: localStorage.getItem("__language:type__") || "zh",
    fallbackLocale: "zh",
    globalInjection: true,
    messages: {
        zh: { ...auto_zh },
        en: { ...auto_en },
        fr: { ...auto_fr },
        id: { ...auto_id },
        ja: { ...auto_ja },
        ko: { ...auto_ko },
        "zh-CHT": { ...auto_zhHant },
    },
    warnHtmlMessage: false,
})

export default i18n

export const I18nGlobal = i18n.global
