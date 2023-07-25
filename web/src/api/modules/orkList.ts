import http from "../index"
import { PageReq } from "../interface/base"


export const getMyList = (data) => {
    return http.get<PageReq>("okr/my/list",data)
}
