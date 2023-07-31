import { computed } from "vue";
import { createPinia, defineStore } from "pinia"
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import { ConfigProviderProps, createDiscreteApi, darkTheme, useOsTheme } from "naive-ui";
import { GlobalState } from "./interface"
import piniaPersistConfig from "./config/pinia-persist"
import { I18nGlobal } from "@/lang"
import utils from "@/utils/utils"

const screenOrientation = utils.screenOrientation()
const windowLandscape = utils.screenOrientation() === 'landscape'
const windowPortrait = utils.screenOrientation() === 'portrait'

export const GlobalStore = defineStore({
    id: "GlobalState",
    state: (): GlobalState => ({
        isLoading: 0,
        language: "zh",
        themeName: "",
        timer: {},
        // 浏览器窗口方向
        windowOrientation: computed(() => { return screenOrientation }),
        windowLandscape: computed(() => { return windowLandscape }), // 横屏
        windowPortrait: computed(() => { return windowPortrait; }),   // 竖屏
        windowActive: true,
        windowScrollY: 0,
        keyboardType: null,
        keyboardHeight: 0,  // 键盘高度
        windowTouch: "ontouchend" in document,
        okrDetail: {
            show: false,
            id: 0,
        },
        addMultipleShow: false,
        multipleId: 0,
        addMultipleData: null,
    }),
    actions: {
        async init() {
            this.isLoading = 0;
            if (["light", "dark"].indexOf(this.themeName) === -1) {
                this.themeName = useOsTheme().value;
            }
        },
        openOkrDetails(id) {
            this.okrDetail = {
                show: true,
                id: id
            }
        },
        setLoading() {
            this.isLoading++
        },
        cancelLoading() {
            this.isLoading--
        },
        setThemeName(name: string) {
            this.themeName = name
        },
        setLanguage(language: any) {
            localStorage.setItem("lang", (I18nGlobal.locale.value = this.language = language));
        },
        setVues(vues: any) {
            window.Vues = vues;
        },
        appSetup() {
            return {
                themeName: computed(() => {
                    return this.themeName;
                }),
                theme: computed(() => {
                    return this.themeName === "dark" ? darkTheme : null;
                }),
            };
        },
        multipleSetup() {
            return {
                addMultipleShow: computed(() => {
                    return this.addMultipleShow;
                }),
                multipleId: computed(() => {
                    return this.multipleId ;
                }),
                addMultipleData: computed(() => {
                    return this.addMultipleData ;
                }),
            };
        },
        dialogProvider() {
            const { dialog } = createDiscreteApi(["dialog"], {
                configProviderProps: computed<ConfigProviderProps>(() => ({
                    theme: this.themeName === "dark" ? darkTheme : null,
                })),
            });
            return dialog;
        },
        timeout(ms: number, key: string, ...name) {
            return new Promise((resolve) => {
                key = `${key}-${name.join("-")}`;
                this.timer[key] && clearTimeout(this.timer[key]);
                if (typeof this.timer[key] !== "undefined") {
                    clearTimeout(this.timer[key]);
                    delete this.timer[key];
                }
                this.timer[key] = setTimeout(resolve, ms);
            });
        },
    },
    persist: piniaPersistConfig("GlobalState"),
})

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
export default pinia
