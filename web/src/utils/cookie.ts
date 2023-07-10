import Cookies from "js-cookie";

const cookie = {
    /**
     * 获取cookie
     * @param name
     * @param defaultVal
     * @returns {string}
     * @constructor
     */
    get(name, defaultVal = "") {
        return decodeURIComponent(Cookies.get(name)) || defaultVal
    },

    /**
     * 设置cookie
     * @param name
     * @param value
     * @constructor
     */
    set(name, value) {
        Cookies.set(name, encodeURIComponent(value))
    },

    /**
     * 删除cookie
     * @param name
     * @constructor
     */
    remove(name) {
        Cookies.remove(name)
    },
}

export default cookie
