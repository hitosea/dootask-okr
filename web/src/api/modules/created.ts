import http from "../index"
import { PageReq } from "../interface/base"
import * as Created from "../interface/created"


export const addOkr = (params: Created.addOkr) => {
    return http.post("okr/create",params)
}

export const getAlignList = (data) => {
    return http.get<PageReq>("okr/align/list",data)
}


export const getProjectList = (data) => {
    return http.get<any>("okr/project/list",data)
}
export const getUserList = (data) => {
    return http.get<any>("okr/user/list",data)
}
