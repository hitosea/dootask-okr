package interfaces

type StatisticsAllS struct {
	Uncompleted     int64   `json:"uncompleted"`       	//个人当前未完成目标数量，减去完成减去已取消
	Completed 		int64	`json:"completed"` 			//个人当前所有完成目标数量加取消目标数量
}

type OkrOverall struct {
	CompletionSum int   `json:"completion_sum"` //个人当前所有进度度
	ScoreSum      float32   `json:"score_sum"`      //个人当前所有评分
	Total         int64 `json:"total"`          //个人当前所有目标数量,除去已取消
}