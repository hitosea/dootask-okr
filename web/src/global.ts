declare module '@vue/runtime-core' {
    export interface ComponentCustomProperties {
        $globalStore: any
    }
}

const initGlobal = (app:any, globalStore:any) => {
    // 全局状态
    app.config.globalProperties.$globalStore = globalStore
}

export default initGlobal


