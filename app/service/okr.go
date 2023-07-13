package service

import (
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/app/utils/common"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

var OkrService = okrService{}

type okrService struct{}

// 创建目标
func (s *okrService) Create(user *interfaces.UserInfoResp, param interfaces.OkrCreateReq) (*model.Okr, error) {
	tx := core.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	startAt, err := common.ParseTime(param.StartAt)
	if err != nil {
		return nil, err
	}

	endAt, err := common.ParseTime(param.EndAt)
	if err != nil {
		return nil, err
	}

	obj := &model.Okr{
		Userid:       user.Userid,
		DepartmentId: user.Department[0],
		Title:        param.Title,
		Type:         param.Type,
		Priority:     param.Priority,
		Ascription:   param.Ascription,
		VisibleRange: param.VisibleRange,
		ProjectId:    param.ProjectId,
		StartAt:      startAt,
		EndAt:        endAt,
	}
	if err := tx.Create(obj).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if param.AlignObjective != "" {
		if err := s.updateAlignment(tx, obj.Id, param.AlignObjective); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	for _, kr := range param.KeyResults {
		keyResult, err := s.createKeyResult(tx, user, obj.Id, kr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		obj.KeyResults = append(obj.KeyResults, keyResult)
	}

	return obj, tx.Commit().Error
}

// 更新目标
func (s *okrService) Update(param interfaces.OkrUpdateReq) (*model.Okr, error) {
	tx := core.DB.Begin()

	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("目标不存在")
	}

	startAt, err := common.ParseTime(param.StartAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	endAt, err := common.ParseTime(param.EndAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	obj.Title = param.Title
	obj.Type = param.Type
	obj.Priority = param.Priority
	obj.VisibleRange = param.VisibleRange
	obj.ProjectId = param.ProjectId
	obj.StartAt = startAt
	obj.EndAt = endAt

	if err := tx.Save(obj).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if param.AlignObjective != "" {
		if err := s.updateAlignment(tx, obj.Id, param.AlignObjective); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	for _, kr := range param.KeyResults {
		keyResult, err := s.updateKeyResult(tx, obj.Id, kr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		obj.KeyResults = append(obj.KeyResults, keyResult)
	}

	return obj, tx.Commit().Error
}

// 创建关键结果
func (s *okrService) createKeyResult(tx *gorm.DB, user *interfaces.UserInfoResp, parentId int, kr *interfaces.OkrKeyResultCreateReq) (*model.Okr, error) {
	startAt, err := common.ParseTime(kr.StartAt)
	if err != nil {
		return nil, err
	}

	endAt, err := common.ParseTime(kr.EndAt)
	if err != nil {
		return nil, err
	}

	keyResult := &model.Okr{
		Userid:       user.Userid,
		DepartmentId: user.Department[0],
		ParentId:     parentId,
		Participant:  kr.Participant,
		Title:        kr.Title,
		Confidence:   kr.Confidence,
		StartAt:      startAt,
		EndAt:        endAt,
	}
	if err := tx.Create(keyResult).Error; err != nil {
		return nil, err
	}

	return keyResult, nil
}

// 更新关键结果
func (s *okrService) updateKeyResult(tx *gorm.DB, parentId int, kr *interfaces.OkrKeyResultUpdateReq) (*model.Okr, error) {
	startAt, err := common.ParseTime(kr.StartAt)
	if err != nil {
		return nil, err
	}

	endAt, err := common.ParseTime(kr.EndAt)
	if err != nil {
		return nil, err
	}

	keyResult, err := s.GetObjectiveById(kr.Id)
	if err != nil {
		return nil, errors.New("目标不存在")
	}

	keyResult.ProjectId = parentId
	keyResult.Participant = kr.Participant
	keyResult.Title = kr.Title
	keyResult.Confidence = kr.Confidence
	keyResult.StartAt = startAt
	keyResult.EndAt = endAt

	if err := tx.Save(keyResult).Error; err != nil {
		return nil, err
	}

	return keyResult, nil
}

// 更新对齐目标
func (s *okrService) updateAlignment(tx *gorm.DB, objId int, alignObjective string) error {
	if err := tx.Where("okr_id = ?", objId).Delete(&model.OkrAlign{}).Error; err != nil {
		return err
	}

	alignmentIDs := common.ExplodeInt(",", alignObjective, true)
	for _, alignmentID := range alignmentIDs {
		alignmentObj := &model.OkrAlign{
			OkrId:      objId,
			AlignOkrId: alignmentID,
		}
		if err := tx.Create(alignmentObj).Error; err != nil {
			return err
		}
	}

	return nil
}

// 根据id获取目标
func (s *okrService) GetObjectiveById(id int) (*model.Okr, error) {
	var obj model.Okr
	if err := core.DB.Where("id = ?", id).First(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

// 根据id获取目标及其关键结果
func (s *okrService) GetObjectiveByIdWithKeyResults(id int) (*model.Okr, error) {
	obj, err := s.GetObjectiveById(id)
	if err != nil {
		return nil, err
	}

	var krs []*model.Okr
	if err := core.DB.Where("parent_id = ?", obj.Id).Order("create_at desc").Find(&krs).Error; err != nil {
		return nil, err
	}
	obj.KeyResults = krs

	return obj, nil
}

// 获取所有目标及其关键结果
func (s *okrService) GetAllObjectivesWithKeyResults() ([]*model.Okr, error) {
	var objs []*model.Okr
	if err := core.DB.Order("create_at desc").Find(&objs).Error; err != nil {
		return nil, err
	}

	for _, obj := range objs {
		var krs []*model.Okr
		if err := core.DB.Where("parent_id = ?", obj.Id).Order("create_at desc").Find(&krs).Error; err != nil {
			return nil, err
		}
		obj.KeyResults = krs
	}

	return objs, nil
}

// 获取我的OkR列表
// 1.可通过目标（O）搜索 2.仅显示我发起的OKR（个人仅能看到个人的）3.按创建时间倒序显示
func (s *okrService) GetMyList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrListResp, error) {
	var objs []*model.Okr
	db := core.DB.Where("userid = ?", user.Userid).Where("parent_id = 0").Order("create_at desc")

	if objective != "" {
		db = db.Where("title like ?", "%"+objective+"%")
	}

	if err := db.Find(&objs).Error; err != nil {
		return nil, err
	}

	var okrs []*interfaces.OkrListResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrListResp{
			Okr: obj,
		})
	}

	okrs, err := s.GetObjectivesWithDetails(okrs, user)
	if err != nil {
		return nil, err
	}

	return okrs, nil
}

// 获取okr列表额外信息
func (s *okrService) GetObjectivesWithDetails(objs []*interfaces.OkrListResp, user *interfaces.UserInfoResp) ([]*interfaces.OkrListResp, error) {
	for _, obj := range objs {
		var krs []*model.Okr
		if err := core.DB.Where("parent_id = ?", obj.Id).Order("create_at desc").Find(&krs).Error; err != nil {
			return nil, err
		}
		s.GetObjectiveExt(obj, krs, user)
	}
	return objs, nil
}

// 额外信息
func (s *okrService) GetObjectiveExt(obj *interfaces.OkrListResp, krs []*model.Okr, user *interfaces.UserInfoResp) *interfaces.OkrListResp {
	obj.KeyResults = krs                                                       // 关键结果
	obj.IsFollow = s.isFollow(user.Userid, obj.Id)                             // 是否被关注
	obj.KrCount = len(krs)                                                     // kr总数量
	obj.KrFinishCount = s.getFinishCount(krs)                                  // kr完成数量
	obj.AlignCount = s.getAlignCount(obj.Id)                                   // 对齐目标数量
	obj.Alias = s.getAlias(obj.Ascription, user.DepartmentName, user.Nickname) // 目标别名
	return obj
}

// 是否被关注
func (s *okrService) isFollow(userid, objectiveId int) bool {
	var count int64
	if err := core.DB.Model(&model.OkrFollow{}).Where("userid = ? and okr_id = ?", userid, objectiveId).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// kr完成数量
func (s *okrService) getFinishCount(krs []*model.Okr) int {
	var count int
	for _, kr := range krs {
		if kr.Progress >= 100 {
			count++
		}
	}
	return count
}

// 对齐目标数量
func (s *okrService) getAlignCount(objId int) int {
	var count int64
	if err := core.DB.Model(&model.OkrAlign{}).Where("okr_id = ?", objId).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}

// 目标别名
func (s *okrService) getAlias(ascription int, departmentName, nickname string) string {
	// Ascription 1-部门 2-个人
	if ascription == 1 {
		return "部门-" + departmentName
	}
	return "个人-" + nickname
}

// 获取参与的OKR列表
// 1.可通过目标（O）搜索 2.被其他OKR选为协助人的OKR（可能是其他部门/其他人的OKR）3.按创建时间倒序显示
func (s *okrService) GetParticipantList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrListResp, error) {
	objectiveIds, err := getObjectiveIdsByParticipant(user.Userid)
	if err != nil {
		return nil, err
	}

	objs, err := getObjectives(objectiveIds, objective)
	if err != nil {
		return nil, err
	}

	// 列表响应
	var okrs []*interfaces.OkrListResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrListResp{
			Okr: obj,
		})
	}

	for _, obj := range okrs {
		krs, err := getKeyResults(obj.Id, user.Userid)
		if err != nil {
			return nil, err
		}
		obj.KeyResults = krs
		s.GetObjectiveExt(obj, krs, user)
	}

	return okrs, nil
}

// 获取与特定用户id相关的objective_ids
func getObjectiveIdsByParticipant(userid int) ([]int, error) {
	var objectiveIds []int
	if err := core.DB.Model(&model.Okr{}).
		Where("participant like ?", "%,"+strconv.Itoa(userid)+",%").
		Pluck("DISTINCT parent_id", &objectiveIds).Error; err != nil {
		return nil, err
	}

	return objectiveIds, nil
}

// 获取满足目标搜索条件的objs
func getObjectives(objectiveIds []int, objective string) ([]*model.Okr, error) {
	db := core.DB.Model(&model.Okr{}).
		Where("id in (?)", objectiveIds).
		Where("parent_id = 0").
		Order("create_at desc")

	if objective != "" {
		db = db.Where("title like ?", "%"+objective+"%")
	}

	var objs []*model.Okr
	if err := db.Find(&objs).Error; err != nil {
		return nil, err
	}

	return objs, nil
}

// 获取特定objective_id和用户id相关的key_results
func getKeyResults(objectiveId, userid int) ([]*model.Okr, error) {
	var krs []*model.Okr
	if err := core.DB.Model(&model.Okr{}).
		Where("parent_id = ? and participant like ?", objectiveId, "%,"+strconv.Itoa(userid)+",%").
		Order("create_at desc").
		Find(&krs).Error; err != nil {
		return nil, err
	}

	return krs, nil
}

// 部门的OKR列表
// 1.高级搜索 2.仅包含部门/及部门其他人员的OKR（通过可见范围控制是否看到部门其他同级人员的）3.按创建时间倒序显示 4.部门的OKR置顶（多个的时候多个都置顶，按创建时间倒序）
func (s *okrService) GetDepartmentList(user *interfaces.UserInfoResp, param interfaces.OkrDepartmentListReq) ([]*interfaces.OkrListResp, error) {

	var objs []*model.Okr
	db := core.DB.Order("ascription asc, create_at desc")

	// 超级管理员可以通过部门筛选
	departmentId := param.DepartmentId
	if departmentId != 0 {
		db = db.Where("department_id = ?", departmentId)
	}

	// 部门负责人可以通过人员筛选
	userid := param.Userid
	if userid != 0 {
		db = db.Where("userid = ?", userid)
	}

	// 目标筛选
	objective := param.Objective
	if objective != "" {
		db = db.Where("title like ?", "%"+objective+"%")
	}

	// 时间筛选
	startAtStr := param.StartAt
	endAtStr := param.EndAt
	if startAtStr != "" && endAtStr != "" {
		startAt, err := common.ParseTime(startAtStr)
		if err != nil {
			return nil, err
		}
		endAt, err := common.ParseTime(endAtStr)
		if err != nil {
			return nil, err
		}
		db = db.Where("start_at >= ? and end_at <= ?", startAt, endAt)
	}

	// 类型筛选
	typeInt := param.Type
	if typeInt != 0 {
		db = db.Where("type = ?", typeInt)
	}

	if err := db.Find(&objs).Error; err != nil {
		return nil, err
	}

	var okrs []*interfaces.OkrListResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrListResp{
			Okr: obj,
		})
	}

	okrs, err := s.GetObjectivesWithDetails(okrs, user)
	if err != nil {
		return nil, err
	}

	return okrs, nil
}

// 关注的OKR列表
// 1.可通过目标（O）搜索 2.按关注时间倒序
func (s *okrService) GetFollowList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrListResp, error) {
	var objs []*model.Okr

	OkrTable := core.DBTableName(&model.Okr{})
	OkrFollowTable := core.DBTableName(&model.OkrFollow{})

	db := core.DB.Table(OkrFollowTable+" AS f").
		Select("o.*").
		Joins(fmt.Sprintf("LEFT JOIN %s AS o ON f.okr_id = o.id", OkrTable)).
		Where("f.userid = ?", user.Userid).
		Order("f.create_at desc")

	if objective != "" {
		db = db.Where("o.title like ?", "%"+objective+"%")
	}

	if err := db.Find(&objs).Error; err != nil {
		return nil, err
	}

	var okrs []*interfaces.OkrListResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrListResp{
			Okr: obj,
		})
	}

	okrs, err := s.GetObjectivesWithDetails(okrs, user)
	if err != nil {
		return nil, err
	}

	return okrs, nil
}

// 获取复盘列表
func (s *okrService) GetReplayList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrListResp, error) {
	var objs []*model.Okr

	OkrTable := core.DBTableName(&model.Okr{})
	OkrReplayTable := core.DBTableName(&model.OkrReplay{})

	db := core.DB.Table(OkrReplayTable+" AS r").
		Select("o.*").
		Joins(fmt.Sprintf("LEFT JOIN %s AS o ON r.okr_id = o.id", OkrTable)).
		Where("r.userid = ?", user.Userid).
		Order("r.create_at desc")

	if objective != "" {
		db = db.Where("o.title like ?", "%"+objective+"%")
	}

	if err := db.Find(&objs).Error; err != nil {
		return nil, err
	}

	var okrs []*interfaces.OkrListResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrListResp{
			Okr: obj,
		})
	}

	okrs, err := s.GetObjectivesWithDetails(okrs, user)
	if err != nil {
		return nil, err
	}

	return okrs, nil
}

// 关注或取消关注目标
func (s *okrService) FollowObjective(userid, objectiveId int, follow bool) error {
	// 检查目标是否存在
	if _, err := s.GetObjectiveById(objectiveId); err != nil {
		return errors.New("目标不存在")
	}

	// 检查是否已关注
	var f model.OkrFollow
	if err := core.DB.Where("userid = ? and okr_id = ?", userid, objectiveId).First(&f).Error; err != nil {
		if !errors.Is(err, core.ErrRecordNotFound) {
			return err
		}
	}

	// 如果已关注且需要取消关注，则删除关注记录
	if !follow && f.Id != 0 {
		if err := core.DB.Delete(&f).Error; err != nil {
			return err
		}
		return nil
	}

	// 如果未关注且需要关注，则创建关注记录
	if follow && f.Id == 0 {
		if err := core.DB.Create(&model.OkrFollow{
			Userid: userid,
			OkrId:  objectiveId,
		}).Error; err != nil {
			return err
		}
		return nil
	}

	return nil
}

// 更新进度和进度状态
func (s *okrService) UpdateProgressAndStatus(id, progress, status int) error {
	// 检查目标是否存在
	if _, err := s.GetObjectiveById(id); err != nil {
		return errors.New("目标不存在")
	}

	// 更新进度和进度状态
	if err := core.DB.Model(&model.Okr{}).Where("id = ?", id).Updates(map[string]interface{}{
		"progress":        progress,
		"progress_status": getStatus(progress, status),
	}).Error; err != nil {
		return err
	}

	return nil
}

// 根据进度和状态获取新的状态
func getStatus(progress, status int) int {
	switch status {
	case model.OkrKeyResultStatusNotStart, model.OkrKeyResultStatusHasProblem, model.OkrKeyResultStatusEnd:
		// 当状态为 0-未开始 3-有难度 4-已结束 时，不更新状态
		return status
	default:
		if progress >= 100 {
			// 当进度为 100 时，状态为已完成
			return model.OkrKeyResultStatusFinished
		} else {
			// 从进度不足 100 时，状态跟随为进行中
			return model.OkrKeyResultStatusInProgress
		}
	}
}

// 更新评分
func (s *okrService) UpdateScore(krId int, score float64, userid int) error {
	// 检查目标是否存在
	kr, err := s.GetObjectiveById(krId)
	if err != nil {
		return errors.New("目标不存在")
	}

	// 检查进度是否为100%
	if kr.Progress < 100 {
		return errors.New("进度不足100%")
	}

	// 检查用户是否为目标负责人或上级 false-负责人 true-上级
	superior, err := s.IsObjectiveManagerOrSuperior(kr.ParentId, userid)
	if err != nil {
		return err
	}
	if !superior {
		// 负责人评分
		if err := core.DB.Model(&model.Okr{}).Where("id = ?", krId).Update("score", score).Error; err != nil {
			return err
		}
	} else {
		// 需要负责人评分才可以上级评分
		if kr.Score == 0 {
			return errors.New("负责人未评分")
		}
		// 上级评分
		if err := core.DB.Model(&model.Okr{}).Where("id = ?", krId).Update("superior_score", score).Error; err != nil {
			return err
		}
	}

	return nil
}

// 检查用户是否为目标负责人或上级
func (s *okrService) IsObjectiveManagerOrSuperior(objectiveId, userid int) (bool, error) {
	// 获取目标负责人
	objective, err := s.GetObjectiveById(objectiveId)
	if err != nil {
		return false, err
	}

	// 检查用户是否为目标负责人
	if objective.Userid == userid {
		return false, nil
	}

	// // todo 获取用户信息
	// user, err := s.GetUserById(userid)
	// if err != nil {
	// 	return false, err
	// }

	// // todo 检查用户是否为上级
	// if user.IsSuperiorOf(objective.Userid) {
	// 	return true, nil
	// }

	return false, nil
}

// 创建复盘
func (s *okrService) CreateReplay(id int, comment string, value string, problem string) error {
	// 检查目标是否存在
	if _, err := s.GetObjectiveById(id); err != nil {
		return errors.New("目标不存在")
	}

	// 创建复盘
	if err := core.DB.Create(&model.OkrReplay{
		OkrId:   id,
		Comment: comment,
		Value:   value,
		Problem: problem,
	}).Error; err != nil {
		return err
	}

	return nil
}

// 获取复盘
func (s *okrService) GetReplayById(id int) (*model.OkrReplay, error) {
	var replay model.OkrReplay
	if err := core.DB.Where("id = ?", id).First(&replay).Error; err != nil {
		return nil, err
	}
	return &replay, nil
}

// 返回评分规定的分数列表
func (s *okrService) GetScoreList() []string {
	return []string{
		"0分 未达成目标，态度问题",
		"1分 未达成目标",
		"2分 目标达成效果较差，未及时调整",
		"3分 目标达成效果不佳",
		"4分 目标勉强达成，还需努力",
		"5分 目标基本达成",
		"6分 目标达成，效果一般",
		"7分 目标达成，效果良好",
		"8分 100%达成目标",
		"9分 110%达成目标",
		"10分 超出预期",
	}
}

// todo 取消对齐目标
