package model

import (
	"time"

	"gorm.io/gorm"
)

type OkrPushLog struct {
	Id        int            `gorm:"primary_key" json:"id"`
	Userid    int            `gorm:"default:0;comment:用户Id" json:"userid"`
	OkrId     int            `gorm:"default:0;comment:okrId" json:"okr_id"`
	Type      int            `gorm:"default:0;comment:提醒类型：0-okr开始提醒，1-距离到期提醒，2-到期超时提醒" json:"type"`
	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at" swaggerignore:"true"`
}
