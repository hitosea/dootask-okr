package model

import (
	"dootask-okr/app/core"
	"time"
)

type OkrObjective struct {
	core.BaseIdModels
	Userid         int             `gorm:"default:0;comment:'用户id'" json:"userid"`
	DepartmentId   int             `gorm:"default:0;comment:'部门id'" json:"department_id"`
	ProjectId      int             `gorm:"default:0;comment:'项目id'" json:"project_id"`
	Objective      string          `gorm:"type:varchar(255);comment:'目标'" json:"objective"`
	Type           int             `gorm:"default:1;comment:'类型 1-承诺型 2-挑战型'" json:"type"`
	Priority       string          `gorm:"type:varchar(10);comment:'优先级'" json:"priority"`
	Ascription     int             `gorm:"default:1;comment:'归属 1-部门 2-个人'" json:"ascription"`
	VisibleRange   int             `gorm:"default:1;comment:'可见范围  1-全公司 2-仅相关成员 3-仅部门成员'" json:"visible_range"`
	AlignObjective string          `gorm:"type:varchar(255);comment:'对齐目标'" json:"align_objective"`
	Completed      int             `gorm:"default:0;comment:'整个O是否完成 0-未完成 1-已完成'" json:"completed"`
	StartAt        time.Time       `gorm:"comment:'开始时间' " json:"start_at"`
	EndAt          time.Time       `gorm:"comment:'结束时间'" json:"end_at"`
	CreateAt       time.Time       `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt       time.Time       `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
	KeyResults     []*OkrKeyResult `gorm:"foreignKey:ObjectiveId;references:Id" json:"key_results,omitempty"`
}

var OkrObjectiveModel = OkrObjective{}
