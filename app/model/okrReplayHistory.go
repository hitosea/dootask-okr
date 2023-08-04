package model

import "time"

type OkrReplayHistory struct {
	Id         int       `gorm:"primary_key" json:"id"`
	ReplayId   int       `gorm:"default:0;comment:复盘Id" json:"replay_id"`
	OkrId      int       `gorm:"default:0;comment:okrId" json:"okr_id"`
	Title      string    `gorm:"type:varchar(255);comment:标题" json:"title"`
	ParentId   int       `gorm:"default:0;comment:父级Id" json:"parent_id"`
	Progress   int       `gorm:"default:0;comment:完成度" json:"progress"`
	Confidence int       `gorm:"default:0;comment:信心指数" json:"confidence"`
	Score      float64   `gorm:"default:0;comment:总评分" json:"score"`
	Comment    int       `gorm:"default:0;comment:复盘评价 1-做得好的 2-可提升的" json:"comment"`
	CreatedAt  time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}
