package migrations

import (
	"dootask-okr/app/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// 添加目标表 AddTableOkrObject
var AddTableOkrObject = &gormigrate.Migration{
	ID: "2023071001-add-table-okr",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Okr{})
	},
}

// 添加关注表 AddTableOkrFollow
var AddTableOkrFollow = &gormigrate.Migration{
	ID: "2023071002-add-table-okr-follow",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrFollow{})
	},
}

// 添加关键结果表 AddTableOkrKeyResult
var AddTableOkrKeyResult = &gormigrate.Migration{
	ID: "2023071003-add-table-okr-key-result",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrAlign{})
	},
}

// 添加复盘表 AddTableOkrReplay
var AddTableOkrReplay = &gormigrate.Migration{
	ID: "2023071004-add-table-okr-replay",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrReplay{})
	},
}

// 添加日志表 AddTableOkrLog
var AddTableOkrLog = &gormigrate.Migration{
	ID: "2023071005-add-table-okr-log",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrLog{})
	},
}
