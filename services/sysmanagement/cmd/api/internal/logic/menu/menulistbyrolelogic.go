package menu

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据角色获取菜单列表
func NewMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListByRoleLogic {
	return &MenuListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListByRoleLogic) MenuListByRole(req *types.ComIDPathReq) (resp *types.MenuListResp, err error) {
	var (
		menuIds  []*string
		menus    []*model.Menu
		treeList []*model.TreeMenu
		menuTree []*types.MenuTreeResp
	)

	menuIds, err = l.svcCtx.RolePermissionModel.ListByRoleID(l.ctx, req.ID)
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
