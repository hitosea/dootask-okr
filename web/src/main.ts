import { createApp } from "vue"
import "dootask-tools"
import App from "./layout/index.vue"
import pinia, { GlobalStore } from "./store"
import { routes } from "./routes/routes"
import I18n from "./lang/index"
import createDemoRouter from "./routes"
import "./assets/styles/index.less"
import directives from "@/directives/index"

const app = createApp(App)
const route = createDemoRouter(routes)
app.use(route)
app.use(I18n)
app.use(pinia)
app.use(directives)

// 加载
const globalStore = GlobalStore()
globalStore.init().then(() => {
    // 翻译
    window.$t = I18n.global.t

    // 初始化
    route.isReady().then(() => {
        app.mount("#vite-app")
    })
})
