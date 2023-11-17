package seeders

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"
)

// SeedOkrReplayTable okr复盘表演示数据
func SeedOkrReplayTable() {
	db := core.DB
	if db.Migrator().HasTable(&model.OkrReplay{}) {
		var count int64
		db.Model(&model.OkrReplay{}).Count(&count)
		if count == 0 {
			okrReplays := []model.OkrReplay{
				{
					Id:              1,
					Userid:          1,
					OkrId:           10,
					OkrAscription:   2,
					OkrUserid:       1,
					OkrDepartmentId: "",
					OkrTitle:        "在本季度提高产品质量",
					OkrProgress:     100,
					OkrPriority:     "P0",
					Review:          `1、提升了产品质量：通过实施全面的产品测试计划，我们能够发现并解决许多潜在的问题和缺陷，从而提高了产品的质量水平。2、提高了迭代的测试覆盖率：达到每个迭代中的测试覆盖率达 90%，意味着我们对核心功能和代码路径进行了更全面和深入的测试，减少了潜在的风险。3、提升用户满意度：通过产品的改进和优化，我们能够满足用户的需求和期望，提高用户的满意度，进而增加用户的忠诚度和留存率。4、增加用户口碑和推荐：用户满意度的提高往往也会带来用户的口碑传播和积极的推荐，吸引更多的新用户加入。`,
					Problem:         `1、测试计划的制定不够详细：如果我们对测试计划的制定过程更加细致和详尽，能够更好地规划测试的范围、方法和资源，提高测试的效果。2、测试自动化程度不够高：如果我们能够进一步提高测试的自动化程度，减少人工测试的工作量，既能提高测试的效率，又能减少人为错误的发生。3、用户反馈收集不全面：如果我们能够更加主动地收集用户的反馈和意见，包括用户调研、用户体验测试、用户反馈渠道等，可以更准确地了解用户需求，从而针对性地改进产品。`,
				},
			}
			db.CreateInBatches(okrReplays, 1)
		}
	}
}
