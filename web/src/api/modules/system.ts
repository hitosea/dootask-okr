import http from "../index"
import { ResultData } from "../interface/base"
import { System } from "../interface/system"

export const systemSetting = (params?: System.systemSetting) => {
    return http.post<ResultData>("system/setting", params)
}

export const systemSettingMail = (params?: System.systemSettingEmail) => {
    return http.post<ResultData>("system/setting/mail", params)
}
