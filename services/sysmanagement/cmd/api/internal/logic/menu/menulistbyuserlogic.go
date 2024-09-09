package menu

import (
	"asense/common/components"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"github.com/jinzhu/copier"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListByUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据当前用户获取菜单列表
func NewMenuListByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListByUserLogic {
	return &MenuListByUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListByUserLogic) MenuListByUser() (resp *types.MenuListResp, err error) {
	var (
		roleIds  []*string
		menuIds  []*string
		menus    []*model.Menu
		treeList []*model.TreeMenu
		menuTree []*types.MenuTreeResp
	)
	userId := components.GetAuthKeyJwtUserID(l.ctx)
	roleIds, err = l.svcCtx.UserRoleModel.ListByUserID(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	if len(roleIds) == 0 {
		return &types.MenuListResp{Items: menuTree}, nil
	}
	roleIds, err = l.svcCtx.RoleModel.ListByIdsToIds(l.ctx, roleIds)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	if len(roleIds) == 0 {
		return &types.MenuListResp{Items: menuTree}, nil
	}

	menuIds, err = l.svcCtx.RolePermissionModel.ListByRoleIds(l.ctx, roleIds)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	menus, err = l.svcCtx.MenuModel.ListByIds(l.ctx, menuIds)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	treeList, err = l.svcCtx.MenuModel.ListTree(l.ctx, menus)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}

	_ = copier.Copy(&menuTree, &treeList)
	return &types.MenuListResp{Items: menuTree}, nil
}
