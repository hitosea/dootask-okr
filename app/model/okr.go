package model

import (
	"dootask-okr/app/core"
)

type Okr struct {
	core.BaseIdModels
	Userid int `gorm:"default:0;comment:'成员ID'" json:"userid"`
	core.BaseAtModels
}

var OkrModel = Okr{}
