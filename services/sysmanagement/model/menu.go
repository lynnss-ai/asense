// Package model
// @File    : menu.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 15:06
// @Desc    :
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

const (
	MenuTypeMenu   MenuTypeEnum = iota + 1 // 菜单
	MenuTypeButton                         // 按钮
)

type (
	MenuModel interface {
		WithTrans(ctx context.Context) MenuModel
		Insert(ctx context.Context, arg *Menu) error
		BatchInsert(ctx context.Context, args []*Menu) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id int) error
		FindOne(ctx context.Context, id string) (*Menu, error)
		ListByIds(ctx context.Context, ids []string) ([]*Menu, error)
		ListAll(ctx context.Context, filter *string, isEnable *bool) ([]*Menu, error)
		ListTree(ctx context.Context, items []*Menu) ([]*TreeMenu, error)
		ExistByMenuCode(ctx context.Context, menuCode string) (bool, error)
		Count(ctx context.Context) (int64, error)
		CountByPid(ctx context.Context, pid string) (int64, error)
		Enable(ctx context.Context, id string) error
	}

	MenuTypeEnum int

	Menu struct {
		ID                   int            `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`                                    // 菜单ID
		PID                  string         `json:"pid" gorm:"column:pid;type:varchar(32);not null"`                                    // 父ID[根节点模型为字符串0]
		MenuName             string         `json:"menu_name" gorm:"column:menu_name;type:varchar(32);not null"`                        // 菜单名称
		MenuCode             string         `json:"menu_code" gorm:"column:menu_code;type:varchar(32);not null"`                        // 菜单编码
		MenuDesc             string         `json:"menu_desc" gorm:"column:menu_desc;type:varchar(255)"`                                // 菜单描述
		MenuIcon             string         `json:"menu_icon" gorm:"column:menu_icon;type:varchar(255)"`                                // 菜单图标
		MenuPath             string         `json:"menu_path" gorm:"column:menu_path;type:varchar(255)"`                                // 菜单路径
		MenuComponent        string         `json:"menu_component" gorm:"column:menu_component;type:varchar(255)"`                      // 菜单组件
		MenuType             MenuTypeEnum   `json:"menu_type" gorm:"column:menu_type;type:int;not null"`                                // 菜单类型[1:菜单,2:按钮]
		IsHideInMenu         bool           `json:"is_hide_in_menu" gorm:"column:is_hide_in_menu;type:bool;not null"`                   // 是否在菜单中隐藏[true:隐藏,false:显示]
		IsHideChildrenInMenu bool           `json:"is_hide_children_in_menu" gorm:"column:is_hide_children_in_menu;type:bool;not null"` // 是否在菜单中隐藏子菜单[true:隐藏,false:显示]
		IsHideInBreadcrumb   bool           `json:"is_hide_in_breadcrumb" gorm:"column:is_hide_in_breadcrumb;type:bool;not null"`       // 是否在面包屑中隐藏[true:隐藏,false:显示]
		IsEnable             bool           `json:"is_enable" gorm:"column:is_enable;type:bool;not null"`                               // 是否启用[true:启用,false:禁用]
		Sort                 int            `json:"sort" gorm:"column:sort;type:int;not null"`                                          // 排序
		CreatedAt            int64          `json:"created_at" gorm:"autoCreateTime:milli"`                                             // 创建时间
		UpdatedAt            int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`                                             // 更新时间
		DeletedAt            gorm.DeletedAt `json:"deleted_at"`                                                                         // 删除时间
	}

	TreeMenu struct {
		Menu
		Children []*TreeMenu
	}
	idMapTreeMenuType map[string]*TreeMenu

	defaultMenuModel struct {
		db *gorm.DB
	}
)

func NewMenuModel(isMigration bool, db *gorm.DB) MenuModel {
	if isMigration {
		if err := db.AutoMigrate(&Menu{}); err != nil {
			panic(err)
		}
	}
	return &defaultMenuModel{db: db}
}

func (m *defaultMenuModel) WithTrans(ctx context.Context) MenuModel {
	return &defaultMenuModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultMenuModel) Insert(ctx context.Context, arg *Menu) error {
	return m.db.Create(&arg).Error
}

func (m *defaultMenuModel) BatchInsert(ctx context.Context, args []*Menu) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultMenuModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Menu{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultMenuModel) Delete(ctx context.Context, id int) error {
	return m.db.Delete(&Menu{}, id).Error
}

func (m *defaultMenuModel) FindOne(ctx context.Context, id string) (*Menu, error) {
	var result *Menu
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultMenuModel) ListByIds(ctx context.Context, ids []string) ([]*Menu, error) {
	var items []*Menu
	err := m.db.Where("id in (?)", ids).Find(&items).Error

	return items, err
}

func (m *defaultMenuModel) ListAll(ctx context.Context, filter *string, isEnable *bool) ([]*Menu, error) {
	var items []*Menu
	query := m.db.Model(&Menu{})
	if filter != nil {
		_filter := components.Filter(*filter)
		query = query.Where("menu_name like ? OR menu_code like ?", _filter, _filter)
	}

	if isEnable != nil {
		query = query.Where("is_enable = ?", *isEnable)
	}

	err := query.Order("sort asc").Find(&items).Error

	return items, err
}

func (m *defaultMenuModel) ListTree(ctx context.Context, items []*Menu) ([]*TreeMenu, error) {
	treeMenus := make([]*TreeMenu, 0, len(items))
	idMap := make(map[string]*TreeMenu)

	for _, item := range items {
		treeMenu := &TreeMenu{
			Menu: Menu{
				ID:                   item.ID,
				PID:                  item.PID,
				MenuName:             item.MenuName,
				MenuCode:             item.MenuCode,
				MenuDesc:             item.MenuDesc,
				MenuIcon:             item.MenuIcon,
				MenuPath:             item.MenuPath,
				MenuComponent:        item.MenuComponent,
				MenuType:             item.MenuType,
				IsHideInMenu:         item.IsHideInMenu,
				IsHideChildrenInMenu: item.IsHideChildrenInMenu,
				IsHideInBreadcrumb:   item.IsHideInBreadcrumb,
				IsEnable:             item.IsEnable,
				Sort:                 item.Sort,
				CreatedAt:            item.CreatedAt,
				UpdatedAt:            item.UpdatedAt,
			},
			Children: nil,
		}

		if item.PID == "0" {
			treeMenus = append(treeMenus, treeMenu)
		} else {
			parent := idMap[item.PID]
			parent.Children = append(parent.Children, treeMenu)
		}
		idMap[item.PID] = treeMenu
	}

	return treeMenus, nil
}

func (m *defaultMenuModel) ExistByMenuCode(ctx context.Context, menuCode string) (bool, error) {
	var count int64
	err := m.db.Model(&Menu{}).Where("menu_code = ?", menuCode).Count(&count).Error

	return count > 0, err
}

func (m *defaultMenuModel) Count(ctx context.Context) (int64, error) {
	var count int64
	err := m.db.Model(&Menu{}).Count(&count).Error

	return count, err
}

func (m *defaultMenuModel) CountByPid(ctx context.Context, pid string) (int64, error) {
	var count int64
	err := m.db.Model(&Menu{}).Where("pid = ?", pid).Count(&count).Error

	return count, err
}

func (m *defaultMenuModel) Enable(ctx context.Context, id string) error {
	var result *Menu
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&Menu{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
