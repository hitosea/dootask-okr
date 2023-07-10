export const routes = [
    {
        name: "index",
        path: "/",
        component: () => import('../views/index.vue')
    },
    {
        name: "main",
        path: "/main",
        component:  () => import('../views/main.vue'),
    }
];
