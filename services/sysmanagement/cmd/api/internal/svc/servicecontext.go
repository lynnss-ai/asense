package svc

import (
	"asense/common/dbcore"
	"asense/services/sysmanagement/cmd/api/internal/config"
	"asense/services/sysmanagement/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config                config.Config
	Tx                    dbcore.Transaction
	AttachmentModel       model.AttachmentModel
	DictionaryModel       model.DictionaryModel
	MenuModel             model.MenuModel
	OrganizationModel     model.OrganizationModel
	OrganizationUserModel model.OrganizationUserModel
	PositionModel         model.PositionModel
	RoleModel             model.RoleModel
	RolePermissionModel   model.RolePermissionModel
	UserModel             model.UserModel
	UserRoleModel         model.UserRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open(postgres.Open(c.Database.Postgres.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表名前缀
			SingularTable: true, //使用单数表名
		},
		SkipDefaultTransaction: true, //禁用默认事务
		PrepareStmt:            true, //缓存预编译语句
	})

	if err != nil {
		panic(err)
	}

	isMigration := c.Database.Postgres.IsMigration

	return &ServiceContext{
		Config:                c,
		Tx:                    dbcore.NewTransaction(db),
		AttachmentModel:       model.NewAttachmentModel(isMigration, db),
		DictionaryModel:       model.NewDictionaryModel(isMigration, db),
		MenuModel:             model.NewMenuModel(isMigration, db),
		OrganizationModel:     model.NewOrganizationModel(isMigration, db),
		OrganizationUserModel: model.NewOrganizationUserModel(isMigration, db),
		PositionModel:         model.NewPositionModel(isMigration, db),
		RoleModel:             model.NewRoleModel(isMigration, db),
		RolePermissionModel:   model.NewRolePermissionModel(isMigration, db),
		UserModel:             model.NewUserModel(isMigration, db),
		UserRoleModel:         model.NewUserRoleModel(isMigration, db),
	}
}
