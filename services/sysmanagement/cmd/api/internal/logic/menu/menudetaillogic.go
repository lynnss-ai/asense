package menu

import (
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
	// todo: add your logic here and delete this line

	return
}
