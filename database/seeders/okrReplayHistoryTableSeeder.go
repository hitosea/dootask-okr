package seeders

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
)

// SeedOkrReplayHistoryTable okr复盘子表演示数据
func SeedOkrReplayHistoryTable() {
	db := core.DB
	if db.Migrator().HasTable(&model.OkrReplayHistory{}) {
		var count int64
		db.Model(&model.OkrReplayHistory{}).Count(&count)
		if count == 0 {
			okrReplayHistory := []model.OkrReplayHistory{
				{
					Id:         1,
					ReplayId:   1,
					OkrId:      11,
					Title:      "实施全面的产品测试计划，并达到每个迭代中的测试覆盖率达 90%",
					ParentId:   10,
					Progress:   100,
					Confidence: 100,
					Score:      9,
					Comment:    1,
				},
				{
					Id:         2,
					ReplayId:   1,
					OkrId:      12,
					Title:      "提高产品的用户满意度评分至少 5%",
					ParentId:   10,
					Progress:   100,
					Confidence: 95,
					Score:      8,
					Comment:    2,
				},
			}
			db.CreateInBatches(okrReplayHistory, 1)
		}
	}
}
