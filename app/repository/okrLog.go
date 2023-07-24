package repository

import (
	"dootask-okr/app/core"
	"dootask-okr/app/model"

	"gorm.io/gorm"
)

type OkrLogRepository interface {
	CreateTx(tx *gorm.DB, log *model.OkrLog) error
	CountByOkrId(okrId int) (int64, error)
	FindByOkrId(okrId, page, pageSize int) ([]*model.OkrLog, error)
}

type okrLogRepository struct{}

func NewOkrLogRepository() OkrLogRepository {
	return &okrLogRepository{}
}

// tx创建OKR日志
func (r *okrLogRepository) CreateTx(tx *gorm.DB, log *model.OkrLog) error {
	return tx.Create(log).Error
}

// 统计某个OKR的日志数量
func (r *okrLogRepository) CountByOkrId(okrId int) (int64, error) {
	var count int64
	if err := core.DB.Model(&model.OkrLog{}).Where("okr_id = ?", okrId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 获取某个OKR的日志列表
func (r *okrLogRepository) FindByOkrId(okrId, page, pageSize int) ([]*model.OkrLog, error) {
	var logs []*model.OkrLog
	offset := (page - 1) * pageSize
	if err := core.DB.Model(&model.OkrLog{}).Where("okr_id = ?", okrId).Order("created_at asc").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
