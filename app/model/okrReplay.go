package model

import (
	"time"
)

type OkrReplay struct {
	Id       int       `gorm:"primary_key" json:"id"`
	OkrId    int       `gorm:"default:0;comment:'目标id'" json:"okr_id"`
	Userid   int       `gorm:"default:0;comment:'用户id'" json:"userid"`
	Value    string    `gorm:"type:varchar(255);comment:'价值与收获'" json:"value"`
	Problem  string    `gorm:"type:varchar(255);comment:'问题与不足'" json:"problem"`
	CreateAt time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}

var OkrReplayModel = OkrReplay{}
