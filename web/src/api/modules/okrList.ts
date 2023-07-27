import http from "../index"
import { PageReq } from "../interface/base"


export const getMyList = (data) => {
    return http.get<PageReq>("okr/my/list",data)
}

//更新对齐目标
export const alignUpdate = (data) => {
    return http.post<PageReq>("okr/align/update",data)
}

//获取对齐的目标通过id
export const getAlignDetail = (data) => {
    return http.get<PageReq>("okr/align/detail",data)
}

//取消对齐的目标
export const getAlignCancel = (data) => {
    return http.get<PageReq>("okr/align/cancel",data)
}

//获取OKR详情
export const getOkrDetail = (data) => {
    return http.get<PageReq>("okr/detail",data)
}

//获取OKR详情
export const okrFollow = (data) => {
    return http.get<PageReq>("/okr/follow",data)
}

//更新进度
export const updateProgress = (data) => {
    return http.post<PageReq>("okr/update/progress",data)
}
