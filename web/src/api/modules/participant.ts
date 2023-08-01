import http from "../index"
import { PageReq } from "../interface/base"

export const getParticipantList = (data) => {
    return http.get<PageReq>("okr/participant/list",data)
}