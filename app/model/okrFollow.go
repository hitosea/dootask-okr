package model

import (
	"time"
)

type OkrFollow struct {
	Id       int       `gorm:"primary_key" json:"id"`
	OkrId    int       `gorm:"default:0;comment:'目标id'" json:"okr_id"`
	Userid   int       `gorm:"default:0;comment:'用户id'" json:"userid"`
	CreateAt time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}

var OkrFollowModel = OkrFollow{}
