// Package model
// @File    : organizationuser.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/27 09:48
// @Desc    : 组织机构用户
package model

import (
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	OrganizationUserModel interface {
		WithTrans(ctx context.Context) OrganizationUserModel
		Insert(ctx context.Context, arg *OrganizationUser) error
		BatchInsert(ctx context.Context, args []*OrganizationUser) error
		DeleteByOrganizationID(ctx context.Context, organizationID string) error
		DeleteByUserID(ctx context.Context, userID string) error
		ListByUserID(ctx context.Context, userID string) ([]*OrganizationUser, error)
		ListByUserIDToOrgID(ctx context.Context, userID string) ([]*string, error)
		ListByOrganizationID(ctx context.Context, organizationID string) ([]*OrganizationUser, error)
		ListByOrganizationIDToUserID(ctx context.Context, organizationID string) ([]*string, error)
		ExistByOrganizationIDAndUserID(ctx context.Context, organizationID, userID string) (bool, error)
		CountByOrganizationID(ctx context.Context, organizationID string) (int64, error)
		CountByUserID(ctx context.Context, userID string) (int64, error)
	}

	OrganizationUser struct {
		ID             string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                         // 组织架构用户ID
		OrganizationID string         `json:"organization_id" gorm:"column:organization_id;type:varchar(32);not null"` // 组织架构ID
		UserID         string         `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"`                 // 用户ID
		PositionID     string         `json:"position_id" gorm:"column:position_id;type:varchar(32)"`                  // 职位ID
		IsDeptManager  bool           `json:"is_dept_manager" gorm:"column:is_dept_manager;type:bool;not null"`        // 是否部门负责人[true:是,false:否]
		CreatedAt      int64          `json:"created_at" gorm:"autoCreateTime:milli"`                                  // 创建时间
		UpdatedAt      int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`                                  // 更新时间
		DeletedAt      gorm.DeletedAt `json:"deleted_at"`                                                              // 删除时间
	}

	defaultOrganizationUserModel struct {
		db *gorm.DB
	}
)

func NewOrganizationUserModel(isMigration bool, db *gorm.DB) OrganizationUserModel {
	if isMigration {
		if err := db.AutoMigrate(&OrganizationUser{}); err != nil {
			panic(err)
		}
	}
	return &defaultOrganizationUserModel{db: db}
}

func (m *defaultOrganizationUserModel) WithTrans(ctx context.Context) OrganizationUserModel {
	return &defaultOrganizationUserModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultOrganizationUserModel) Insert(ctx context.Context, arg *OrganizationUser) error {
	return m.db.Create(&arg).Error
}

func (m *defaultOrganizationUserModel) BatchInsert(ctx context.Context, args []*OrganizationUser) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultOrganizationUserModel) DeleteByOrganizationID(ctx context.Context, organizationID string) error {
	return m.db.Delete(&OrganizationUser{}, "organization_id = ?", organizationID).Error
}

func (m *defaultOrganizationUserModel) DeleteByUserID(ctx context.Context, userID string) error {
	return m.db.Delete(&OrganizationUser{}, "user_id = ?", userID).Error
}

func (m *defaultOrganizationUserModel) ListByUserID(ctx context.Context, userID string) ([]*OrganizationUser, error) {
	var items []*OrganizationUser
	err := m.db.Where("user_id = ?", userID).Find(&items).Error

	return items, err
}

func (m *defaultOrganizationUserModel) ListByUserIDToOrgID(ctx context.Context, userID string) ([]*string, error) {
	var orgIds []*string
	err := m.db.Where("user_id = ?", userID).Pluck("organization_id", &orgIds).Error

	return orgIds, err
}

func (m *defaultOrganizationUserModel) ListByOrganizationID(ctx context.Context, organizationID string) ([]*OrganizationUser, error) {
	var items []*OrganizationUser
	err := m.db.Where("organization_id = ?", organizationID).Find(&items).Error

	return items, err
}

func (m *defaultOrganizationUserModel) ListByOrganizationIDToUserID(ctx context.Context, organizationID string) ([]*string, error) {
	var userIds []*string
	err := m.db.Where("organization_id = ?", organizationID).Pluck("user_id", &userIds).Error

	return userIds, err
}

func (m *defaultOrganizationUserModel) ExistByOrganizationIDAndUserID(ctx context.Context, organizationID, userID string) (bool, error) {
	var count int64
	err := m.db.Model(&OrganizationUser{}).Where("organization_id = ? and user_id = ?", organizationID, userID).Count(&count).Error

	return count > 0, err
}

func (m *defaultOrganizationUserModel) CountByOrganizationID(ctx context.Context, organizationID string) (int64, error) {
	var count int64
	err := m.db.Model(&OrganizationUser{}).Where("organization_id = ?", organizationID).Count(&count).Error
	return count, err
}

func (m *defaultOrganizationUserModel) CountByUserID(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := m.db.Model(&OrganizationUser{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
