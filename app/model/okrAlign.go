package model

import (
	"time"
)

type OkrAlign struct {
	Id         int       `gorm:"primary_key" json:"id"`
	OkrId      int       `gorm:"default:0;comment:'当前okrId'" json:"okr_id"`
	AlignOkrId int       `gorm:"default:0;comment:'对齐目标okrId'" json:"align_okr_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"updated_at"`
}

var OkrAlignModel = OkrAlign{}
