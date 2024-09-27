// Package model
// @File    : attachment.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/27 11:54
// @Desc    : 附件
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	AttachmentModel interface {
		WithTrans(ctx context.Context) AttachmentModel
		Insert(ctx context.Context, arg *Attachment) error
		BatchInsert(ctx context.Context, args []*Attachment) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id string) error
		FindOne(ctx context.Context, id string) (*Attachment, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*Attachment, error)
		Count(ctx context.Context) (int64, error)
		Enable(ctx context.Context, id string) error
	}

	Attachment struct {
		ID        string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`      // 附件ID
		FileName  string         `json:"file_name" gorm:"column:file_name;type:varchar(255)"`  // 文件名
		FileSize  int64          `json:"file_size" gorm:"column:file_size;type:bigint"`        // 文件大小
		FileType  string         `json:"file_type" gorm:"column:file_type;type:varchar(32)"`   // 文件类型
		FilePath  string         `json:"file_path" gorm:"column:file_path;type:varchar(255)"`  // 文件路径
		IsEnable  bool           `json:"is_enable" gorm:"column:is_enable;type:bool;not null"` // 是否启用[true:启用,false:禁用]
		CreatedAt int64          `json:"created_at" gorm:"autoCreateTime:milli"`               // 创建时间
		UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`               // 更新时间
		DeletedAt gorm.DeletedAt `json:"deleted_at"`                                           // 删除时间
	}

	defaultAttachmentModel struct {
		db *gorm.DB
	}
)

func NewAttachmentModel(isMigration bool, db *gorm.DB) AttachmentModel {
	if isMigration {
		if err := db.AutoMigrate(&Attachment{}); err != nil {
			panic(err)
		}
	}
	return &defaultAttachmentModel{db: db}
}

func (m *defaultAttachmentModel) WithTrans(ctx context.Context) AttachmentModel {
	return &defaultAttachmentModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultAttachmentModel) Insert(ctx context.Context, arg *Attachment) error {
	return m.db.Create(&arg).Error
}

func (m *defaultAttachmentModel) BatchInsert(ctx context.Context, args []*Attachment) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultAttachmentModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Attachment{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultAttachmentModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&Attachment{}, id).Error
}

func (m *defaultAttachmentModel) FindOne(ctx context.Context, id string) (*Attachment, error) {
	var result *Attachment
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultAttachmentModel) FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*Attachment, error) {
	var (
		total int64
		items []*Attachment
	)
	p := components.PageHandle(page, pageSize, filter)
	query := m.db.Model(&Attachment{})

	if filter != nil {
		query = query.Where("file_name like ?", p.Filter)
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

func (m *defaultAttachmentModel) Count(ctx context.Context) (int64, error) {
	var count int64
	err := m.db.Model(&Attachment{}).Count(&count).Error

	return count, err
}

func (m *defaultAttachmentModel) Enable(ctx context.Context, id string) error {
	var result *Attachment
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&Attachment{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
