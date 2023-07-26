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