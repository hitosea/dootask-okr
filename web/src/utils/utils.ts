import localforage from "localforage"
import webTs from "./web"
localforage.config({ name: 'DooTask', storeName: 'common' })
const utils = {

    /**
     * 简单判断IPv4地址
     * @param value
     */
    apiUrl(url: string) {
        url = '../apps/okr/src' + url;
        return window.$ && window.$.apiUrl ? window.$.apiUrl(url) : url
    },

    /**
     * 简单判断IPv4地址
     * @param value
     */
    isIpv4(value: string) {
        return /^(\d+)\.(\d+)\.(\d+)\.(\d+)$/.test(value)
    },

    /**
     * 判断是否为空
     * @param name
     */
    isEmpty(name) {
        return (
            name === null ||
            name === undefined ||
            name === "null" ||
            name === "undefined" ||
            name.replace(/\s/g, "") === ""
        )
    },

    /**
     * 是否数组
     * @param obj
     * @returns {boolean}
     */
    isArray(obj) {
        return (
            typeof obj == "object" &&
            Object.prototype.toString.call(obj).toLowerCase() == "[object array]" &&
            typeof obj.length == "number"
        )
    },
    /**
     * 获取元素属性
     * @param el
     * @param attrName
     * @param def
     * @returns {Property<any>|string|string}
     */
    getAttr(el, attrName, def = "") {
        return el ? el.getAttribute(attrName) : def;
    },
    /**
     * 是否数组对象
     * @param obj
     * @returns {boolean}
     */
    isJson(obj) {
        return (
            typeof obj == "object" &&
            Object.prototype.toString.call(obj).toLowerCase() == "[object object]" &&
            typeof obj.length == "undefined"
        )
    },

    /**
     * 获取对象值
     * @param obj
     * @param key
     * @returns {*}
     */
    getObject(obj, key) {
        const keys = key.replace(/,/g, "|").replace(/\./g, "|").split("|")
        while (keys.length > 0) {
            const k = keys.shift()
            if (utils.isArray(obj)) {
                obj = obj[utils.parseInt(k)] || ""
            } else if (utils.isJson(obj)) {
                obj = obj[k] || ""
            } else {
                break
            }
        }
        return obj
    },

    /**
     * 转成数字
     * @param param
     * @returns {number|number}
     */
    parseInt(param) {
        const num = parseInt(param)
        return isNaN(num) ? 0 : num
    },

    /**
     * 是否在数组里
     * @param key
     * @param array
     * @param regular
     * @returns {boolean|*}
     */
    inArray(key, array, regular = false) {
        if (!utils.isArray(array)) {
            return false
        }
        if (regular) {
            return !!array.find((item) => {
                if (item && item.indexOf("*")) {
                    const rege = new RegExp(
                        "^" + item.replace(/[-\/\\^$+?.()|[\]{}]/g, "\\$&").replace(/\*/g, ".*") + "$",
                        "g",
                    )
                    if (rege.test(key)) {
                        return true
                    }
                }
                return item == key
            })
        } else {
            return array.includes(key)
        }
    },

    /**
     * 克隆对象
     * @param myObj
     * @returns {*}
     */
    cloneJSON(myObj) {
        if (typeof myObj !== "object") return myObj
        if (myObj === null) return myObj
        //
        return utils.jsonParse(utils.jsonStringify(myObj))
    },

    /**
     * 将一个 JSON 字符串转换为对象（已try）
     * @param str
     * @param defaultVal
     * @returns {*}
     */
    jsonParse(str, defaultVal = undefined) {
        if (str === null) {
            return defaultVal ? defaultVal : {}
        }
        if (typeof str === "object") {
            return str
        }
        try {
            return JSON.parse(str.replace(/\n/g, "\\n").replace(/\r/g, "\\r"))
        } catch (e) {
            return defaultVal ? defaultVal : {}
        }
    },

    /**
     * 将 JavaScript 值转换为 JSON 字符串（已try）
     * @param json
     * @param defaultVal
     * @returns {string}
     */
    jsonStringify(json, defaultVal = undefined) {
        if (typeof json !== "object") {
            return json
        }
        try {
            return JSON.stringify(json)
        } catch (e) {
            return defaultVal ? defaultVal : ""
        }
    },

    /**
     * 字符串是否包含
     * @param string
     * @param find
     * @param lower
     * @returns {boolean}
     */
    strExists(string, find, lower = false) {
        string += ""
        find += ""
        if (lower !== true) {
            string = string.toLowerCase()
            find = find.toLowerCase()
        }
        return string.indexOf(find) !== -1
    },

    /**
     * 字符串是否左边包含
     * @param string
     * @param find
     * @param lower
     * @returns {boolean}
     */
    leftExists(string, find, lower = false) {
        string += ""
        find += ""
        if (lower !== true) {
            string = string.toLowerCase()
            find = find.toLowerCase()
        }
        return string.substring(0, find.length) === find
    },

    /**
     * 删除左边字符串
     * @param string
     * @param find
     * @param lower
     * @returns {string}
     */
    leftDelete(string, find, lower = false) {
        string += ""
        find += ""
        if (utils.leftExists(string, find, lower)) {
            string = string.substring(find.length)
        }
        return string ? string : ""
    },

    /**
     * 字符串是否右边包含
     * @param string
     * @param find
     * @param lower
     * @returns {boolean}
     */
    rightExists(string, find, lower = false) {
        string += ""
        find += ""
        if (lower !== true) {
            string = string.toLowerCase()
            find = find.toLowerCase()
        }
        return string.substring(string.length - find.length) === find
    },

    /**
     * 删除右边字符串
     * @param string
     * @param find
     * @param lower
     * @returns {string}
     */
    rightDelete(string, find, lower = false) {
        string += ""
        find += ""
        if (utils.rightExists(string, find, lower)) {
            string = string.substring(0, string.length - find.length)
        }
        return string ? string : ""
    },

    /**
     * 随机字符串
     * @param len
     */
    randomString(len) {
        len = len || 32
        const $chars = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678oOLl9gqVvUuI1"
        const maxPos = $chars.length
        let pwd = ""
        for (let i = 0; i < len; i++) {
            pwd += $chars.charAt(Math.floor(Math.random() * maxPos))
        }
        return pwd
    },

    /**
     * 检测手机号码格式
     * @param str
     * @returns {boolean}
     */
    isPhone(str) {
        return /^1([3456789])\d{9}$/.test(str)
    },

    /**
     * 检测邮箱地址格式
     * @param email
     * @returns {boolean}
     */
    isEmail(email) {
        return /^([0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*@([0-9a-zA-Z][-\w]*\.)+[a-zA-Z]*)$/i.test(email)
    },

    /**
     * 指定键获取url参数
     * @param key
     * @returns {*}
     */
    urlParameter(key) {
        const params = utils.urlParameterAll()
        return typeof key === "undefined" ? params : params[key]
    },

    urlParameterAll() {
        let search = window.location.search || window.location.hash || ""
        const index = search.indexOf("?")
        if (index !== -1) {
            search = search.substring(index + 1)
        }
        const arr = search.split("&")
        const params = {}
        arr.forEach((item) => {
            // 遍历数组
            const index = item.indexOf("=")
            if (index === -1) {
                params[item] = ""
            } else {
                params[item.substring(0, index)] = item.substring(index + 1)
            }
        })
        return params
    },

    /**
     * 删除地址中的参数
     * @param url
     * @param parameter
     * @returns {string|*}
     */
    removeURLParameter(url, parameter) {
        if (parameter instanceof Array) {
            parameter.forEach((key) => {
                url = utils.removeURLParameter(url, key)
            })
            return url
        }
        const urlParts = url.split("?")
        if (urlParts.length >= 2) {
            //参数名前缀
            const prefix = encodeURIComponent(parameter) + "="
            const pars = urlParts[1].split(/[&;]/g)

            //循环查找匹配参数
            for (let i = pars.length; i-- > 0;) {
                if (pars[i].lastIndexOf(prefix, 0) !== -1) {
                    //存在则删除
                    pars.splice(i, 1)
                }
            }

            return urlParts[0] + (pars.length > 0 ? "?" + pars.join("&") : "")
        }
        return url
    },

    /**
     * 连接加上参数
     * @param url
     * @param params
     * @returns {*}
     */
    urlAddParams(url, params) {
        if (utils.isJson(params)) {
            if (url) {
                url = utils.removeURLParameter(url, Object.keys(params))
            }
            url += ""
            url += url.indexOf("?") === -1 ? "?" : ""
            for (const key in params) {
                if (!params.hasOwnProperty(key)) {
                    continue
                }
                url += "&" + key + "=" + params[key]
            }
        } else if (params) {
            url += (url.indexOf("?") === -1 ? "?" : "&") + params
        }
        if (!url) {
            return ""
        }
        return utils.rightDelete(url.replace("?&", "?"), "?")
    },

    /**
     * 返回10位数时间戳
     * @param v
     * @returns {number}
     * @constructor
     */
    Time(v = undefined) {
        let time
        if (typeof v === "string" && utils.strExists(v, "-")) {
            v = v.replace(/-/g, "/")
            time = new Date(v).getTime()
        } else {
            time = new Date().getTime()
        }
        return Math.round(time / 1000)
    },

    /**
     * 返回 时间对象|时间戳
     * @param v
     * @param stamp 是否返回时间戳
     * @returns {Date|number}
     * @constructor
     */
    Date(v, stamp = false) {
        if (typeof v === "string" && utils.strExists(v, "-")) {
            v = v.replace(/-/g, "/")
        }
        if (stamp === true) {
            return Math.round(new Date(v).getTime() / 1000)
        }
        return new Date(v)
    },

    /**
     * 时间戳转时间格式
     * @param format
     * @param v
     * @returns {string}
     */
    formatDate(format = undefined, v = undefined) {
        if (typeof format === "undefined" || format === "") {
            format = "Y-m-d H:i:s"
        }
        let dateObj
        if (v instanceof Date) {
            dateObj = v
        } else {
            if (typeof v === "undefined") {
                v = new Date().getTime()
            } else if (/^(-)?\d{1,10}$/.test(v)) {
                v = v * 1000
            } else if (/^(-)?\d{1,13}$/.test(v)) {
                v = v * 1000
            } else if (/^(-)?\d{1,14}$/.test(v)) {
                v = v * 100
            } else if (/^(-)?\d{1,15}$/.test(v)) {
                v = v * 10
            } else if (/^(-)?\d{1,16}$/.test(v)) {
                v = v * 1
            } else {
                return v
            }
            dateObj = utils.Date(v)
        }
        //
        format = format.replace(/Y/g, dateObj.getFullYear())
        format = format.replace(/m/g, utils.zeroFill(dateObj.getMonth() + 1, 2))
        format = format.replace(/d/g, utils.zeroFill(dateObj.getDate(), 2))
        format = format.replace(/H/g, utils.zeroFill(dateObj.getHours(), 2))
        format = format.replace(/i/g, utils.zeroFill(dateObj.getMinutes(), 2))
        format = format.replace(/s/g, utils.zeroFill(dateObj.getSeconds(), 2))
        return format
    },

    /**
     * 时间处理
     * @param format
     * @param v
     * @returns {string}
     */
    TimeHandle (time:any,type:number=1):string  {
        if(!time){
            return "";
        }
        if( typeof time == 'number' ){
            return utils.formatDate('Y-m-d 00:00:00', time / 1000)
        }
        return utils.formatDate('Y-m-d',time) + (type == 2 ? ' 23:59:00' : ' 00:00:00')
    },

    /**
     * 接口返回时间转正常格式
     * @param format
     * @param v
     * @returns {string}
     */
    GoDate(format) {
        // 创建一个Date对象
        const dateObj = new Date(format);
        const year = dateObj.getFullYear();
        const month = String(dateObj.getMonth() + 1).padStart(2, "0");
        const day = String(dateObj.getDate()).padStart(2, "0");
        const formattedDate = `${year}/${month}/${day}`;
        return (formattedDate)
    },

    /**
     * 接口返回时间转正常格式
     * @param format
     * @param v
     * @returns {string}
     */
    GoDateHMS(format) {
        const dateObj = new Date(format);
        const year = dateObj.getFullYear();
        const month = String(dateObj.getMonth() + 1).padStart(2, "0");
        const day = String(dateObj.getDate()).padStart(2, "0");
        const hours = String(dateObj.getHours()).padStart(2, "0");
        const minutes = String(dateObj.getMinutes()).padStart(2, "0");
        const seconds = String(dateObj.getSeconds()).padStart(2, "0");
        const formattedDate = `${year}-${month}-${day}`;
        const formattedTime = `${hours}:${minutes}:${seconds}`;
        return (formattedDate + ' ' + formattedTime)
    },

    /**
     * 补零
     * @param str
     * @param length
     * @param after
     * @returns {*}
     */
    zeroFill(str, length, after = false) {
        str += ""
        if (str.length >= length) {
            return str
        }
        let _str = "",
            _ret = ""
        for (let i = 0; i < length; i++) {
            _str += "0"
        }
        if (after) {
            _ret = `${str}${_str}`
            return _ret.substring(0, length)
        } else {
            _ret = `${_str}${str}`
            return _ret.substring(_ret.length - length)
        }
    },
    /**
     * 正则提取域名
     * @param weburl
     * @returns {string|string}
     */
    getDomain(weburl) {
        const urlReg = /http(s)?:\/\/([^\/]+)/i
        const domain = (weburl + "").match(urlReg)
        return domain != null && domain.length > 0 ? domain[2] : ""
    },

    /**
     * 获取屏幕方向
     * @returns {string}
     */
    screenOrientation() {
        try {
            if (typeof window.screen.orientation === "object") {
                return utils.strExists(window.screen.orientation.type, 'portrait') ? 'portrait' : 'landscape'
            }
        } catch (e) {
            //
        }
        return window.width() > window.height() ? "landscape" : "portrait"
    },

    /**
     * 获取数组长度（处理数组不存在）
     * @param array
     * @returns {number|*}
     */
    arrayLength(array) {
        if (array) {
            try {
                return array.length;
            } catch (e) {
                return 0
            }
        }
        return 0;
    },
    /**
 * 相当于 intval
 * @param str
 * @param fixed
 * @returns {number}
 */
    runNum(str, fixed = null) {
        let _s = Number(str);
        if (_s + "" === "NaN") {
            _s = 0;
        }
        if (fixed && /^[0-9]*[1-9][0-9]*$/.test(fixed)) {
            _s = Number(_s.toFixed(fixed));
            let rs = Number(_s.toString().indexOf('.'));
            if (rs < 0) {
                _s = Number(_s.toString() + ".")
                for (let i = 0; i < fixed; i++) {
                    _s = Number(_s.toString() + "0");
                }
            }
        }
        return _s;
    },
    /**
         * 字节转换
         * @param bytes
         * @returns {string}
         */
    bytesToSize(bytes) {
        if (bytes === 0) return '0 B';
        let k = 1024;
        let sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
        let i = Math.floor(Math.log(bytes) / Math.log(k));
        if (typeof sizes[i] === "undefined") {
            return '0 B';
        }
        return utils.runNum((bytes / Math.pow(k, i)), 2) + ' ' + sizes[i];
    },

    /**
         * 动态加载js文件
         * @param url
         * @returns {Promise<unknown>}
         */
    loadScript(url) {
        return new Promise(async (resolve, reject) => {
            url = webTs.originUrl(url)
            if (utils.rightExists(url, '.css')) {
                return resolve(utils.loadCss(url))
            }
            //
            let i = 0
            while (utils.__loadScript[url] === "loading") {
                await new Promise(r => setTimeout(r, 1000))
                i++
                if (i > 30) {
                    return reject("加载超时")
                }
            }
            if (utils.__loadScript[url] === "loaded") {
                return resolve(false)
            }
            utils.__loadScript[url] = "loading"
            //
            const script: any = document.createElement("script")
            script.type = "text/javascript"
            if (script.readyState) {
                script.onreadystatechange = () => {
                    if (script.readyState === "loaded" || script.readyState === "complete") {
                        script.onreadystatechange = null
                        utils.__loadScript[url] = "loaded"
                        resolve(true)
                    }
                }
            } else {
                script.onload = () => {
                    utils.__loadScript[url] = "loaded"
                    resolve(true)
                }
                script.onerror = (e) => {
                    utils.__loadScript[url] = "error"
                    reject(e)
                }
            }
            if (utils.rightExists(url, '.js')) {
                script.src = url + "?hash=" + window.systemInfo.version
            } else {
                script.src = url
            }
            document.body.appendChild(script)
        })
    },
    loadScriptS(urls) {
        return new Promise<void>(resolve => {
            let i = 0
            const recursiveCallback = () => {
                if (++i < urls.length) {
                    utils.loadScript(urls[i]).finally(recursiveCallback)
                } else {
                    resolve()
                }
            }
            utils.loadScript(urls[0]).finally(recursiveCallback)
        })
    },
    __loadScript: {},
    /**
         * 动态加载css文件
         * @param url
         * @returns {Promise<unknown>}
         */
    loadCss(url) {
        return new Promise(async (resolve, reject) => {
            url = webTs.originUrl(url)
            if (utils.rightExists(url, '.js')) {
                return resolve(utils.loadScript(url))
            }
            //
            let i = 0
            while (utils.__loadCss[url] === "loading") {
                await new Promise(r => setTimeout(r, 1000))
                i++
                if (i > 30) {
                    return reject("加载超时")
                }
            }
            if (utils.__loadCss[url] === "loaded") {
                return resolve(false)
            }
            utils.__loadCss[url] = "loading"
            //
            const script: any = document.createElement('link')
            if (script.readyState) {
                script.onreadystatechange = () => {
                    if (script.readyState == 'loaded' || script.readyState == 'complete') {
                        script.onreadystatechange = null
                        utils.__loadCss[url] = "loaded"
                        resolve(true)
                    }
                }
            } else {
                script.onload = () => {
                    utils.__loadCss[url] = "loaded"
                    resolve(true)

                }
                script.onerror = (e) => {
                    utils.__loadCss[url] = "error"
                    reject(e)
                }
            }
            script.rel = 'stylesheet'
            if (utils.rightExists(url, '.css')) {
                script.href = url + "?hash=" + window.systemInfo.version
            } else {
                script.href = url
            }
            document.getElementsByTagName('head').item(0).appendChild(script)
        })
    },
    __loadCss: {},
    /**
        * =============================================================================
        * *****************************   localForage   ******************************
        * =============================================================================
        */

    __IDBTimer: {},
    IDBSave(key, value, delay = 100) {
        if (typeof utils.__IDBTimer[key] !== "undefined") {
            clearTimeout(utils.__IDBTimer[key])
            delete utils.__IDBTimer[key]
        }
        utils.__IDBTimer[key] = setTimeout(async _ => {
            await localforage.setItem(key, value)
        }, delay)
    },

    IDBDel(key) {
        localforage.removeItem(key).then(_ => { })
    },

    IDBSet(key, value) {
        return localforage.setItem(key, value)
    },

    IDBRemove(key) {
        return localforage.removeItem(key)
    },

    IDBClear() {
        return localforage.clear()
    },

    IDBValue(key) {
        return localforage.getItem(key)
    },

    async IDBString(key, def = "") {
        const value = await utils.IDBValue(key)
        return typeof value === "string" || typeof value === "number" ? value : def;
    },

    async IDBInt(key, def = 0) {
        const value = await utils.IDBValue(key)
        return typeof value === "number" ? value : def;
    },

    async IDBBoolean(key, def = false) {
        const value = await utils.IDBValue(key)
        return typeof value === "boolean" ? value : def;
    },

    async IDBArray(key, def = []) {
        const value = await utils.IDBValue(key)
        return utils.isArray(value) ? value : def;
    },

    async IDBJson(key, def = {}) {
        const value = await utils.IDBValue(key)
        return utils.isJson(value) ? value : def;
    },

    /**
             * 刷新当前地址
             * @returns {string}
             */
    reloadUrl() {
        if (window.isEEUiApp && utils.isAndroid()) {
            let url = window.location.href;
            let key = '_='
            let reg = new RegExp(key + '\\d+');
            let timestamp = utils.Time();
            if (url.indexOf(key) > -1) {
                url = url.replace(reg, key + timestamp);
            } else {
                if (url.indexOf('\?') > -1) {
                    let urlArr = url.split('\?');
                    if (urlArr[1]) {
                        url = urlArr[0] + '?' + key + timestamp + '&' + urlArr[1];
                    } else {
                        url = urlArr[0] + '?' + key + timestamp;
                    }
                } else {
                    if (url.indexOf('#') > -1) {
                        url = url.split('#')[0] + '?' + key + timestamp + location.hash;
                    } else {
                        url = url + '?' + key + timestamp;
                    }
                }
            }
        } else {
            window.location.reload();
        }
    },
    /**
     * 是否安卓
     * @returns {boolean|string}
     */
    isAndroid() {
        let ua = typeof window !== 'undefined' && window.navigator.userAgent.toLowerCase();
        return ua && ua.indexOf('android') > 0;
    },


    /**
     * 等比缩放尺寸
     * @param width
     * @param height
     * @param maxWidth
     * @param maxHeight
     * @returns {{width, height}|{width: number, height: number}}
     */
    scaleToScale(width, height, maxWidth, maxHeight) {
        let tempWidth;
        let tempHeight;
        if (width > 0 && height > 0) {
            if (width / height >= maxWidth / maxHeight) {
                if (width > maxWidth) {
                    tempWidth = maxWidth;
                    tempHeight = (height * maxWidth) / width;
                } else {
                    tempWidth = width;
                    tempHeight = height;
                }
            } else {
                if (height > maxHeight) {
                    tempHeight = maxHeight;
                    tempWidth = (width * maxHeight) / height;
                } else {
                    tempWidth = width;
                    tempHeight = height;
                }
            }
            return { width: parseInt(tempWidth), height: parseInt(tempHeight) };
        }
        return { width, height };
    },

    /**
     * 随机获取范围
     * @param Min
     * @param Max
     * @returns {*}
     */
    randNum(Min, Max) {
        let Range = Max - Min;
        let Rand = Math.random();
        return Min + Math.round(Rand * Range); //四舍五入
    },
    /**
     * 获取文本长度
     * @param string
     * @returns {number}
     */
    stringLength(string) {
        if (typeof string === "number" || typeof string === "string") {
            return (string + "").length
        }
        return 0;
    },

    /**
     * 阻止滑动穿透
     * @param el
     */
    scrollPreventThrough(el) {
        if (!el) {
            return;
        }
        if (el.getAttribute("data-prevent-through") === "yes") {
            return;
        }
        el.setAttribute("data-prevent-through", "yes")
        //
        let targetY = null;
        el.addEventListener('touchstart', function (e) {
            targetY = Math.floor(e.targetTouches[0].clientY);
        });
        el.addEventListener('touchmove', function (e) {
            // 检测可滚动区域的滚动事件，如果滑到了顶部或底部，阻止默认事件
            let NewTargetY = Math.floor(e.targetTouches[0].clientY),    //本次移动时鼠标的位置，用于计算
                sTop = el.scrollTop,        //当前滚动的距离
                sH = el.scrollHeight,       //可滚动区域的高度
                lyBoxH = el.clientHeight;   //可视区域的高度
            if (sTop <= 0 && NewTargetY - targetY > 0) {
                // 下拉页面到顶
                e.preventDefault();
            } else if (sTop >= sH - lyBoxH && NewTargetY - targetY < 0) {
                // 上翻页面到底
                e.preventDefault();
            }
        }, false);
    },
    /**
         * =============================================================================
         * *****************************   sessionStorage   ****************************
         * =============================================================================
         */

    setSessionStorage(key?: any, value?: any) {
        return utils.__operationSessionStorage(key, value);
    },

    getSessionStorageValue(key?: any) {
        return utils.__operationSessionStorage(key);
    },

    getSessionStorageString(key: any, def = '') {
        let value = utils.__operationSessionStorage(key);
        return typeof value === "string" || typeof value === "number" ? value : def;
    },

    getSessionStorageInt(key, def = 0) {
        let value = utils.__operationSessionStorage(key);
        return typeof value === "number" ? value : def;
    },

    __operationSessionStorage(key?: any, value?: any) {
        if (!key) {
            return;
        }
        let keyName = '__state__';
        if (key.substring(0, 5) === 'cache') {
            keyName = '__state:' + key + '__';
        }
        if (typeof value === 'undefined') {
            return utils.__loadFromlSession(key, '', keyName);
        } else {
            utils.__savaToSession(key, value, keyName);
        }
    },

    __savaToSession(key?: any, value?: any, keyName?: any) {
        try {
            if (typeof keyName === 'undefined') keyName = '__seller__';
            let seller: any = window.sessionStorage.getItem(keyName);
            if (!seller) {
                seller = {};
            } else {
                seller = JSON.parse(seller);
            }
            seller[key] = value;
            window.sessionStorage.setItem(keyName, JSON.stringify(seller))
        } catch (e) {
        }
    },

    __loadFromlSession(key?: any, def?: any, keyName?: any) {
        try {
            if (typeof keyName === 'undefined') keyName = '__seller__';
            let seller = window.sessionStorage.getItem(keyName);
            if (!seller) {
                return def;
            }
            seller = JSON.parse(seller);
            if (!seller || typeof seller[key] === 'undefined') {
                return def;
            }
            return seller[key];
        } catch (e) {
            return def;
        }
    },



}

export default utils
