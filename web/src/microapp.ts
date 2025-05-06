import { GlobalStore } from "@/store"
import { UserStore } from "@/store/user"
import { addDataListener, getAppData } from "@/utils/app"

// 与基座进行数据交互
export const handleMicroData = () => {
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

    // 打开窗口
    addDataListener(({initialData}) => {
        if (initialData.type == "details") {
            globalStore.openOkrDetails(initialData.id || 0)
        }
    }, true)
}
