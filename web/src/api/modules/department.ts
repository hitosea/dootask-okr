import http from "../index"
import { PageReq } from "../interface/base"

export const getDepartmentOkrList = (data) => {
    return http.get<PageReq>("okr/department/list",data)
}

export const getDepartmentList = () => {
    return http.get<PageReq>("okr/department/search")
}