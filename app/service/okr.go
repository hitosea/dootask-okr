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
	if core.DB.Model(&model.UserDepartment{}).Where("parent_id = 0").Where("owner_userid = ?", user.Userid).First(&model.UserDepartment{}).Error != nil && param.Ascription == 1 {
		return nil, e.New(constant.ErrNoPermission)
	}

	// 时间格式化
	startAt, err := common.ParseTime(param.StartAt)
	if err != nil {
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	endAt, err := common.ParseTime(param.EndAt)
	if err != nil {
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	// 开始时间不能大于结束时间，请重新选择合适的时间段
	if startAt.After(endAt) {
		return nil, e.New(constant.ErrOkrTimeInvalid)
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
			if err := s.updateAlignment(obj, user.Userid, param.AlignObjective, true, tx); err != nil {
				return err
			}
		}

		// 关键结果
		for _, kr := range param.KeyResults {
			// 去掉kr.Participant中的0或0,
			kr.Participant = strings.TrimLeft(kr.Participant, "0,")
			kr.Participant = strings.TrimLeft(kr.Participant, "0")

			keyResult, err := s.createKeyResult(tx, kr, user, obj)
			if err != nil {
				return err
			}
			obj.KeyResults = append(obj.KeyResults, keyResult)
			participantIds = append(participantIds, common.ExplodeInt(",", kr.Participant, true)...)
		}

		// 动态日志
		if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "add", "创建OKR", nil); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 创建O时（通知发起/所有KR参与人）
	participantIds = common.ArrayUniqueInt(participantIds)
	if len(participantIds) > 0 {
		go DootaskService.DialogOkrPush(obj, user.Token, 1, participantIds)
	}

	return obj, nil
}

// 更新目标
func (s *okrService) Update(user *interfaces.UserInfoResp, param interfaces.OkrUpdateReq) (*model.Okr, error) {
	obj, err := s.GetObjectiveByIdWithKeyResults(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	err = s.CheckObjectiveOperation(obj, user.Userid)
	if err != nil {
		return nil, err
	}

	startAt, err := common.ParseTime(param.StartAt)
	if err != nil {
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	endAt, err := common.ParseTime(param.EndAt)
	if err != nil {
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	// 开始时间不能大于结束时间，请重新选择合适的时间段
	if startAt.After(endAt) {
		return nil, e.New(constant.ErrOkrTimeInvalid)
	}

	// 至少有一条关键结果
	if len(param.KeyResults) == 0 {
		return nil, e.New(constant.ErrOkrKeyResultAtLeastOne)
	}

	// 比较 obj.KeyResults 和 param.KeyResults 中的 id 差值，然后删除
	var krIds []int
	var diffIds []int
	for _, kr := range obj.KeyResults {
		krIds = append(krIds, kr.Id)
	}

	for _, kr := range param.KeyResults {
		if kr.Id != 0 {
			diffIds = append(diffIds, kr.Id)
		}
	}

	diffIds = common.ArrayDifferenceProcessing(krIds, diffIds)
	obj.VisibleRange = param.VisibleRange // 可见范围优先赋值
	obj.Type = param.Type                 // 类型优先赋值
	obj.Priority = param.Priority         // 优先级优先赋值

	var participantIds []int // 所有参与人
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		if len(diffIds) > 0 {
			if err := tx.Where("id in (?)", diffIds).Delete(&model.Okr{}).Error; err != nil {
				return err
			}
		}
		obj.KeyResults = nil
		for _, kr := range param.KeyResults {
			// 去掉kr.Participant中的0
			kr.Participant = strings.TrimLeft(kr.Participant, "0,")
			kr.Participant = strings.TrimLeft(kr.Participant, "0")
			if kr.Id == 0 {
				// 新增kr
				var addKr interfaces.OkrKeyResultCreateReq
				addKr.Title = kr.Title
				addKr.Participant = kr.Participant
				addKr.Confidence = kr.Confidence
				addKr.StartAt = kr.StartAt
				addKr.EndAt = kr.EndAt
				keyResult, err := s.createKeyResult(tx, &addKr, user, obj)
				if err != nil {
					return err
				}
				// 新增kr时，发送提示消息
				keyResult.ParentTitle = obj.Title
				go DootaskService.DialogOkrPush(keyResult, user.Token, 4, common.ExplodeInt(",", kr.Participant, true))
				obj.KeyResults = append(obj.KeyResults, keyResult)
			} else {
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
			if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", "修改O目标标题", interfaces.OkrLogParams{
				TitleChange: []string{obj.Title, param.Title},
			}); err != nil {
				return err
			}
			if len(participantIds) > 0 {
				go DootaskService.DialogOkrPush(obj, user.Token, 2, participantIds)
			}
		}

		// O时间变动时，发送提示消息
		if obj.StartAt != startAt || obj.EndAt != endAt {
			if err := s.InsertOkrLogTx(tx, obj.Id, user.Userid, "update", "修改O目标周期", interfaces.OkrLogParams{
				TimeChange: []string{common.FormatDate2(obj.StartAt) + "~" + common.FormatDate2(obj.EndAt), common.FormatDate2(startAt) + "~" + common.FormatDate2(endAt)},
			}); err != nil {
				return err
			}
			if len(participantIds) > 0 {
				go DootaskService.DialogOkrPush(obj, user.Token, 5, participantIds)
			}
		}

		obj.Title = param.Title
		obj.ProjectId = param.ProjectId
		obj.StartAt = startAt
		obj.EndAt = endAt

		// 重新计算总目标进度值
		progress := s.CalculateProgressTx(tx, obj)
		obj.Progress = progress
		if progress >= 100 {
			obj.Completed = 1
			// O目标完成时，重新计算KR总评分
			score := s.CalculateScoreTx(tx, obj)
			if score > 0 {
				obj.Score = score
			}
		} else {
			obj.Completed = 0
		}

		if err := tx.Save(obj).Error; err != nil {
			return err
		}

		// 更新对齐目标
		if param.AlignObjective != "" {
			if err := s.updateAlignment(obj, user.Userid, param.AlignObjective, false, tx); err != nil {
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

// 重新计算总目标进度值
func (s *okrService) CalculateProgressTx(tx *gorm.DB, obj *model.Okr) int {
	krs := obj.KeyResults
	sumProgress := 0
	for _, item := range krs {
		// 计算总进度值
		sumProgress += item.Progress
	}

	// 更新总目标进度值，kr 进度值相加/kr 数量
	progress := 0
	if len(krs) > 0 {
		progress = int(math.Round(float64(sumProgress) / float64(len(krs))))
	}

	return progress
}

// O目标完成时，重新计算KR总评分
func (s *okrService) CalculateScoreTx(tx *gorm.DB, obj *model.Okr) float64 {
	// 是否满足计算KR总评分条件
	completed := true
	for _, item := range obj.KeyResults {
		if item.Score == -1 || item.SuperiorScore == -1 {
			completed = false
			break
		}
	}

	if !completed {
		return 0
	}

	// 计算KR总评分
	score := s.getObjectiveScore(obj)

	return score
}

// 创建关键结果
func (s *okrService) createKeyResult(tx *gorm.DB, kr *interfaces.OkrKeyResultCreateReq, user *interfaces.UserInfoResp, obj *model.Okr) (*model.Okr, error) {
	// KR标题
	if kr.Title == "" {
		return nil, e.New(constant.ErrOkrKeyResultTitleEmpty)
	}

	// 标题长度 255
	if !common.IsChineseCharCountValid(kr.Title) {
		return nil, e.New(constant.ErrOkrTitleLengthInvalid)
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
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	endAt, err := common.ParseTime(kr.EndAt)
	if err != nil {
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	// 开始时间不能大于结束时间，请重新选择合适的时间段
	if startAt.After(endAt) {
		return nil, e.New(constant.ErrOkrTimeInvalid)
	}

	keyResult := &model.Okr{
		Userid:       user.Userid,
		DepartmentId: common.ArrayImplode(user.Department),
		ParentId:     obj.Id,
		ProjectId:    obj.ProjectId,
		DialogId:     obj.DialogId,
		Type:         obj.Type,
		Priority:     obj.Priority,
		Ascription:   obj.Ascription,
		VisibleRange: obj.VisibleRange,
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

	// 标题长度 255
	if !common.IsChineseCharCountValid(kr.Title) {
		return nil, e.New(constant.ErrOkrTitleLengthInvalid)
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
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	endAt, err := common.ParseTime(kr.EndAt)
	if err != nil {
		return nil, e.New(constant.ErrOkrTimeFormat)
	}

	// 开始时间不能大于结束时间，请重新选择合适的时间段
	if startAt.After(endAt) {
		return nil, e.New(constant.ErrOkrTimeInvalid)
	}

	keyResult, err := s.GetObjectiveById(kr.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	// 判断传入的KR是否跟数据库中的KR是否有改变
	if keyResult.Type != obj.Type || keyResult.Priority != obj.Priority || keyResult.Title != kr.Title || keyResult.Confidence != kr.Confidence || keyResult.StartAt != startAt || keyResult.EndAt != endAt || keyResult.Participant != kr.Participant {
		err = s.CheckObjectiveOperation(keyResult, user.Userid)
		if err != nil {
			return nil, err
		}
	}

	// 父级目标标题
	keyResult.ParentTitle = obj.Title

	oldParticipant := common.ExplodeInt(",", keyResult.Participant, true)
	newParticipant := common.ExplodeInt(",", kr.Participant, true)
	diffParticipant := common.ArrayDifferenceProcessing(newParticipant, oldParticipant)

	// KR变动时，发送提示消息
	if keyResult.Title != kr.Title {
		if err := s.InsertOkrLogTx(core.DB, keyResult.ParentId, user.Userid, "update", "修改KR标题", interfaces.OkrLogParams{
			TitleChange: []string{keyResult.Title, kr.Title},
		}); err != nil {
			return nil, err
		}
		go DootaskService.DialogOkrPush(keyResult, user.Token, 3, newParticipant)
	}

	// KR时间变动时，发送提示消息
	if keyResult.StartAt != startAt || keyResult.EndAt != endAt {
		if err := s.InsertOkrLogTx(core.DB, keyResult.ParentId, user.Userid, "update", "修改KR周期", interfaces.OkrLogParams{
			Title:      keyResult.Title,
			TimeChange: []string{common.FormatDate2(keyResult.StartAt) + "~" + common.FormatDate2(keyResult.EndAt), common.FormatDate2(startAt) + "~" + common.FormatDate2(endAt)},
		}); err != nil {
			return nil, err
		}
		go DootaskService.DialogOkrPush(keyResult, user.Token, 6, newParticipant)
	}

	// 为KR添加新参与人时，发送提示消息
	if len(diffParticipant) > 0 {
		if err := s.InsertOkrLogTx(core.DB, keyResult.ParentId, user.Userid, "update", "修改KR参与人", nil); err != nil {
			return nil, err
		}

		addDiffParticipant := common.ArrayDifferenceAddProcessing(newParticipant, oldParticipant)
		if len(addDiffParticipant) > 0 {
			go DootaskService.DialogOkrPush(keyResult, user.Token, 4, diffParticipant)
		}
	}

	// KR信息指数变动时，新增动态信息
	if keyResult.Confidence != kr.Confidence {
		if err := s.InsertOkrLogTx(core.DB, keyResult.ParentId, user.Userid, "update", "修改KR信心指数", interfaces.OkrLogParams{
			Title:            kr.Title,
			ConfidenceChange: []int{keyResult.Confidence, kr.Confidence},
		}); err != nil {
			return nil, err
		}
	}
	keyResult.Type = obj.Type
	keyResult.Priority = obj.Priority
	keyResult.ProjectId = obj.Id
	keyResult.Participant = kr.Participant
	keyResult.Title = kr.Title
	keyResult.Confidence = kr.Confidence
	keyResult.VisibleRange = obj.VisibleRange
	keyResult.StartAt = startAt
	keyResult.EndAt = endAt

	if err := tx.Save(keyResult).Error; err != nil {
		return nil, err
	}

	return keyResult, nil
}

// 更新对齐目标
func (s *okrService) updateAlignment(obj *model.Okr, userid int, alignObjective string, isCreate bool, tx ...*gorm.DB) error {
	db := core.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	var ids []int
	if err := db.Model(&model.OkrAlign{}).Where("okr_id = ?", obj.Id).Pluck("align_okr_id", &ids).Error; err != nil {
		return err
	}

	alignmentIds := common.ExplodeInt(",", alignObjective, true)

	if !common.IsEqual(ids, alignmentIds) {
		if !isCreate {
			if err := s.InsertOkrLogTx(db, obj.Id, userid, "update", "修改对齐目标", nil); err != nil {
				return err
			}
		}
	}

	// 计算新增的差集
	alignmentDiffIds := common.ArrayDifferenceAddProcessing(alignmentIds, ids)

	// 批量插入新的对齐目标
	var alignmentObjs []*model.OkrAlign
	for _, alignmentId := range alignmentDiffIds {
		if err := db.Where("id = ?", alignmentId).First(&model.Okr{}).Error; err != nil {
			continue
		}
		alignmentObjs = append(alignmentObjs, &model.OkrAlign{
			OkrId:      obj.Id,
			AlignOkrId: alignmentId,
		})
	}

	if len(alignmentObjs) > 0 {
		if err := db.Create(alignmentObjs).Error; err != nil {
			return err
		}
	}

	return nil
}

// 根据id获取目标
func (s *okrService) GetObjectiveById(id int) (*model.Okr, error) {
	var obj model.Okr
	if err := core.DB.Where("id = ?", id).First(&obj).Error; err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, e.New(constant.ErrOkrNoData)
		}
		return nil, err
	}
	return &obj, nil
}

// 根据id获取是关键结果的目标
func (s *okrService) GetObjectiveByIdIsKeyResult(id int) (*model.Okr, error) {
	var obj model.Okr
	if err := core.DB.Where("id = ? and parent_id > 0", id).First(&obj).Error; err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, e.New(constant.ErrOkrNoData)
		}
		return nil, err
	}
	return &obj, nil
}

// 根据id获取目标及其关键结果
func (s *okrService) GetObjectiveByIdWithKeyResults(id int) (*model.Okr, error) {
	var obj model.Okr
	if err := core.DB.Preload("KeyResults").Where("id = ?", id).First(&obj).Error; err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, e.New(constant.ErrOkrNoData)
		}
		return nil, err
	}
	return &obj, nil
}

// 获取我的OkR列表
// 1.可通过目标（O）搜索 2.仅显示我发起的OKR（个人仅能看到个人的）3.按创建时间倒序显示
func (s *okrService) GetMyList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	var objs []*model.Okr
	db := core.DB.Preload("KeyResults").Where("userid = ?", user.Userid).Where("parent_id = 0").Order("canceled,completed asc, created_at desc")
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

		// KR总评分
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

	// 用户头像
	users, _ := DootaskService.GetUserBasic(user.Token, []int{obj.Userid})
	if len(users) > 0 {
		obj.UserAvatar = users[0].Userimg
	}

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
		if kr.Score == -1 || kr.SuperiorScore == -1 {
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
		).Order("canceled,completed asc, created_at desc")

	if objective != "" {
		db = db.Where("title LIKE ?", "%"+objective+"%")
	}

	var count int64
	if err := db.Model(&model.Okr{}).Count(&count).Error; err != nil {
		return nil, err
	}

	db = db.Preload("KeyResults") // 显示全部KR
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&objs).Error; err != nil {
		return nil, err
	}

	// 列表响应
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

// 获取对齐目标列表
// 1. 显示所在部门（即我可见的）+　我参与的
func (s *okrService) GetAlignList(user *interfaces.UserInfoResp, objective string, page, pageSize int) (*interfaces.Pagination, error) {
	var (
		departmentSql  strings.Builder
		participantSql strings.Builder
		participantIds []int
		okrs           []*model.Okr
		count          int64
		err            error
		okrTable       = core.DBTableName(&model.Okr{})
	)

	// 已取消和已完成的OKR不显示 测试提出的需求
	db := core.DB.Table(okrTable + " AS okr").Select("DISTINCT okr.*").Where("okr.canceled = 0 AND okr.completed = 0").Order("okr.created_at desc")
	participantDb := core.DB.Table(okrTable + " AS okr2").Where("okr2.canceled = 0 AND okr2.completed = 0").Order("okr2.created_at desc")

	// 用户不是超级管理员时，只能看到自己所在部门的OKR
	if !user.IsAdmin() {
		if len(user.Department) == 0 {
			return interfaces.PaginationRsp(page, pageSize, 0, nil), nil
		}

		departments, err := s.GetDepartmentsBySearchDeptId(user.Department)
		if err != nil {
			return nil, err
		}

		for i, departmentId := range departments {
			if i > 0 {
				departmentSql.WriteString(" OR ")
				participantSql.WriteString(" AND ")
			}
			fmt.Fprintf(&departmentSql, "FIND_IN_SET(%d, okr.department_id) > 0", departmentId)
			fmt.Fprintf(&participantSql, "FIND_IN_SET(%d, okr2.department_id) = 0", departmentId)
		}

		// 判断是否是普通组员
		var department model.UserDepartment
		core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departments).Where("owner_userid = ?", user.Userid).First(&department)
		if department.Id == 0 {
			// 1.自己发布的 2.可见范围 1-全公司 2-仅相关成员 3-仅部门成员
			db = db.Where("okr.visible_range IN (1, 3) OR okr.userid = ?", user.Userid)
		} else {
			// 判断是否是同级部门负责人
			var departmentSameLevel []model.UserDepartment
			core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departments).Where("parent_id > 0").Where("owner_userid = ?", user.Userid).Find(&departmentSameLevel)
			if len(departmentSameLevel) > 0 {
				// 部门负责人可以看到自己所在部门的所有OKR
				var sqlSame []string
				for _, department := range departmentSameLevel {
					sqlSame = append(sqlSame, fmt.Sprintf("FIND_IN_SET(%d, okr.department_id) > 0", department.Id))
				}
				db = db.Where("okr.visible_range IN (1, 3) OR (okr.visible_range = 2 AND ("+strings.Join(sqlSame, " OR ")+")) OR okr.userid = ?", user.Userid)
			}
		}
	}

	// 标题筛选
	if objective != "" {
		// 部门
		db = db.Joins(fmt.Sprintf(`LEFT JOIN %s AS son ON okr.id = son.parent_id`, okrTable))
		db = db.Where(`(okr.title LIKE ? OR son.title LIKE ?)`, "%"+objective+"%", "%"+objective+"%")

		// 我参与的
		participantDb = participantDb.Joins(fmt.Sprintf(`LEFT JOIN %s AS parent ON okr2.parent_id = parent.id`, okrTable))
		participantDb = participantDb.Where(`(okr2.title LIKE ? OR parent.title LIKE ?)`, "%"+objective+"%", "%"+objective+"%")
	}

	db = db.Where(departmentSql.String())
	db = db.Where("okr.parent_id = 0")

	if err = participantDb.
		Where("FIND_IN_SET(?, okr2.participant) > 0", user.Userid).
		Where(participantSql.String()).
		Pluck("DISTINCT okr2.parent_id", &participantIds).Error; err != nil {
		return nil, err
	}

	db = db.Or("okr.id IN (?)", common.ArrayUniqueInt(participantIds))

	if err = db.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("KeyResults").Offset((page - 1) * pageSize).Limit(pageSize).Find(&okrs).Error; err != nil {
		return nil, err
	}

	return interfaces.PaginationRsp(page, pageSize, count, okrs), nil
}

// 部门的OKR列表
// 1.高级搜索 2.仅包含部门/及部门其他人员的OKR（通过可见范围控制是否看到部门其他同级人员的）3.按创建时间倒序显示 4.部门的OKR置顶（多个的时候多个都置顶，按创建时间倒序）
func (s *okrService) GetDepartmentList(user *interfaces.UserInfoResp, param interfaces.OkrDepartmentListReq, page, pageSize int) (*interfaces.Pagination, error) {
	var objs []*model.Okr
	db := core.DB.Model(&model.Okr{}).Where("parent_id = 0").Order("canceled,completed asc, ascription asc, created_at desc")

	// 用户不是超级管理员时，只能看到自己所在部门的OKR
	if !user.IsAdmin() {
		if len(user.Department) == 0 {
			return interfaces.PaginationRsp(page, pageSize, 0, nil), nil
		}

		departments, err := s.GetDepartmentsBySearchDeptId(user.Department)
		if err != nil {
			return nil, err
		}
		var sql []string
		for _, departmentId := range departments {
			sql = append(sql, fmt.Sprintf("FIND_IN_SET(%d, department_id) > 0", departmentId))
		}
		db = db.Where(strings.Join(sql, " OR "))

		// 判断是否是普通组员
		var department model.UserDepartment
		core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departments).Where("owner_userid = ?", user.Userid).First(&department)
		if department.Id == 0 {
			// 1.自己发布的 2.可见范围 1-全公司 2-仅相关成员 3-仅部门成员
			db = db.Where("visible_range IN (1, 3) OR userid = ?", user.Userid)
		} else {
			// 判断是否是同级部门负责人
			var departmentSameLevel []model.UserDepartment
			core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departments).Where("parent_id > 0").Where("owner_userid = ?", user.Userid).Find(&departmentSameLevel)
			if len(departmentSameLevel) > 0 {
				// 部门负责人可以看到自己所在部门的所有OKR
				var sqlSame []string
				for _, department := range departmentSameLevel {
					sqlSame = append(sqlSame, fmt.Sprintf("FIND_IN_SET(%d, department_id) > 0", department.Id))
				}
				db = db.Where("visible_range IN (1, 3) OR (visible_range = 2 AND ("+strings.Join(sqlSame, " OR ")+")) OR userid = ?", user.Userid)
			}
		}
	}

	// 超级管理员可以通过部门筛选
	departmentId := param.DepartmentId
	if departmentId != 0 {
		admDepartments, err := s.GetDepartmentsBySearchDeptId([]int{departmentId})
		if err != nil {
			return nil, err
		}
		var admSql []string
		for _, departmentId := range admDepartments {
			admSql = append(admSql, fmt.Sprintf("FIND_IN_SET(%d, department_id) > 0", departmentId))
		}
		db = db.Where(strings.Join(admSql, " OR "))
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
			return nil, e.New(constant.ErrOkrTimeFormat)
		}
		endAt, err := common.ParseTime(endAtStr)
		if err != nil {
			return nil, e.New(constant.ErrOkrTimeFormat)
		}
		db = db.Where("(start_at >= ? AND start_at <= ?) OR (end_at >= ? AND end_at <= ?) OR (start_at <= ? AND end_at >= ?)", startAt, endAt, startAt, endAt, startAt, endAt)
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
			db = db.Where("progress >= 100 and score = -1")
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
		Order("o.canceled,o.completed asc, f.created_at desc")

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

	objResp := &interfaces.OkrResp{
		Okr:          obj,
		SuperiorUser: s.GetSuperiorUserIds(obj, user.Userid),
	}

	s.GetObjectiveExt(objResp, obj.KeyResults, user)

	return objResp, nil
}

// 关注或取消关注目标
func (s *okrService) FollowObjective(userid, objectiveId int) (*model.OkrFollow, error) {
	// 只要顶级目标才能被关注
	var obj model.Okr
	if err := core.DB.Where("id = ? and parent_id = 0", objectiveId).First(&obj).Error; err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, e.New(constant.ErrOkrNoData)
		}
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
	kr, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	err = s.CheckObjectiveOperation(kr, user.Userid)
	if err != nil {
		return nil, err
	}

	// 开始事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		// 如果传值更新进度有值，则更新进度
		if param.Progress != kr.Progress {
			if err := s.InsertOkrLogTx(tx, kr.ParentId, user.Userid, "update", "修改KR进度", interfaces.OkrLogParams{
				Title:          kr.Title,
				ProgressChange: []int{kr.Progress, param.Progress},
			}); err != nil {
				return err
			}
			kr.Progress = param.Progress
		}

		// 如果传值更新状态有值，则更新状态
		if param.Status != kr.ProgressStatus {
			if err := s.InsertOkrLogTx(tx, kr.ParentId, user.Userid, "update", "修改KR状态", interfaces.OkrLogParams{
				Title:                kr.Title,
				ProgressStatusChange: []string{model.ProgressStatusMap[kr.ProgressStatus], model.ProgressStatusMap[param.Status]},
			}); err != nil {
				return err
			}
			kr.ProgressStatus = param.Status
		}

		// 更新目标
		if err := tx.Save(kr).Error; err != nil {
			return err
		}

		// 检查所在目标的 KR 是否全部完成
		objWithKrs, err := s.GetObjectiveByIdWithKeyResults(kr.ParentId)
		if err != nil {
			return err
		}
		krs := objWithKrs.KeyResults

		allCompleted := true
		sumProgress := 0
		for _, item := range krs {
			// 更新 KR 进度值
			if param.Id == item.Id {
				item.Progress = param.Progress
			}
			// 进度全部完成 100%
			if item.Progress < 100 {
				allCompleted = false
			}
			// 计算总进度值
			sumProgress += item.Progress
		}

		// 更新总目标进度值，kr 进度值相加/kr 数量
		progress := 0
		if len(krs) > 0 {
			progress = int(math.Round(float64(sumProgress) / float64(len(krs))))
		}

		// 更新总目标的状态是否完成
		if allCompleted {
			if err := tx.Model(&model.Okr{}).Where("id = ?", kr.ParentId).Updates(map[string]interface{}{
				"Completed": 1,
				"Progress":  progress,
			}).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&model.Okr{}).Where("id = ?", kr.ParentId).Updates(map[string]interface{}{
				"Completed": 0,
				"Progress":  progress,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return kr, nil
}

// 更新评分
func (s *okrService) UpdateScore(user *interfaces.UserInfoResp, param interfaces.OkrScoreReq) (*model.Okr, error) {
	kr, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}
	// 检查进度是否为100%
	if kr.Progress < 100 {
		return nil, e.New(constant.ErrOkrProgressNotEnough)
	}

	// 检查用户是否为目标负责人或上级 false-负责人 true-上级
	superior := s.IsObjectiveManager(kr, user)
	if !superior {
		// 检查是否已评分
		if kr.Score > -1 {
			return nil, e.New(constant.ErrOkrOwnerScored)
		}
		// 检查用户是否为目标负责人
		if kr.Userid != user.Userid {
			return nil, e.New(constant.ErrOkrOwnerNotCancel)
		}
		// 负责人评分
		err = core.DB.Transaction(func(tx *gorm.DB) error {
			// 如果是超管评分，则更新目标评分
			if user.IsAdmin() {
				if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
					"score":          param.Score,
					"superior_score": param.Score,
				}).Scan(&kr).Error; err != nil {
					return err
				}
				// 更新O评分
				var obj *model.Okr
				tx.Preload("KeyResults").Where("id = ?", kr.ParentId).First(&obj)
				if err != nil {
					return err
				}
				err = s.UpdateObjectiveScoreTx(tx, obj)
				if err != nil {
					return err
				}
			} else {
				if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Update("score", param.Score).Scan(&kr).Error; err != nil {
					return err
				}
			}

			// 新增动态日志
			if err := s.InsertOkrLogTx(tx, kr.ParentId, user.Userid, "update", "责任人打分", interfaces.OkrLogParams{
				Title: kr.Title,
			}); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 需要负责人评分才可以上级评分
		if kr.Score == -1 {
			return nil, e.New(constant.ErrOkrOwnerNotScore)
		}
		// 检查是否已评分
		if kr.SuperiorScore > -1 {
			return nil, e.New(constant.ErrOkrSuperiorScored)
		}
		// 上级评分
		err = core.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&model.Okr{}).Where("id = ?", param.Id).Update("superior_score", param.Score).Scan(&kr).Error; err != nil {
				return err
			}

			// 新增动态日志
			if err := s.InsertOkrLogTx(tx, kr.ParentId, user.Userid, "update", "上级打分", interfaces.OkrLogParams{
				Title: kr.Title,
			}); err != nil {
				return err
			}

			// 更新O评分
			var obj *model.Okr
			tx.Preload("KeyResults").Where("id = ?", kr.ParentId).First(&obj)
			if err != nil {
				return err
			}
			err = s.UpdateObjectiveScoreTx(tx, obj)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return kr, nil
}

// 检查用户是否为目标上级
func (s *okrService) IsObjectiveManager(kr *model.Okr, user *interfaces.UserInfoResp) bool {
	// 负责人 = 部门负责人
	if kr.Userid == user.Userid && kr.Score == -1 {
		return false
	}

	// 是否超管评分 1.顶级部门负责人已评分 3.顶级部门负责人评分后，超管可评分
	var topDepartment model.UserDepartment
	core.DB.Model(&model.UserDepartment{}).Where("parent_id = 0").Where("owner_userid = ?", kr.Userid).First(&topDepartment)
	if user.IsAdmin() && topDepartment.Id > 0 {
		return true
	}

	// 目标负责人的部门
	depIds := common.ExplodeInt(",", kr.DepartmentId, true)
	if len(depIds) == 0 {
		return false
	}

	// 获取部门小组负责人
	deptAllIds, _ := s.GetDepartmentsBySearchDeptId(depIds)
	var deptAllOwnerIds []int
	core.DB.Model(&model.UserDepartment{}).Where("id IN (?)", deptAllIds).Where("parent_id > 0").Pluck("owner_userid", &deptAllOwnerIds)

	var count int64
	db := core.DB.Model(&model.UserDepartment{}).
		Where("owner_userid = ?", user.Userid)

	if common.InArrayInt(kr.Userid, deptAllOwnerIds) {
		db = db.Where("id IN (?)", deptAllIds)
		db = db.Where("parent_id = 0")
	} else {
		db = db.Where("id IN (?)", depIds)
		// db = db.Where("parent_id > 0") // 注释后，自己添加OKR时的上级都可评分
	}

	if err := db.Count(&count).Error; err != nil {
		return false
	}

	if count > 0 {
		// 负责人评分 1.不是超管 2.kr负责人不是当前用户 3.当前用户是部门负责人
		if !user.IsAdmin() && kr.Userid != user.Userid {
			return true
		}
	}

	return false
}

// 获取目标评分上级用户ids
func (s *okrService) GetSuperiorUserIds(obj *model.Okr, userid int) []int {
	var userids []int

	// 如果是超管评分，则返回所有超管用户
	var topDepartment model.UserDepartment
	core.DB.Model(&model.UserDepartment{}).Where("parent_id = 0").Where("owner_userid = ?", obj.Userid).First(&topDepartment)
	if topDepartment.Id > 0 {
		core.DB.Model(&model.User{}).Where("identity LIKE ?", "%,admin,%").Pluck("userid", &userids)
		return userids
	}

	// 目标负责人的部门
	depIds := common.ExplodeInt(",", obj.DepartmentId, true)
	if len(depIds) == 0 {
		return nil
	}

	// 获取部门小组负责人
	deptAllIds, _ := s.GetDepartmentsBySearchDeptId(depIds)
	var deptAllOwnerIds []int
	core.DB.Model(&model.UserDepartment{}).Where("id IN (?)", deptAllIds).Where("parent_id > 0").Pluck("owner_userid", &deptAllOwnerIds)

	db := core.DB.Model(&model.UserDepartment{})

	if common.InArrayInt(obj.Userid, deptAllOwnerIds) {
		db = db.Where("id IN (?)", deptAllIds)
		db = db.Where("parent_id = 0")
	} else {
		db = db.Where("id IN (?)", depIds)
		// db = db.Where("parent_id > 0") // 注释后，自己添加OKR时的上级都可评分
	}

	if err := db.Pluck("DISTINCT owner_userid", &userids).Error; err != nil {
		return nil
	}

	return userids
}

// 取消/重启目标
func (s *okrService) CancelObjective(userid, okrId int) (*model.Okr, error) {
	kr, err := s.GetObjectiveById(okrId)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	if kr.Userid != userid {
		return nil, e.New(constant.ErrOkrOwnerNotCancel)
	}

	// 更新取消状态
	var record interfaces.OkrLogParams
	if kr.Canceled == 0 {
		kr.Canceled = 1
		record = interfaces.OkrLogParams{
			StatusChange: []string{model.CanceledMap[0], model.CanceledMap[1]},
		}

	} else if kr.Canceled == 1 {
		kr.Canceled = 0
		record = interfaces.OkrLogParams{
			StatusChange: []string{model.CanceledMap[1], model.CanceledMap[0]},
		}
	}

	if err := core.DB.Save(kr).Error; err != nil {
		return nil, err
	}

	// 日志
	if err := s.InsertOkrLogTx(core.DB, kr.Id, userid, "update", "修改O目标状态", record); err != nil {
		return nil, err
	}

	return kr, nil
}

// 更新参与人
func (s *okrService) UpdateParticipant(user *interfaces.UserInfoResp, param interfaces.OkrParticipantUpdateReq) (*model.Okr, error) {
	kr, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	err = s.CheckObjectiveOperation(kr, user.Userid)
	if err != nil {
		return nil, err
	}

	// 去掉Participant中的0
	if strings.Contains(param.Participant, "0,") {
		param.Participant = strings.Trim(param.Participant, "0,")
	}

	oldParticipant := common.ExplodeInt(",", kr.Participant, true)
	newParticipant := common.ExplodeInt(",", param.Participant, true)
	diffParticipant := common.ArrayDifferenceProcessing(newParticipant, oldParticipant)

	// 为KR添加新参与人时，发送提示消息
	if len(diffParticipant) > 0 {
		if err := s.InsertOkrLogTx(core.DB, kr.ParentId, user.Userid, "update", "修改KR参与人", nil); err != nil {
			return nil, err
		}

		addDiffParticipant := common.ArrayDifferenceAddProcessing(newParticipant, oldParticipant)
		if len(addDiffParticipant) > 0 {
			obj, err := s.GetObjectiveByIdWithKeyResults(kr.ParentId)
			if err != nil {
				return nil, e.New(constant.ErrOkrNoData)
			}
			kr.ParentTitle = obj.Title
			go DootaskService.DialogOkrPush(kr, user.Token, 4, diffParticipant)
		}
	}

	kr.Participant = param.Participant
	if err := core.DB.Save(kr).Error; err != nil {
		return nil, err
	}

	return kr, nil
}

// 更新信心指数
func (s *okrService) UpdateConfidence(userid int, param interfaces.OkrConfidenceUpdateReq) (*model.Okr, error) {
	kr, err := s.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}

	err = s.CheckObjectiveOperation(kr, userid)
	if err != nil {
		return nil, err
	}

	if kr.Confidence != param.Confidence {
		if err := s.InsertOkrLogTx(core.DB, kr.ParentId, userid, "update", "修改KR信心指数", interfaces.OkrLogParams{
			Title:            kr.Title,
			ConfidenceChange: []int{kr.Confidence, param.Confidence},
		}); err != nil {
			return nil, err
		}
	}

	kr.Confidence = param.Confidence
	if err := core.DB.Save(kr).Error; err != nil {
		return nil, err
	}

	return kr, nil
}

// 检查目标是否可以操作
func (s *okrService) CheckObjectiveOperation(okr *model.Okr, userid int) error {
	// 以下只能O的负责人可操作（其他人仅能查看）
	// 详情：编辑OKR、修改参与人、取消目标、重启目标、修改进度、信心、对齐目标、添加复盘
	if okr.Userid != userid {
		return e.New(constant.ErrOkrOwnerNotCancel)
	}
	if okr.ParentId == 0 {
		// O已取消
		if okr.Canceled == 1 {
			return e.New(constant.ErrOkrCanceled)
		}
		// O已完成
		if okr.Completed == 1 {
			return e.New(constant.ErrOkrCompleted)
		}
	} else {
		// 父级O已取消
		var parentOkr model.Okr
		if err := core.DB.Where("id = ?", okr.ParentId).First(&parentOkr).Error; err != nil {
			return err
		}
		if parentOkr.Canceled == 1 {
			return e.New(constant.ErrOkrCanceled)
		}
		// 父级O已完成
		if parentOkr.Completed == 1 {
			return e.New(constant.ErrOkrCompleted)
		}

		if okr.Score > -1 || okr.SuperiorScore > -1 {
			return e.New(constant.ErrOkrScored)
		}
	}
	return nil
}

// 创建 OKR 复盘记录
func (s *okrService) CreateOkrReplay(userid int, req interfaces.OkrReplayCreateReq) (*model.OkrReplay, error) {
	// 获取 OKR 目标及其关键结果
	obj, err := s.GetObjectiveByIdWithKeyResults(req.OkrId)
	if err != nil {
		return nil, err
	}

	if obj.Userid != userid {
		return nil, e.New(constant.ErrOkrOwnerNotCancel)
	}

	// 检查关键结果是否已评分
	for _, kr := range obj.KeyResults {
		if kr.Score == -1 || kr.SuperiorScore == -1 {
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
		Problem:         req.Problem,
	}

	// 创建 KR 复盘历史记录
	var histories []*model.OkrReplayHistory
	commentsMap := make(map[int]interfaces.OkrReplayComment)
	for _, comment := range req.Comments {
		// 复盘评论判断
		if !common.InArrayInt(comment.Comment, []int{1, 2}) {
			return nil, e.New(constant.ErrOkrReplayCommentInvalid)
		}
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
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, e.New(constant.ErrOkrNoData)
		}
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
	var alias []string
	if departmentId != "" && ascription == 1 {
		departmentIds := common.ExplodeInt(",", departmentId, true)
		if len(departmentIds) == 0 {
			return nil
		}

		// 遍历departmentId获取顶级的部门名称，及parent_id=0的部门名称
		topLevelDepartments, err := s.GetUniqueTopLevelDepartments(departmentIds)
		if err != nil {
			return nil
		}

		// 获取所有顶级部门名称
		for _, department := range topLevelDepartments {
			alias = append(alias, department.Name)
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

// 通过顶级部门查询所有子部门
func (s *okrService) GetDepartmentsBySearchDeptId(ids []int) ([]int, error) {
	var departments []int

	depts, err := s.GetUniqueTopLevelDepartments(ids)
	if err != nil {
		return nil, err
	}

	for _, dept := range depts {
		departments = append(departments, dept.Id)
	}

	allDeptIds, err := s.GetDepartmentsByTopLevelIds(departments)
	if err != nil {
		return nil, err
	}

	return allDeptIds, nil
}

// 获取顶级部门
func (s *okrService) GetUniqueTopLevelDepartments(ids []int) ([]model.UserDepartment, error) {
	var departments []model.UserDepartment

	// 查询所有子部门
	if err := core.DB.Where("id IN (?)", ids).Preload("ParentDepartment").Find(&departments).Error; err != nil {
		return nil, err
	}

	// 获取所有顶级部门
	encountered := make(map[int]bool)
	var uniqueTopLevelDepartments []model.UserDepartment
	for _, department := range departments {
		// 判断是否为顶级部门
		if department.ParentDepartment == nil {
			uniqueTopLevelDepartments = append(uniqueTopLevelDepartments, department)
			encountered[department.Id] = true
			continue
		}

		// 找到顶级部门
		topDepartment := department
		for topDepartment.ParentDepartment != nil {
			topDepartment = *topDepartment.ParentDepartment
		}

		// 判断是否已经添加过
		if !encountered[topDepartment.Id] {
			uniqueTopLevelDepartments = append(uniqueTopLevelDepartments, topDepartment)
			encountered[topDepartment.Id] = true
		}
	}

	return uniqueTopLevelDepartments, nil
}

// 通过顶级部门ids获取所有子部门信息包括顶级部门
func (s *okrService) GetDepartmentsByTopLevelIds(ids []int) ([]int, error) {
	var departments []int

	// 查询所有子部门
	if err := core.DB.Model(&model.UserDepartment{}).Where("parent_id IN (?)", ids).Pluck("id", &departments).Error; err != nil {
		return nil, err
	}

	// 包括顶级部门ids
	departments = common.ArrayUniqueInt(append(departments, ids...))

	return departments, nil
}

// UpdateAlignObjective 更新对齐目标
func (s *okrService) UpdateAlignObjective(userid, okrId int, alignObjective string) (*model.Okr, error) {
	obj, err := s.GetObjectiveById(okrId)
	if err != nil {
		return nil, err
	}

	err = s.CheckObjectiveOperation(obj, userid)
	if err != nil {
		return nil, err
	}

	if err := s.updateAlignment(obj, userid, alignObjective, false); err != nil {
		return nil, err
	}

	return obj, nil
}

// 取消对齐目标
func (s *okrService) CancelAlignObjective(userid, okrId, alignOkrId int) error {
	obj, err := s.GetObjectiveById(okrId)
	if err != nil {
		return err
	}

	err = s.CheckObjectiveOperation(obj, userid)
	if err != nil {
		return err
	}

	var align model.OkrAlign
	if err := core.DB.Where("okr_id = ? and align_okr_id = ?", okrId, alignOkrId).First(&align).Error; err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return e.New(constant.ErrOkrNoData)
		}
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
func (s *okrService) InsertOkrLogTx(tx *gorm.DB, okrId, userId int, operation, content string, record any) error {
	var recordJson string
	if record == nil {
		recordJson = ""
	} else {
		recordJson = common.StructToJson(record)
	}

	log := &model.OkrLog{
		OkrId:     okrId,
		Userid:    userId,
		Operation: operation,
		Content:   content,
		Record:    recordJson,
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

		if log.Record != "" {
			log.Records, _ = common.StrToMap(log.Record)
		}
	}

	return interfaces.PaginationRsp(page, pageSize, count, logs), nil
}

// 获取部门列表
func (s *okrService) GetDepartmentSearch(page, pageSize int) (*interfaces.Pagination, error) {
	var departments []*model.UserDepartment
	var count int64

	query := core.DB.Model(&model.UserDepartment{}).Where("parent_id = 0")
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&departments).Error; err != nil {
		return nil, err
	}

	return interfaces.PaginationRsp(page, pageSize, count, departments), nil
}
