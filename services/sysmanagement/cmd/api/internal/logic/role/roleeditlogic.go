package role

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑角色
func NewRoleEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleEditLogic {
	return &RoleEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleEditLogic) RoleEdit(req *types.RoleEditReq) error {
	// todo: add your logic here and delete this line

	return nil
}
