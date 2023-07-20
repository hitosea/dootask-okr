import Index from '../views/index.vue'
import Analysis from '../views/analysis.vue'

export const routes = [
    {
        name: "index",
        path: "/",
        component: Index
    },
    {
        name: "analysis",
        path: "/analysis",
        component: Analysis
    }
];
