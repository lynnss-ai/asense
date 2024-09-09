// Package model
// @File    : dictionary.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 15:07
// @Desc    :
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	DictionaryModel interface {
		WithTrans(ctx context.Context) DictionaryModel
		Insert(ctx context.Context, arg *Dictionary) error
		BatchInsert(ctx context.Context, args []*Dictionary) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id string) error
		FindOne(ctx context.Context, id string) (*Dictionary, error)
		FindOneByDicCode(ctx context.Context, dicCode string) (*Dictionary, error)
		List(ctx context.Context, pid *string, isEnable *bool, isEdit *bool, isHide *bool, filter *string) ([]*Dictionary, error)
		CountByPid(ctx context.Context, id *string) (int64, error)
		ExistByDicCode(ctx context.Context, dicCode string) (bool, error)
		Enable(ctx context.Context, id string) error
	}

	Dictionary struct {
		ID        string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`           // ID
		PID       string         `json:"pid" gorm:"column:pid;type:varchar(32);not null"`           // 父ID[根节点模型为字符串0]
		DicName   string         `json:"dic_name" gorm:"column:dic_name;type:varchar(32);not null"` // 字典名称
		DicCode   string         `json:"dic_code" gorm:"column:dic_code;type:varchar(32);not null"` // 字典编码
		DicValue  string         `json:"dic_value" gorm:"column:dic_value;type:varchar(2048)"`      // 字典值
		DicValue2 string         `json:"dic_value2" gorm:"column:dic_value2;type:varchar(2048)"`    // 字典值2
		DicValue3 string         `json:"dic_value3" gorm:"column:dic_value3;type:varchar(2048)"`    // 字典值3
		DicDesc   string         `json:"dic_desc" gorm:"column:dic_desc;type:varchar(255)"`         // 字典描述
		Sort      int            `json:"sort" gorm:"column:sort;type:int;not null"`                 // 排序
		IsEnable  bool           `json:"is_enable" gorm:"column:is_enable;type:bool;not null"`      // 是否启用[true:启用,false:禁用]
		IsEdit    bool           `json:"is_edit" gorm:"column:is_edit;type:bool;not null"`          // 是否可编辑[true:可编辑,false:不可编辑]
		IsHide    bool           `json:"is_hide" gorm:"column:is_hide;type:bool;not null"`          // 是否隐藏[true:隐藏,false:显示]
		CreatedAt int64          `json:"created_at" gorm:"autoCreateTime:milli"`                    // 创建时间
		UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`                    // 更新时间
		DeletedAt gorm.DeletedAt `json:"deleted_at"`                                                // 删除时间
	}

	defaultDictionaryModel struct {
		db *gorm.DB
	}
)

func NewDictionaryModel(isMigration bool, db *gorm.DB) DictionaryModel {
	if isMigration {
		if err := db.AutoMigrate(&Dictionary{}); err != nil {
			panic(err)
		}
	}
	return &defaultDictionaryModel{db: db}
}

func (m *defaultDictionaryModel) WithTrans(ctx context.Context) DictionaryModel {
	return &defaultDictionaryModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultDictionaryModel) Insert(ctx context.Context, arg *Dictionary) error {
	return m.db.Create(&arg).Error
}

func (m *defaultDictionaryModel) BatchInsert(ctx context.Context, args []*Dictionary) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultDictionaryModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Dictionary{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultDictionaryModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&Dictionary{}, id).Error
}

func (m *defaultDictionaryModel) FindOne(ctx context.Context, id string) (*Dictionary, error) {
	var result *Dictionary
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultDictionaryModel) FindOneByDicCode(ctx context.Context, dicCode string) (*Dictionary, error) {
	var result *Dictionary
	err := m.db.Where("dic_code = ?", dicCode).First(&result).Error

	return result, err
}

func (m *defaultDictionaryModel) List(ctx context.Context, pid *string, isEnable *bool, isEdit *bool, isHide *bool, filter *string) ([]*Dictionary, error) {
	var items []*Dictionary
	query := m.db.Model(&Dictionary{})
	if pid != nil {
		query = query.Where("pid = ?", *pid)
	} else {
		query = query.Where("pid = ?", "0")
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", *isEnable)
	}
	if isEdit != nil {
		query = query.Where("is_edit = ?", *isEdit)
	}
	if isHide != nil {
		query = query.Where("is_hide = ?", *isHide)
	}

	if filter != nil {
		_filter := components.Filter(*filter)
		query = query.Where("dic_name like ? OR dic_value like ?", _filter, _filter)
	}

	err := query.Order("sort").Find(&items).Error

	return items, err
}

func (m *defaultDictionaryModel) CountByPid(ctx context.Context, id *string) (int64, error) {
	var count int64
	query := m.db.Model(&Dictionary{})
	if id != nil {
		query = query.Where("pid = ?", *id)
	}

	err := query.Count(&count).Error

	return count, err
}

func (m *defaultDictionaryModel) ExistByDicCode(ctx context.Context, dicCode string) (bool, error) {
	var count int64
	err := m.db.Model(&Dictionary{}).Where("dic_code = ?", dicCode).Count(&count).Error

	return count > 0, err
}

func (m *defaultDictionaryModel) Enable(ctx context.Context, id string) error {
	var result *Dictionary
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&Dictionary{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
