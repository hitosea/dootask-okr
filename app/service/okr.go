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
	db := core.DB.Where("parent_id = 0").Order("ascription asc, create_at desc")

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

	// 已完成未评分筛选 Completed 0-未完成 1-已完成
	completed := param.Completed
	if completed != 0 {
		if completed == 1 {
			db = db.Where("progress >= 100 and score = 0")
		} else {
			db = db.Where("progress < 100")
		}
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
func (s *okrService) FollowObjective(userid, objectiveId int) (interface{}, error) {
	// 只要顶级目标才能被关注
	var obj model.Okr
	if err := core.DB.Where("id = ? and parent_id = 0", objectiveId).First(&obj).Error; err != nil {
		return nil, err
	}

	// 检查是否已关注
	var f model.OkrFollow
	if err := core.DB.Where("userid = ? and okr_id = ?", userid, objectiveId).First(&f).Error; err != nil {
		if !errors.Is(err, core.ErrRecordNotFound) {
			return nil, err
		}
	}

	// 如果已关注且需要取消关注，则删除关注记录
	if f.Id != 0 {
		if err := core.DB.Delete(&f).Error; err != nil {
			return nil, err
		}
		return nil, nil
	}

	// 如果未关注且需要关注，则创建关注记录
	if f.Id == 0 {
		if err := core.DB.Create(&model.OkrFollow{
			Userid: userid,
			OkrId:  objectiveId,
		}).Error; err != nil {
			return nil, err
		}
		return map[string]int{
			"userid": userid,
			"okrId":  objectiveId,
		}, nil
	}

	return nil, nil
}

// 更新进度和进度状态
func (s *okrService) UpdateProgressAndStatus(param interfaces.OkrUpdateProgressReq) error {
	// 检查目标是否存在
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return fmt.Errorf("目标不存在：%w", err)
	}

	// 如果传值更新进度有值，则更新进度
	if param.Progress != 0 {
		obj.Progress = param.Progress
	}

	// 如果传值更新状态有值，则更新状态
	if param.Status != 0 {
		obj.ProgressStatus = param.Status
	}

	// 更新目标
	if err := core.DB.Save(obj).Error; err != nil {
		return fmt.Errorf("更新目标失败：%w", err)
	}

	// 检查所在目标的 KR 是否全部完成
	objWithKrs, err := s.GetObjectiveByIdWithKeyResults(obj.ParentId)
	if err != nil {
		return fmt.Errorf("获取所在目标的 KR 失败：%w", err)
	}
	krs := objWithKrs.KeyResults

	allCompleted := true
	var sumProgress int
	for _, kr := range krs {
		// 进度全部完成 100%
		if kr.Progress < 100 {
			allCompleted = false
			break
		}
		sumProgress += kr.Progress
	}

	// 未完成，则更新总目标进度值，kr 进度值相加/kr 数量
	progress := 0
	if len(krs) > 0 {
		progress = sumProgress / len(krs)
	}

	// 更新总目标的状态是否完成
	if allCompleted {
		if err := core.DB.Model(&model.Okr{}).Where("id = ?", obj.ParentId).Updates(map[string]interface{}{
			"Completed": 1,
			"Progress":  progress,
		}).Error; err != nil {
			return fmt.Errorf("更新总目标失败：%w", err)
		}
	} else {
		if err := core.DB.Model(&model.Okr{}).Where("id = ?", obj.ParentId).Update("progress", progress).Error; err != nil {
			return err
		}
	}

	return nil
}

// 更新评分
func (s *okrService) UpdateScore(user *interfaces.UserInfoResp, param interfaces.OkrScoreReq) error {
	// 检查目标是否存在
	kr, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return errors.New("目标不存在")
	}

	// 检查进度是否为100%
	if kr.Progress < 100 {
		return errors.New("进度不足100%")
	}

	// 检查用户是否为目标负责人或上级 false-负责人 true-上级
	superior, err := s.IsObjectiveManager(user)
	if err != nil {
		return err
	}
	if !superior {
		// 检查用户是否为目标负责人
		if kr.Userid != user.Userid {
			return errors.New("暂无权限评分")
		}
		// 负责人评分
		if err := core.DB.Model(&model.Okr{}).Where("id = ?", param.Id).Update("score", param.Score).Error; err != nil {
			return err
		}
	} else {
		// 需要负责人评分才可以上级评分
		if kr.Score == 0 {
			return errors.New("负责人未评分")
		}
		// 上级评分
		if err := core.DB.Model(&model.Okr{}).Where("id = ?", param.Id).Update("superior_score", param.Score).Error; err != nil {
			return err
		}
	}

	return nil
}

// 检查用户是否为目标上级
func (s *okrService) IsObjectiveManager(user *interfaces.UserInfoResp) (bool, error) {
	// todo 检查用户是否为上级
	return false, nil
}

// 结束/重启目标
func (s *okrService) FinishObjective(param interfaces.OkrFinishReq) error {
	// 检查目标是否存在
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return errors.New("目标不存在")
	}

	obj.Finished = param.Finished
	if err := core.DB.Save(obj).Error; err != nil {
		return err
	}

	return nil
}

// 更新参与人
func (s *okrService) UpdateParticipant(param interfaces.OkrParticipantUpdateReq) error {
	// 检查目标是否存在
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return errors.New("目标不存在")
	}

	obj.Participant = param.Participant
	if err := core.DB.Save(obj).Error; err != nil {
		return err
	}

	return nil
}

// 更新信心指数
func (s *okrService) UpdateConfidence(param interfaces.OkrConfidenceUpdateReq) error {
	// 检查目标是否存在
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return errors.New("目标不存在")
	}

	obj.Confidence = param.Confidence
	if err := core.DB.Save(obj).Error; err != nil {
		return err
	}

	return nil
}

// 创建复盘
func (s *okrService) CreateReplay(userid int, param interfaces.OkrReplayCreateReq) error {
	tx := core.DB.Begin()

	_, err := s.GetObjectiveById(param.Id)
	if err != nil {
		tx.Rollback()
		return errors.New("目标不存在")
	}

	replay := model.OkrReplay{
		OkrId:   param.Id,
		Userid:  userid,
		Value:   param.Value,
		Problem: param.Problem,
	}
	if err := tx.Create(&replay).Error; err != nil {
		tx.Rollback()
		return err
	}

	// todo 添加到复盘历史表
	for _, kr := range param.Comments {
		if err := tx.Model(&model.Okr{}).Where("id = ?", kr.Id).Update("replay_comment", kr.Comment).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

// 复盘详情
func (s *okrService) GetReplayDetail(id int) (*model.OkrReplay, error) {
	var replay model.OkrReplay
	if err := core.DB.Where("id = ?", id).First(&replay).Error; err != nil {
		return nil, err
	}
	return &replay, nil
}

// 评分规定的分数列表
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
