package database

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
	"dootask-okr/config"
	"dootask-okr/database/migrations"
	"dootask-okr/database/seeders"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Init() error {
	options := &gormigrate.Options{
		TableName:                 config.CONF.System.Prefix + "okr_migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: true,
	}

	db := core.DB
	// 判断是否是mysql
	if db.Dialector.Name() == "mysql" {
		db = db.Set("gorm:table_options", "CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	}

	// 删除旧的迁移文件
	exists := db.Migrator().HasTable("migrations")
	if exists {
		var count int64
		db.Raw("SELECT COUNT(*) FROM migrations WHERE id = ?", "2023071001-add-table-okr").Count(&count)
		if count > 0 {
			db.Exec("DROP TABLE IF EXISTS migrations")
		}
	}

	one := db.Migrator().HasTable(&model.Okr{})

	// 清理可能存在的冲突迁移记录
	migrationTableName := config.CONF.System.Prefix + "okr_migrations"
	if db.Migrator().HasTable(migrationTableName) {
		// 获取数据库中所有的迁移记录
		var existingMigrations []string
		db.Raw("SELECT id FROM " + migrationTableName).Scan(&existingMigrations)

		// 定义当前代码中存在的迁移ID
		validMigrations := map[string]bool{
			"2023071001-add-table-okr":                        true,
			"2023071002-add-table-okr-follow":                 true,
			"2023071003-add-table-okr-key-align":              true,
			"2023071004-add-table-okr-replay":                 true,
			"2023071005-add-table-okr-replay-history":         true,
			"2023071015-add-table-okr-log":                    true,
			"2023071006-add-table-okr-replay-history-comment": true,
			"2023071007-add-table-okr-score":                  true,
			"2023071008-add-table-okr-log-record":             true,
			"2023081408-add-table-okr-replay-problem":         true,
			"2023120709-add-table-okr-setting":                true,
			"2023120710-add-table-okr-push-log":               true,
			"2023120811-add-table-okr-11version":              true,
		}

		// 删除不存在于代码中的迁移记录
		for _, migrationID := range existingMigrations {
			if !validMigrations[migrationID] {
				db.Exec("DELETE FROM "+migrationTableName+" WHERE id = ?", migrationID)
			}
		}
	}

	m := gormigrate.New(db, options, []*gormigrate.Migration{
		migrations.AddTableOkr,
		migrations.AddTableOkrFollow,
		migrations.AddTableOkrAlign,
		migrations.AddTableOkrReplay,
		migrations.AddTableOkrReplayHistory,
		migrations.AddTableOkrLog,
		migrations.AddTableOkrReplayHistoryComment,
		migrations.AddTableOkrScore,
		migrations.AddTableOkrLogRecord,
		migrations.AddTableOkrReplayProblem,
		migrations.AddTableOkrSetting,
		migrations.AddTableOkrPushLog,
		migrations.AddTableOkr11Version,
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	two := db.Migrator().HasTable(&model.Okr{})
	IsInitOkrEmpty := false
	if !one && two {
		IsInitOkrEmpty = true
	}

	// 演示数据
	if os.Getenv("DEMO_DATA") == "true" && IsInitOkrEmpty {
		if _, err := os.Stat(".cache/demo_data"); os.IsNotExist(err) {
			seeders.SeedOkrTable()
			seeders.SeedOkrLogTable()
			seeders.SeedOkrFollowTable()
			seeders.SeedOkrAlignTable()
			seeders.SeedOkrReplayTable()
			seeders.SeedOkrReplayHistoryTable()

			_, err := os.Create(".cache/demo_data")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
