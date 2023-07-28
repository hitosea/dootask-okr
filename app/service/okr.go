package service

import (
	"dootask-okr/app/constant"
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/app/repository"
	"dootask-okr/app/utils/common"
	e "dootask-okr/app/utils/error"
	"errors"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

var OkrService = NewOkrService()

type okrService struct {
	okrLogRepo repository.OkrLogRepository
}

func NewOkrService() *okrService {
	return &okrService{
		okrLogRepo: repository.NewOkrLogRepository(),
	}
}

// 创建目标
func (s *okrService) Create(user *interfaces.UserInfoResp, param interfaces.OkrCreateReq) (*model.Okr, error) {
	// 至少有一条关键结果
	if len(param.KeyResults) == 0 {
		return nil, e.New(constant.ErrOkrKeyResultAtLeastOne)
	}

	// 创建部门OKR权限
	if core.DB.Model(&model.UserDepartment{}).Where("owner_userid = ?", user.Userid).First(&model.UserDepartment{}).Error != nil && param.Ascription == 1 {
		return nil, e.New(constant.ErrNoPermission)
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
			if err := s.updateAlignment(obj.Id, param.AlignObjective, tx); err != nil {
				return err
			}
		}

		// 关键结果
		for _, kr := range param.KeyResults {
			keyResult, err := s.createKeyResult(tx, kr, user, obj)
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
		return nil, e.New(constant.ErrOkrNoData)
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
		return nil, e.New(constant.ErrOkrKeyResultAtLeastOne)
	}

	var participantIds []int // 所有参与人
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		for _, kr := range param.KeyResults {
			if kr.Id == 0 {
				// 新增kr
				var addKr *interfaces.OkrKeyResultCreateReq
				addKr.Title = kr.Title
				addKr.Participant = kr.Participant
				addKr.Confidence = kr.Confidence
				addKr.StartAt = kr.StartAt
				addKr.EndAt = kr.EndAt
				keyResult, err := s.createKeyResult(tx, addKr, user, obj)
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
			if err := s.updateAlignment(obj.Id, param.AlignObjective, tx); err != nil {
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
func (s *okrService) createKeyResult(tx *gorm.DB, kr *interfaces.OkrKeyResultCreateReq, user *interfaces.UserInfoResp, obj *model.Okr) (*model.Okr, error) {
	// KR标题
	if kr.Title == "" {
		return nil, e.New(constant.ErrOkrKeyResultTitleEmpty)
	}

	// 信心指数 范围0-100
	if kr.Confidence < 0 || kr.Confidence > 100 {
		return nil, e.New(constant.ErrOkrConfidenceInvalid)
	}

	// 时间
	if kr.StartAt == "" || kr.EndAt == "" {
		return nil, e.New(constant.ErrOkrTimeEmpty)
	}

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
		ParentId:     obj.Id,
		ProjectId:    obj.ProjectId,
		DialogId:     obj.DialogId,
		Priority:     obj.Priority,
		Ascription:   obj.Ascription,
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
	// KR标题
	if kr.Title == "" {
		return nil, e.New(constant.ErrOkrKeyResultTitleEmpty)
	}

	// 信心指数 范围0-100
	if kr.Confidence < 0 || kr.Confidence > 100 {
		return nil, e.New(constant.ErrOkrConfidenceInvalid)
	}

	// 时间
	if kr.StartAt == "" || kr.EndAt == "" {
		return nil, e.New(constant.ErrOkrTimeEmpty)
	}

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
		return nil, e.New(constant.ErrOkrNoData)
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
func (s *okrService) updateAlignment(objId int, alignObjective string, tx ...*gorm.DB) error {
	db := core.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	if err := db.Where("okr_id = ?", objId).Delete(&model.OkrAlign{}).Error; err != nil {
		return err
	}

	alignmentIds := common.ExplodeInt(",", alignObjective, true)
	for _, alignmentId := range alignmentIds {
		if err := db.Where("id = ?", alignmentId).First(&model.Okr{}).Error; err != nil {
			continue
		}
		alignmentObj := &model.OkrAlign{
			OkrId:      objId,
			AlignOkrId: alignmentId,
		}
		if err := db.Create(alignmentObj).Error; err != nil {
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

// 根据id获取是关键结果的目标
func (s *okrService) GetObjectiveByIdIsKeyResult(id int) (*model.Okr, error) {
	var obj model.Okr
	if err := core.DB.Where("id = ? and parent_id > 0", id).First(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

// 根据id获取目标及其关键结果
func (s *okrService) GetObjectiveByIdWithKeyResults(id int) (*model.Okr, error) {
	var obj model.Okr
	if err := core.DB.Preload("KeyResults").Where("id = ?", id).First(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

// 获取我的OkR列表
// 1.可通过目标（O）搜索 2.仅显示我发起的OKR（个人仅能看到个人的）3.按创建时间倒序显示
func (s *okrService) GetMyList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	var objs []*model.Okr
	db := core.DB.Preload("KeyResults").Where("userid = ?", user.Userid).Where("parent_id = 0").Order("created_at desc")
	if objective != "" {
		db = db.Where("title LIKE ?", "%"+objective+"%")
	}

	var count int64
	if err := db.Model(&model.Okr{}).Count(&count).Error; err != nil {
		return nil, err
	}

	var okrs []*interfaces.OkrResp
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&objs).Error; err != nil {
		return nil, err
	}
	for _, obj := range objs {
		okrs = append(okrs, &interfaces.OkrResp{
			Okr: obj,
		})
	}

	okrs, err := s.GetObjectivesWithDetails(okrs, user)
	if err != nil {
		return nil, err
	}

	return interfaces.PaginationRsp(page, pageSize, count, okrs), nil
}

// 获取okr列表额外信息
func (s *okrService) GetObjectivesWithDetails(objs []*interfaces.OkrResp, user *interfaces.UserInfoResp) ([]*interfaces.OkrResp, error) {
	for _, obj := range objs {
		krs := obj.KeyResults
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
	obj.IsFollow = s.isFollow(user.Userid, obj.Id)                             // 是否被关注
	obj.KrCount = len(krs)                                                     // kr总数量
	obj.KrFinishCount = s.getFinishCount(krs)                                  // kr完成数量
	aliasIds := s.getAlignObjectiveIds(obj.Id)                                 // 对齐目标ids
	obj.AlignObjective = aliasIds                                              // 对齐目标
	obj.AlignCount = len(aliasIds)                                             // 对齐目标数量
	obj.Alias = s.getOwningAlias(obj.Ascription, obj.Userid, obj.DepartmentId) // 目标所属名称

	return obj
}

// KR总评分 KR评分=【自评*40%+上级评分*60%】
func (s *okrService) getKrScore(obj *model.Okr) float64 {
	if obj.Score == -1 || obj.SuperiorScore == -1 {
		return 0
	}
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
		return e.New(constant.ErrOkrInvalidKrScore)
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

// 对齐目标ids
func (s *okrService) getAlignObjectiveIds(objId int) []int {
	var ids []int
	if err := core.DB.Model(&model.OkrAlign{}).Where("okr_id = ?", objId).Pluck("align_okr_id", &ids).Error; err != nil {
		return nil
	}
	return ids
}

// 获取参与的OKR列表
// 1.可通过目标（O）搜索 2.被其他OKR选为协助人的OKR（可能是其他部门/其他人的OKR）3.按创建时间倒序显示
func (s *okrService) GetParticipantList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	var objs []*model.Okr
	db := core.DB.Model(&model.Okr{}).
		Where("id IN (?)", core.DB.Model(&model.Okr{}).
			Select("DISTINCT parent_id").
			Where("FIND_IN_SET(?, participant) > 0", user.Userid),
		)

	if objective != "" {
		db = db.Where("title LIKE ?", "%"+objective+"%")
	}

	var count int64
	if err := db.Model(&model.Okr{}).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("KeyResults", "FIND_IN_SET(?, participant) > 0", user.Userid).Offset((page - 1) * pageSize).Limit(pageSize).Find(&objs).Error; err != nil {
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
		s.GetObjectiveExt(obj, obj.KeyResults, user)
	}

	return interfaces.PaginationRsp(page, pageSize, count, okrs), nil
}

// 获取对齐目标列表
// 1. 显示所在部门（即我可见的）+　我参与的
func (s *okrService) GetAlignList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	// 构造部门查询条件
	departmentSql := make([]string, 0, len(user.Department))
	krSql := make([]string, 0, len(user.Department))
	for _, departmentId := range user.Department {
		departmentSql = append(departmentSql, fmt.Sprintf("FIND_IN_SET(%d, department_id) > 0", departmentId))
		krSql = append(krSql, fmt.Sprintf("FIND_IN_SET(%d, department_id) = 0", departmentId))
	}

	db := core.DB.Model(&model.Okr{}).Where("parent_id = 0")
	// 标题筛选
	if objective != "" {
		krSql = append(krSql, fmt.Sprintf("title LIKE '%%%s%%'", objective))
		db = db.Where(fmt.Sprintf("title LIKE '%%%s%%'", objective))
	}
	db = db.Where(strings.Join(departmentSql, " OR "))
	db = db.Or("id IN (?)", core.DB.Model(&model.Okr{}).
		Select("DISTINCT parent_id").
		Where("FIND_IN_SET(?, participant) > 0", user.Userid).
		Where(strings.Join(krSql, " AND ")),
	)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	var okrs []*model.Okr
	if err := db.Preload("KeyResults").Offset((page - 1) * pageSize).Limit(pageSize).Find(&okrs).Error; err != nil {
		return nil, err
	}

	// 参与的O中，只显示我参与的KR
	for _, obj := range okrs {
		krs := make([]*model.Okr, 0, len(obj.KeyResults))
		for _, kr := range obj.KeyResults {
			if strings.Contains(kr.Participant, fmt.Sprintf("%d", user.Userid)) {
				krs = append(krs, kr)
			}
		}
		obj.KeyResults = krs
	}

	return interfaces.PaginationRsp(page, pageSize, count, okrs), nil
}

// 部门的OKR列表
// 1.高级搜索 2.仅包含部门/及部门其他人员的OKR（通过可见范围控制是否看到部门其他同级人员的）3.按创建时间倒序显示 4.部门的OKR置顶（多个的时候多个都置顶，按创建时间倒序）
func (s *okrService) GetDepartmentList(user *interfaces.UserInfoResp, param interfaces.OkrDepartmentListReq, page, pageSize int) (*interfaces.Pagination, error) {
	var objs []*model.Okr
	db := core.DB.Model(&model.Okr{}).Where("parent_id = 0").Order("ascription asc, created_at desc")

	// 用户不是超级管理员时，只能看到自己所在部门的OKR
	if !common.InArray("admin", user.Identity) {
		var sql []string
		for _, departmentId := range user.Department {
			sql = append(sql, fmt.Sprintf("FIND_IN_SET(%d, department_id) > 0", departmentId))
		}
		db = db.Where(strings.Join(sql, " OR "))

		// 可见范围 1-全公司 2-仅相关成员 3-仅部门成员
		db = db.Where("visible_range = 3")
	}

	// 超级管理员可以通过部门筛选
	departmentId := param.DepartmentId
	if departmentId != 0 {
		db = db.Where("FIND_IN_SET(?, department_id) > 0", departmentId)
	}

	// 部门负责人可以通过人员筛选
	userid := param.Userid
	if userid != 0 {
		db = db.Where("userid = ?", userid)
	}

	// 目标筛选
	objective := param.Objective
	if objective != "" {
		db = db.Where("title LIKE ?", "%"+objective+"%")
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

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("KeyResults").Offset((page - 1) * pageSize).Limit(pageSize).Find(&objs).Error; err != nil {
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

	return interfaces.PaginationRsp(page, pageSize, count, okrs), nil
}

// 关注的OKR列表
// 1.可通过目标（O）搜索 2.当前账号关注的OKR 3.按关注时间倒序
func (s *okrService) GetFollowList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	var objs []*model.Okr

	okrTable := core.DBTableName(&model.Okr{})
	okrFollowTable := core.DBTableName(&model.OkrFollow{})

	db := core.DB.Table(okrTable+" AS o").
		Joins(fmt.Sprintf("LEFT JOIN %s AS f ON f.okr_id = o.id", okrFollowTable)).
		Where("f.userid = ?", user.Userid).
		Order("f.created_at desc")

	if objective != "" {
		db = db.Where("o.title LIKE ?", "%"+objective+"%")
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("KeyResults").Offset((page - 1) * pageSize).Limit(pageSize).Find(&objs).Error; err != nil {
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

	return interfaces.PaginationRsp(page, pageSize, count, okrs), nil
}

// 获取复盘列表
func (s *okrService) GetReplayList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	var replays []*model.OkrReplay

	db := core.DB.Model(&model.OkrReplay{}).
		Where("userid = ?", user.Userid).
		Order("created_at DESC")

	if objective != "" {
		db = db.Where("okr_title LIKE ?", "%"+objective+"%")
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("KrHistory").Offset((page - 1) * pageSize).Limit(pageSize).Find(&replays).Error; err != nil {
		return nil, err
	}

	// 获取所属别名
	for _, replay := range replays {
		replay.OkrAlias = s.getOwningAlias(replay.OkrAscription, replay.Userid, replay.OkrDepartmentId)
	}

	return interfaces.PaginationRsp(page, pageSize, count, replays), nil
}

// 获取复盘列表by目标id
func (s *okrService) GetReplayListByOkrId(okrId, page, pageSize int) (*interfaces.Pagination, error) {
	var replays []*model.OkrReplay

	db := core.DB.Model(&model.OkrReplay{}).
		Where("okr_id = ?", okrId).
		Order("created_at ASC")

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&replays).Error; err != nil {
		return nil, err
	}

	return interfaces.PaginationRsp(page, pageSize, count, replays), nil
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
func (s *okrService) FollowObjective(userid, objectiveId int) (*model.OkrFollow, error) {
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
		return &f, nil
	}

	// 如果未关注且需要关注，则创建关注记录
	var okrFollow model.OkrFollow
	if f.Id == 0 {
		if err := core.DB.Create(&model.OkrFollow{
			Userid: userid,
			OkrId:  objectiveId,
		}).Scan(&okrFollow).Error; err != nil {
			return nil, err
		}
	}

	return &okrFollow, nil
}

// 更新进度和进度状态
func (s *okrService) UpdateProgressAndStatus(user *interfaces.UserInfoResp, param interfaces.OkrUpdateProgressReq) (*model.Okr, error) {
	obj, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	// 评分后不允许修改进度
	if obj.Score > -1 {
		return nil, e.New(constant.ErrOkrScoredNotUpdateProgress)
	}

	// 开始事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		// 如果传值更新进度有值，则更新进度
		if param.Progress != 0 {
			logContent := fmt.Sprintf("更新OKR: %s 进度：%d=>%d", obj.Title, obj.Progress, param.Progress)
			if err := s.InsertOkrLogTx(tx, obj.ParentId, user.Userid, "update", logContent); err != nil {
				return err
			}
			obj.Progress = param.Progress
		}

		// 如果传值更新状态有值，则更新状态
		if param.Status != 0 {
			logContent := fmt.Sprintf("更新OKR: %s 状态：%s=>%s", obj.Title, model.ProgressStatusMap[obj.ProgressStatus], model.ProgressStatusMap[param.Status])
			if err := s.InsertOkrLogTx(tx, obj.ParentId, user.Userid, "update", logContent); err != nil {
				return err
			}
			obj.ProgressStatus = param.Status
		}

		// 更新目标
		if err := tx.Save(obj).Error; err != nil {
			return err
		}

		// 检查所在目标的 KR 是否全部完成
		objWithKrs, err := s.GetObjectiveByIdWithKeyResults(obj.ParentId)
		if err != nil {
			return err
		}
		krs := objWithKrs.KeyResults

		allCompleted := true
		sumProgress := 0
		for _, kr := range krs {
			// 更新 KR 进度值
			if param.Id == kr.Id {
				kr.Progress = param.Progress
			}
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
				return err
			}
		} else {
			if err := tx.Model(&model.Okr{}).Where("id = ?", obj.ParentId).Update("progress", progress).Error; err != nil {
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

// 更新评分
func (s *okrService) UpdateScore(user *interfaces.UserInfoResp, param interfaces.OkrScoreReq) (*model.Okr, error) {
	obj, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}
	// 检查进度是否为100%
	if obj.Progress < 100 {
		return nil, e.New(constant.ErrOkrProgressNotEnough)
	}

	// 检查用户是否为目标负责人或上级 false-负责人 true-上级
	superior := s.IsObjectiveManager(obj, user)
	if !superior {
		// 检查用户是否为目标负责人
		if obj.Userid != user.Userid {
			return nil, e.New(constant.ErrOkrNoPermissionScore)
		}
		// 检查是否已评分
		if obj.Score > -1 {
			return nil, e.New(constant.ErrOkrOwnerScored)
		}
		// 负责人评分
		err = core.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Update("score", param.Score).Scan(&obj).Error; err != nil {
				return err
			}

			// 新增动态日志
			logContent := fmt.Sprintf("责任人打分: %s", obj.Title)
			if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", logContent); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 需要负责人评分才可以上级评分
		if obj.Score == -1 {
			return nil, e.New(constant.ErrOkrOwnerNotScore)
		}
		// 检查是否已评分
		if obj.SuperiorScore > -1 {
			return nil, e.New(constant.ErrOkrSuperiorScored)
		}
		// 上级评分
		err = core.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Update("superior_score", param.Score).Scan(&obj).Error; err != nil {
				return err
			}

			// 更新O评分
			obj, err := s.GetObjectiveByIdWithKeyResults(obj.ParentId)
			if err != nil {
				return err
			}
			err = s.UpdateObjectiveScoreTx(tx, obj)
			if err != nil {
				return err
			}

			// 新增动态日志
			logContent := fmt.Sprintf("上级打分: %s", obj.Title)
			if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", logContent); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// 检查用户是否为目标上级
func (s *okrService) IsObjectiveManager(kr *model.Okr, user *interfaces.UserInfoResp) bool {
	// 目标负责人的部门
	depIds := common.ExplodeInt(",", kr.DepartmentId, true)
	if len(depIds) == 0 {
		return false
	}

	// 负责人 = 部门负责人
	if kr.Userid == user.Userid && kr.Score == 0 {
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
func (s *okrService) CancelObjective(okrId int) (*model.Okr, error) {
	obj, err := s.GetObjectiveById(okrId)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	// 更新取消状态
	if obj.Canceled == 0 {
		obj.Canceled = 1
	} else if obj.Canceled == 1 {
		obj.Canceled = 0
	}

	if err := core.DB.Save(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

// 更新参与人
func (s *okrService) UpdateParticipant(param interfaces.OkrParticipantUpdateReq) (*model.Okr, error) {
	obj, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	obj.Participant = param.Participant
	if err := core.DB.Save(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

// 更新信心指数
func (s *okrService) UpdateConfidence(param interfaces.OkrConfidenceUpdateReq) (*model.Okr, error) {
	obj, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	obj.Confidence = param.Confidence
	if err := core.DB.Save(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

// 创建 OKR 复盘记录
func (s *okrService) CreateOkrReplay(userid int, req interfaces.OkrReplayCreateReq) (*model.OkrReplay, error) {
	// 获取 OKR 目标及其关键结果
	obj, err := s.GetObjectiveByIdWithKeyResults(req.OkrId)
	if err != nil {
		return nil, err
	}

	// 检查关键结果是否已评分
	for _, kr := range obj.KeyResults {
		if kr.Score == 0 || kr.SuperiorScore == 0 {
			return nil, e.New(constant.ErrOkrKrScoreNotComplete)
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
		OkrAscription:   obj.Ascription,
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
		return nil, err
	}

	return &replay, nil
}

// 获取复盘详情
func (s *okrService) GetReplayDetail(id int) (*model.OkrReplay, error) {
	var replay model.OkrReplay
	if err := core.DB.Where("id = ?", id).First(&replay).Error; err != nil {
		return nil, err
	}

	// 获取目标所属别名
	replay.OkrAlias = s.getOwningAlias(replay.OkrAscription, replay.OkrUserid, replay.OkrDepartmentId)

	// 获取 KR 复盘历史记录
	var history []*model.OkrReplayHistory
	if err := core.DB.Where("replay_id = ?", id).Find(&history).Error; err != nil {
		return nil, err
	}

	replay.KrHistory = history

	return &replay, nil
}

// 获取目标所属名称
func (s *okrService) getOwningAlias(ascription int, userid int, departmentId string) []string {
	alias := []string{""}
	if departmentId != "" && ascription == 1 {
		departmentIds := common.ExplodeInt(",", departmentId, true)
		if len(departmentIds) == 0 {
			return nil
		}

		// 获取部门名称
		if err := core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departmentIds).Pluck("name", &alias).Error; err != nil {
			return nil
		}

		return alias
	}

	// 获取用户名称
	if userid > 0 && ascription == 2 {
		var user model.User
		if err := core.DB.Model(&model.User{}).Where("userid = ?", userid).Find(&user).Error; err != nil {
			return nil
		}
		alias = []string{user.GetNickname()}
	}
	return alias
}

// UpdateAlignObjective 更新对齐目标
func (s *okrService) UpdateAlignObjective(okrId int, alignObjective string) (*model.Okr, error) {
	obj, err := s.GetObjectiveById(okrId)
	if err != nil {
		return nil, err
	}

	if err := s.updateAlignment(obj.Id, alignObjective); err != nil {
		return nil, err
	}

	return obj, nil
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
			Okr: obj,
			Alias: func() []string {
				if obj.ParentId == 0 {
					// 个人或部门
					return s.getOwningAlias(obj.Ascription, obj.Userid, obj.DepartmentId)
				}
				// 参与人
				ids := common.ExplodeInt(",", obj.Participant, true)
				var nicknames []string
				if err := core.DB.Model(&model.User{}).Where("userid in (?)", ids).Pluck("nickname", &nicknames).Error; err != nil {
					return nil
				}
				return nicknames
			}(),
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

// tx新增动态日志
func (s *okrService) InsertOkrLogTx(tx *gorm.DB, okrId, userId int, operation, content string) error {

	log := &model.OkrLog{
		OkrId:     okrId,
		Userid:    userId,
		Operation: operation,
		Content:   content,
	}

	return s.okrLogRepo.CreateTx(tx, log)
}

// 获取动态列表
func (s *okrService) GetOkrLogList(user *interfaces.UserInfoResp, okrId, page, pageSize int) (*interfaces.Pagination, error) {
	count, err := s.okrLogRepo.CountByOkrId(okrId)
	if err != nil {
		return nil, err
	}

	logs, err := s.okrLogRepo.FindByOkrId(okrId, page, pageSize)
	if err != nil {
		return nil, err
	}

	// 获取用户头像、昵称
	var userids []int
	err = core.DB.Model(&model.OkrLog{}).Where("okr_id = ?", okrId).Pluck("DISTINCT userid", &userids).Error
	if err != nil {
		return nil, err
	}
	userList, err := DootaskService.GetUserBasic(user.Token, userids)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		for _, user := range userList {
			if user.Userid == log.Userid {
				log.UserAvatar = user.Userimg
				log.UserNickname = user.Nickname
			}
		}
	}

	return interfaces.PaginationRsp(page, pageSize, count, logs), nil
}

// 获取部门列表
func (s *okrService) GetDepartmentSearch(page, pageSize int) (*interfaces.Pagination, error) {
	var departments []*model.UserDepartment
	var count int64

	if err := core.DB.Model(&model.UserDepartment{}).Count(&count).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize
	if err := core.DB.Model(&model.UserDepartment{}).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&departments).Error; err != nil {
		return nil, err
	}

	return interfaces.PaginationRsp(page, pageSize, count, departments), nil
}
