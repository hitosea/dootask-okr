package model

import (
	"time"
)

type OkrSetting struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255);comment:名称" json:"name"`
	Desc      string    `gorm:"type:varchar(255);comment:描述" json:"desc"`
	Setting   string    `gorm:"type:longtext;comment:设置" json:"setting"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}

var (
	SettingOkrKey = "okr" // okr设置key
)
