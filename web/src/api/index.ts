import axios, {AxiosInstance, AxiosRequestConfig} from 'axios'
import utils from "../utils/utils";
import {GlobalStore} from "../store";
import {ResultData} from "./interface/base";
import {CODE} from "./constant";

const config = {
    baseURL: import.meta.env.VITE_API_URL as string || '/microapp/okr/api/v1', // 所有的请求地址前缀部分
    timeout: 60000, // 请求超时时间毫秒
    withCredentials: true, // 异步请求携带cookie
    headers: {
        // 设置后端需要的传参类型
        'Content-Type': 'application/json',     
    },
}

class RequestHttp {
    // 定义成员变量并指定类型
    service: AxiosInstance;

    public constructor(config: AxiosRequestConfig) {
        // 实例化axios
        this.service = axios.create(config);

        /**
         * 请求拦截器
         * 客户端发送请求 -> [请求拦截器] -> 服务器
         * token校验(JWT) : 接受服务器返回的token,存储到vuex/pinia/本地储存当中
         */
        this.service.interceptors.request.use(
            function (config) {
                const userInfo = JSON.parse(localStorage.getItem("UserState"))
                config.headers.Token = userInfo?.info?.token
                return config
            },
            function (error) {
                // 对请求错误做些什么
                // console.log(error)
                return Promise.reject(error)
            }
        )

        /**
         * 响应拦截器
         * 服务器换返回信息 -> [拦截统一处理] -> 客户端JS获取到信息
         */
        this.service.interceptors.response.use(
            function (response) {
                // console.log(response)
                // 2xx 范围内的状态码都会触发该函数。
                // 对响应数据做点什么
                // dataAxios 是 axios 返回数据中的 data
                const dataAxios = response.data
                //
                if (!utils.isJson(dataAxios)) {
                    return Promise.reject({code: CODE.StatusInternalServerError, msg: $t('数据格式错误'), data: dataAxios})
                }
                if (dataAxios.code !== CODE.StatusOK) {
                    return Promise.reject(dataAxios)
                }
                return dataAxios
            },
            function (error) {
                // 超出 2xx 范围的状态码都会触发该函数。
                // 对响应错误做点什么
                // console.log(error)
                return Promise.reject({code: CODE.StatusInternalServerError, msg: $t('请求失败'), data: error})
            }
        )
    }

    // 自定义方法封装（常用请求）
    get<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.get(url, {params});
    }

    post<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.post(url, params);
    }

    put<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.put(url, params);
    }

    delete<T>(url: string, params?: object): Promise<ResultData<T>> {
        return this.service.delete(url, {params});
    }
}

export default new RequestHttp(config);

export function ResultDialog({code, msg, data}, dialogOptions = {}) {
    let title = $t('提示')
    let content = msg
    if (code !== CODE.StatusOK) {
        title = $t('错误')
        if (utils.isJson(data) && (data.err || data.error)) {
            title = msg
            content = data.err || data.error
        }
    }
    let options = {
        title,
        content,
        positiveText: $t('确定'),
    }
    if (utils.isJson(data) && utils.isJson(data.dialog)) {
        options = Object.assign(options, data.dialog)
    }
    if (utils.isJson(dialogOptions)) {
        options = Object.assign(options, dialogOptions)
    }
    const clobalStore = GlobalStore()
    if (code === CODE.StatusOK) {
        return clobalStore.dialogProvider().success(options)
    } else {
        return clobalStore.dialogProvider().error(options)
    }
}
