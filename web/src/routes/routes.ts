import Index from '../views/index.vue'
import Analysis from '../views/analysis.vue'
import Main from '../views/main.vue'
import addOkr from '../views/addOkr.vue'

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
    },
    {
        name: "addOkr",
        path: "/addOkr",
        component: addOkr
    }
];
