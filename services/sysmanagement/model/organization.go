// Package model
// @File    : organization.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/27 09:33
// @Desc    : 组织架构模型
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
)

const (
	OrganizationTypeCompany    OrganizationTypeEnum = iota + 1 // 公司
	OrganizationTypeDepartment                                 // 部门
)

type (
	OrganizationModel interface {
		WithTrans(ctx context.Context) OrganizationModel
		Insert(ctx context.Context, arg *Organization) error
		BatchInsert(ctx context.Context, args []*Organization) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id string) error
		FindOne(ctx context.Context, id string) (*Organization, error)
		ListByIds(ctx context.Context, ids []*string) ([]*Organization, error)
		ListAll(ctx context.Context, filter *string, isEnable *bool) ([]*Organization, error)
		ListTree(ctx context.Context, items []*Organization) ([]*TreeOrganization, error)
		ExistByOrgCode(ctx context.Context, id *string, orgCode string) (bool, error)
		Count(ctx context.Context) (int64, error)
		CountByPid(ctx context.Context, pid string) (int64, error)
		Enable(ctx context.Context, id string) error
	}

	OrganizationTypeEnum int

	Organization struct {
		ID        string               `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`           // 组织架构ID
		PID       string               `json:"pid" gorm:"column:pid;type:varchar(32);not null"`           // 父ID[根节点模型为字符串0]
		OrgName   string               `json:"org_name" gorm:"column:org_name;type:varchar(64);not null"` // 组织架构名称
		OrgCode   string               `json:"org_code" gorm:"column:org_code;type:varchar(32);not null"` // 组织架构编码
		OrgDesc   string               `json:"org_desc" gorm:"column:org_desc;type:varchar(255)"`         // 组织架构描述
		OrgType   OrganizationTypeEnum `json:"org_type" gorm:"column:org_type;type:int;not null"`         // 组织架构类型[1:公司,2:部门]
		IsEnable  bool                 `json:"is_enable" gorm:"column:is_enable;type:bool;not null"`      // 是否启用[true:启用,false:禁用]
		Sort      int                  `json:"sort" gorm:"column:sort;type:int;not null"`                 // 排序
		CreatedAt int64                `json:"created_at" gorm:"autoCreateTime:milli"`                    // 创建时间
		UpdatedAt int64                `json:"updated_at" gorm:"autoUpdateTime:milli"`                    // 更新时间
		DeletedAt gorm.DeletedAt       `json:"deleted_at"`                                                // 删除时间
	}

	TreeOrganization struct {
		Organization
		Children []*TreeOrganization
	}
	idMapTreeOrganizationType map[string]*TreeOrganization

	defaultOrganizationModel struct {
		db *gorm.DB
	}
)

func NewOrganizationModel(isMigration bool, db *gorm.DB) OrganizationModel {
	if isMigration {
		if err := db.AutoMigrate(&Organization{}); err != nil {
			panic(err)
		}
	}
	return &defaultOrganizationModel{db: db}
}

// WithTrans 开启事务
func (m *defaultOrganizationModel) WithTrans(ctx context.Context) OrganizationModel {
	return &defaultOrganizationModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultOrganizationModel) Insert(ctx context.Context, arg *Organization) error {
	return m.db.Create(&arg).Error
}

func (m *defaultOrganizationModel) BatchInsert(ctx context.Context, args []*Organization) error {
	return m.db.CreateInBatches(&args, 500).Error
}

func (m *defaultOrganizationModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Organization{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultOrganizationModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&Organization{}, id).Error
}

func (m *defaultOrganizationModel) FindOne(ctx context.Context, id string) (*Organization, error) {
	var result *Organization
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultOrganizationModel) ListByIds(ctx context.Context, ids []*string) ([]*Organization, error) {
	var items []*Organization
	err := m.db.Where("id in (?)", ids).Find(&items).Error

	return items, err
}

func (m *defaultOrganizationModel) ListAll(ctx context.Context, filter *string, isEnable *bool) ([]*Organization, error) {
	var items []*Organization
	query := m.db.Model(&Organization{})
	if filter != nil {
		_filter := components.Filter(*filter)
		query = query.Where("org_name like ? OR org_code like ?", _filter, _filter)
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", *isEnable)
	}
	err := query.Order("sort asc").Find(&items).Error

	return items, err
}

func (m *defaultOrganizationModel) ListTree(ctx context.Context, items []*Organization) ([]*TreeOrganization, error) {
	TreeOrganizations := make([]*TreeOrganization, 0, len(items))
	idMap := make(map[string]*TreeOrganization)

	for _, item := range items {
		treeOrganizations := &TreeOrganization{
			Organization: Organization{
				ID:        item.ID,
				PID:       item.PID,
				OrgName:   item.OrgName,
				OrgCode:   item.OrgCode,
				OrgDesc:   item.OrgDesc,
				OrgType:   item.OrgType,
				IsEnable:  item.IsEnable,
				Sort:      item.Sort,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			},
			Children: nil,
		}

		if item.PID == "0" {
			TreeOrganizations = append(TreeOrganizations, treeOrganizations)
		} else {
			parent := idMap[item.PID]
			parent.Children = append(parent.Children, treeOrganizations)
		}
		idMap[item.ID] = treeOrganizations
	}
	return TreeOrganizations, nil
}

func (m *defaultOrganizationModel) ExistByOrgCode(ctx context.Context, id *string, orgCode string) (bool, error) {
	var count int64
	query := m.db.Model(&Organization{}).Where("org_code = ?", orgCode)
	if id != nil {
		query = query.Where("id != ?", *id)
	}
	err := query.Count(&count).Error

	return count > 0, err
}

func (m *defaultOrganizationModel) Count(ctx context.Context) (int64, error) {
	var count int64
	err := m.db.Model(&Organization{}).Count(&count).Error

	return count, err
}

func (m *defaultOrganizationModel) CountByPid(ctx context.Context, pid string) (int64, error) {
	var count int64
	err := m.db.Model(&Organization{}).Where("pid = ?", pid).Count(&count).Error

	return count, err
}

func (m *defaultOrganizationModel) Enable(ctx context.Context, id string) error {
	var result *Organization
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&Organization{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
