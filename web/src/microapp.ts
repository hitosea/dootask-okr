import { GlobalStore } from "@/store"
import { UserStore } from "@/store/user"
import { addDataListener, getAppData, removeDataListener } from "@/utils/app"

// 与基座进行数据交互
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
    globalStore.setIsElectron(initialData.isElectron)
    globalStore.setThemeName(initialData.themeName == "auto" ? "light" : initialData.themeName)
    globalStore.setLanguage(initialData.languages.languageName)
    userStore.setUserInfo(initialData.userInfo)

    // 打开窗口监听器
    const dataListener = ({initialData}) => {
        if (initialData.type == "details") {
            globalStore.openOkrDetails(initialData.id || 0)
        }
    }
    addDataListener(dataListener, true)

    // 返回清理函数，以便可以手动调用
    return () => {
        removeDataListener(dataListener)
    }
}
