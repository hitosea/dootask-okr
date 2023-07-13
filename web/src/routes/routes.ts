import Index from '../views/index.vue'
import Main from '../views/main.vue'

export const routes = [
    {
        name: "index",
        path: "/",
        component: Index
    },
    {
        name: "main",
        path: "/main",
        component: Main
    }
];
