package seeders

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
)

// SeedOkrFollowTable okr关注表演示数据
func SeedOkrFollowTable() {
	db := core.DB
	if db.Migrator().HasTable(&model.OkrFollow{}) {
		var count int64
		db.Model(&model.OkrFollow{}).Count(&count)
		if count == 0 {
			okrFollow := []model.OkrFollow{
				{
					Id:     1,
					OkrId:  7,
					Userid: 1,
				},
				{
					Id:     2,
					OkrId:  10,
					Userid: 1,
				},
				{
					Id:     3,
					OkrId:  1,
					Userid: 2,
				},
				{
					Id:     4,
					OkrId:  13,
					Userid: 2,
				},
			}
			db.CreateInBatches(okrFollow, 1)
		}
	}
}
