package service

import (
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/config"
	"fmt"
	"strings"
)

var OkrAnalyzeService = okrAnalyzeService{}

type okrAnalyzeService struct{}

// 获取OKR整体平均完成度
func (s *okrAnalyzeService) GetOverallCompleteness(userid any) (*interfaces.OkrAnalyzeOverall, error) {
	var okrAnalyzeOverall interfaces.OkrAnalyzeOverall
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0")
	if err := db.Session(&core.Session).Count(&okrAnalyzeOverall.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where(&model.Okr{Completed: 1}).Count(&okrAnalyzeOverall.Completed).Error; err != nil {
		return nil, err
	}
	return &okrAnalyzeOverall, nil
}

// 获取OKR各部门平均完成度
func (s *okrAnalyzeService) GetDeptCompleteness(userid any) (*[]interfaces.OkrAnalyzeDept, error) {
	var okrAnalyzeDepts []interfaces.OkrAnalyzeDept
	okrTable := core.DBTableName(&model.Okr{})
	userDepartmentTable := core.DBTableName(&model.UserDepartment{})

	db := core.DB.Table(okrTable + " AS a").
		Joins(fmt.Sprintf(`
			LEFT JOIN (
				SELECT department_id, COUNT(*) as total, SUM(CASE WHEN completed != 0 THEN 1 ELSE 0 END) as completed FROM %s 
				GROUP BY department_id
			) b on a.department_id = b.department_id
		`, okrTable))

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		db = db.Joins(fmt.Sprintf("LEFT JOIN %s AS dept ON dept.id = a.department_id", userDepartmentTable))
		db = db.Select("DISTINCT(a.department_id), b.total, b.completed, dept.name as department_name")
	} else {
		db = db.Select("DISTINCT(a.department_id), b.total, b.completed, a.department_id as department_name")
	}

	db = db.Where("a.parent_id = 0").Where("a.department_id <> 0").Find(&okrAnalyzeDepts)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &okrAnalyzeDepts, nil
}

// 获取OKR评分分布
func (s *okrAnalyzeService) GetScore(userid any) (*interfaces.OkrAnalyzeScore, error) {
	var okrAnalyzeScore interfaces.OkrAnalyzeScore
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0")
	if err := db.Session(&core.Session).Count(&okrAnalyzeScore.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score = 0").Count(&okrAnalyzeScore.Unscored).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score > 0 and score <= 3").Count(&okrAnalyzeScore.ZeroToThree).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score > 3 and score <= 7").Count(&okrAnalyzeScore.ThreeToSeven).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score > 7 and score <= 10").Count(&okrAnalyzeScore.SevenToTen).Error; err != nil {
		return nil, err
	}
	return &okrAnalyzeScore, nil
}

// 获取OKR各部门评分分布
func (s *okrAnalyzeService) GetDeptScore(userid any) (*[]interfaces.OkrAnalyzeScoreDept, error) {
	var okrAnalyzeScoreDepts []interfaces.OkrAnalyzeScoreDept
	okrTable := core.DBTableName(&model.Okr{})
	userDepartmentTable := core.DBTableName(&model.UserDepartment{})

	db := core.DB.Table(okrTable + " AS a").
		Joins(fmt.Sprintf(`
			LEFT JOIN (
				SELECT department_id, 
					COUNT(*) as total, 
					SUM(CASE WHEN score = 0 THEN 1 ELSE 0 END) as unscored, 
					SUM(CASE WHEN score > 0 and score <= 3 THEN 1 ELSE 0 END) as zero_to_three, 
					SUM(CASE WHEN score > 3 and score <= 7 THEN 1 ELSE 0 END) as three_to_seven, 
					SUM(CASE WHEN score > 7 and score <= 10 THEN 1 ELSE 0 END) as seven_to_ten
				FROM %s 
				GROUP BY department_id
			) b on a.department_id = b.department_id
		`, okrTable))

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		db = db.Joins(fmt.Sprintf("LEFT JOIN %s AS dept ON dept.id = a.department_id", userDepartmentTable))
		db = db.Select("DISTINCT(a.department_id), b.total, b.*, dept.name as department_name")
	} else {
		db = db.Select("DISTINCT(a.department_id), b.total, b.*, a.department_id as department_name")
	}

	db = db.Where("a.parent_id = 0").Where("a.department_id <> 0").Find(&okrAnalyzeScoreDepts)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &okrAnalyzeScoreDepts, nil
}

// 获取OKR评分分布
func (s *okrAnalyzeService) GetPersonnel(userid any) (*interfaces.OkrAnalyzePersonnel, error) {
	var okrAnalyzePersonnel interfaces.OkrAnalyzePersonnel

	// 总人员数量
	if err := core.DB.Model(model.Okr{}).Where("parent_id > 0").Select("Count(DISTINCT(userid)) as total").Find(&okrAnalyzePersonnel.Total).Error; err != nil {
		return nil, err
	}

	// 已完成的人员数量： 个人必须全部KR都完成评分，才计为已完成评分人员（部门负责人已为自己的目标评分，而未为下级评分时，部门负责人的评分工作属于未完成)
	okrTable := core.DBTableName(&model.Okr{})
	db := core.DB.Table(okrTable + " AS a")

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		userDepartmentTable := core.DBTableName(&model.UserDepartment{})
		db = db.Joins(fmt.Sprintf(`
			LEFT JOIN (
				SELECT u.userid, 
					COUNT(*) as total, 
					SUM(CASE WHEN u.score > 0 AND (dept.owner_userid is null or dept.owner_userid != u.userid OR uu.total = uu.completed) THEN 1 ELSE 0 END) as completed 
				FROM %s u 
				LEFT JOIN %s dept ON dept.id = u.department_id
				LEFT JOIN (
					SELECT department_id, 
						COUNT(*) as total, 
						SUM(CASE WHEN superior_score > 0 THEN 1 ELSE 0 END) as completed 
					FROM %s WHERE parent_id = 0 
					GROUP BY department_id
				) uu on u.department_id = uu.department_id
				WHERE u.parent_id > 0  AND ( dept.owner_userid != u.userid OR uu.total = uu.completed )
				GROUP BY userid
			) b on a.userid = b.userid 
		`, okrTable, userDepartmentTable, okrTable))
	} else {
		db = db.Joins(fmt.Sprintf(`
			LEFT JOIN (
				SELECT userid, COUNT(*) as total, SUM(CASE WHEN score > 0 THEN 1 ELSE 0 END) as completed
				FROM %s WHERE parent_id > 0
				GROUP BY userid
			) b on a.userid = b.userid
		`, okrTable))
	}

	db = db.Where("a.parent_id > 0").Where("b.total = b.completed").Select("Count(DISTINCT(a.userid)) as completed")

	if err := db.Find(&okrAnalyzePersonnel.Completed).Error; err != nil {
		return nil, err
	}

	return &okrAnalyzePersonnel, nil
}
