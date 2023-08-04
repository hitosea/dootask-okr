package model

import (
	"time"
)

type OkrFollow struct {
	Id        int       `gorm:"primary_key" json:"id"`
	OkrId     int       `gorm:"default:0;comment:目标Id" json:"okr_id"`
	Userid    int       `gorm:"default:0;comment:用户Id" json:"userid"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}

var OkrFollowModel = OkrFollow{}
