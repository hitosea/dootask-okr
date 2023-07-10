package model

import (
	"dootask-okr/app/core"
	"time"
)

type OkrFollow struct {
	core.BaseIdModels
	ObjectiveId int       `gorm:"default:0;comment:'目标id'" json:"objective_id"`
	Userid      int       `gorm:"default:0;comment:'用户id'" json:"userid"`
	CreateAt    time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt    time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}

var OkrFollowModel = OkrFollow{}
