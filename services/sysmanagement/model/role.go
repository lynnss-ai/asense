// Package model
// @File    : role.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 15:10
// @Desc    :
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

type (
	RoleModel interface {
		WithTrans(ctx context.Context) RoleModel
		Insert(ctx context.Context, arg *Role) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id string) error
		FindOne(ctx context.Context, id string) (*Role, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*Role, error)
		ListBySetPermission(ctx context.Context, filter *string, isEnable *bool) ([]*Role, error)
		ListByIds(ctx context.Context, ids []*string) ([]*Role, error)
		ListByIdsToIds(ctx context.Context, ids []*string) ([]*string, error)
		ExistByRoleCode(ctx context.Context, roleCode string) (bool, error)
		Enable(ctx context.Context, id string) error
	}

	Role struct {
		ID              string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                      // ID
		RoleName        string         `json:"role_name" gorm:"column:role_name;type:varchar(32);not null"`          // 角色名称
		RoleCode        string         `json:"role_code" gorm:"column:role_code;type:varchar(32);not null"`          // 角色编码
		RoleDesc        string         `json:"role_desc" gorm:"column:role_desc;type:varchar(255)"`                  // 角色描述
		IsSetPermission bool           `json:"is_set_permission" gorm:"column:is_set_permission;type:bool;not null"` // 是否设置权限[true:设置,false:未设置]
		SelectMenuIds   string         `json:"select_menu_ids" gorm:"column:select_menu_ids;type:varchar(2048)"`     // 选择的菜单ID
		IsEnable        bool           `json:"is_enable" gorm:"column:is_enable;type:bool;not null"`                 // 是否启用[true:启用,false:禁用]
		IsAdmin         bool           `json:"is_admin" gorm:"column:is_admin;type:bool;not null"`                   // 是否管理员[true:是,false:否]
		CreatedAt       int64          `json:"created_at" gorm:"autoCreateTime:milli"`                               // 创建时间
		UpdatedAt       int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`                               // 更新时间
		DeletedAt       gorm.DeletedAt `json:"deleted_at"`                                                           // 删除时间
	}

	defaultRoleModel struct {
		db *gorm.DB
	}
)

func NewRoleModel(isMigration bool, db *gorm.DB) RoleModel {
	if isMigration {
		if err := db.AutoMigrate(&Role{}); err != nil {
			panic(err)
		}
	}
	return &defaultRoleModel{db: db}
}

func (m *defaultRoleModel) WithTrans(ctx context.Context) RoleModel {
	return &defaultRoleModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultRoleModel) Insert(ctx context.Context, arg *Role) error {
	return m.db.Create(&arg).Error
}

func (m *defaultRoleModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Role{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultRoleModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&Role{}, id).Error
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id string) (*Role, error) {
	var result *Role
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultRoleModel) FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*Role, error) {
	var (
		total int64
		items []*Role
	)
	p := components.PageHandle(page, pageSize, filter)
	query := m.db.Model(&Role{})
	if filter != nil {
		query = query.Where("role_name like ? OR role_code like ?", p.Filter, p.Filter)
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", isEnable)
	}
	err := query.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = query.Limit(p.PageSize).Offset(p.Page).Order("created_at desc").Find(&items).Error

	return total, items, err
}

func (m *defaultRoleModel) ListBySetPermission(ctx context.Context, filter *string, isEnable *bool) ([]*Role, error) {
	var items []*Role
	query := m.db.Model(&Role{}).Where("is_set_permission = ?", true)
	if filter != nil {
		_filter := components.Filter(*filter)
		query = query.Where("role_name like ? OR role_code like ?", _filter, _filter)
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", isEnable)
	}
	err := query.Order("created_at desc").Find(&items).Error

	return items, err
}

func (m *defaultRoleModel) ListByIds(ctx context.Context, ids []*string) ([]*Role, error) {
	var items []*Role
	err := m.db.Where("id in (?)", ids).Find(&items).Error

	return items, err
}

func (m *defaultRoleModel) ListByIdsToIds(ctx context.Context, ids []*string) ([]*string, error) {
	var items []*string
	err := m.db.Model(&Role{}).Where("id in (?) AND is_enable = TRUE", ids).Pluck("id", &items).Error

	return items, err
}

func (m *defaultRoleModel) ExistByRoleCode(ctx context.Context, roleCode string) (bool, error) {
	var count int64
	err := m.db.Model(&Role{}).Where("role_code = ?", roleCode).Count(&count).Error

	return count > 0, err
}

func (m *defaultRoleModel) Enable(ctx context.Context, id string) error {
	var result *Role
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&Role{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
