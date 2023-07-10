package interfaces

// OkrAddReq 创建OKR请求
type OkrAddReq struct {
	Objective      string               `json:"objective" binding:"required"`     // 目标
	Type           int                  `json:"type" binding:"required"`          // 类型 1-承诺型 2-挑战型
	Priority       string               `json:"priority" binding:"required"`      // 优先级
	Ascription     int                  `json:"ascription" binding:"required"`    // 归属 1-部门 2-个人
	VisibleRange   int                  `json:"visible_range" binding:"required"` // 可见范围  1-全公司 2-仅相关成员 3-仅部门成员
	AlignObjective string               `json:"align_objective"`                  // 对齐目标,多个用逗号隔开
	StartAt        string               `json:"start_at" binding:"required"`      // 开始时间
	EndAt          string               `json:"end_at" binding:"required"`        // 结束时间
	KeyResults     []OkrKeyResultAddReq `json:"key_results" binding:"required"`   // 关键结果
}

// OkrKeyResultAddReq 创建OKR关键结果请求
type OkrKeyResultAddReq struct {
	Participant string `json:"participant" binding:"required"` // 参与人,多个用逗号隔开
	KeyResult   string `json:"key_result" binding:"required"`  // 关键结果
	Confidence  int    `json:"confidence" binding:"required"`  // 信心指数
	StartAt     string `json:"start_at" binding:"required"`    // 开始时间
	EndAt       string `json:"end_at" binding:"required"`      // 结束时间
}
