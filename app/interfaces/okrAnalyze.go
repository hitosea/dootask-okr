package interfaces

type OkrAnalyzeDepartment struct {
	Id    int64  `json:"id"`    // 部门id
	Name  string `json:"name"`  // 部门名称
	Owner bool   `json:"owner"` // 是否部门负责人
}

// OKR整体平均完成度
type OkrAnalyzeOverall struct {
	Total     int64 `json:"total"`    // okr 总数
	Completed int64 `json:"complete"` // okr 完成数
}

// OKR各部门平均完成度
type OkrAnalyzeDept struct {
	DepartmentId   int    `json:"department_id"`   //部门id
	DepartmentName string `json:"department_name"` //部门名称
	*OkrAnalyzeOverall
}

// OKR评分分布
type OkrAnalyzeScore struct {
	Total        int64 `json:"total"`          // okr 总数
	Unscored     int64 `json:"unscored"`       //未评分
	ZeroToThree  int64 `json:"zero_to_three"`  //0-3分
	ThreeToSeven int64 `json:"three_to_seven"` //3-7分
	SevenToTen   int64 `json:"seven_to_ten"`   //7-10分
}

// OKR各部门平均完成度
type OkrAnalyzeScoreDept struct {
	DepartmentId   int    `json:"department_id"`   //部门id
	DepartmentName string `json:"department_name"` //部门名称
	*OkrAnalyzeScore
}

// OKR人员评分率
type OkrAnalyzePersonnelScoreRate struct {
	Total     int64 `json:"total"`    // okr 总数
	Completed int64 `json:"complete"` // okr 完成数
}

// OKR部门评分占比
type OkrAnalyzeDeptScoreProportion struct {
	DepartmentId    int    `json:"department_id"`    //部门id
	DepartmentName  string `json:"department_name"`  //部门名称
	Total           int64  `json:"total"`            //okr 总数
	Unscored        int64  `json:"unscored"`         //未评分
	AlreadyReviewed int64  `json:"already_reviewed"` //已评分
}
