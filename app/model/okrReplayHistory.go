package model

import "time"

type OkrReplayHistory struct {
	Id         int       `gorm:"primary_key" json:"id"`
	ReplayId   int       `gorm:"default:0;comment:'复盘id'" json:"replay_id"`
	Title      string    `gorm:"type:varchar(255);comment:'标题'" json:"title"`
	ParentId   int       `gorm:"default:0;comment:'父级id'" json:"parent_id"`
	Progress   int       `gorm:"default:0;comment:'完成度'" json:"progress"`
	Confidence int       `gorm:"default:0;comment:'信心指数'" json:"confidence"`
	Score      float64   `gorm:"default:0;comment:'总评分'" json:"score"`
	Comment    string    `gorm:"type:varchar(255);comment:'复盘评价'" json:"comment"`
	CreateAt   time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt   time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}
