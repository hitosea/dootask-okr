import Index from '../views/index.vue'
import Analysis from '../views/analysis.vue'
import Main from '../views/main.vue'
import addOkr from '../views/addOkr.vue'
import okrDetails from '../views/okrDetails.vue'
import addMultiple from '../views/addMultiple.vue'
import I18n from "@/lang/index"

export const routes = [
    {
        name: "OkrDetails",
        path: "/",
        meta: { title:  I18n.global.t('OKR管理')},
        component: Main
    },
    {
        name: "List",
        path: "/list",
        meta: { title:  I18n.global.t('OKR管理')},
        component: Index
    },
    {
        name: "Analysis",
        path: "/analysis",
        meta: { title:  I18n.global.t('OKR结果分析')},
        component: Analysis
    },

    {
        name: "addOkr",
        path: "/addOkr",
        meta: { title:  I18n.global.t('添加/编辑OKR')},
        component: addOkr
    },
    {
        name: "okrDetails",
        path: "/okrDetails",
        meta: { title:  I18n.global.t('OKR详情')},
        component: okrDetails
    },
    {
        name: "addMultiple",
        path: "/addMultiple",
        meta: { title:  I18n.global.t('添加/编辑复盘')},
        component: addMultiple
    }
];
