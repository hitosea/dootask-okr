import Index from '../views/index.vue'
import Analysis from '../views/analysis.vue'
import Main from '../views/main.vue'

export const routes = [
    {
        name: "OkrDetails",
        path: "/",
        component: Main
    },
    {
        name: "List",
        path: "/list",
        component: Index
    },
    {
        name: "Analysis",
        path: "/analysis",
        component: Analysis
    }
];
