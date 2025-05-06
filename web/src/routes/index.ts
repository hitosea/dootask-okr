import { createRouter, createWebHistory } from "vue-router"
import { ref } from "vue"
import Main from "../views/main.vue"

export const loadingBarApiRef = ref(null)

export default function createDemoRouter(routes: any) {
    routes.push({
        name: '/:catchAll(.*)',
        path: '/:catchAll(.*)',
        component: Main
    })

    return createRouter({
        history: createWebHistory(),
        routes
    })
}
