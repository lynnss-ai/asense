package menu

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 菜单详情
func NewMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDetailLogic {
	return &MenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDetailLogic) MenuDetail(req *types.ComIDPathReq) (resp *types.MenuResp, err error) {
	var menu *model.Menu

	menu, err = l.svcCtx.MenuModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	resp = &types.MenuResp{
		ID:                   menu.ID,
		PID:                  menu.PID,
		MenuName:             menu.MenuName,
		MenuCode:             menu.MenuCode,
		MenuDesc:             menu.MenuDesc,
		MenuIcon:             menu.MenuIcon,
		MenuPath:             menu.MenuPath,
		MenuType:             int(menu.MenuType),
		MenuComponent:        menu.MenuComponent,
		IsHideInMenu:         menu.IsHideInMenu,
		IsHideChildrenInMenu: menu.IsHideChildrenInMenu,
		IsHideInBreadcrumb:   menu.IsHideInBreadcrumb,
		IsEnable:             menu.IsEnable,
		Sort:                 menu.Sort,
	}
	return resp, nil
}
