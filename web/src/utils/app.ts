/**
 * 检查当前应用是否为微前端应用
 * @returns {boolean} 如果当前应用是微前端应用，返回true；否则返回false
 */
export const isMicroApp = (): boolean => {
    return window.microApp && window.microApp.getData
}

/**
 * 获取下一个可用的模态框 z-index 值
 * @returns {number} 返回一个递增的 z-index 值
 */
export const nextModalIndex = (): number => {
    const appData = window.microApp?.getData()
    if (!appData) return 1000

    if (typeof appData.nextModalIndex === "function") {
        return appData.nextModalIndex()
    }
    return 1000
}

/**
 * 关闭微前端应用
 * @param destroy - 可选参数，布尔值，表示是否销毁应用。默认为false。
 */
export const handleCloseApp = (destroy = false): void => {
    const appData = window.microApp?.getData()
    if (!appData) return

    if (typeof appData.handleClose === "function") {
        appData.handleClose(destroy)
    }
}

/**
 * 逐步返回上一个页面
 * @description 类此似于浏览器的后退按钮，返回到最后一个页面时会关闭应用。
 */
export const handleBackApp = (): void => {
    const appData = window.microApp?.getData()
    if (!appData) return

    if (typeof appData.handleBack === "function") {
        appData.handleBack()
    }
}

/**
 * 添加数据监听器
 * @param callback - 回调函数，当数据发生变化时调用
 * @param autoTrigger - 在初次绑定监听函数时如果有缓存数据，是否需要主动触发一次
 */
export const addDataListener = (callback: Function, autoTrigger = false): void => {
    window.microApp?.addDataListener(callback, autoTrigger)
}

/**
 * 移除数据监听器
 * @param callback - 回调函数，之前添加的监听器
 */
export const removeDataListener = (callback: Function): void => {
    window.microApp?.removeDataListener(callback)
}

/**
 * 从微前端父应用获取共享数据
 * @param key - 可选参数，指定要获取的数据键名。支持：
 *             1. 普通键名（如'user'）
 *             2. 带点符号的路径（如'user.info.name'）用于获取嵌套值
 *             3. 数组索引（如'items.0'）用于获取数组元素
 * @returns 当不传key时返回全部共享数据；传key时返回对应值（如果不存在则返回null）
 * @example
 * // 获取全部数据
 * const allData = getAppData();
 *
 * // 获取对象嵌套值
 * const userName = getAppData('user.name');
 *
 * // 获取数组元素
 * const firstItem = getAppData('items.0');
 *
 * // 获取不存在的键值
 * const notExist = getAppData('some.nonexistent.key'); // 返回null
 */
export const getAppData = (key: string | null = null): any => {
    const appData = window.microApp?.getData()
    if (!appData) return null
    if (!key) return appData

    return key.split(".").reduce((obj, k) => {
        if (obj && typeof obj === "object") {
            // 处理数组索引（如 items.0）
            const arrayIndex = /^\d+$/.test(k) ? parseInt(k) : k
            return obj[arrayIndex]
        }
        return null
    }, appData)
}
