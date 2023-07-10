import http from "../index"
import { ResultData } from "../interface/base"
import { User } from "../interface/user"

export const getUserInfo = () => {
    return http.get<User.Info>("user/info")
}

export const userLogin = (params: User.LoginReq) => {
    return http.post<User.Info>("user/login", params)
}

export const userReg = (params: User.RegReq) => {
    return http.post<User.Info>("user/reg", params)
}

export const needCode = (params: User.needCode) => {
    return http.get<boolean>("user/login/needcode", params)
}

export const codeImg = () => {
    return http.get<User.codeImg>("user/login/codeimg")
}

export const getQrcodeStatus = (params: User.qrcode) => {
    return http.post<User.Info>("user/login/qrcode", params)
}

export const editPassword = (params: User.resetPassword) => {
    return http.post<ResultData>("user/edit/password", params)
}

export const sendEmail = (params: User.sendEmailCode) => {
    return http.post<ResultData>("user/edit/password", params)
}

export const changeEmail = (params: User.changeEmail) => {
    return http.post<ResultData>("user/edit/password", params)
}

