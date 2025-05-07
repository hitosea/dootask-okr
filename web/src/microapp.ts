import { GlobalStore } from "@/store"
import { UserStore } from "@/store/user"
import { addDataListener, getAppData, removeDataListener } from "@/utils/app"
import utils from "@/utils/utils";

export const initAppData = () => {
    const appData = getAppData()
    if (!appData) {
        return
    }
    const {initialData} = appData
    const globalStore = GlobalStore()
    const userStore = UserStore()

    // 初始化
    globalStore.setBaseUrl(initialData.baseUrl)
    globalStore.setThemeName(initialData.themeName == "auto" ? "light" : initialData.themeName)
    globalStore.setLanguage(initialData.languages.languageName)
    userStore.setUserInfo(initialData.userInfo)

    // 窗口监听器
    const dataListener = ({type, initialData}) => {
        switch (type) {
            case "beforeClose":
                return utils.beforeClose();

            default:
                if (initialData.type == "details") {
                    globalStore.openOkrDetails(initialData.id || 0)
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
