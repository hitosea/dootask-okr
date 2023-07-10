/// <reference types="vite/client" />

declare module "*.vue" {
    import type { DefineComponent } from "vue"
    // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
    const component: DefineComponent<{}, {}, any>
    export default component
}

declare const $t: any

declare interface Window {
    systemInfo: any
    $t: any
    width: any
    height: any
    isEEUiApp: any
    eventCenterForAppNameVite: any
    __MICRO_APP_NAME__: string
    __MICRO_APP_ENVIRONMENT__: string
    __MICRO_APP_BASE_APPLICATION__: string
}
