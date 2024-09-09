package menu

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"github.com/jinzhu/copier"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAllListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取所有菜单
func NewMenuAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAllListLogic {
	return &MenuAllListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAllListLogic) MenuAllList() (resp *types.MenuListResp, err error) {
	var (
		menus    []*model.Menu
		treeList []*model.TreeMenu
		menuTree []*types.MenuTreeResp
	)
	menus, err = l.svcCtx.MenuModel.ListAll(l.ctx, nil, nil)
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
