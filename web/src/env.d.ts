/// <reference types="vite/client" />

declare module "*.vue" {
    import { DefineComponent } from "vue"
    // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
    const component: DefineComponent<{}, {}, any>
    export default component
}

declare const $t: any

declare interface Window {
    $t: any
    width: any
    height: any
    microApp: any,
}
