package model

import (
	"dootask-okr/app/core"
	"time"
)

type OkrKeyResult struct {
	core.BaseIdModels
	Participant    string    `gorm:"type:varchar(255);comment:'参与人'" json:"participant"`
	ObjectiveId    int       `gorm:"default:0;comment:'目标id'" json:"objective_id"`
	KeyResult      string    `gorm:"type:varchar(255);comment:'关键结果'" json:"key_result"`
	Progress       int       `gorm:"default:0;comment:'进度指数'" json:"progress"`
	ProgressStatus int       `gorm:"default:0;comment:'进度状态 0-未开始 1-已完成 2-进行中 3-有难度 4-已结束'" json:"progress_status"`
	Confidence     int       `gorm:"default:0;comment:'信心指数'" json:"confidence"`
	Score          float64   `gorm:"default:0;comment:'个人评分'" json:"score"`
	SuperiorScore  float64   `gorm:"default:0;comment:'上级评分'" json:"superior_score"`
	StartAt        time.Time `gorm:"comment:'开始时间' " json:"start_at"`
	EndAt          time.Time `gorm:"comment:'结束时间'" json:"end_at"`
	CreateAt       time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt       time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}

var (
	OkrKeyResultModel            = OkrKeyResult{}
	OkrKeyResultStatusNotStart   = 0 // 未开始
	OkrKeyResultStatusFinished   = 1 // 已完成
	OkrKeyResultStatusInProgress = 2 // 进行中
	OkrKeyResultStatusHasProblem = 3 // 有难度
	OkrKeyResultStatusEnd        = 4 // 已结束
)
