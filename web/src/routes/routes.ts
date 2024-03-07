import Index from '../views/index.vue'
import Analysis from '../views/analysis.vue'
import addOkr from '../views/childpage/addOkr.vue'
import okrDetails from '../views/childpage/okrDetails.vue'
import addMultiple from '../views/childpage/addMultiple.vue'
import dimission from '../views/childpage/dimission.vue'
import archive from '../views/childpage/archive.vue'
import setting from '../views/childpage/setting.vue'
import multipleDetails from '../views/childpage/multipleDetails.vue'
import I18n from "@/lang/index"

export const routes = [
    {
        name: "List",
        path: "/:catchAll(.*)/list",
        meta: { title: 'OKR ' + I18n.global.t('管理')},
        component: Index
    },
    {
        name: "Analysis",
        path: "/:catchAll(.*)/analysis",
        meta: { title: 'OKR ' + I18n.global.t('结果分析')},
        component: Analysis
    },
    {
        name: "addOkr",
        path: "/:catchAll(.*)/addOkr",
        meta: { title:  I18n.global.t('添加/编辑') + ' OKR'},
        component: addOkr
    },
    {
        name: "okrDetails",
        path: "/:catchAll(.*)/okrDetails",
        meta: { title:  I18n.global.t('OKR详情')},
        component: okrDetails
    },
    {
        name: "addMultiple",
        path: "/:catchAll(.*)/addMultiple",
        meta: { title:  I18n.global.t('复盘添加/详情')},
        component: addMultiple
    },
    {
        name: "deletePersonnel",
        path: "/:catchAll(.*)/deletePersonnel",
        meta: { title:  I18n.global.t('离职/删除人员') + ' OKR'},
        component: dimission
    },
    {
        name: "archive",
        path: "/:catchAll(.*)/archive",
        meta: { title:  I18n.global.t('已归档') + ' OKR'},
        component: archive
    },
    {
        name: "setting",
        path: "/:catchAll(.*)/setting",
        meta: { title: 'OKR ' + I18n.global.t('设置')},
        component: setting
    },
    {
        name: "multipleDetails",
        path: "/:catchAll(.*)/multipleDetails",
        meta: { title:  I18n.global.t('复盘详情')},
        component: multipleDetails
    }
]
