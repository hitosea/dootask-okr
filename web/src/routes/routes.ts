import Main from '@/views/main.vue'
import Index from '@/views/index.vue'
import Analysis from '@/views/analysis.vue'
import okrDetails from '@/views/single/okrDetails.vue'
import I18n from '@/lang/index'

export const routes = [
    {
        name: 'main',
        path: '/:catchAll(.*)',
        component: Main
    },
    {
        name: "list",
        path: "/:catchAll(.*)/list",
        meta: { title: 'OKR ' + I18n.global.t('管理')},
        component: Index
    },
    {
        name: "analysis",
        path: "/:catchAll(.*)/analysis",
        meta: { title: 'OKR ' + I18n.global.t('结果分析')},
        component: Analysis
    },
    {
        name: "okrDetails",
        path: "/:catchAll(.*)/okrDetails",
        meta: { title:  I18n.global.t('OKR详情')},
        component: okrDetails
    }
]
