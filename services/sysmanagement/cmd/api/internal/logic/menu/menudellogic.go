package menu

import (
	"asense/common/errorx"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除菜单
func NewMenuDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDelLogic {
	return &MenuDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDelLogic) MenuDel(req *types.ComIDPathReq) error {
	var (
		count int64
		err   error
	)
	count, err = l.svcCtx.MenuModel.CountByPid(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if count > 0 {
		return errorx.NewDefaultError("该菜单下存在子菜单，不能删除")
	}
	count, err = l.svcCtx.RolePermissionModel.CountByMenuID(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if count > 0 {
		return errorx.NewDefaultError("该菜单下存在角色，不能删除")
	}
	err = l.svcCtx.MenuModel.Delete(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
