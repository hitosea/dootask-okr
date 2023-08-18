import { createRouter, createWebHashHistory } from 'vue-router'
import { ref } from "vue";
import { UserStore } from "@/store/user"

export const loadingBarApiRef = ref(null)

export default function createDemoRouter(app, routes) {
    const router = createRouter({
        history: createWebHashHistory(),
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
        if (!from || to.path !== from.path) {
            if (loadingBarApiRef.value) {
                loadingBarApiRef.value.finish()
            }
        }
    })

    return router
}
