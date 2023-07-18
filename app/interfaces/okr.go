package interfaces

import "dootask-okr/app/model"

// 基础OKR请求
type OkrBaseReq struct {
	Title          string `json:"title" binding:"required"`           // 目标
	Type           int    `json:"type" binding:"required"`            // 类型 1-承诺型 2-挑战型
	Priority       string `json:"priority" binding:"required"`        // 优先级
	VisibleRange   int    `json:"visible_range" binding:"required"`   // 可见范围  1-全公司 2-仅相关成员 3-仅部门成员
	AlignObjective string `json:"align_objective" binding:"required"` // 对齐目标
	ProjectId      int    `json:"project_id" binding:"required"`      // 项目id
	StartAt        string `json:"start_at" binding:"required"`        // 开始时间
	EndAt          string `json:"end_at" binding:"required"`          // 结束时间
}

// 创建OKR请求
type OkrCreateReq struct {
	OkrBaseReq
	Ascription int                      `json:"ascription" binding:"required"`  // 归属 1-部门 2-个人
	KeyResults []*OkrKeyResultCreateReq `json:"key_results" binding:"required"` // 关键结果
}

// 基础OKR关键结果请求
type OkrKeyResultBaseReq struct {
	Participant string `json:"participant" binding:"required"` // 参与人,多个用逗号隔开
	Title       string `json:"title" binding:"required"`       // 关键结果
	Confidence  int    `json:"confidence" binding:"required"`  // 信心指数
	StartAt     string `json:"start_at" binding:"required"`    // 开始时间
	EndAt       string `json:"end_at" binding:"required"`      // 结束时间
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
	Id       int  `json:"id" binding:"required"` // okr id
	IsDelete bool `json:"is_delete"`             // 是否删除 true-删除 false-不删除
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
	DepartmentId int    `json:"department_id"` // 部门id
	Userid       int    `json:"userid"`        // 用户id
	Objective    string `json:"objective"`     // 目标
	StartAt      string `json:"start_at"`      // 开始时间
	EndAt        string `json:"end_at"`        // 结束时间
	Type         int    `json:"type"`          // 类型 1-承诺型 2-挑战型
	Completed    int    `json:"completed"`     // 是否已完成未评分 0-未完成 1-已完成
}

// OKR更新进度和进度状态请求
type OkrUpdateProgressReq struct {
	Id       int `form:"id" binding:"required"` // okr id
	Progress int `form:"progress"`              // 进度
	Status   int `form:"status"`                // 进度状态 0-默认 1-正常 2-有风险 3-延期 4-已结束
}

// OKR评分请求
type OkrScoreReq struct {
	Id    int     `form:"id" binding:"required"` // okr id
	Score float64 `form:"score"`                 // 个人评分
}

// 取消/重启目标
type OkrCanceledReq struct {
	Id       int `form:"id" binding:"required"` // okr id
	Canceled int `form:"canceled"`              // 状态 0-重启 1-结束
}

// 更新参与人请求
type OkrParticipantUpdateReq struct {
	Id          int    `form:"id" binding:"required"`          // okr id
	Participant string `form:"participant" binding:"required"` // 参与人,多个用逗号隔开
}

// 更新信心指数请求
type OkrConfidenceUpdateReq struct {
	Id         int `form:"id" binding:"required"`         // okr id
	Confidence int `form:"confidence" binding:"required"` // 信心指数
}

// 添加复盘请求
type OkrReplayCreateReq struct {
	OkrId    int                 `json:"okr_id" binding:"required"` // okr id
	Comments []*OkrReplayComment `json:"comments"`                  // 复盘评价
	Review   string              `json:"review"`                    // 回顾
}

// 复盘评价
type OkrReplayComment struct {
	OkrId   int    `json:"okr_id" binding:"required"` // okr id
	Comment string `json:"comment"`                   // 评价
}
