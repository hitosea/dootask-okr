import http from "../index"
import { Plugin } from "../interface/plugin"

export const getPluginMenu = () => {
    return http.get<Plugin.Menu>("plugin/menu")
}
