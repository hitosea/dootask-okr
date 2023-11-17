package seeders

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
)

// SeedOkrLogTable okr日志表演示数据
func SeedOkrLogTable() {
	db := core.DB
	if db.Migrator().HasTable(&model.OkrLog{}) {
		var count int64
		db.Model(&model.OkrLog{}).Count(&count)
		if count == 0 {
			okrLogs := []model.OkrLog{
				{
					Id:        1,
					OkrId:     1,
					Userid:    1,
					Operation: "add",
					Content:   "创建OKR",
					Record:    "",
				},
				{
					Id:        2,
					OkrId:     4,
					Userid:    1,
					Operation: "add",
					Content:   "创建OKR",
					Record:    "",
				},
				{
					Id:        3,
					OkrId:     7,
					Userid:    1,
					Operation: "add",
					Content:   "创建OKR",
					Record:    "",
				},
				{
					Id:        4,
					OkrId:     10,
					Userid:    1,
					Operation: "add",
					Content:   "创建OKR",
					Record:    "",
				},
				{
					Id:        5,
					OkrId:     13,
					Userid:    1,
					Operation: "add",
					Content:   "创建OKR",
					Record:    "",
				},

				{
					Id:        6,
					OkrId:     10,
					Userid:    1,
					Operation: "update",
					Content:   "修改KR进度",
					Record:    `{"title":"实施全面的产品测试计划，并达到每个迭代中的测试覆盖率达 90%","progress_change":[0,100]}`,
				},
				{
					Id:        7,
					OkrId:     10,
					Userid:    1,
					Operation: "update",
					Content:   "修改KR状态",
					Record:    `{"title":"实施全面的产品测试计划，并达到每个迭代中的测试覆盖率达 90%","progress_status_change":["未开始","正常"]}`,
				},
				{
					Id:        8,
					OkrId:     10,
					Userid:    1,
					Operation: "update",
					Content:   "责任人打分",
					Record:    `{"title":"实施全面的产品测试计划，并达到每个迭代中的测试覆盖率达 90%"}`,
				},
				{
					Id:        9,
					OkrId:     10,
					Userid:    1,
					Operation: "update",
					Content:   "修改KR进度",
					Record:    `{"title":"提高产品的用户满意度评分至少 5%","progress_change":[0,100]}`,
				},
				{
					Id:        10,
					OkrId:     10,
					Userid:    1,
					Operation: "update",
					Content:   "责任人打分",
					Record:    `{"title":"提高产品的用户满意度评分至少 5%"}`,
				},
				{
					Id:        11,
					OkrId:     13,
					Userid:    1,
					Operation: "update",
					Content:   "修改O目标状态",
					Record:    `{"status_change":["正常","已取消"]}`,
				},
				{
					Id:        12,
					OkrId:     7,
					Userid:    1,
					Operation: "update",
					Content:   "修改KR进度",
					Record:    `{"title":"实施市场推广活动，使市场份额增长 10%","progress_change":[0,80]}`,
				},
				{
					Id:        13,
					OkrId:     1,
					Userid:    1,
					Operation: "update",
					Content:   "修改KR进度",
					Record:    `{"title":"增加每日活跃用户数","progress_change":[0,100]}`,
				},
			}
			db.CreateInBatches(okrLogs, 1)
		}
	}
}
