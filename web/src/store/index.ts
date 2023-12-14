import { computed } from "vue";
import { createPinia, defineStore } from "pinia"
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import { ConfigProviderProps, createDiscreteApi, darkTheme, useOsTheme } from "naive-ui";
import { GlobalState } from "./interface"
import piniaPersistConfig from "./config/pinia-persist"
import { I18nGlobal } from "@/lang"

export const GlobalStore = defineStore({
    id: "GlobalState",
    state: (): GlobalState => ({
        appName: "",
        baseUrl: "",
        baseRoute: "",
        isLoading: 0,
        language: "zh",
        themeName: "",
        timer: {},
        electron: null,
        //
        modalList: [],
        // 浏览器窗口方向
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
        addMultipleChange: false,
        multipleId: 0,
        addMultipleData: null,
        okrEditData: null,
        okrEdit: false,
        doubleSkip: null,
    }),
    actions: {
        async init() {
            this.isLoading = 0;
            if (["light", "dark"].indexOf(this.themeName) === -1) {
                this.themeName = useOsTheme().value;
            }
        },
        setAppName(val) {
            this.appName = val
        },
        openOkrDetails(id) {
            this.okrDetail = { show: false, id: id  }
            this.okrDetail = { show: true, id: id }
        },
        setBaseUrl(url) {
            if(url){
                if(url.indexOf('http://127.0.0.1:5567') !== -1){
                    url = 'http://127.0.0.1:5566'
                }
                this.baseUrl = (new URL(url)).origin
            }
        },
        setBaseRoute(route) {
            if(route){
                this.baseRoute = '/' + route
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
        setElectron(electron: any) {
            this.electron = electron;
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
        okrSetup() {
            return {
                okrEditData: computed(() => {
                    return this.okrEditData;
                }),
                okrEdit: computed(() => {
                    return this.okrEdit;
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
        isSingle(){
            return  document.querySelector('.electron-single-micro-apps') ? 1 : 0 ;
        },
        isPortrait(){
            return document.querySelector('.window-portrait') ? 1 : 0;
        }
    },
    persist: piniaPersistConfig("GlobalState"),
});

//
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
export default pinia
