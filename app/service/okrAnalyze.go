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
* @param user 用户
 */
func (s *okrAnalyzeService) GetOverallCompleteness(user *interfaces.UserInfoResp) (*interfaces.OkrAnalyzeOverall, error) {
	var data interfaces.OkrAnalyzeOverall
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0")
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
* @param user 用户
 */
func (s *okrAnalyzeService) GetDeptCompleteness(user *interfaces.UserInfoResp) (*[]interfaces.OkrAnalyzeDept, error) {
	var data []interfaces.OkrAnalyzeDept

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(departmentTable+" AS dept").Joins(fmt.Sprintf(`
				LEFT JOIN (
					SELECT u.department, 
						COUNT(*) as total, 
						SUM(CASE WHEN okr.completed != 0 THEN 1 ELSE 0 END) as completed 
					FROM %s okr
					LEFT JOIN %s u on okr.userid = u.userid
					where u.userid > 0 and u.department <> '' and okr.parent_id = 0
					GROUP BY u.department
				) b on find_in_set(dept.id,b.department)
			`, okrTable, userTable)).
			Select("dept.id as department_id, dept.name as department_name, SUM(ifnull(b.total,0)) total, SUM(ifnull(b.completed,0)) completed").
			Group("department_id").
			Where("b.total > ?", 0)
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
* @param user 用户
 */
func (s *okrAnalyzeService) GetScore(user *interfaces.UserInfoResp) (*interfaces.OkrAnalyzeScore, error) {
	var data interfaces.OkrAnalyzeScore
	db := core.DB.Model(model.Okr{}).Where("parent_id = 0")
	if err := db.Session(&core.Session).Count(&data.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score = 0").Count(&data.Unscored).Error; err != nil {
		return nil, err
	}
	if err := db.Session(&core.Session).Where("score > 0 and score <= 3").Count(&data.ZeroToThree).Error; err != nil {
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
* @param user 用户
 */
func (s *okrAnalyzeService) GetDeptScore(user *interfaces.UserInfoResp) (*[]interfaces.OkrAnalyzeScoreDept, error) {
	var data []interfaces.OkrAnalyzeScoreDept

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(departmentTable+" AS dept").Joins(fmt.Sprintf(`
				LEFT JOIN (
					SELECT u.department, 
						COUNT(*) as total, 
						SUM(CASE WHEN score = 0 THEN 1 ELSE 0 END) as unscored, 
						SUM(CASE WHEN score > 0 and score <= 3 THEN 1 ELSE 0 END) as zero_to_three, 
						SUM(CASE WHEN score > 3 and score <= 7 THEN 1 ELSE 0 END) as three_to_seven, 
						SUM(CASE WHEN score > 7 and score <= 10 THEN 1 ELSE 0 END) as seven_to_ten
					FROM %s okr
					LEFT JOIN %s u on okr.userid = u.userid
					where u.userid > 0 and u.department <> '' and okr.parent_id = 0
					GROUP BY u.department
				) b on find_in_set(dept.id,b.department)
			`, okrTable, userTable)).
			Select(`
				dept.id as department_id,
				dept.name as department_name,
				SUM(ifnull(b.total,0)) total,
				SUM(ifnull(b.unscored,0)) unscored,
				SUM(ifnull(b.zero_to_three,0)) zero_to_three,
				SUM(ifnull(b.three_to_seven,0)) three_to_seven,
				SUM(ifnull(b.seven_to_ten,0)) seven_to_ten
			`).
			Group("department_id").
			Where("b.total > ?", 0)

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
* 获取OKR人员评分分布
* @param user 用户
 */
func (s *okrAnalyzeService) GetPersonnelScoreRate(user *interfaces.UserInfoResp) (*interfaces.OkrAnalyzePersonnelScoreRate, error) {
	var data interfaces.OkrAnalyzePersonnelScoreRate

	// 总人员数量
	if err := core.DB.Model(model.Okr{}).Select("Count(DISTINCT(userid)) as total").Find(&data.Total).Error; err != nil {
		return nil, err
	}

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		// 已完成的人员数量： 个人必须全部KR都完成评分，才计为已完成评分人员（部门负责人已为自己的目标评分，而未为下级评分时，部门负责人的评分工作属于未完成)
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(okrTable+" AS a").Joins(fmt.Sprintf(`
			LEFT JOIN (
				select 
					okr.userid,
					COUNT(*) as total, 
					SUM(CASE WHEN okr.completed > 0 and (dept.owner_userid is null OR  uu.total = uu.completed ) THEN 1 ELSE 0 END) as completed
				from %s okr
				LEFT JOIN %s users on okr.userid = users.userid
				LEFT JOIN %s dept on find_in_set(dept.id, users.department) and dept.owner_userid != okr.userid
				LEFT JOIN (
					SELECT dept.id, 
						COUNT(*) as total, 
						SUM(CASE WHEN okr.superior_score > 0 THEN 1 ELSE 0 END) as completed 
					FROM %s okr
					LEFT JOIN %s users on okr.userid = users.userid
					LEFT JOIN %s dept on find_in_set(dept.id, users.department) 
					WHERE okr.parent_id = 0 
					GROUP BY dept.id
				) uu on dept.id = uu.id
				GROUP BY okr.userid
			) b on a.userid = b.userid 
		`, okrTable, userTable, departmentTable, okrTable, userTable, departmentTable)).
			Where("a.parent_id", 0).
			Where("b.total = b.completed").
			Select("Count(DISTINCT(a.userid)) as completed")
		if err := db.Find(&data.Completed).Error; err != nil {
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
* @param user 用户
 */
func (s *okrAnalyzeService) GetDeptScoreProportion(user *interfaces.UserInfoResp) (*[]interfaces.OkrAnalyzeDeptScoreProportion, error) {
	var data []interfaces.OkrAnalyzeDeptScoreProportion

	// 检查部门表是否存在
	if !strings.Contains(config.CONF.System.Dsn, "sqlite") {
		okrTable := core.DBTableName(&model.Okr{})
		userTable := core.DBTableName(&model.User{})
		departmentTable := core.DBTableName(&model.UserDepartment{})
		db := core.DB.Table(departmentTable+" AS dept").Joins(fmt.Sprintf(`
				LEFT JOIN (
					SELECT u.department, 
						COUNT(*) as total, 
						SUM(CASE WHEN score = 0 THEN 1 ELSE 0 END) as unscored, 
						SUM(CASE WHEN score > 0 THEN 1 ELSE 0 END) as already_reviewed
					FROM %s okr
					LEFT JOIN %s u on okr.userid = u.userid
					where u.userid > 0 and u.department <> '' and okr.parent_id = 0
					GROUP BY u.department
				) b on find_in_set(dept.id,b.department)
			`, okrTable, userTable)).
			Select(`
				dept.id as department_id,
				dept.name as department_name,
				SUM(ifnull(b.total,0)) total,
				SUM(ifnull(b.unscored,0)) unscored,
				SUM(ifnull(b.already_reviewed,0)) already_reviewed
			`).
			Group("department_id").
			Where("b.total > ?", 0)

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
