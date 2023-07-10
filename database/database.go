package database

import (
	"dootask-okr/app/core"
	"dootask-okr/database/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Init() error {
	m := gormigrate.New(core.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.AddTableOkrObject,
		migrations.AddTableOkrFollow,
		migrations.AddTableOkrKeyResult,
		migrations.AddTableOkrReplay,
	})
	if err := m.Migrate(); err != nil {
		return err
	}
	return nil
}
