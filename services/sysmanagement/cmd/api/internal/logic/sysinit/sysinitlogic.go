package sysinit

import (
	"asense/common/components"
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/common/utils/encryptutil"
	"asense/common/utils/randomutil"
	"asense/services/sysmanagement/model"
	"context"
	"strings"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SysinitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 系统初始化
func NewSysinitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysinitLogic {
	return &SysinitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysinitLogic) Sysinit() error {
	count, err := l.svcCtx.UserModel.Count(l.ctx)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if count > 0 {
		return errorx.NewDefaultError("系统已经初始化，无需再次初始化")
	}

	dictionaries := []*model.Dictionary{
		{ID: "10000001", PID: "0", DicName: "平台默认参数", DicCode: "D10000001", DicDesc: "", DicValue: "", DicValue2: "", DicValue3: "", Sort: 1, IsEdit: false, IsEnable: true, IsHide: false},
		{ID: "1000000101", PID: "10000001", DicName: "默认头像", DicCode: "D1000000101", DicDesc: "", DicValue: "", DicValue2: "", DicValue3: "", Sort: 1, IsEdit: false, IsEnable: true, IsHide: false},
	}

	menus := []*model.Menu{
		{ID: "100000001", PID: "0", MenuName: "首页", MenuCode: "", MenuDesc: "", MenuIcon: "", MenuPath: "/workbench", MenuComponent: "@pages/Workbench", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 1},

		{ID: "800000001", PID: "0", MenuName: "基础设施管理", MenuCode: "M800000001", MenuDesc: "", MenuIcon: "", MenuPath: "/infra", MenuComponent: "@pages/Infra/Dictionary", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 2},
		{ID: "800000001001", PID: "0", MenuName: "字典数据管理", MenuCode: "M800000001001", MenuDesc: "", MenuIcon: "", MenuPath: "/infra/dictionary", MenuComponent: "@pages/Infra/Dictionary", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 1},

		{ID: "900000001", PID: "0", MenuName: "系统管理", MenuCode: "M900000001", MenuDesc: "", MenuIcon: "", MenuPath: "/system", MenuComponent: "@pages/System/User", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 3},
		{ID: "900000001001", PID: "0", MenuName: "用户管理", MenuCode: "M900000001001", MenuDesc: "", MenuIcon: "", MenuPath: "/system/user", MenuComponent: "@pages/System/User", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 1},
		{ID: "900000001002", PID: "0", MenuName: "角色管理", MenuCode: "M900000001002", MenuDesc: "", MenuIcon: "", MenuPath: "/system/role", MenuComponent: "@pages/System/Role", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 2},
		{ID: "900000001003", PID: "0", MenuName: "菜单管理", MenuCode: "M900000001003", MenuDesc: "", MenuIcon: "", MenuPath: "/system/menu", MenuComponent: "@pages/System/Menu", MenuType: model.MenuTypeMenu, IsHideInMenu: false, IsHideChildrenInMenu: false, IsHideInBreadcrumb: false, IsEnable: true, Sort: 3},
	}

	salt := randomutil.GetRandomNumStr(32)
	password, _ := encryptutil.GeneratePassword(components.DefaultPassword, salt)
	email := components.DefaultEmail
	phone := components.DefaultPhone
	user := model.User{
		ID:            dbcore.NewId(),
		Name:          "超级管理员",
		UserName:      "admin",
		Phone:         &phone,
		Password:      password,
		Email:         &email,
		Salt:          salt,
		Avatar:        components.DefaultAvatar,
		IsEnable:      true,
		IsAdmin:       true,
		Remark:        "",
		LastLoginTime: nil,
	}

	role := model.Role{
		ID:              dbcore.NewId(),
		RoleName:        "超级管理员",
		RoleCode:        components.AdministratorRoleCode,
		RoleDesc:        "超级管理员",
		IsSetPermission: true,
		SelectedMenuIds: "",
		IsEnable:        true,
		IsAdmin:         true,
	}

	userRole := model.UserRole{
		ID:     dbcore.NewId(),
		UserID: user.ID,
		RoleID: role.ID,
	}

	var rolePermissions []*model.RolePermission
	var selectedMenuIds []string
	for _, menu := range menus {
		rolePermissions = append(rolePermissions, &model.RolePermission{
			ID:     dbcore.NewId(),
			RoleID: role.ID,
			MenuID: menu.ID,
		})

		if menu.PID != "0" {
			selectedMenuIds = append(selectedMenuIds, menu.ID)
		}
	}
	role.SelectedMenuIds = strings.Join(selectedMenuIds, ",")

	err = l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.DictionaryModel.WithTrans(ctx).BatchInsert(l.ctx, dictionaries); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.MenuModel.WithTrans(ctx).BatchInsert(l.ctx, menus); err != nil {
			return errorx.NewDataBaseError(err)
		}

		if err := l.svcCtx.UserModel.WithTrans(ctx).Insert(l.ctx, &user); err != nil {
			return errorx.NewDataBaseError(err)
		}

		if err := l.svcCtx.RoleModel.WithTrans(ctx).Insert(l.ctx, &role); err != nil {
			return errorx.NewDataBaseError(err)
		}

		if err := l.svcCtx.UserRoleModel.WithTrans(ctx).Insert(l.ctx, &userRole); err != nil {
			return errorx.NewDataBaseError(err)
		}

		return nil
	})
	return err
}
