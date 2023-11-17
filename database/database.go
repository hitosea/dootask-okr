package database

import (
	"dootask-okr/app/core"
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

	// 演示数据
	if os.Getenv("DEMO_DATA") == "true" {
		seeders.SeedOkrTable()
		seeders.SeedOkrLogTable()
		seeders.SeedOkrFollowTable()
		seeders.SeedOkrAlignTable()
		seeders.SeedOkrReplayTable()
		seeders.SeedOkrReplayHistoryTable()
	}

	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
