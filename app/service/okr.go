package service

import (
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/app/utils/common"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"gorm.io/gorm"
)

var OkrService = okrService{}

type okrService struct{}

// 创建目标
func (s *okrService) Create(user *interfaces.UserInfoResp, param interfaces.OkrCreateReq) (*model.Okr, error) {
	// 至少有一条关键结果
	if len(param.KeyResults) == 0 {
		return nil, errors.New("至少有一条关键结果")
	}

	// 时间格式化
	startAt, err := common.ParseTime(param.StartAt)
	if err != nil {
		return nil, err
	}

	endAt, err := common.ParseTime(param.EndAt)
	if err != nil {
		return nil, err
	}

	// 创建目标
	obj := &model.Okr{
		Userid:       user.Userid,
		DepartmentId: common.ArrayImplode(user.Department),
		Title:        param.Title,
		Type:         param.Type,
		Priority:     param.Priority,
		Ascription:   param.Ascription,
		VisibleRange: param.VisibleRange,
		ProjectId:    param.ProjectId,
		StartAt:      startAt,
		EndAt:        endAt,
	}

	var participantIds []int
	participantIds = append(participantIds, user.Userid) // 通知发起/所有KR参与人
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(obj).Error; err != nil {
			return err
		}

		// 创建对话
		dialogId, err := DootaskService.DialogOkrAdd(user.Token, obj)
		if err != nil {
			return err
		}
		obj.DialogId = dialogId
		if err := tx.Save(obj).Error; err != nil {
			return err
		}

		// 对齐目标
		if param.AlignObjective != "" {
			if err := s.updateAlignment(tx, obj.Id, param.AlignObjective); err != nil {
				return err
			}
		}

		// 关键结果
		for _, kr := range param.KeyResults {
			keyResult, err := s.createKeyResult(tx, user, obj.Id, kr)
			if err != nil {
				return err
			}
			obj.KeyResults = append(obj.KeyResults, keyResult)
			participantIds = append(participantIds, common.ExplodeInt(",", kr.Participant, true)...)
		}

		// 动态日志
		if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "add", "创建OKR"); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 创建O时（通知发起/所有KR参与人）
	participantIds = common.ArrayUniqueInt(participantIds)
	go DootaskService.DialogOkrPush(obj, user.Token, 1, participantIds)

	return obj, nil
}

// 更新目标
func (s *okrService) Update(user *interfaces.UserInfoResp, param interfaces.OkrUpdateReq) (*model.Okr, error) {
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return nil, errors.New("目标不存在")
	}

	startAt, err := common.ParseTime(param.StartAt)
	if err != nil {
		return nil, err
	}

	endAt, err := common.ParseTime(param.EndAt)
	if err != nil {
		return nil, err
	}

	// 至少有一条关键结果
	if len(param.KeyResults) == 0 {
		return nil, errors.New("至少有一条关键结果")
	}

	var participantIds []int // 所有参与人
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		for _, kr := range param.KeyResults {
			if kr.Id == 0 {
				// 新增kr
				var addKr interfaces.OkrKeyResultCreateReq
				addKr.Title = kr.Title
				addKr.Participant = kr.Participant
				addKr.Confidence = kr.Confidence
				addKr.StartAt = kr.StartAt
				addKr.EndAt = kr.EndAt
				keyResult, err := s.createKeyResult(tx, user, obj.Id, &addKr)
				if err != nil {
					return err
				}
				// 新增kr时，发送提示消息
				keyResult.ParentTitle = obj.Title
				go DootaskService.DialogOkrPush(keyResult, user.Token, 4, common.ExplodeInt(",", kr.Participant, true))
				obj.KeyResults = append(obj.KeyResults, keyResult)
			} else {
				// 是否删除 true-删除 false-不删除
				if kr.IsDelete {
					// 删除kr
					if err := tx.Where("id = ?", kr.Id).Delete(&model.Okr{}).Error; err != nil {
						return err
					}
					continue
				}
				// 更新kr
				keyResult, err := s.updateKeyResult(tx, kr, user, obj)
				if err != nil {
					return err
				}
				obj.KeyResults = append(obj.KeyResults, keyResult)
			}
			participantIds = append(participantIds, common.ExplodeInt(",", kr.Participant, true)...)
		}

		// O目标变动时，发送提示消息
		participantIds = common.ArrayUniqueInt(participantIds)
		if obj.Title != param.Title {
			go DootaskService.DialogOkrPush(obj, user.Token, 2, participantIds)
		}

		// O时间变动时，发送提示消息
		if obj.StartAt != startAt || obj.EndAt != endAt {
			go DootaskService.DialogOkrPush(obj, user.Token, 5, participantIds)
		}

		// 新增动态日志
		if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", "更新OKR目标"); err != nil {
			return err
		}

		obj.Title = param.Title
		obj.Type = param.Type
		obj.Priority = param.Priority
		obj.VisibleRange = param.VisibleRange
		obj.ProjectId = param.ProjectId
		obj.StartAt = startAt
		obj.EndAt = endAt

		if err := tx.Save(obj).Error; err != nil {
			return err
		}

		// 更新对齐目标
		if param.AlignObjective != "" {
			if err := s.updateAlignment(tx, obj.Id, param.AlignObjective); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return obj, nil
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
		DepartmentId: common.ArrayImplode(user.Department),
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
func (s *okrService) updateKeyResult(tx *gorm.DB, kr *interfaces.OkrKeyResultUpdateReq, user *interfaces.UserInfoResp, obj *model.Okr) (*model.Okr, error) {
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

	// 父级目标标题
	keyResult.ParentTitle = obj.Title

	oldParticipant := common.ExplodeInt(",", keyResult.Participant, true)
	newParticipant := common.ExplodeInt(",", kr.Participant, true)
	diffParticipant := common.ArrayDifferenceProcessing(newParticipant, oldParticipant)

	// KR变动时，发送提示消息
	if keyResult.Title != kr.Title {
		go DootaskService.DialogOkrPush(keyResult, user.Token, 3, newParticipant)
	}

	// KR时间变动时，发送提示消息
	if keyResult.StartAt != startAt || keyResult.EndAt != endAt {
		go DootaskService.DialogOkrPush(keyResult, user.Token, 6, newParticipant)
	}

	// 为KR添加新参与人时，发送提示消息
	if len(diffParticipant) > 0 {
		go DootaskService.DialogOkrPush(keyResult, user.Token, 4, diffParticipant)
	}

	keyResult.ProjectId = obj.Id
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
func (s *okrService) GetMyList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrResp, error) {
	var objs []*model.Okr
	db := core.DB.Where("userid = ?", user.Userid).Where("parent_id = 0").Order("create_at desc")

	if objective != "" {
		db = db.Where("title like ?", "%"+objective+"%")
	}

	if err := db.Find(&objs).Error; err != nil {
		return nil, err
	}

	var okrs []*interfaces.OkrResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrResp{
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
func (s *okrService) GetObjectivesWithDetails(objs []*interfaces.OkrResp, user *interfaces.UserInfoResp) ([]*interfaces.OkrResp, error) {
	for _, obj := range objs {
		var krs []*model.Okr
		if err := core.DB.Where("parent_id = ?", obj.Id).Order("create_at desc").Find(&krs).Error; err != nil {
			return nil, err
		}

		// krs KR总评分
		for _, kr := range krs {
			krScore := s.getKrScore(kr)
			kr.KrScore = krScore
		}

		s.GetObjectiveExt(obj, krs, user)
	}
	return objs, nil
}

// 额外信息
func (s *okrService) GetObjectiveExt(obj *interfaces.OkrResp, krs []*model.Okr, user *interfaces.UserInfoResp) *interfaces.OkrResp {
	obj.KeyResults = krs                                       // 关键结果
	obj.IsFollow = s.isFollow(user.Userid, obj.Id)             // 是否被关注
	obj.KrCount = len(krs)                                     // kr总数量
	obj.KrFinishCount = s.getFinishCount(krs)                  // kr完成数量
	obj.AlignCount = s.getAlignCount(obj.Id)                   // 对齐目标数量
	obj.Alias = s.getOwningAlias(obj.Userid, obj.DepartmentId) // 目标所属名称
	return obj
}

// KR总评分 KR评分=【自评*40%+上级评分*60%】
func (s *okrService) getKrScore(obj *model.Okr) float64 {
	return math.Round((obj.Score*0.4+obj.SuperiorScore*0.6)*10) / 10
}

// O总评分 O的评分=所有KR总分之和/KR数量
func (s *okrService) getObjectiveScore(obj *model.Okr) float64 {
	if len(obj.KeyResults) == 0 {
		return 0
	}
	var score float64
	for _, kr := range obj.KeyResults {
		krScore := s.getKrScore(kr)
		score += krScore
	}
	return math.Round((score/float64(len(obj.KeyResults)))*10) / 10
}

// 所有KR更新评分是否完成，完成则更新O评分，否则不更新
func (s *okrService) UpdateObjectiveScoreTx(tx *gorm.DB, obj *model.Okr) error {
	// 检查所有KR评分是否完成
	for _, kr := range obj.KeyResults {
		if kr.Score == 0 || kr.SuperiorScore == 0 {
			return nil
		}
	}

	// 计算O评分
	score := s.getObjectiveScore(obj)
	if math.IsNaN(score) {
		return errors.New("invalid KR score")
	}

	// 更新O评分
	obj.Score = score
	if err := tx.Save(obj).Error; err != nil {
		return err
	}

	return nil
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

// 获取参与的OKR列表
// 1.可通过目标（O）搜索 2.被其他OKR选为协助人的OKR（可能是其他部门/其他人的OKR）3.按创建时间倒序显示
func (s *okrService) GetParticipantList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrResp, error) {
	objectiveIds, err := getObjectiveIdsByParticipant(user.Userid)
	if err != nil {
		return nil, err
	}

	objs, err := getParticipantObjectives(objectiveIds, objective)
	if err != nil {
		return nil, err
	}

	// 列表响应
	var okrs []*interfaces.OkrResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrResp{
			Okr: obj,
		})
	}

	for _, obj := range okrs {
		krs, err := getParticipantKeyResults(obj.Id, user.Userid)
		if err != nil {
			return nil, err
		}
		obj.KeyResults = krs
		s.GetObjectiveExt(obj, krs, user)
	}

	return okrs, nil
}

// 获取对齐目标列表
// 1. 显示所在部门（即我可见的）+　我参与的
func (s *okrService) GetAlignList(user *interfaces.UserInfoResp, keywords string) ([]*model.Okr, error) {
	// 获取我所在部门的OKR
	departmentIds := common.ExplodeInt(",", user.Department, true)
	var departmentObjectiveIds []int
	var err error
	for _, departmentId := range departmentIds {
		departmentObjectiveIds, err = getObjectiveIdsByDepartment(departmentId)
		if err != nil {
			return nil, err
		}
	}

	depObjs := []*model.Okr{}
	err = core.DB.Model(&model.Okr{}).Where("id in (?)", departmentObjectiveIds).Where("parent_id = 0").Order("create_at desc").Find(&depObjs).Error
	if err != nil {
		return nil, err
	}

	for _, obj := range depObjs {
		krs := []*model.Okr{}
		if err := core.DB.Where("parent_id = ?", obj.Id).Order("create_at desc").Find(&krs).Error; err != nil {
			return nil, err
		}
		obj.KeyResults = krs
	}

	// 我参与的OKR
	participantObjectiveIds, err := getObjectiveIdsByParticipant(user.Userid)
	if err != nil {
		return nil, err
	}

	// 差集处理
	objectiveIds := common.ArrayDifferenceProcessing(participantObjectiveIds, departmentObjectiveIds)
	if len(objectiveIds) > 0 {
		objs, err := getParticipantObjectives(objectiveIds, "")
		if err != nil {
			return nil, err
		}

		for _, obj := range objs {
			krs, err := getParticipantKeyResults(obj.Id, user.Userid)
			if err != nil {
				return nil, err
			}
			obj.KeyResults = krs
		}

		// 合并目标和关键结果
		depObjs = append(depObjs, objs...)
	}

	// 过滤标题中包含关键字的目标和关键结果
	filteredObjs := []*model.Okr{}
	for _, obj := range depObjs {
		if strings.Contains(obj.Title, keywords) {
			filteredObjs = append(filteredObjs, obj)
		}
		for _, kr := range obj.KeyResults {
			if strings.Contains(kr.Title, keywords) {
				filteredObjs = append(filteredObjs, obj)
			}
		}
	}

	// 按照创建时间排序
	sort.Slice(filteredObjs, func(i, j int) bool {
		return filteredObjs[i].CreateAt.After(filteredObjs[j].CreateAt)
	})

	return filteredObjs, nil
}

// 获取与特定部门id相关的objective_ids
func getObjectiveIdsByDepartment(departmentId int) ([]int, error) {
	var objectiveIds []int
	if err := core.DB.Model(&model.Okr{}).
		Where("parent_id = 0").
		Where("FIND_IN_SET(?, department_id) > 0", departmentId).
		Pluck("id", &objectiveIds).Error; err != nil {
		return nil, err
	}

	return objectiveIds, nil
}

// 获取与特定用户id相关的objective_ids
func getObjectiveIdsByParticipant(userid int) ([]int, error) {
	var objectiveIds []int
	if err := core.DB.Model(&model.Okr{}).
		Where("FIND_IN_SET(?, participant) > 0", userid).
		Pluck("DISTINCT parent_id", &objectiveIds).Error; err != nil {
		return nil, err
	}

	return objectiveIds, nil
}

// 获取参与的目标搜索列表
func getParticipantObjectives(objectiveIds []int, objective string) ([]*model.Okr, error) {
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

// 获取参与的KR列表
func getParticipantKeyResults(objectiveId, userid int) ([]*model.Okr, error) {
	var krs []*model.Okr
	if err := core.DB.Model(&model.Okr{}).
		Where("FIND_IN_SET(?, participant) > 0", userid).
		Where("parent_id = ?", objectiveId).
		Order("create_at desc").
		Find(&krs).Error; err != nil {
		return nil, err
	}

	return krs, nil
}

// 部门的OKR列表
// 1.高级搜索 2.仅包含部门/及部门其他人员的OKR（通过可见范围控制是否看到部门其他同级人员的）3.按创建时间倒序显示 4.部门的OKR置顶（多个的时候多个都置顶，按创建时间倒序）
func (s *okrService) GetDepartmentList(user *interfaces.UserInfoResp, param interfaces.OkrDepartmentListReq) ([]*interfaces.OkrResp, error) {
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

	var okrs []*interfaces.OkrResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrResp{
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
// 1.可通过目标（O）搜索 2.当前账号关注的OKR 3.按关注时间倒序
func (s *okrService) GetFollowList(user *interfaces.UserInfoResp, objective string) ([]*interfaces.OkrResp, error) {
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

	var okrs []*interfaces.OkrResp
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrResp{
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
func (s *okrService) GetReplayList(user *interfaces.UserInfoResp, objective string) ([]*model.OkrReplay, error) {
	var replays []*model.OkrReplay

	db := core.DB.Model(&model.OkrReplay{}).
		Where("userid = ?", user.Userid).
		Order("create_at DESC")

	if objective != "" {
		db = db.Where("okr_title LIKE ?", "%"+objective+"%")
	}

	if err := db.Preload("KrHistory").Find(&replays).Error; err != nil {
		return nil, err
	}

	// 获取所属别名
	for _, replay := range replays {
		replay.OkrAlias = s.getOwningAlias(replay.Userid, replay.OkrDepartmentId)
	}

	return replays, nil
}

// 获取复盘列表by目标id
func (s *okrService) GetReplayListByOkrId(okrId int) ([]*model.OkrReplay, error) {
	var replays []*model.OkrReplay

	db := core.DB.Model(&model.OkrReplay{}).
		Where("okr_id = ?", okrId).
		Order("create_at ASC")

	if err := db.Find(&replays).Error; err != nil {
		return nil, err
	}

	return replays, nil
}

// 获取OKR详情
func (s *okrService) GetOkrDetail(user *interfaces.UserInfoResp, okrId int) (*interfaces.OkrResp, error) {
	obj, err := s.GetObjectiveByIdWithKeyResults(okrId)
	if err != nil {
		return nil, err
	}

	// KR总评分
	for _, kr := range obj.KeyResults {
		krScore := s.getKrScore(kr)
		kr.KrScore = krScore
	}

	objResp := &interfaces.OkrResp{Okr: obj}

	s.GetObjectiveExt(objResp, obj.KeyResults, user)

	return objResp, nil
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
func (s *okrService) UpdateProgressAndStatus(user *interfaces.UserInfoResp, param interfaces.OkrUpdateProgressReq) error {
	// 检查目标是否存在
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return fmt.Errorf("目标不存在：%w", err)
	}

	// 开始事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		// 如果传值更新进度有值，则更新进度
		if param.Progress != 0 {
			logContent := fmt.Sprintf("更新OKR: %s 进度：%d=>%d", obj.Title, obj.Progress, param.Progress)
			if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", logContent); err != nil {
				return err
			}
			obj.Progress = param.Progress
		}

		// 如果传值更新状态有值，则更新状态
		if param.Status != 0 {
			logContent := fmt.Sprintf("更新OKR: %s 状态：%s=>%s", obj.Title, model.ProgressStatusMap[obj.ProgressStatus], model.ProgressStatusMap[param.Status])
			if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", logContent); err != nil {
				return err
			}
			obj.ProgressStatus = param.Status
		}

		// 更新目标
		if err := tx.Save(obj).Error; err != nil {
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
			if err := tx.Model(&model.Okr{}).Where("id = ?", obj.ParentId).Updates(map[string]interface{}{
				"Completed": 1,
				"Progress":  progress,
			}).Error; err != nil {
				return fmt.Errorf("更新总目标失败：%w", err)
			}
		} else {
			if err := tx.Model(&model.Okr{}).Where("id = ?", obj.ParentId).Update("progress", progress).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
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
	superior := s.IsObjectiveManager(kr, user)
	if !superior {
		// 检查用户是否为目标负责人
		if kr.Userid != user.Userid {
			return errors.New("暂无权限评分")
		}
		// 负责人评分
		err = core.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Update("score", param.Score).Error; err != nil {
				return err
			}

			// 新增动态日志
			logContent := fmt.Sprintf("责任人打分: %s", kr.Title)
			if err := s.InsertOkrLogTx(tx, kr.Id, user.Userid, "update", logContent); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	} else {
		// 需要负责人评分才可以上级评分
		if kr.Score == 0 {
			return errors.New("负责人未评分")
		}
		// 上级评分
		err = core.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Update("superior_score", param.Score).Error; err != nil {
				return err
			}

			// 更新O评分
			obj, err := s.GetObjectiveByIdWithKeyResults(kr.ParentId)
			if err != nil {
				return errors.New("目标不存在")
			}
			err = s.UpdateObjectiveScoreTx(tx, obj)
			if err != nil {
				return err
			}

			// 新增动态日志
			logContent := fmt.Sprintf("上级打分: %s", kr.Title)
			if err := s.InsertOkrLogTx(tx, kr.Id, user.Userid, "update", logContent); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// 检查用户是否为目标上级
func (s *okrService) IsObjectiveManager(kr *model.Okr, user *interfaces.UserInfoResp) bool {
	// 拿到目标负责人的部门
	depIds := common.ExplodeInt(",", kr.DepartmentId, true)
	if len(depIds) == 0 {
		return false
	}
	// 检查用户是否为部门负责人
	departmentTable := core.DBTableName(&model.UserDepartment{})
	userTable := core.DBTableName(&model.User{})
	var count int64

	if err := core.DB.Table(departmentTable+" AS d").
		Select("u.userid").
		Joins(fmt.Sprintf("LEFT JOIN %s AS u ON d.owner_userid = u.userid", userTable)).
		Where("d.id in (?)", depIds).
		Where("u.userid = ?", user.Userid).
		Count(&count).Error; err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

// 取消/重启目标
func (s *okrService) CancelObjective(param interfaces.OkrCanceledReq) error {
	// 检查目标是否存在
	obj, err := s.GetObjectiveById(param.Id)
	if err != nil {
		return errors.New("目标不存在")
	}

	obj.Canceled = param.Canceled
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

// 创建 OKR 复盘记录
func (s *okrService) CreateOkrReplay(userid int, req interfaces.OkrReplayCreateReq) error {
	// 获取 OKR 目标及其关键结果
	obj, err := s.GetObjectiveByIdWithKeyResults(req.OkrId)
	if err != nil {
		return errors.New("目标不存在")
	}

	// 检查关键结果是否已评分
	for _, kr := range obj.KeyResults {
		if kr.Score == 0 || kr.SuperiorScore == 0 {
			return errors.New("KR评分未完成")
		} else {
			// 计算关键结果总评分
			krScore := s.getKrScore(kr)
			kr.KrScore = krScore
		}
	}

	// 创建 OKR 复盘记录
	replay := model.OkrReplay{
		Userid:          userid,
		OkrId:           req.OkrId,
		OkrUserid:       obj.Userid,
		OkrDepartmentId: obj.DepartmentId,
		OkrTitle:        obj.Title,
		OkrProgress:     obj.Progress,
		OkrPriority:     obj.Priority,
		Review:          req.Review,
	}

	// 创建 KR 复盘历史记录
	var histories []*model.OkrReplayHistory
	commentsMap := make(map[int]interfaces.OkrReplayComment)
	for _, comment := range req.Comments {
		commentsMap[comment.OkrId] = *comment
	}
	for _, kr := range obj.KeyResults {
		comment, ok := commentsMap[kr.Id]
		if !ok {
			continue
		}
		history := &model.OkrReplayHistory{
			OkrId:      kr.Id,
			Title:      kr.Title,
			ParentId:   kr.ParentId,
			Progress:   kr.Progress,
			Confidence: kr.Confidence,
			Score:      kr.KrScore,
			Comment:    comment.Comment,
		}
		histories = append(histories, history)
	}

	// 开始事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&replay).Error; err != nil {
			return err
		}

		for _, history := range histories {
			history.ReplayId = replay.Id
			if err := tx.Create(history).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// 获取复盘详情
func (s *okrService) GetReplayDetail(id int) (*model.OkrReplay, error) {
	var replay model.OkrReplay
	if err := core.DB.Where("id = ?", id).First(&replay).Error; err != nil {
		return nil, err
	}

	// 获取目标所属别名
	replay.OkrAlias = s.getOwningAlias(replay.OkrUserid, replay.OkrDepartmentId)

	// 获取 KR 复盘历史记录
	var history []*model.OkrReplayHistory
	if err := core.DB.Where("replay_id = ?", id).Find(&history).Error; err != nil {
		return nil, err
	}

	replay.KrHistory = history

	return &replay, nil
}

// 获取目标所属名称
func (s *okrService) getOwningAlias(userid int, departmentId string) []string {
	if departmentId != "" {
		departmentIds := common.ExplodeInt(",", departmentId, true)
		if len(departmentIds) == 0 {
			return nil
		}

		// 获取部门名称
		var departmentNames []string
		if err := core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departmentIds).Pluck("name", &departmentNames).Error; err != nil {
			return nil
		}

		return departmentNames
	}

	// 获取用户名称
	var nickname string
	if err := core.DB.Model(&model.User{}).Where("userid = ?", userid).Pluck("nickname", &nickname).Error; err != nil {
		return nil
	}
	return []string{nickname}
}

// 取消对齐目标
func (s *okrService) CancelAlignObjective(okrId, alignOkrId int) error {
	var align model.OkrAlign
	if err := core.DB.Where("okr_id = ? and align_okr_id = ?", okrId, alignOkrId).First(&align).Error; err != nil {
		return err
	}

	if err := core.DB.Delete(&align).Error; err != nil {
		return err
	}

	return nil
}

// 获取对齐目标列表by目标id
func (s *okrService) GetAlignListByOkrId(user *interfaces.UserInfoResp, okrId int) ([]*interfaces.OkrAlignResp, error) {
	var alignOkrId []int
	if err := core.DB.Model(&model.OkrAlign{}).Where("okr_id = ?", okrId).Pluck("align_okr_id", &alignOkrId).Error; err != nil {
		return nil, err
	}
	// 获取对齐目标
	var alignOkrs []*model.Okr
	if err := core.DB.Unscoped().Model(&model.Okr{}).Where("id in (?)", alignOkrId).Find(&alignOkrs).Error; err != nil {
		return nil, err
	}

	var okrAlignResps []*interfaces.OkrAlignResp
	for _, obj := range alignOkrs {
		okrAlignResps = append(okrAlignResps, &interfaces.OkrAlignResp{
			Okr:   obj,
			Alias: s.getOwningAlias(obj.Userid, obj.DepartmentId),
			Prefix: func() string {
				if obj.ParentId == 0 {
					return "O"
				}
				return "KR"
			}(),
			AlignObjective: func() string {
				if obj.ParentId > 0 {
					o, err := s.GetObjectiveById(obj.ParentId)
					if err != nil {
						return ""
					}
					return o.Title
				}
				return ""
			}(),
		})
	}

	return okrAlignResps, nil
}

// 新增动态日志
func (s *okrService) InsertOkrLog(okrId, userId int, operation, content string) error {
	return s.InsertOkrLogTx(core.DB, okrId, userId, operation, content)
}
func (s *okrService) InsertOkrLogTx(tx *gorm.DB, okrId, userId int, operation, content string) error {
	log := &model.OkrLog{
		OkrId:     okrId,
		Userid:    userId,
		Operation: operation,
		Content:   content,
	}

	if err := tx.Create(log).Error; err != nil {
		return err
	}

	return nil
}

// 获取动态列表
func (s *okrService) GetOkrLogList(okrId int) ([]*model.OkrLog, error) {
	var logs []*model.OkrLog
	if err := core.DB.Where("okr_id = ?", okrId).Order("create_at asc").Find(&logs).Error; err != nil {
		return nil, err
	}

	return logs, nil
}
