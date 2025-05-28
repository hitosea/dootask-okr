package migrations

import (
	"dootask-okr/app/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// 添加目标表 AddTableOkr
var AddTableOkr = &gormigrate.Migration{
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

// 添加关键结果表 AddTableOkrAlign
var AddTableOkrAlign = &gormigrate.Migration{
	ID: "2023071003-add-table-okr-key-align",
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

// 添加复盘历史表 AddTableOkrReplayHistory
var AddTableOkrReplayHistory = &gormigrate.Migration{
	ID: "2023071005-add-table-okr-replay-history",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrReplayHistory{})
	},
}

// 添加日志表 AddTableOkrLog
var AddTableOkrLog = &gormigrate.Migration{
	ID: "2023071015-add-table-okr-log",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrLog{})
	},
}

// 更改复盘历史表 comment 字段 AddTableOkrReplayHistoryComment
var AddTableOkrReplayHistoryComment = &gormigrate.Migration{
	ID: "2023071006-add-table-okr-replay-history-comment",
	Migrate: func(tx *gorm.DB) error {
		if !tx.Migrator().HasColumn(&model.OkrReplayHistory{}, "comment") {
			if err := tx.Migrator().AlterColumn(&model.OkrReplayHistory{}, "comment"); err != nil {
				return err
			}
			return nil
		}
		return nil
	},
}

// 更改okr表 score和superior_score字段 AddTableOkrScore
var AddTableOkrScore = &gormigrate.Migration{
	ID: "2023071007-add-table-okr-score",
	Migrate: func(tx *gorm.DB) error {
		if tx.Migrator().HasColumn(&model.Okr{}, "score") {
			if err := tx.Migrator().AlterColumn(&model.Okr{}, "score"); err != nil {
				return err
			}
			return nil
		}
		if tx.Migrator().HasColumn(&model.Okr{}, "superior_score") {
			if err := tx.Migrator().AlterColumn(&model.Okr{}, "superior_score"); err != nil {
				return err
			}
			return nil
		}
		return nil
	},
}

// 更改日志表 新增record字段 AddTableOkrLogRecord
var AddTableOkrLogRecord = &gormigrate.Migration{
	ID: "2023071008-add-table-okr-log-record",
	Migrate: func(tx *gorm.DB) error {
		if !tx.Migrator().HasColumn(&model.OkrLog{}, "record") {
			if err := tx.Migrator().AddColumn(&model.OkrLog{}, "record"); err != nil {
				return err
			}
			return nil
		}
		return nil
	},
}

// 更改复盘表 新增problem字段 AddTableOkrReplayProblem
var AddTableOkrReplayProblem = &gormigrate.Migration{
	ID: "2023081408-add-table-okr-replay-problem",
	Migrate: func(tx *gorm.DB) error {
		if tx.Migrator().HasColumn(&model.OkrReplay{}, "review") {
			if err := tx.Migrator().AlterColumn(&model.OkrReplay{}, "review"); err != nil {
				return err
			}
			return nil
		}
		if !tx.Migrator().HasColumn(&model.OkrReplay{}, "problem") {
			if err := tx.Migrator().AddColumn(&model.OkrReplay{}, "problem"); err != nil {
				return err
			}
			return nil
		}
		return nil
	},
}

// 新增设置表 AddTableOkrSetting
var AddTableOkrSetting = &gormigrate.Migration{
	ID: "2023120709-add-table-okr-setting",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrSetting{})
	},
}

// 新增okr发送记录表 AddTableOkrPushLog
var AddTableOkrPushLog = &gormigrate.Migration{
	ID: "2023120710-add-table-okr-push-log",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.OkrPushLog{})
	},
}

// 更改okr表 新增1.1版本需求字段 AddTableOkr11Version
var AddTableOkr11Version = &gormigrate.Migration{
	ID: "2023120811-add-table-okr-11version",
	Migrate: func(tx *gorm.DB) error {
		columns := []string{"score_num", "score_completed_at", "status", "archive_userid", "archive_at"}
		for _, column := range columns {
			if !tx.Migrator().HasColumn(&model.Okr{}, column) {
				if err := tx.Migrator().AddColumn(&model.Okr{}, column); err != nil {
					return err
				}
			}
		}

		columns2 := []string{"replay"}
		for _, column := range columns2 {
			if !tx.Migrator().HasColumn(&model.OkrReplay{}, column) {
				if err := tx.Migrator().AddColumn(&model.OkrReplay{}, column); err != nil {
					return err
				}
			}
		}
		return nil
	},
}
