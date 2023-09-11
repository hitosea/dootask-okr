import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'
import { ref } from "vue";
import { UserStore } from "@/store/user"
import { GlobalStore } from "@/store"
import Empty from '../views/empty.vue'
import Main from '../views/main.vue'

export const loadingBarApiRef = ref(null)

export default function createDemoRouter(app, routes) {

    const isDetails = window.eventCenterForAppNameVite?.appName == 'okr-details' ? true : false
    const isElectron = !!(window && window.process && window.process?.type);
    const isEEUiApp = window && window.navigator && /eeui/i.test(window.navigator.userAgent);

    routes.push({
        name: '/:catchAll(.*)',
        path: '/:catchAll(.*)',
        component: isDetails ? Main : Empty
    })

    const router = createRouter({
        history: isElectron || isEEUiApp ? createWebHashHistory() :  createWebHistory(),
        routes
    })

    router.beforeEach(function (to, from, next) {
        if(to.query.token){
            UserStore().setToken(to.query.token)
        }
        if (!from || to.path !== from.path) {
            if (loadingBarApiRef.value) {
                loadingBarApiRef.value.start()
            }
        }
        next()
    })

    router.afterEach(function (to, from) {
        if(to?.meta?.title && window.eventCenterForAppNameVite?.appName == "micro-app"){
            document.title = to.meta.title
        }
        GlobalStore().setBaseRoute(to.params?.catchAll || '')
        if (!from || to.path !== from.path) {
            if (loadingBarApiRef.value) {
                loadingBarApiRef.value.finish()
            }
        }
    })

    return router
}
