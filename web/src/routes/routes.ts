export const routes = [
    {
        name: "index",
        path: "/",
        component: () => import("../views/index.vue"),
    },
    {
        name: "login",
        path: "/login",
        component: () => import("../views/login.vue"),
    },
    {
        name: "manage",
        path: "/manage",
        component: () => import("../views/manage/index.vue"),
        children: [
            {
                name: 'manage-microapp',
                path: 'microapp/:page*', 
                component: () => import("../views/manage/microapp/menu.vue"),
            },
            {
                name: 'manage-dashboard',
                path: 'dashboard', 
                component: () => import("../views/manage/dashboard/index.vue"),
            },
            {
                name: 'manage-calendar',
                path: 'calendar', 
                component: () => import("../views/manage/calendar/index.vue"),
            },
            {
                name: 'manage-messenger',
                path: 'messenger', 
                component: () => import("../views/manage/messenger/index.vue"),
            },
            {
                name: 'manage-file',
                path: 'file', 
                component: () => import("../views/manage/file/index.vue"),
            },
            {
                name: "manage-setting",
                path: "setting",
                component: () => import("../views/manage/setting/index.vue"),
                children: [
                    {
                        name: "manage-setting-personal",
                        path: "personal",
                        component: () => import("../views/manage/setting/personal.vue"),
                    },
                    {
                        name: "manage-setting-language",
                        path: "language",
                        component: () => import("../views/manage/setting/language.vue"),
                    },
                    {
                        name: "manage-setting-theme",
                        path: "theme",
                        component: () => import("../views/manage/setting/theme.vue"),
                    },
                    {
                        name: "manage-setting-password",
                        path: "password",
                        component: () => import("../views/manage/setting/password.vue"),
                    },
                    {
                        name: "manage-setting-email",
                        path: "email",
                        component: () => import("../views/manage/setting/email.vue"),
                    },
                    {
                        name: "manage-setting-system",
                        path: "system",
                        component: () => import("../views/manage/setting/system.vue"),
                    },
                ],
            },
        ],
    },
];
