package role

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleEnableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 启用|禁用角色
func NewRoleEnableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleEnableLogic {
	return &RoleEnableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleEnableLogic) RoleEnable(req *types.ComIDPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
