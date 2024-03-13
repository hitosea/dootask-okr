import http from "../index"
import { PageReq } from "../interface/base"

export const getCompanyOkrList = (data) => {
    return http.get<PageReq>("/okr/company/list",data)
}

