import Index from '../views/index.vue'
import Analysis from '../views/analysis.vue'
import Main from '../views/main.vue'
import addOkr from '../views/addOkr.vue'
import okrDetails from '../views/okrDetails.vue'
import addMultiple from '../views/addMultiple.vue'

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
    },
    {
        name: "okrDetails",
        path: "/okrDetails",
        component: okrDetails
    },
    {
        name: "addMultiple",
        path: "/addMultiple",
        component: addMultiple
    }
];
