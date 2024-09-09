package menu

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuEnableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 启用|禁用菜单
func NewMenuEnableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuEnableLogic {
	return &MenuEnableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuEnableLogic) MenuEnable(req *types.ComIDPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
