package model

import (
	"dootask-okr/app/core"
	"time"
)

type OkrAlign struct {
	core.BaseIdModels
	OkrId      int       `gorm:"default:0;comment:'当前okrId'" json:"okr_id"`
	AlignOkrId int       `gorm:"default:0;comment:'对齐目标okrId'" json:"align_okr_id"`
	CreateAt   time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt   time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}

var OkrAlignModel = OkrAlign{}
