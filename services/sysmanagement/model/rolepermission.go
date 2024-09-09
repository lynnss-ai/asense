// Package model
// @File    : rolepermission.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/9 11:12
// @Desc    :
package model

import (
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	RolePermissionModel interface {
		WithTrans(ctx context.Context) RolePermissionModel
		Insert(ctx context.Context, arg *RolePermission) error
		BatchInsert(ctx context.Context, args []*RolePermission) error
		DeleteByRoleID(ctx context.Context, roleID string) error
		DeleteByPermissionID(ctx context.Context, permissionID string) error
		ListByRoleID(ctx context.Context, roleID string) ([]*string, error)
		ListByRoleIds(ctx context.Context, roleIDs []*string) ([]*string, error)
		CountByRoleID(ctx context.Context, roleID string) (int64, error)
		CountByMenuID(ctx context.Context, menuID string) (int64, error)
	}

	RolePermission struct {
		ID        string `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`         // ID
		RoleID    string `json:"role_id" gorm:"column:role_id;type:varchar(32);not null"` // 角色ID
		MenuID    string `json:"menu_id" gorm:"column:menu_id;type:varchar(32);not null"` // 菜单ID
		CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`                  // 创建时间
		UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`                  // 更新时间
	}

	defaultRolePermissionModel struct {
		db *gorm.DB
	}
)

func NewRolePermissionModel(isMigration bool, db *gorm.DB) RolePermissionModel {
	if isMigration {
		if err := db.AutoMigrate(&RolePermission{}); err != nil {
			panic(err)
		}
	}
	return &defaultRolePermissionModel{db: db}
}

func (m *defaultRolePermissionModel) WithTrans(ctx context.Context) RolePermissionModel {
	return &defaultRolePermissionModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultRolePermissionModel) Insert(ctx context.Context, arg *RolePermission) error {
	return m.db.Create(&arg).Error
}

func (m *defaultRolePermissionModel) BatchInsert(ctx context.Context, args []*RolePermission) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultRolePermissionModel) DeleteByRoleID(ctx context.Context, roleID string) error {
	return m.db.Where("role_id = ?", roleID).Delete(&RolePermission{}).Error
}

func (m *defaultRolePermissionModel) DeleteByPermissionID(ctx context.Context, permissionID string) error {
	return m.db.Where("permission_id = ?", permissionID).Delete(&RolePermission{}).Error
}

func (m *defaultRolePermissionModel) ListByRoleID(ctx context.Context, roleID string) ([]*string, error) {
	var items []*string
	err := m.db.Model(&RolePermission{}).Where("role_id = ?", roleID).Pluck("menu_id", &items).Error

	return items, err
}

func (m *defaultRolePermissionModel) ListByRoleIds(ctx context.Context, roleIDs []*string) ([]*string, error) {
	var items []*string
	err := m.db.Model(&RolePermission{}).Where("role_id in (?)", roleIDs).Pluck("menu_id", &items).Error

	return items, err
}

func (m *defaultRolePermissionModel) CountByRoleID(ctx context.Context, roleID string) (int64, error) {
	var total int64
	err := m.db.Model(&RolePermission{}).Where("role_id = ?", roleID).Count(&total).Error

	return total, err
}

func (m *defaultRolePermissionModel) CountByMenuID(ctx context.Context, menuID string) (int64, error) {
	var total int64
	err := m.db.Model(&RolePermission{}).Where("menu_id = ?", menuID).Count(&total).Error

	return total, err
}
