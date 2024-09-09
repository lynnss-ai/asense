package menu

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
