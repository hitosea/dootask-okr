import { GlobalStore } from "@/store"
import { UserStore } from "@/store/user"
import { getAppData } from "@/utils/app"

const initGlobalStore = ({ initialData }: any) => {
    const globalStore = GlobalStore()
    const userStore = UserStore()
    globalStore.setBaseUrl(initialData.url)
    globalStore.setIsElectron(initialData.isElectron)
    globalStore.setThemeName(initialData.themeName == "auto" ? "light" : initialData.themeName)
    globalStore.setLanguage(initialData.languages?.languageName)
    userStore.setUserInfo(initialData.userInfo)
}

// 与基座进行数据交互
export const handleMicroData = () => {
    const appData = getAppData()
    if (!appData) {
        return
    }

    // 初始化
    initGlobalStore(appData)

    // 打开窗口
    if (appData.type == "details") {
        const globalStore = GlobalStore()
        globalStore.openOkrDetails(appData.initialData.id || 0)
    }
}
