declare module '@vue/runtime-core' {
    export interface ComponentCustomProperties {
        $openChildPage: any;
        $globalStore: any;
    }
}

const initGlobal = (app:any, route:any, globalStore:any) => {
    // 打开子页面
    app.config.globalProperties.$openChildPage = (path: string, query:any = {}) => {
        if (window.innerWidth < 768 && globalStore.isPortrait()) {
            route.push({
                path: globalStore.baseRoute + path,
                query: query,
            })
            return false;
        }else {
            return true;
        }
    }
    // 全局状态
    app.config.globalProperties.$globalStore = globalStore
}

export default initGlobal


