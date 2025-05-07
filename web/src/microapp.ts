import { GlobalStore } from "@/store"
import { UserStore } from "@/store/user"
import { addDataListener, getAppData, removeDataListener } from "@/utils/app"
import utils from "@/utils/utils";

export const initAppData = () => {
    const appData = getAppData()
    if (!appData) {
        return
    }
    const {props} = appData
    const globalStore = GlobalStore()
    const userStore = UserStore()

    // 初始化
    globalStore.setBaseUrl(props.baseUrl)
    globalStore.setThemeName(props.themeName == "auto" ? "light" : props.themeName)
    globalStore.setLanguage(props.languageName)
    userStore.setUserInfo(props.userInfo)

    // 窗口监听器
    const dataListener = ({type, props}) => {
        switch (type) {
            case "beforeClose":
                return utils.beforeClose();

            default:
                if (props.type == "details") {
                    globalStore.openOkrDetails(props.id || 0)
                }
                break
        }
    }
    addDataListener(dataListener, true)

    // 返回清理函数，以便可以手动调用
    return () => {
        removeDataListener(dataListener)
    }
}
