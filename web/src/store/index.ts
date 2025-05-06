import { computed } from "vue"
import { createPinia, defineStore } from "pinia"
import piniaPluginPersistedstate from "pinia-plugin-persistedstate"
import { ConfigProviderProps, createDiscreteApi, darkTheme, useOsTheme } from "naive-ui"
import { GlobalState } from "./interface"
import { I18nGlobal } from "@/lang"
import { getAppData } from "@/utils/app"

export const GlobalStore = defineStore({
    id: "GlobalState",
    state: (): GlobalState => ({
        baseUrl: "",
        baseRoute: "",
        language: "zh",
        themeName: "",
        isElectron: false,
        timer: {},
        //
        okrDetail: {
            show: false,
            id: 0,
        },
        addMultipleShow: false,
        addMultipleChange: false,
        addMultipleData: null,
        multipleId: 0,
        okrEditData: null,
        okrEdit: false,
        doubleSkip: null,
    }),
    actions: {
        async init() {
            if (["light", "dark"].indexOf(this.themeName) === -1) {
                this.themeName = useOsTheme().value
            }
        },
        openOkrDetails(id) {
            this.okrDetail = { show: false, id: id }
            this.okrDetail = { show: true, id: id }
        },
        setBaseUrl(url) {
            this.baseUrl = url
        },
        setBaseRoute(route) {
            if (route) {
                this.baseRoute = "/" + route
            }
        },
        setThemeName(name: string) {
            this.themeName = name
        },
        setLanguage(language: any) {
            I18nGlobal.locale.value = this.language = language
        },
        setIsElectron(isElectron: boolean) {
            this.isElectron = isElectron
        },
        appSetup() {
            return {
                themeName: computed(() => {
                    return this.themeName
                }),
                theme: computed(() => {
                    return this.themeName === "dark" ? darkTheme : null
                }),
            }
        },
        okrSetup() {
            return {
                okrEditData: computed(() => {
                    return this.okrEditData
                }),
                okrEdit: computed(() => {
                    return this.okrEdit
                }),
            }
        },
        multipleSetup() {
            return {
                addMultipleShow: computed(() => {
                    return this.addMultipleShow
                }),
                multipleId: computed(() => {
                    return this.multipleId
                }),
                addMultipleData: computed(() => {
                    return this.addMultipleData
                }),
            }
        },
        dialogProvider() {
            const { dialog } = createDiscreteApi(["dialog"], {
                configProviderProps: computed<ConfigProviderProps>(() => ({
                    theme: this.themeName === "dark" ? darkTheme : null,
                })),
            })
            return dialog
        },
        timeout(ms: number, key: string, ...name) {
            return new Promise((resolve) => {
                key = `${key}-${name.join("-")}`
                this.timer[key] && clearTimeout(this.timer[key])
                if (typeof this.timer[key] !== "undefined") {
                    clearTimeout(this.timer[key])
                    delete this.timer[key]
                }
                this.timer[key] = setTimeout(resolve, ms)
            })
        },
        isSingle() {
            return getAppData('initialData.isSubElectron') ? 1 : 0
        },
        isPortrait() {
            return getAppData('instance.store.state.windowPortrait') ? 1 : 0
        },
    },
})

//
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
export default pinia
