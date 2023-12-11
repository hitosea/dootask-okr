
import { Router } from 'vue-router'
import { GlobalStore } from "@/store"
import { UserStore } from "@/store/user"
import utils from "@/utils/utils"

const initGlobaStore = (data:any) => {
    const globalStore = GlobalStore()
    const userStore = UserStore()
    globalStore.setBaseUrl(data.url)
    globalStore.setThemeName(data.theme == 'auto' ? 'light' : data.theme)
    globalStore.setLanguage(data.languages?.languageType)
    globalStore.setVues(data.vues)
    globalStore.setElectron(data.electron)
    userStore.setUserInfo(data.userInfo)
}

// 数据处理
const handleData = (router: Router,appName:string, data:any) => {
    const globalStore = GlobalStore()
    if(!data){
        return false;
    }
    // 初始化
    if(data.type == "init"){
        initGlobaStore(data)
    }
    // 打开窗口
    if(data.type == "open"){
        if(data.model == 'details' && data.show){
            globalStore.openOkrDetails(data.id || 0)
        }
    }
    // 默认路径
    if (data.path && typeof data.path === 'string') {
        data.path = data.path.replace(/^#/, '')
        // 当基座下发path时进行跳转
        if (data.path) {
            if( data.path == '/app-vite#' || data.path == '/app-vite#/' ){
                router.back()
            }else{
                router.replace(data.path as string)
            }
        }
    }
    // 路由
    if(data.type == 'route'){
        if(data.action == 'back'){
            let is = false;
            if(utils.closeLastModel()){
                is = true;
            }else if(appName == window.eventCenterForAppNameVite?.appName &&  appName == 'micro-app'){
                is = true;
                router.back()
            }
            if(data.callback) data.callback(appName,is);
        }
    }
    // modal可见性
    if(data.type == 'modalVisible'){
        if(data.callback) data.callback(appName,utils.closeLastModel(false));
    }

}

// 与基座进行数据交互
export const handleMicroData = (router: Router) => {
    const eventCenterForMicroApp = window.eventCenterForAppNameVite;
    //
    if (eventCenterForMicroApp) {
        // 设置名称
        GlobalStore().setAppName(eventCenterForMicroApp.appName)
        // 主动获取基座下发的数据
        let info = eventCenterForMicroApp.getData();
        if (info?.type == "init") {
            initGlobaStore(info)
        }
        // 监听基座下发的数据变化
        eventCenterForMicroApp.addDataListener((data: any) => {
            handleData(router, eventCenterForMicroApp.appName, data)
        });
        // 监听基座下发的数据变化 - 全局
        eventCenterForMicroApp.addGlobalDataListener((data: any) => {
            handleData(router, eventCenterForMicroApp.appName, data)
        });
    }
}

/**
 * 用于解决主应用和子应用都是vue-router4时相互冲突，导致点击浏览器返回按钮，路由错误的问题。
 * 相关issue：https://github.com/micro-zoe/micro-app/issues/155
 * 当前vue-router版本：4.0.12
 */
export const fixBugForVueRouter4 = (router: Router) =>{
        // 判断主应用是main-vue3或main-vite，因为这这两个主应用是 vue-router4
        // if (window.location.href.includes('/main-vue3') || window.location.href.includes('/main-vite')) {
        /**
         * 重要说明：
         * 1、这里主应用下发的基础路由为：`/main-xxx/app-vite`，其中 `/main-xxx` 是主应用的基础路由，需要去掉，我们只取`/app-vite`，不同项目根据实际情况调整
         * 2、因为vite关闭了沙箱，又是hash路由，我们这里写死realBaseRoute为：/app-vite#
         */
        // const realBaseRoute = '/app-vite#'

        // router.beforeEach((to,from) => {
        //     if (typeof window.history.state?.current === 'string') {
        //         window.history.state.current = window.history.state.current.replace(new RegExp(realBaseRoute, 'g'), '')
        //     }
        // })

        // router.afterEach(_ => {
        //     if (typeof window.history.state === 'object') {
        //         window.history.state.current = realBaseRoute + (window.history.state.current || '')
        //     }
        // })
    // }
}

