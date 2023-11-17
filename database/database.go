package database

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
	"dootask-okr/database/migrations"
	"dootask-okr/database/seeders"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Init() error {
	options := &gormigrate.Options{
		TableName:                 "migrations",
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

	one := db.Migrator().HasTable(&model.Okr{})

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
