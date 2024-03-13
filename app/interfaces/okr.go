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
	AutoSync       int    `json:"auto_sync"`                   // 是否自动同步
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
	Confidence  int    `json:"confidence"`                  // 信心指数 0-100
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
	IsFollow       bool     `json:"is_follow"`       // 是否被关注
	KrCount        int      `json:"kr_count"`        // kr总数量
	KrFinishCount  int      `json:"kr_finish_count"` // kr完成数量
	AlignObjective []int    `json:"align_objective"` // 对齐目标ids
	AlignCount     int      `json:"align_count"`     // 对齐目标数量
	Alias          []string `json:"alias"`           // 目标别名
	UserAvatar     string   `json:"user_avatar"`     // 用户头像
	SuperiorUser   []int    `json:"superior_user"`   // 上级用户
}

// OKR对齐目标响应
type OkrAlignResp struct {
	*model.Okr
	Prefix         string   `json:"prefix"`          // 前缀
	AlignObjective string   `json:"align_objective"` // 对齐目标
	Alias          []string `json:"alias"`           // 目标别名
}

// OKR被对齐目标响应
type OkrByAlignResp struct {
	*model.Okr
	Prefix         string   `json:"prefix"`          // 前缀
	AlignObjective string   `json:"align_objective"` // 对齐目标
	Alias          []string `json:"alias"`           // 目标别名
}

// OKR部门列表请求
type OkrDepartmentListReq struct {
	Userid    int    `form:"userid" json:"userid"`       // 用户id
	Objective string `form:"objective" json:"objective"` // 目标（O）
	StartAt   string `form:"start_at" json:"start_at"`   // 开始时间
	EndAt     string `form:"end_at" json:"end_at"`       // 结束时间
	Type      int    `form:"type" json:"type"`           // 类型 1-承诺型 2-挑战型
	Completed int    `form:"completed" json:"completed"` // 是否已完成未评分 0-未完成 1-已完成
	*Pages
}

// OKR全公司列表请求
type OkrCompanyListReq struct {
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
	Progress int `form:"progress" json:"progress"`        // 进度 0-100
	Status   int `form:"status" json:"status"`            // 进度状态 0-默认 1-正常 2-有风险 3-延期
}

// OKR评分请求
type OkrScoreReq struct {
	Id    int     `form:"id" binding:"required" json:"id"` // okr id
	Score float64 `form:"score" json:"score"`              // 个人评分 0-10
	Type  int     `form:"type" json:"type"`                // 评分类型 1-个人评分 2-上级评分
}

// 取消/重启目标
type OkrCanceledReq struct {
	Id int `form:"id" binding:"required" json:"id"` // okr id
}

// 更新参与人请求
type OkrParticipantUpdateReq struct {
	Id          int    `form:"id" binding:"required" json:"id"` // okr id
	Participant string `form:"participant" json:"participant"`  // 参与人,多个用逗号隔开
}

// 更新信心指数请求
type OkrConfidenceUpdateReq struct {
	Id         int `form:"id" binding:"required" json:"id"` // okr id
	Confidence int `form:"confidence" json:"confidence"`    // 信心指数 0-100
}

// 添加复盘请求
type OkrReplayCreateReq struct {
	OkrId    int                 `json:"okr_id" binding:"required"`   // okr id
	Comments []*OkrReplayComment `json:"comments" binding:"required"` // 复盘评价
	Review   string              `json:"review"`                      // 价值与收获
	Problem  string              `json:"problem"`                     // 问题与不足
}

// 复盘上级评价
type OkrReplaySuperiorReviewReq struct {
	Id             int    `json:"id" binding:"required"` // 复盘 id
	SuperiorReview string `json:"superior_review"`       // 上级评价
}

// 复盘评价
type OkrReplayComment struct {
	OkrId   int `json:"okr_id" binding:"required"`  // okr id
	Comment int `json:"comment" binding:"required"` // 评价 1-做得好的 2-可提升的
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
	Objective  string `form:"objective" json:"objective"`     // 目标（O）
	ExcludeIds string `form:"exclude_ids" json:"exclude_ids"` // 排除的id
	*Pages
}

// 对齐目标列表请求
type OkrAlignListReq struct {
	OkrId      int    `form:"okr_id" json:"okr_id"`                   // id
	Ascription int    `form:"ascription,default=2" json:"ascription"` // 归属 1-部门 2-个人
	Objective  string `form:"objective" json:"objective"`             // 目标（O）
	*Pages
}

// 对齐目标更新请求
type OkrAlignUpdateReq struct {
	Id             int    `form:"id" binding:"required" json:"id"`        // okr id
	AlignObjective string `form:"align_objective" json:"align_objective"` // 对齐目标
}

// 获取用户列表请求
type OkrUserListReq struct {
	DeptOnly bool   `form:"dept_only,default=false" json:"dept_only"` // 是否只获取部门用户
	Keyword  string `form:"keyword" json:"keyword"`                   // 关键词
	*Pages
}

// OKR日志参数
type OkrLogParams struct {
	Title                string   `json:"title,omitempty"`                  // 标题
	ParentTitle          string   `json:"parent_title,omitempty"`           // 父级标题
	TitleChange          []string `json:"title_change,omitempty"`           // 标题变更
	ProgressChange       []int    `json:"progress_change,omitempty"`        // 进度变更
	ProgressStatusChange []string `json:"progress_status_change,omitempty"` // 进度状态变更
	ConfidenceChange     []int    `json:"confidence_change,omitempty"`      // 信心指数变更
	TimeChange           []string `json:"time_change,omitempty"`            // 时间变更
	StatusChange         []string `json:"status_change,omitempty"`          // 状态变更
	UserChange           []string `json:"user_change,omitempty"`            // 人员变更
}

// 离职/删除人员列表分页请求
type OkrLeaveUpdateListReq struct {
	Objective string `form:"objective" json:"objective"` // 目标（O）
	Userid    int    `form:"userid" json:"userid"`       // 用户id
	*Pages
}

// OKR负责人列表
type OkrOwnerListReq struct {
	KeyWord string `form:"keyword" json:"keyword"`         // 关键词
	Status  int    `form:"status,default=0" json:"status"` // 状态 0-正常 1-已归档 2-离职/删除人员
	*Pages
}

// 更新离职/删除人员OKR负责人请求
type OkrLeaveUpdateReq struct {
	OkrIds []int `form:"okr_ids" json:"okr_ids" binding:"required"` // okr ids
	Userid int   `form:"userid" json:"userid" binding:"required"`   // 新负责人id
}

// OKR复盘列表请求
type OkrReplayListReq struct {
	OkrId          int    `form:"okr_id" json:"okr_id"`                   // 目标id(兼容手机端详情)
	DepartmentId   int    `form:"department_id" json:"department_id"`     // 部门id
	Userid         int    `form:"userid" json:"userid"`                   // 用户id
	Objective      string `form:"objective" json:"objective"`             // 目标（O）
	StartAt        string `form:"start_at" json:"start_at"`               // 开始时间
	EndAt          string `form:"end_at" json:"end_at"`                   // 结束时间
	SelectReplayed bool   `form:"select_replayed" json:"select_replayed"` // 选中已评分未复盘 true-选中 false-未选中
	*Pages
}

// OKR复盘列表请求
type OkrReplayResp struct {
	*model.OkrReplay
	ReplayIds    string `json:"replay_ids"`    // 复盘集合
	SuperiorUser []int  `json:"superior_user"` // 上级用户
}
