// Package model
// @File    : user.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/7 15:07
// @Desc    : 用户
package model

import (
	"asense/common/components"
	"asense/common/dbcore"
	"context"
	"gorm.io/gorm"
	"time"
)

type (
	UserModel interface {
		WithTrans(ctx context.Context) UserModel
		Insert(ctx context.Context, arg *User) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		Delete(ctx context.Context, id string) error
		FindOne(ctx context.Context, id string) (*User, error)
		FindOneByUserName(ctx context.Context, userName string) (*User, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*User, error)
		ExistByUserName(ctx context.Context, userName string) (bool, error)
		Count(ctx context.Context) (int64, error)
		Enable(ctx context.Context, id string) error
	}

	User struct {
		ID            string          `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`             // 用户ID
		Name          string          `json:"name" gorm:"column:name;type:varchar(32);not null"`           // 姓名
		UserName      string          `json:"user_name" gorm:"column:user_name;type:varchar(32);not null"` // 用户名
		Phone         *string         `json:"phone" gorm:"column:phone;type:varchar(20)"`                  // 手机号
		Password      string          `json:"password" gorm:"column:password;type:varchar(64);not null"`   // 密码
		Email         *string         `json:"email" gorm:"column:email;type:varchar(255)"`                 // 邮箱
		Salt          string          `json:"salt" gorm:"column:salt;type:varchar(256);not null"`          // 盐值
		Avatar        string          `json:"avatar" gorm:"column:avatar;type:varchar(255)"`               // 头像
		IsEnable      bool            `json:"is_enable" gorm:"column:is_enable;type:bool;not null"`        // 是否启用[true:启用,false:禁用]
		IsAdmin       bool            `json:"is_admin" gorm:"column:is_admin;type:bool;not null"`          // 是否管理员[true:是,false:否]
		Remark        string          `json:"remark" gorm:"column:remark;type:varchar(255)"`               // 备注
		LastLoginTime *time.Time      `json:"last_login_time" gorm:"column:last_login_time;type:datetime"` // 最后登录时间
		CreatedAt     int64           `json:"created_at" gorm:"autoCreateTime:milli"`                      // 创建时间
		UpdatedAt     int64           `json:"updated_at" gorm:"autoUpdateTime:milli"`                      // 更新时间
		DeletedAt     *gorm.DeletedAt `json:"deleted_at"`                                                  // 删除时间
	}

	defaultUserModel struct {
		db *gorm.DB
	}
)

func NewUserModel(isMigration bool, db *gorm.DB) UserModel {
	if isMigration {
		if err := db.AutoMigrate(&User{}); err != nil {
			panic(err)
		}
	}
	return &defaultUserModel{db: db}
}

func (m *defaultUserModel) WithTrans(ctx context.Context) UserModel {
	return &defaultUserModel{db: dbcore.GetDB(ctx, m.db)}
}

func (m *defaultUserModel) Insert(ctx context.Context, arg *User) error {
	return m.db.Create(&arg).Error
}

func (m *defaultUserModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&User{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultUserModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&User{}, id).Error
}

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	var result *User
	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultUserModel) FindOneByUserName(ctx context.Context, userName string) (*User, error) {
	var result *User
	err := m.db.Where("user_name = ?", userName).First(&result).Error

	return result, err
}

func (m *defaultUserModel) FindPage(ctx context.Context, page, pageSize int, filter *string, isEnable *bool) (int64, []*User, error) {
	var (
		total int64
		items []*User
	)
	p := components.PageHandle(page, pageSize, filter)
	query := m.db.Model(&User{})
	if filter != nil {
		query = query.Where("name like ? OR user_name like ?", p.Filter, p.Filter)
	}
	if isEnable != nil {
		query = query.Where("is_enable = ?", isEnable)
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

func (m *defaultUserModel) ExistByUserName(ctx context.Context, userName string) (bool, error) {
	var total int64
	err := m.db.Model(&User{}).Where("user_name = ?", userName).Count(&total).Error

	return total > 0, err
}

func (m *defaultUserModel) Count(ctx context.Context) (int64, error) {
	var total int64
	err := m.db.Model(&User{}).Count(&total).Error

	return total, err
}

func (m *defaultUserModel) Enable(ctx context.Context, id string) error {
	var result *User
	err := m.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}

	return m.db.Model(&User{}).Where("id = ?", id).Update("is_enable", !result.IsEnable).Error
}
