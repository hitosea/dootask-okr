package core

import (
	"dootask-okr/app/utils/common"

	"gorm.io/gorm"
)

type BaseIdModels struct {
	Id int `gorm:"primary_key" json:"id"`
}

type BaseAtModels struct {
	CreatedAt  int64  `gorm:"autoCreateTime;comment:'创建时间'" json:"-"`
	UpdatedAt  int64  `gorm:"autoUpdateTime;comment:'更新时间'" json:"-"`
	CreatedAts string `gorm:"-" json:"created_at,omitempty"`
	UpdatedAts string `gorm:"-" json:"updated_at,omitempty"`
}

// 自动转换时间格式
func (m *BaseAtModels) AfterFind(tx *gorm.DB) (err error) {
	if m.CreatedAt > 0 {
		m.CreatedAts = common.UnixTimeToStr(m.CreatedAt)
	}
	if m.UpdatedAt > 0 {
		m.UpdatedAts = common.UnixTimeToStr(m.UpdatedAt)
	}
	return
}
