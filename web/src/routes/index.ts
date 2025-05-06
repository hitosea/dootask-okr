import { createRouter, createWebHistory } from "vue-router"
import { ref } from "vue"
import Empty from "../views/empty.vue"
import Main from "../views/main.vue"
import { getAppData } from "@/utils/app"
import { GlobalStore } from "@/store"

export const loadingBarApiRef = ref(null)

export default function createDemoRouter(routes: any) {
    routes.push({
        name: '/:catchAll(.*)',
        path: '/:catchAll(.*)',
        component: getAppData('initialData.empty') === true ? Empty : Main
    })

    const router = createRouter({
        history: createWebHistory(),
        routes
    })

    router.afterEach(function (to, from) {
        GlobalStore().setBaseRoute(to.params?.catchAll || '')
    })

    return router
}
