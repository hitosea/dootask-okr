package model

import (
	"time"
)

type Okr struct {
	Id             int       `gorm:"primary_key" json:"id"`
	ParentId       int       `gorm:"default:0;comment:'父级目标id'" json:"parent_id"`
	Userid         int       `gorm:"default:0;comment:'用户id'" json:"userid"`
	DepartmentId   int       `gorm:"default:0;comment:'部门id'" json:"department_id"`
	ProjectId      int       `gorm:"default:0;comment:'项目id'" json:"project_id"`
	Title          string    `gorm:"type:varchar(255);comment:'标题内容'" json:"title"`
	Type           int       `gorm:"default:1;comment:'类型 1-承诺型 2-挑战型'" json:"type"`
	Priority       string    `gorm:"type:varchar(10);comment:'优先级'" json:"priority"`
	Ascription     int       `gorm:"default:1;comment:'归属 1-部门 2-个人'" json:"ascription"`
	VisibleRange   int       `gorm:"default:1;comment:'可见范围  1-全公司 2-仅相关成员 3-仅部门成员'" json:"visible_range"`
	Completed      int       `gorm:"default:0;comment:'整个O是否完成 0-未完成 1-已完成'" json:"completed"`
	Participant    string    `gorm:"type:varchar(255);comment:'参与人'" json:"participant"`
	Progress       int       `gorm:"default:0;comment:'进度指数0-100'" json:"progress"`
	ProgressStatus int       `gorm:"default:0;comment:'进度状态 0-未开始 1-已完成 2-进行中 3-有难度 4-已结束'" json:"progress_status"`
	Confidence     int       `gorm:"default:0;comment:'信心指数0-100'" json:"confidence"`
	Score          float64   `gorm:"default:0;comment:'个人评分'" json:"score"`
	SuperiorScore  float64   `gorm:"default:0;comment:'上级评分'" json:"superior_score"`
	StartAt        time.Time `gorm:"comment:'开始时间' " json:"start_at"`
	EndAt          time.Time `gorm:"comment:'结束时间'" json:"end_at"`
	CreateAt       time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt       time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
	KeyResults     []*Okr    `gorm:"-" json:"key_results"`
}

var (
	OkrModel                     = Okr{}
	OkrKeyResultStatusNotStart   = 0 // 未开始
	OkrKeyResultStatusFinished   = 1 // 已完成
	OkrKeyResultStatusInProgress = 2 // 进行中
	OkrKeyResultStatusHasProblem = 3 // 有难度
	OkrKeyResultStatusEnd        = 4 // 已结束
)
