package role

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除角色
func NewRoleDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDelLogic {
	return &RoleDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDelLogic) RoleDel(req *types.ComIDPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
