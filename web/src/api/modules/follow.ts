import http from "../index"
import { PageReq } from "../interface/base"

export const getFollowList = (data) => {
    return http.get<PageReq>("okr/follow/list",data)
}