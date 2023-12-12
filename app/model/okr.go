package model

import (
	"dootask-okr/app/core"
	"dootask-okr/app/utils/common"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Okr struct {
	Id             int            `gorm:"primary_key" json:"id"`
	ParentId       int            `gorm:"default:0;comment:父级目标Id" json:"parent_id"`
	Userid         int            `gorm:"default:0;comment:用户Id" json:"userid"`
	DepartmentId   string         `gorm:"type:varchar(100);default:'';comment:部门Id" json:"department_id"`
	ProjectId      int            `gorm:"default:0;comment:项目Id" json:"project_id"`
	DialogId       int            `gorm:"default:0;comment:聊天Id" json:"dialog_id"`
	Title          string         `gorm:"type:varchar(255);comment:标题内容" json:"title"`
	Type           int            `gorm:"default:1;comment:类型 1-承诺型 2-挑战型" json:"type"`
	Priority       string         `gorm:"type:varchar(10);comment:优先级" json:"priority"`
	Ascription     int            `gorm:"default:1;comment:归属 1-部门 2-个人" json:"ascription"`
	VisibleRange   int            `gorm:"default:1;comment:可见范围 1-全公司 2-仅相关成员 3-仅部门成员" json:"visible_range"`
	Completed      int            `gorm:"default:0;comment:整个O是否完成 0-未完成 1-已完成" json:"completed"`
	Canceled       int            `gorm:"default:0;comment:整个O是否取消 0-未完成 1-已取消" json:"canceled"`
	Participant    string         `gorm:"type:varchar(255);comment:参与人" json:"participant"`
	Progress       int            `gorm:"default:0;comment:进度指数0-100" json:"progress"`
	ProgressStatus int            `gorm:"default:0;comment:进度状态 0-未开始 1-正常 2-有风险 3-已延期" json:"progress_status"`
	Confidence     int            `gorm:"default:0;comment:信心指数0-100" json:"confidence"`
	Score          float64        `gorm:"default:-1;comment:个人评分" json:"score"`          // 个人评分和O总评分
	SuperiorScore  float64        `gorm:"default:-1;comment:上级评分" json:"superior_score"` // 上级评分
	ScoreNum       int            `gorm:"default:0;comment:评分次数" json:"score_num"`
	Status         int            `gorm:"default:0;comment:状态 0-正常 1-已归档 2-人员离职/删除" json:"status"`
	ArchiveUserid  int            `gorm:"default:0;comment:归档人员Id" json:"archive_userid"`
	ArchiveAt      *time.Time     `gorm:"comment:归档时间" json:"archive_at"`
	StartAt        time.Time      `gorm:"comment:开始时间" json:"start_at"`
	EndAt          time.Time      `gorm:"comment:结束时间" json:"end_at"`
	CreatedAt      time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at" swaggerignore:"true"`
	KeyResults     []*Okr         `gorm:"ForeignKey:ParentId" json:"key_results,omitempty"`
	ParentOKr      *Okr           `gorm:"ForeignKey:ParentId" json:"parent_okr,omitempty"`
	User           *User          `gorm:"ForeignKey:Userid;References:Userid" json:"user,omitempty"`
	KrScore        float64        `gorm:"-" json:"kr_score"`               // KR总评分
	ParentTitle    string         `gorm:"-" json:"parent_title,omitempty"` // 父级目标标题
	CanUpdateScore bool           `gorm:"-" json:"can_update_score"`       // KR是否能评分
	ObjectiveNum   string         `gorm:"-" json:"objective_num"`          // O的数字编号
}

var (
	OkrModel                     = Okr{}
	OkrKeyResultStatusNotStart   = 0 // 未开始
	OkrKeyResultStatusFinished   = 1 // 正常
	OkrKeyResultStatusInProgress = 2 // 有风险
	OkrKeyResultStatusHasProblem = 3 // 已延期
	ProgressStatusMap            = map[int]string{
		0: "未开始",
		1: "正常",
		2: "有风险",
		3: "已延期",
	}
	CanceledMap = map[int]string{
		0: "正常",
		1: "已取消",
	}
	SelfScoreWeight     = 30 // 个人评分权重30
	SuperiorScoreWeight = 70 // 上级评分权重70
	DefaultScoreNum     = 3  // 默认评分次数3次
)

// 获取所有部门ids
func (m Okr) GetAllDeptIds() []int {
	var departmentId []int
	var departmentIds []string
	if err := core.DB.Model(m).Where("deleted_at IS NULL").Pluck("DISTINCT(department_id)", &departmentIds).Error; err != nil {
		return departmentId
	}
	for _, v := range departmentIds {
		vs := strings.Split(v, ",")
		for _, vv := range vs {
			num, err := strconv.Atoi(vv)
			if err != nil {
				return departmentId
			}
			departmentId = append(departmentId, num)
		}
		fmt.Println(strings.Split(v, ","))
	}
	return common.ArrayUniqueInt(departmentId)
}
