package common

import (
	"FastFiber/global"
	"FastFiber/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug   bool
	Preload []string // 新增预加载字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{
			Logger: global.MysqlLog,
		})
	}

	query := DB.Model(&model).Where(&model)

	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	// 获取总记录数（优化后的方式）
	query.Count(&count)

	// 仅在需要分页时应用Limit/Offset
	if option.Limit > 0 {
		offset := (option.Page - 1) * option.Limit
		if offset < 0 {
			offset = 0
		}
		query = query.Limit(option.Limit).Offset(offset)
	}

	// 添加排序条件
	if option.Sort == "" {
		query = query.Order("created_at DESC")
	} else {
		query = query.Order(option.Sort)
	}

	err = query.Find(&list).Error
	return list, count, err
}
