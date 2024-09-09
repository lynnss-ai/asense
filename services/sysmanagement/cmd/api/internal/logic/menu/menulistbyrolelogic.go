package menu

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
