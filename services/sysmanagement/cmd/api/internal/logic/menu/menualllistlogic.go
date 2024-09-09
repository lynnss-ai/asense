package menu

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
