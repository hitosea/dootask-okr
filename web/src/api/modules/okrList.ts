import http from "../index"
import { PageReq ,replayData} from "../interface/base"


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

//更新信心
export const confidenceUpdate = (data) => {
    return http.post<PageReq>("okr/confidence/update",data)
}

//评分
export const okrScore = (data) => {
    return http.post<PageReq>("okr/score",data)
}

//获取动态
export const getLogList = (data) => {
    return http.get<PageReq>("okr/log/list",data)
}

//获取复盘
export const getReplayList = (data) => {
    return http.get<PageReq>("okr/replay/okr/list",data)
}

//添加复盘
export const replayCreate = (data) => {
    return http.post<PageReq>("okr/replay/create",data)
}
//复盘详情
export const replayDetail = (data) => {
    return http.post<replayData>("okr/replay/detail",data)
}

//复盘上级评论
export const superiorReview = (data) => {
    return http.post<PageReq>("/okr/replay/superior/review",data)
}

//更新参与人
export const participantUpdate = (data) => {
    return http.post<PageReq>("okr/participant/update",data)
}

//取消重启目标
export const okrCancel = (data) => {
    return http.get<PageReq>("okr/cancel",data)
}

//okr离职人员列表
export const getLeaveList= (data) => {
    return http.get<PageReq>("okr/leave/list",data)
}

//okr删除
export const okrDelete = (id) => {
    return http.get<PageReq>("okr/delete",id)
}

//okr分配
export const okrAssign = (data) => {
    return http.post<PageReq>("okr/leave/update",data)
}

//操作归档
export const okrArchive = (id) => {
    return http.get<PageReq>("okr/archive",id)
}

//获取归档列表
export const getOkrArchive = (data) => {
    return http.get<PageReq>("okr/archive/list",data)
}

//还原归档目标
export const okrArchiveRestore = (id) => {
    return http.get<PageReq>("okr/archive/restore",id)
}

