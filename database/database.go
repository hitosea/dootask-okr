package database

import (
	"dootask-okr/app/core"
	"dootask-okr/database/migrations"

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
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
