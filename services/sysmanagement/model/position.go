// Package model
// @File    : position.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/27 11:47
// @Desc    : 职位
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	PositionModel interface {
		WithTrans(ctx context.Context) PositionModel
		Insert(ctx context.Context, arg *Position) error
		BatchInsert(ctx context.Context, args []*Position) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id string) error
		FindOne(ctx context.Context, id string) (*Position, error)
		List(ctx context.Context, filter *string, isEnable *bool) ([]*Position, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*Position, error)
		Count(ctx context.Context) (int64, error)
		Enable(ctx context.Context, id string) error
	}

	Position struct {
		ID           string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                     // 职位ID
		PositionName string         `json:"position_name" gorm:"column:position_name;type:varchar(64);not null"` // 职位名称
		PositionCode string         `json:"position_code" gorm:"column:position_code;type:varchar(32);not null"` // 职位编码
		PositionDesc string         `json:"position_desc" gorm:"column:position_desc;type:varchar(255)"`         // 职位描述
		IsEnable     bool           `json:"is_enable" gorm:"column:is_enable;type:bool;not null"`                // 是否启用[true:启用,false:禁用]
		CreatedAt    int64          `json:"created_at" gorm:"autoCreateTime:milli"`                              // 创建时间
		UpdatedAt    int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`                              // 更新时间
		DeletedAt    gorm.DeletedAt `json:"deleted_at"`                                                          // 删除时间
	}

	defaultPositionModel struct {
		db *gorm.DB
	}
)

func NewPositionModel(isMigration bool, db *gorm.DB) PositionModel {
	if isMigration {
		if err := db.AutoMigrate(&Position{}); err != nil {
			panic(err)
		}
	}
	return &defaultPositionModel{db: db}
}

func (m *defaultPositionModel) WithTrans(ctx context.Context) PositionModel {
	return &defaultPositionModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultPositionModel) Insert(ctx context.Context, arg *Position) error {
	return m.db.Create(&arg).Error
}

func (m *defaultPositionModel) BatchInsert(ctx context.Context, args []*Position) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultPositionModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Position{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultPositionModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&Position{}, id).Error
}

func (m *defaultPositionModel) FindOne(ctx context.Context, id string) (*Position, error) {
	var result *Position
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultPositionModel) List(ctx context.Context, filter *string, isEnable *bool) ([]*Position, error) {
	var items []*Position
	query := m.db.Model(&Position{})
	if filter != nil {
		query = query.Where("position_name like ?", *filter)
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", *isEnable)
	}
	err := query.Find(&items).Error

	return items, err
}

func (m *defaultPositionModel) FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*Position, error) {
	var (
		total int64
		items []*Position
	)
	p := components.PageHandle(page, pageSize, filter)
	query := m.db.Model(&Position{})
	if filter != nil {
		query = query.Where("position_name like ?", p.Filter)
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", *isEnable)
	}
	err := query.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	err = query.Offset(p.Page).Limit(p.PageSize).Order("created_at desc").Find(&items).Error
	if err != nil {
		return 0, nil, err
	}

	return total, items, nil
}

func (m *defaultPositionModel) Count(ctx context.Context) (int64, error) {
	var count int64
	err := m.db.Model(&Position{}).Count(&count).Error

	return count, err
}

func (m *defaultPositionModel) Enable(ctx context.Context, id string) error {
	var result *Position
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&Position{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
