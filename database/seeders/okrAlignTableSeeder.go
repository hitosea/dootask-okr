package seeders

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
)

// SeedOkrAlignTable okr对齐目标表演示数据
func SeedOkrAlignTable() {
	db := core.DB
	if db.Migrator().HasTable(&model.OkrAlign{}) {
		var count int64
		db.Model(&model.OkrAlign{}).Count(&count)
		if count == 0 {
			okrAlign := []model.OkrAlign{
				{
					Id:         1,
					OkrId:      4,
					AlignOkrId: 1,
				},
			}
			db.CreateInBatches(okrAlign, 1)
		}
	}
}
