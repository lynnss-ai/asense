package menu

import (
	"asense/common/errorx"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑菜单
func NewMenuEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuEditLogic {
	return &MenuEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuEditLogic) MenuEdit(req *types.MenuEditReq) error {
	v := map[string]interface{}{
		"pid":                      req.PID,
		"menu_name":                req.MenuName,
		"menu_desc":                req.MenuDesc,
		"menu_icon":                req.MenuIcon,
		"menu_path":                req.MenuPath,
		"menu_type":                req.MenuType,
		"menu_component":           req.MenuComponent,
		"is_hide_in_menu":          req.IsHideInMenu,
		"is_hide_children_in_menu": req.IsHideChildrenInMenu,
		"is_hide_in_breadcrumb":    req.IsHideInBreadcrumb,
		"sort":                     req.Sort,
	}
	err := l.svcCtx.MenuModel.Update(l.ctx, req.ID, v)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
