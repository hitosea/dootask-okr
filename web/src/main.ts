import { createApp } from "vue"
import App from "./layout/index.vue"
import pinia, { GlobalStore } from "./store"
import { routes } from "./routes/routes"
import I18n from "./lang/index"
import createDemoRouter from "./routes"
import "./assets/styles/index.css"
import directives from "@/directives/index"
import { handleMicroData,fixBugForVueRouter4 } from "./microapp"

window.isEEUiApp = window && window.navigator && /eeui/i.test(window.navigator.userAgent)

const app = createApp(App)
const route = createDemoRouter(app, routes)
app.use(route)
app.use(I18n)
app.use(pinia)
app.use(directives)

GlobalStore().init().then(() => {
    route.isReady().then(() => {
        window.$t = I18n.global.t
        app.mount("#vite-app")
        // 与基座进行数据交互
        handleMicroData(route)
        // 用于解决主应用和子应用都是vue-router4时相互冲突，导致点击浏览器返回按钮，路由错误的问题。
        fixBugForVueRouter4(route)
    })
})

// 监听卸载操作
window.addEventListener('unmount', function () {
    app.unmount()
    window.eventCenterForAppNameVite?.clearDataListener()
})
