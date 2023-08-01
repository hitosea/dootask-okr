import http from "../index"
import { PageReq } from "../interface/base"

export const getReplayList = (data) => {
    return http.get<PageReq>("okr/replay/list",data)
}
