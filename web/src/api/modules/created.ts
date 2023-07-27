import http from "../index"
import { PageReq } from "../interface/base"
import * as Created from "../interface/created"

//创建OKR
export const addOkr = (params: Created.addOkr) => {
    return http.post("okr/create",params)
}
//更新OKR
export const upDateOkr = (params: Created.addOkr) => {
    return http.post("okr/update",params)
}

//获取能对齐的目标
export const getAlignList = (data) => {
    return http.get<PageReq>("okr/align/list",data)
}

//项目列表
export const getProjectList = (data) => {
    return http.get<any>("okr/project/list",data)
}

//用户列表
export const getUserList = (data) => {
    return http.get<any>("okr/user/list",data)
}
