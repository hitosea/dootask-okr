import utils from "./utils";

const result = {
    /**
     * 获取返回码
     * @returns {*|number}
     */
    code() {
        return utils.parseInt(utils.urlParameter("result_code") || window['result_code'])
    },

    /**
     * 获取返回消息
     * @returns {*|string|string}
     */
    msg() {
        return decodeURIComponent(utils.urlParameter("result_msg") || window['result_msg']).replace("{{.MSG}}", "")
    },

    /**
     * 获取返回数据
     * @returns {*|string|string}
     */
    data() {
        return decodeURIComponent(utils.urlParameter("result_data") || window['result_data']).replace("{{.DATA}}", "")
    },
}

export default result
