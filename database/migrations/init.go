package migrations

import (
	"dootask-okr/app/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// 添加okr AddTableOkr
var AddTableOkr = &gormigrate.Migration{
	ID: "2023071001-add-table-okr",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Okr{})
	},
}
