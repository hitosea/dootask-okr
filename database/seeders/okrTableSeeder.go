package seeders

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
	"time"
)

// SeedOkrTable okr表演示数据
func SeedOkrTable() {
	db := core.DB
	if db.Migrator().HasTable(&model.Okr{}) {
		var count int64
		db.Model(&model.Okr{}).Count(&count)
		if count == 0 {
			okrs := []model.Okr{
				{
					Id:             1,
					ParentId:       0,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "提升产品用户活跃度：增加每日活跃用户数，优化用户体验， 提高用户满意度",
					Type:           1,
					Priority:       "P0",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       50,
					ProgressStatus: 0,
					Confidence:     0,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             2,
					ParentId:       1,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "增加每日活跃用户数",
					Type:           1,
					Priority:       "P0",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       100,
					ProgressStatus: 0,
					Confidence:     90,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             3,
					ParentId:       1,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "优化用户体验， 提高用户满意度",
					Type:           1,
					Priority:       "P0",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     98,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             4,
					ParentId:       0,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "在本季度增加用户注册量",
					Type:           1,
					Priority:       "P1",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     0,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             5,
					ParentId:       4,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "实施市场营销活动，吸引更多的潜在用户注册",
					Type:           1,
					Priority:       "P1",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     90,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             6,
					ParentId:       4,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "优化用户注册流程，提高转化率至少 30%",
					Type:           1,
					Priority:       "P1",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     100,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             7,
					ParentId:       0,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "在本季度增加市场份额",
					Type:           2,
					Priority:       "P2",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       40,
					ProgressStatus: 0,
					Confidence:     0,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             8,
					ParentId:       7,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "实施市场推广活动，使市场份额增长 10%",
					Type:           2,
					Priority:       "P2",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       80,
					ProgressStatus: 0,
					Confidence:     80,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             9,
					ParentId:       7,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "提高现有客户的留存率至少 5%",
					Type:           2,
					Priority:       "P2",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     70,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             10,
					ParentId:       0,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "在本季度提高产品质量",
					Type:           1,
					Priority:       "P0",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      1,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       100,
					ProgressStatus: 0,
					Confidence:     0,
					Score:          8.5,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             11,
					ParentId:       10,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "实施全面的产品测试计划，并达到每个迭代中的测试覆盖率达 90%",
					Type:           1,
					Priority:       "P0",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       100,
					ProgressStatus: 1,
					Confidence:     100,
					Score:          9,
					SuperiorScore:  9,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             12,
					ParentId:       10,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "提高产品的用户满意度评分至少 5%",
					Type:           1,
					Priority:       "P0",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "2",
					Progress:       100,
					ProgressStatus: 0,
					Confidence:     95,
					Score:          8,
					SuperiorScore:  8,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             13,
					ParentId:       0,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "提升团队协作与沟通效率",
					Type:           2,
					Priority:       "P2",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       1,
					Participant:    "1,2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     0,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             14,
					ParentId:       13,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "实施团队培训计划，提高每个成员的协作与沟通技能",
					Type:           2,
					Priority:       "P2",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "1,2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     60,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
				{
					Id:             15,
					ParentId:       13,
					Userid:         1,
					DepartmentId:   "",
					ProjectId:      0,
					DialogId:       0,
					Title:          "减少团队会议时间，使每个成员每周节省至少 2 小时",
					Type:           2,
					Priority:       "P2",
					Ascription:     2,
					VisibleRange:   1,
					Completed:      0,
					Canceled:       0,
					Participant:    "2",
					Progress:       0,
					ProgressStatus: 0,
					Confidence:     70,
					Score:          -1,
					SuperiorScore:  -1,
					StartAt:        time.Now(),
					EndAt:          time.Now().AddDate(0, 3, 0),
				},
			}
			db.CreateInBatches(okrs, 1)
		}
	}
}
