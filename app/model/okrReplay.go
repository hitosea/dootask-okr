package model

import (
	"time"
)

type OkrReplay struct {
	Id              int                 `gorm:"primary_key" json:"id"`
	Userid          int                 `gorm:"default:0;comment:'复盘用户id'" json:"userid"`
	OkrId           int                 `gorm:"default:0;comment:'目标id'" json:"okr_id"`
	OkrUserid       int                 `gorm:"default:0;comment:'目标所属用户id'" json:"okr_userid"`
	OkrDepartmentId string              `gorm:"type:varchar(100);default:'';comment:'目标部门id'" json:"okr_department_id"`
	OkrTitle        string              `gorm:"type:varchar(255);comment:'目标标题'" json:"okr_title"`
	OkrProgress     int                 `gorm:"default:0;comment:'目标完成度'" json:"okr_progress"`
	OkrPriority     string              `gorm:"type:varchar(10);comment:'优先级'" json:"okr_priority"`
	Review          string              `gorm:"type:varchar(255);comment:'回顾'" json:"review"`
	CreatedAt       time.Time           `gorm:"autoCreateTime;comment:'创建时间'" json:"created_at"`
	UpdatedAt       time.Time           `gorm:"autoUpdateTime;comment:'更新时间'" json:"updated_at"`
	KrHistory       []*OkrReplayHistory `gorm:"foreignkey:ReplayId" json:"kr_history"`
	OkrAlias        []string            `gorm:"-" json:"okr_alias"`
}

var OkrReplayModel = OkrReplay{}
