package model

import (
	"time"
)

type OkrLog struct {
	Id           int       `gorm:"primary_key" json:"id"`
	OkrId        int       `gorm:"default:0;comment:'Okr id'" json:"okr_id"`
	Userid       int       `gorm:"default:0;comment:'用户id'" json:"userid"`
	Operation    string    `gorm:"type:varchar(255);comment:'操作'" json:"operation"`
	Content      string    `gorm:"type:varchar(255);comment:'内容'" json:"content"`
	CreatedAt    time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"created_at"`
	UserAvatar   string    `gorm:"-" json:"user_avatar"`
	UserNickname string    `gorm:"-" json:"user_nickname"`
}

var OkrLogModel = OkrLog{}
