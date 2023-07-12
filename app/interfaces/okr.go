package interfaces

import "time"

// OkrCreateReq 创建OKR请求
type OkrCreateReq struct {
	Title          string                  `json:"Title" binding:"required"`           // 目标
	Type           int                     `json:"type" binding:"required"`            // 类型 1-承诺型 2-挑战型
	Priority       string                  `json:"priority" binding:"required"`        // 优先级
	Ascription     int                     `json:"ascription" binding:"required"`      // 归属 1-部门 2-个人
	VisibleRange   int                     `json:"visible_range" binding:"required"`   // 可见范围  1-全公司 2-仅相关成员 3-仅部门成员
	AlignObjective string                  `json:"align_objective" binding:"required"` // 对齐目标
	StartAt        time.Time               `json:"start_at" binding:"required"`        // 开始时间
	EndAt          time.Time               `json:"end_at" binding:"required"`          // 结束时间
	KeyResults     []OkrKeyResultCreateReq `json:"key_results" binding:"required"`     // 关键结果
}

// OkrKeyResultCreateReq 创建OKR关键结果请求
type OkrKeyResultCreateReq struct {
	Participant string    `json:"participant" binding:"required"` // 参与人,多个用逗号隔开
	Title       string    `json:"title" binding:"required"`       // 关键结果
	Confidence  int       `json:"confidence" binding:"required"`  // 信心指数
	StartAt     time.Time `json:"start_at" binding:"required"`    // 开始时间
	EndAt       time.Time `json:"end_at" binding:"required"`      // 结束时间
}

// 更新OKR请求
type OkrUpdateReq struct {
	Id             int                     `json:"id" binding:"required"`              // id
	Title          string                  `json:"objective" binding:"required"`       // 目标
	Type           int                     `json:"type" binding:"required"`            // 类型 1-承诺型 2-挑战型
	Priority       string                  `json:"priority" binding:"required"`        // 优先级
	VisibleRange   int                     `json:"visible_range" binding:"required"`   // 可见范围  1-全公司 2-仅相关成员 3-仅部门成员
	AlignObjective string                  `json:"align_objective" binding:"required"` // 对齐目标
	StartAt        time.Time               `json:"start_at" binding:"required"`        // 开始时间
	EndAt          time.Time               `json:"end_at" binding:"required"`          // 结束时间
	KeyResults     []OkrKeyResultUpdateReq `json:"key_results" binding:"required"`     // 关键结果
}

// OkrKeyResultUpdateReq 更新OKR关键结果请求
type OkrKeyResultUpdateReq struct {
	Id          int       `json:"id" binding:"required"`          // id
	Participant string    `json:"participant" binding:"required"` // 参与人,多个用逗号隔开
	Title       string    `json:"key_result" binding:"required"`  // 关键结果
	Confidence  int       `json:"confidence" binding:"required"`  // 信心指数
	StartAt     time.Time `json:"start_at" binding:"required"`    // 开始时间
	EndAt       time.Time `json:"end_at" binding:"required"`      // 结束时间
}
