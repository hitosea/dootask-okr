package interfaces

import (
	"dootask-okr/app/model"
)

// 基础OKR请求
type OkrBaseReq struct {
	Title          string `json:"title" binding:"required"`    // 目标 （必填）
	Type           int    `json:"type" binding:"required"`     // 类型 1-承诺型 2-挑战型 （必填）
	Priority       string `json:"priority" binding:"required"` // 优先级 （必填）
	VisibleRange   int    `json:"visible_range"`               // 可见范围  1-全公司 2-仅相关成员 3-仅部门成员
	AlignObjective string `json:"align_objective"`             // 对齐目标
	ProjectId      int    `json:"project_id"`                  // 项目id
	StartAt        string `json:"start_at" binding:"required"` // 开始时间
	EndAt          string `json:"end_at" binding:"required"`   // 结束时间
}

// 创建OKR请求
type OkrCreateReq struct {
	OkrBaseReq
	Ascription int                      `json:"ascription" binding:"required"`  // 归属 1-部门 2-个人
	KeyResults []*OkrKeyResultCreateReq `json:"key_results" binding:"required"` // 关键结果
}

// 基础OKR关键结果请求
type OkrKeyResultBaseReq struct {
	Title       string `json:"title" binding:"required"`    // 关键结果 （必填）
	Participant string `json:"participant"`                 // 参与人,多个用逗号隔开
	Confidence  int    `json:"confidence"`                  // 信心指数 1-100
	StartAt     string `json:"start_at" binding:"required"` // 开始时间 （必填）
	EndAt       string `json:"end_at" binding:"required"`   // 结束时间 （必填）
}

// 创建OKR关键结果请求
type OkrKeyResultCreateReq struct {
	OkrKeyResultBaseReq
}

// 更新OKR请求
type OkrUpdateReq struct {
	Id int `json:"id" binding:"required"` // okr id
	OkrBaseReq
	KeyResults []*OkrKeyResultUpdateReq `json:"key_results" binding:"required"` // 关键结果
}

// 更新OKR关键结果请求
type OkrKeyResultUpdateReq struct {
	Id       int  `json:"id" binding:"required"`        // okr id
	IsDelete bool `json:"is_delete" binding:"required"` // 是否删除 true-删除 false-不删除
	OkrKeyResultBaseReq
}

// OKR列表/详情响应
type OkrResp struct {
	*model.Okr
	IsFollow      bool     `json:"is_follow"`       // 是否被关注
	KrCount       int      `json:"kr_count"`        // kr总数量
	KrFinishCount int      `json:"kr_finish_count"` // kr完成数量
	AlignCount    int      `json:"align_count"`     // 对齐目标数量
	Alias         []string `json:"alias"`           // 目标别名
}

// OKR对齐目标响应
type OkrAlignResp struct {
	*model.Okr
	Prefix         string   `json:"prefix"`          // 前缀
	AlignObjective string   `json:"align_objective"` // 对齐目标
	Alias          []string `json:"alias"`           // 目标别名
}

// OKR部门列表请求
type OkrDepartmentListReq struct {
	DepartmentId int    `form:"department_id" json:"department_id"` // 部门id
	Userid       int    `form:"userid" json:"userid"`               // 用户id
	Objective    string `form:"objective" json:"objective"`         // 目标（O）
	StartAt      string `form:"start_at" json:"start_at"`           // 开始时间
	EndAt        string `form:"end_at" json:"end_at"`               // 结束时间
	Type         int    `form:"type" json:"type"`                   // 类型 1-承诺型 2-挑战型
	Completed    int    `form:"completed" json:"completed"`         // 是否已完成未评分 0-未完成 1-已完成
	*Pages
}

// OKR更新进度和进度状态请求
type OkrUpdateProgressReq struct {
	Id       int `form:"id" binding:"required" json:"id"` // okr id
	Progress int `form:"progress" json:"progress"`        // 进度 1-100
	Status   int `form:"status" json:"status"`            // 进度状态 0-默认 1-正常 2-有风险 3-延期
}

// OKR评分请求
type OkrScoreReq struct {
	Id    int     `form:"id" binding:"required" json:"id"` // okr id
	Score float64 `form:"score" json:"score"`              // 个人评分 0-10
}

// 取消/重启目标
type OkrCanceledReq struct {
	Id int `form:"id" binding:"required" json:"id"` // okr id
}

// 更新参与人请求
type OkrParticipantUpdateReq struct {
	Id          int    `form:"id" binding:"required" json:"id"`                   // okr id
	Participant string `form:"participant" binding:"required" json:"participant"` // 参与人,多个用逗号隔开
}

// 更新信心指数请求
type OkrConfidenceUpdateReq struct {
	Id         int `form:"id" binding:"required" json:"id"`                 // okr id
	Confidence int `form:"confidence" binding:"required" json:"confidence"` // 信心指数 1-100
}

// 添加复盘请求
type OkrReplayCreateReq struct {
	OkrId    int                 `json:"okr_id" binding:"required"`   // okr id
	Comments []*OkrReplayComment `json:"comments" binding:"required"` // 复盘评价
	Review   string              `json:"review" binding:"required"`   // 回顾
}

// 复盘评价
type OkrReplayComment struct {
	OkrId   int    `json:"okr_id" binding:"required"`  // okr id
	Comment string `json:"comment" binding:"required"` // 评价
}

// 取消对齐目标请求
type OkrAlignCancelReq struct {
	OkrId      int `form:"okr_id" binding:"required" json:"okr_id"`             // okr id
	AlignOkrId int `form:"align_okr_id" binding:"required" json:"align_okr_id"` // 对齐目标okr id
}

// OKR id请求
type OkrIdReq struct {
	Id int `form:"id" binding:"required" json:"id"` // okr id
}

// OKR id列表请求
type OkrIdListReq struct {
	Id int `form:"id" binding:"required" json:"id"` // okr id
	*Pages
}

// 复盘 id请求
type OkrReplayIdReq struct {
	Id int `form:"id" binding:"required" json:"id"` // 复盘 id
}

// 复盘 id列表请求
type OkrReplayIdListReq struct {
	Id int `form:"id" binding:"required" json:"id"` // 复盘 id
	*Pages
}

// 列表基础分页请求
type OkrListBaseReq struct {
	Objective string `form:"objective" json:"objective"` // 目标（O）
	*Pages
}

// 对齐目标更新请求
type OkrAlignUpdateReq struct {
	Id             int    `form:"id" binding:"required" json:"id"`                           // okr id
	AlignObjective string `form:"align_objective" binding:"required" json:"align_objective"` // 对齐目标
}
