import { createApp } from "vue"
import App from "./app/index.vue"
import pinia, { GlobalStore } from "./store"
import { routes } from "./routes/routes"
import I18n from "./lang/index"
import createDemoRouter from "./routes"
import "./styles/index.css"
import microappInit from "./microapp"
import directives from "@/directives/index"

const app = createApp(App)
const route = createDemoRouter(app, routes)
app.use(route)
app.use(I18n)
app.use(pinia)
app.use(directives)


window.isEEUiApp = window && window.navigator && /eeui/i.test(window.navigator.userAgent)

GlobalStore().init().then(() => {
    route.isReady().then(() => {
        window.$t = I18n.global.t
        microappInit()
        app.mount("#app")
    })
})