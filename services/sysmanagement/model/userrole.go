// Package model
// @File    : userrole.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 15:07
// @Desc    :
package model

import (
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	UserRoleModel interface {
		WithTrans(ctx context.Context) UserRoleModel
		Insert(ctx context.Context, arg *UserRole) error
		BatchInsert(ctx context.Context, args []*UserRole) error
		DeleteByUserID(ctx context.Context, userID string) error
		DeleteByRoleID(ctx context.Context, roleID string) error
		ListByUserID(ctx context.Context, userID string) ([]*string, error)
		CountByRoleID(ctx context.Context, roleID string) (int64, error)
	}

	UserRole struct {
		ID        string `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`         // 用户角色ID
		UserID    string `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"` // 用户ID
		RoleID    string `json:"role_id" gorm:"column:role_id;type:varchar(32);not null"` // 角色ID
		CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`                  // 创建时间
	}

	defaultUserRoleModel struct {
		db *gorm.DB
	}
)

func NewUserRoleModel(isMigration bool, db *gorm.DB) UserRoleModel {
	if isMigration {
		if err := db.AutoMigrate(&UserRole{}); err != nil {
			panic(err)
		}
	}
	return &defaultUserRoleModel{db: db}
}

func (m *defaultUserRoleModel) WithTrans(ctx context.Context) UserRoleModel {
	return &defaultUserRoleModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultUserRoleModel) Insert(ctx context.Context, arg *UserRole) error {
	return m.db.Create(&arg).Error
}

func (m *defaultUserRoleModel) BatchInsert(ctx context.Context, args []*UserRole) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultUserRoleModel) DeleteByUserID(ctx context.Context, userID string) error {
	return m.db.Where("user_id = ?", userID).Delete(&UserRole{}).Error
}

func (m *defaultUserRoleModel) DeleteByRoleID(ctx context.Context, roleID string) error {
	return m.db.Where("role_id = ?", roleID).Delete(&UserRole{}).Error
}

func (m *defaultUserRoleModel) ListByUserID(ctx context.Context, userID string) ([]*string, error) {
	var roleIds []*string
	err := m.db.Where("user_id = ?", userID).Pluck("role_id", &roleIds).Error

	return roleIds, err
}

func (m *defaultUserRoleModel) CountByRoleID(ctx context.Context, roleID string) (int64, error) {
	var count int64
	err := m.db.Model(&UserRole{}).Where("role_id = ?", roleID).Count(&count).Error

	return count, err
}
