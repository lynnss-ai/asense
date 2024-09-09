package menu

import (
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
	// todo: add your logic here and delete this line

	return nil
}
