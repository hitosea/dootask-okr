package dbs

import (
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"

	"gorm.io/gorm"
)

// 获取分页列表
func GetPageList(db *gorm.DB, page, pageSize int, data any) *interfaces.Pagination {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 50
	}
	var count int64
	if err := db.Session(&core.Session).Offset((page - 1) * pageSize).Limit(pageSize).Find(data).Count(&count).Error; err != nil {
		return nil
	}
	return interfaces.PaginationRsp(page, pageSize, count, data)
}
