package model

import (
	"dootask-okr/app/core"
	"time"
)

type UserDepartment struct {
	core.BaseIdModels
	Name             string          `json:"name"`
	DialogId         int             `json:"dialog_id"`
	ParentId         int             `json:"parent_id"`
	OwnerUserid      int             `json:"owner_userid"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	ParentDepartment *UserDepartment `gorm:"ForeignKey:ParentId" json:"parent_department,omitempty"`
	OwnerUser        *User           `gorm:"ForeignKey:OwnerUserid;References:Userid" json:"owner_user,omitempty"`
}

var UserDepartmentModel = UserDepartment{}
