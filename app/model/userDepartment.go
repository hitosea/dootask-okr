package model

import (
	"dootask-okr/app/core"
	"time"
)

type UserDepartment struct {
	core.BaseIdModels
	Name        string    `gorm:"type:varchar(100);comment:'部门名称'" json:"name"`
	DialogId    int       `gorm:"default:0;comment:'聊天会话ID'" json:"dialog_id"`
	ParentId    int       `gorm:"default:0;comment:'上级部门'" json:"parent_id"`
	OwnerUserid int       `gorm:"default:0;comment:'部门负责人'" json:"owner_userid"`
	CreateAt    time.Time `gorm:"autoCreateTime;comment:'创建时间'" json:"create_at"`
	UpdateAt    time.Time `gorm:"autoUpdateTime;comment:'更新时间'" json:"update_at"`
}

var UserDepartmentModel = UserDepartment{}
