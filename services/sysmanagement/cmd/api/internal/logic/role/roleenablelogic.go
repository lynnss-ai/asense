package role

import (
	"asense/common/errorx"
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
	err := l.svcCtx.RoleModel.Enable(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
