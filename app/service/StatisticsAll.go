package service

import (
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
)

// 获取个人完成与未完成目标数量

func StatisticsAll(userID int) (*interfaces.StatisticsAllS, error) {
	var statisticsAll interfaces.StatisticsAllS

	db := core.DB.Model(&model.Okr{}).
		Where("userid = ?", userID).
		Where("parent_id = ?", 0)

	// 查询当前用户下所有未完成目标、已完成目标和已取消目标的数量
	err := db.Session(&core.Session).
		Select("SUM(CASE WHEN completed = 1 OR canceled = 1 THEN 1 ELSE 0 END) as completed,COUNT(*)-SUM(CASE WHEN completed = 1 OR canceled = 1  THEN 1 ELSE 0 END) AS uncompleted").
		Scan(&statisticsAll).Error

	if err != nil {
		return nil, err
	}

	return &statisticsAll, nil
}

// 获取个人整体评分与完成度

func (s *okrService) GetOkrOverall(userId int) (*interfaces.OkrOverall, error) {
	resp := &interfaces.OkrOverall{}

	db := core.DB.Model(&model.Okr{}).
		Where("userid = ? and parent_id = ?", userId, 0)

	err := db.Session(&core.Session).
		Select("count(*) as Total,SUM(score) as ScoreSum,SUM(progress) as CompletionSum").
		Find(&resp).Error

	if err != nil {
		return nil, err
	}

	return resp, nil
}