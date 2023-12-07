package interfaces

// okr设置
type OkrSettingReq struct {
	Type                string `json:"type"`                  //类型 get: 获取 (默认)，save: 保存，all: 获取所有
	ScoreDepartmentUser int    `json:"score_department_user"` // 单选：指定可评分部门及部分负责人OKR的人员
	SelfScoreWeight     int    `json:"self_score_weight"`     // 自评权重
	SuperiorScoreWeight int    `json:"superior_score_weight"` // 上级评分权重
}
