package model

import (
	"dootask-okr/app/core"
	"time"
)

type UserDepartment struct {
	core.BaseIdModels
	Name        string    `json:"name"`
	DialogId    int       `json:"dialog_id"`
	ParentId    int       `json:"parent_id"`
	OwnerUserid int       `json:"owner_userid"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

var UserDepartmentModel = UserDepartment{}
