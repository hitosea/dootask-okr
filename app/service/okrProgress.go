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

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var OkrProgressService = NewOkrProgressService()

type okrProgressService struct {
	okrLogRepo repository.OkrLogRepository
}

func NewOkrProgressService() *okrProgressService {
	return &okrProgressService{
		okrLogRepo: repository.NewOkrLogRepository(),
	}
}

// 同步所有我(o)对齐的kr的进度
func (s *okrProgressService) SyncAllParentProgress(tx *gorm.DB, okrId, userid int) error {
	if tx == nil {
		tx = core.DB
	}
	// 1.判断我是否为o,并是否开启自动同步
	alignTable := core.DBTableName(&model.OkrAlign{})
	okrTable := core.DBTableName(&model.Okr{})
	var okr *model.Okr
	if err := tx.Where("id = ?", okrId).First(&okr).Error; err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return e.New(constant.ErrOkrNoData)
		}
		return err
	}
	if okr.ParentId != 0 {
		return e.New(constant.ErrInvalidParameter)
	}
	// 2.获取所有我对齐的kr
	var parentKrs []*model.Okr
	err := core.DB.Table(okrTable+" AS okrs").
		Select("okrs.*").
		Joins(fmt.Sprintf(`LEFT JOIN %s align ON okrs.id = align.align_okr_id`, alignTable)).
		Where("align.okr_id = ?", okr.Id).
		Where("okrs.parent_id > 0").
		Where("okrs.canceled = 0").
		Where("okrs.deleted_at IS NULL").
		Find(&parentKrs).Error
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return e.New(constant.ErrOkrNoData)
		}
		return err
	}
	// 更新所有我对齐的kr的进度
	for _, kr := range parentKrs {
		s.SyncKrProgress(nil, kr.Id, userid)
	}
	//
	return err
}

// 同步当前KR的进度
func (s *okrProgressService) SyncKrProgress(tx *gorm.DB, krId, userid int) (*model.Okr, error) {
	if tx == nil {
		tx = core.DB
	}
	//
	var kr *model.Okr
	err := tx.Transaction(func(tx *gorm.DB) error {
		// 1.判断我是否为kr并我的o是否开启自动同步
		alignTable := core.DBTableName(&model.OkrAlign{})
		okrTable := core.DBTableName(&model.Okr{})
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("ParentOKr").Where("id = ?", krId).First(&kr).Error; err != nil {
			if errors.Is(err, core.ErrRecordNotFound) {
				return e.New(constant.ErrOkrNoData)
			}
			return err
		}
		if kr.ParentId == 0 || kr.ParentOKr.AutoSync != 1 {
			return e.New(constant.ErrInvalidParameter)
		}
		// 2.获取所有对齐我的o, 得出进度
		var subsetOKrs []*model.Okr
		err := core.DB.Table(okrTable+" AS okrs").
			Select("okrs.*").
			Joins(fmt.Sprintf(`LEFT JOIN %s align ON okrs.id = align.okr_id`, alignTable)).
			Where("align.align_okr_id = ?", kr.Id).
			Where("okrs.canceled = 0").
			Where("okrs.deleted_at IS NULL").
			Find(&subsetOKrs).Error
		if err != nil {
			if errors.Is(err, core.ErrRecordNotFound) {
				return e.New(constant.ErrOkrNoData)
			}
			return err
		}
		progress := float64(0)
		cardinal := math.Round(float64(100)/float64(len(subsetOKrs))*100) / 100
		for _, okr := range subsetOKrs {
			progress += float64(cardinal * float64(float64(okr.Progress)/100))
		}
		// 3.更新进度
		user := interfaces.UserInfoResp{
			UserBasicResp: &interfaces.UserBasicResp{
				Userid: userid,
			},
		}
		s.UpdateProgressAndStatus(tx, &user, interfaces.OkrUpdateProgressReq{
			Id:       kr.Id,
			Progress: int(math.Floor(progress)),
			Status:   kr.Status,
		})
		//
		return nil
	})
	//
	return kr.ParentOKr, err
}

// 更新进度和进度状态
func (s *okrProgressService) UpdateProgressAndStatus(tx *gorm.DB, user *interfaces.UserInfoResp, param interfaces.OkrUpdateProgressReq) (*model.Okr, error) {
	kr, err := OkrService.GetObjectiveByIdIsKeyResult(param.Id)
	if err != nil {
		return nil, e.New(constant.ErrOkrNoData)
	}
	// 验证权限
	if user != nil && kr.AutoSync == 0 {
		err = OkrService.CheckObjectiveOperation(kr, user.Userid)
		if err != nil {
			return nil, err
		}
	}
	// 开始事务
	if tx == nil {
		tx = core.DB
	}
	err = tx.Transaction(func(tx *gorm.DB) error {
		// 如果传值更新进度有值，则更新进度
		if param.Progress != kr.Progress {
			if user != nil && kr.AutoSync == 0 {
				if err := OkrService.InsertOkrLogTx(tx, kr.ParentId, user.Userid, "update", "修改KR进度", interfaces.OkrLogParams{
					Title:          kr.Title,
					ProgressChange: []int{kr.Progress, param.Progress},
				}); err != nil {
					return err
				}
			}
			kr.Progress = param.Progress
		}

		// 如果传值更新状态有值，则更新状态
		if param.Status != kr.ProgressStatus {
			if user != nil && kr.AutoSync == 0 {
				if err := OkrService.InsertOkrLogTx(tx, kr.ParentId, user.Userid, "update", "修改KR状态", interfaces.OkrLogParams{
					Title:                kr.Title,
					ProgressStatusChange: []string{model.ProgressStatusMap[kr.ProgressStatus], model.ProgressStatusMap[param.Status]},
				}); err != nil {
					return err
				}
			}
			kr.ProgressStatus = param.Status
		}

		// 更新目标
		if err := tx.Save(kr).Error; err != nil {
			return err
		}

		// 检查所在目标的 KR 是否全部完成
		objWithKrs, err := OkrService.GetObjectiveByIdWithKeyResults(kr.ParentId)
		if err != nil {
			return err
		}
		krs := objWithKrs.KeyResults

		allCompleted := 1
		sumProgress := 0
		for _, item := range krs {
			OldProgressKr := item.Progress
			// 更新 KR 进度值
			if param.Id == item.Id {
				item.Progress = param.Progress
			}
			// 进度全部完成 100%
			if item.Progress < 100 {
				allCompleted = 0
			}
			// 计算总进度值
			sumProgress += item.Progress
			// 给参与人发送信息和新增动态信息（开启自动更新）
			if OldProgressKr != item.Progress && kr.AutoSync == 1 {
				item.ParentTitle = objWithKrs.Title
				// 给参与人发送信息
				textTypeKr := 13
				if allCompleted == 1 {
					textTypeKr = 15
				}
				ParticipantIds := common.ExplodeInt(",", item.Participant, true)
				ParticipantIds = append(ParticipantIds, item.Userid)
				// 去重
				go DootaskService.DialogOkrPush(item, "", textTypeKr, ParticipantIds)
				// 动态信息
				if err := OkrService.InsertOkrLogTx(tx, item.ParentId, user.Userid, "update", "进度自动更新", interfaces.OkrLogParams{
					Title:          item.Title,
					ProgressChange: []int{OldProgressKr, param.Progress},
				}); err != nil {
					return err
				}
			}
		}

		// 更新总目标进度值，kr 进度值相加/kr 数量
		oldProgressObj := objWithKrs.Progress
		progress := 0
		if len(krs) > 0 {
			progress = int(math.Floor(float64(sumProgress) / float64(len(krs))))
		}
		// 给负责人发送信息和新增动态信息 （开启自动更新）
		if oldProgressObj != progress && objWithKrs.AutoSync == 1 {
			textTypeObj := 12
			if allCompleted == 1 {
				textTypeObj = 14
			}
			go DootaskService.DialogOkrPush(objWithKrs, "", textTypeObj, []int{objWithKrs.Userid})
			if err := OkrService.InsertOkrLogTx(tx, objWithKrs.Id, user.Userid, "update", "进度自动更新", interfaces.OkrLogParams{
				Title:          objWithKrs.Title,
				ProgressChange: []int{oldProgressObj, progress},
			}); err != nil {
				return err
			}
		}

		// 更新总目标的状态是否完成
		if err := tx.Model(&model.Okr{}).Where("id = ?", kr.ParentId).Updates(map[string]interface{}{
			"Completed": allCompleted,
			"Progress":  progress,
		}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	//
	s.SyncAllParentProgress(tx, kr.ParentId, user.Userid)
	//
	return kr, nil
}
