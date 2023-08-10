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

/**
* 获取OKR整体平均完成度
* 公式：【所有O的完成度之和/O的数量】
* @param user 用户
 */
func (s *okrAnalyzeService) GetOverallCompleteness(user *interfaces.UserInfoResp) (*interfaces.OkrAnalyzeOverall, error) {
	var data interfaces.OkrAnalyzeOverall
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0 and canceled = 0")
	if err := db.Session(&core.Session).Count(&data.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where(&model.Okr{Completed: 1}).Count(&data.Completed).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

/**
* 获取OKR各部门平均完成度
* 公式：各部门所有已完成的O/O的数量（按完成度高到低显示，显示不完可使用滚动条滚动显示）【各部门所有O的完成度之和/各部门O的数量】
* @param user 用户
 */
func (s *okrAnalyzeService) GetDeptCompleteness(user *interfaces.UserInfoResp) (*[]interfaces.OkrAnalyzeDept, error) {
	var data []interfaces.OkrAnalyzeDept

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(departmentTable + " AS dept").Joins(fmt.Sprintf(`
				LEFT JOIN %s dept_two on dept_two.parent_id = dept.id
				LEFT JOIN (
					SELECT u.department, 
						COUNT(*) as total, 
						SUM(CASE WHEN okr.completed != 0 THEN 1 ELSE 0 END) as completed 
					FROM %s okr
					LEFT JOIN %s u on okr.userid = u.userid
					where u.userid > 0 
						and u.department <> '' 
						and okr.parent_id = 0 
						and okr.canceled = 0
					GROUP BY u.department
				) b on find_in_set(dept.id,b.department) or find_in_set(dept_two.id, b.department) 
			`, departmentTable, okrTable, userTable)).
			Select(`
				dept.id as department_id, 
				dept.name as department_name, 
				SUM(ifnull(b.total,0)) total, 
				SUM(ifnull(b.completed,0)) completed
			`).
			Where("dept.parent_id = 0").
			Group("dept.id").
			Order("completed / total desc,b.total desc")
			// Where("b.total > ?", 0)
		if err := db.Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		data = append(data, interfaces.OkrAnalyzeDept{
			DepartmentId:   0,
			DepartmentName: "假数据",
			OkrAnalyzeOverall: &interfaces.OkrAnalyzeOverall{
				Total:     6,
				Completed: 3,
			},
		})
	}

	return &data, nil
}

/**
* 获取OKR评分分布
* 公式：显示O的评分分布
* @param user 用户
 */
func (s *okrAnalyzeService) GetScore(user *interfaces.UserInfoResp) (*interfaces.OkrAnalyzeScore, error) {
	var data interfaces.OkrAnalyzeScore
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0 and canceled = 0")
	if err := db.Session(&core.Session).Count(&data.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score < 0").Count(&data.Unscored).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score >= 0 and score <= 3").Count(&data.ZeroToThree).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score > 3 and score <= 7").Count(&data.ThreeToSeven).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score > 7 and score <= 10").Count(&data.SevenToTen).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

/**
* 获取OKR各部门评分分布
* 公式：显示各部门O的评分分布（O中某个KR完成评分也计入分数）
* @param user 用户
 */
func (s *okrAnalyzeService) GetDeptScore(user *interfaces.UserInfoResp) (*[]interfaces.OkrAnalyzeScoreDept, error) {
	var data []interfaces.OkrAnalyzeScoreDept

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(departmentTable + " AS dept").Joins(fmt.Sprintf(`
				LEFT JOIN (
					SELECT depts.id, 
						COUNT(*) as total, 
						SUM(CASE WHEN score < 0 THEN 1 ELSE 0 END) as unscored, 
						SUM(CASE WHEN score >= 0 and score <= 3 THEN 1 ELSE 0 END) as zero_to_three, 
						SUM(CASE WHEN score > 3 and score <= 7 THEN 1 ELSE 0 END) as three_to_seven, 
						SUM(CASE WHEN score > 7 and score <= 10 THEN 1 ELSE 0 END) as seven_to_ten
					FROM %s okr
					LEFT JOIN %s u on okr.userid = u.userid
					LEFT JOIN %s depts on find_in_set(depts.id,u.department) 
					where u.userid > 0 
						and u.department <> '' 
						and okr.parent_id = 0
						and okr.canceled = 0
					GROUP BY depts.id
				) b on dept.id = b.id OR b.id IN(select id FROM %s where parent_id = dept.id) 
			`, okrTable, userTable, departmentTable, departmentTable)).
			Select(`
				dept.id as department_id,
				dept.name as department_name,
				SUM(ifnull(b.total,0)) total,
				SUM(ifnull(b.unscored,0)) unscored,
				SUM(ifnull(b.zero_to_three,0)) zero_to_three,
				SUM(ifnull(b.three_to_seven,0)) three_to_seven,
				SUM(ifnull(b.seven_to_ten,0)) seven_to_ten
			`).
			Where("dept.parent_id = 0").
			Group("dept.id").
			Order("b.total desc")
		if err := db.Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		data = append(data, interfaces.OkrAnalyzeScoreDept{
			DepartmentId:   0,
			DepartmentName: "假数据",
			OkrAnalyzeScore: &interfaces.OkrAnalyzeScore{
				Total:        6,
				Unscored:     3,
				ZeroToThree:  1,
				ThreeToSeven: 1,
				SevenToTen:   1,
			},
		})
	}

	return &data, nil
}

/**
* 获取OKR员评分率
* 公式：已完成评分的OKR所占比例，一个OKR里负责人与上级都完成评分，才能计为完成评分的OKR【所有已完成评分的O/O的数量】
* @param user 用户
 */
func (s *okrAnalyzeService) GetPersonnelScoreRate(user *interfaces.UserInfoResp) (*interfaces.OkrAnalyzePersonnelScoreRate, error) {
	var data interfaces.OkrAnalyzePersonnelScoreRate
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0 and canceled = 0")
	// 总okr数量
	if err := db.Session(&core.Session).Select("Count(*) as total").Find(&data.Total).Error; err != nil {
		return nil, err
	}
	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		if err := db.Session(&core.Session).Select("Count(*) as completed").Where("score > -1").Find(&data.Completed).Error; err != nil {
			return nil, err
		}
	} else {
		data.Total = 0
		data.Completed = 0
	}
	return &data, nil
}

/**
* 获取OKR部门评分占比
* 公式：各个部门完成OKR评分的所占比例【各部门所有已完成评分的O/各部门内O的数量】
* @param user 用户
 */
func (s *okrAnalyzeService) GetDeptScoreProportion(user *interfaces.UserInfoResp) (*[]interfaces.OkrAnalyzeDeptScoreProportion, error) {
	var data []interfaces.OkrAnalyzeDeptScoreProportion

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(departmentTable + " AS dept").Joins(fmt.Sprintf(`
				LEFT JOIN (
					SELECT depts.id, 
						count(*) as okr_total,
						SUM(CASE WHEN okr.score > -1 THEN 1 ELSE 0 END) as completed 
					FROM %s okr
					LEFT JOIN %s u on okr.userid = u.userid
					LEFT JOIN %s depts on find_in_set(depts.id,u.department) 
					where u.userid > 0 and u.department <> '' and okr.parent_id = 0 and okr.canceled = 0
					GROUP BY depts.id
				) b on dept.id = b.id OR b.id IN(select id FROM %s where parent_id = dept.id) 
			`, okrTable, userTable, departmentTable, departmentTable)).
			Select(`
				dept.id as department_id,
				dept.name as department_name,
				SUM(ifnull(b.okr_total,0)) total,
				SUM(ifnull(b.okr_total,0) - ifnull(b.completed,0)) unscored,
				SUM(ifnull(b.completed,0)) already_reviewed
			`).
			Where("dept.parent_id = 0").
			Group("dept.id").
			Order("total desc")
		if err := db.Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		data = append(data, interfaces.OkrAnalyzeDeptScoreProportion{
			DepartmentId:    0,
			DepartmentName:  "假数据",
			Total:           2,
			Unscored:        1,
			AlreadyReviewed: 1,
		})
	}

	return &data, nil
}
