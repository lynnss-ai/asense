package menu

import (
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
	// todo: add your logic here and delete this line

	return nil
}
