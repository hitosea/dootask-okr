import {createRouter, createWebHistory} from 'vue-router'
import {ref} from "vue";

export const loadingBarApiRef = ref(null)

export default function createDemoRouter(app, routes) {
    const router = createRouter({
        history: createWebHistory(),
        routes
    })
    router.beforeEach(function (to, from, next) {
        if (!from || to.path !== from.path) {
            if (loadingBarApiRef.value) {
                loadingBarApiRef.value.start()
            }
        }
        next()
    })

    router.afterEach(function (to, from) {
        if (!from || to.path !== from.path) {
            if (loadingBarApiRef.value) {
                loadingBarApiRef.value.finish()
            }
        }
    })
    return router
}
