package role

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleSetPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置角色权限
func NewRoleSetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleSetPermissionLogic {
	return &RoleSetPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleSetPermissionLogic) RoleSetPermission(req *types.RoleSetPermissionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
