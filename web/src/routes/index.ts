import { createRouter, createWebHistory } from "vue-router"
import { ref } from "vue"

export const loadingBarApiRef = ref(null)

export default function createDemoRouter(routes: any) {
    return createRouter({
        history: createWebHistory(),
        routes
    })
}
