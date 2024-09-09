package menu

import (
	"asense/common/components"
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/common/utils/characterutil"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增菜单
func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAddLogic) MenuAdd(req *types.MenuReq) error {
	var (
		user    *model.User
		isExist bool
		err     error
	)
	userID := components.GetAuthKeyJwtUserID(l.ctx)
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if !user.IsAdmin {
		return errorx.NewDefaultError("非管理员不能添加菜单")
	}
	isExist, err = l.svcCtx.MenuModel.ExistByMenuCode(l.ctx, req.MenuCode)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if isExist {
		return errorx.NewDefaultError("菜单编码已存在")
	}

	menu := model.Menu{
		ID:                   dbcore.NewId(),
		PID:                  req.PID,
		MenuName:             req.MenuName,
		MenuCode:             characterutil.StitchingBuilderStr("M", req.MenuCode),
		MenuDesc:             req.MenuDesc,
		MenuIcon:             req.MenuIcon,
		MenuPath:             req.MenuPath,
		MenuType:             model.MenuTypeEnum(req.MenuType),
		MenuComponent:        req.MenuComponent,
		IsHideInMenu:         req.IsHideInMenu,
		IsHideChildrenInMenu: req.IsHideChildrenInMenu,
		IsHideInBreadcrumb:   req.IsHideInBreadcrumb,
		IsEnable:             true,
		Sort:                 req.Sort,
	}

	err = l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.MenuModel.WithTrans(ctx).Insert(l.ctx, &menu); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if req.IsJoinAdminist {
			adminRole, err := l.svcCtx.RoleModel.FindOneByCode(l.ctx, components.AdministratorRoleCode)
			if err != nil {
				return errorx.NewDataBaseError(err)
			}
			err = l.svcCtx.RolePermissionModel.Insert(l.ctx, &model.RolePermission{
				ID:     dbcore.NewId(),
				RoleID: adminRole.ID,
				MenuID: menu.ID,
			})
			if err != nil {
				return errorx.NewDataBaseError(err)
			}
		}
		return nil
	})

	return err
}
