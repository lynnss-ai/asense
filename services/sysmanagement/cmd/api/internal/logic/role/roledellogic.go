package role

import (
	"asense/common/errorx"
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
	var (
		count int64
		err   error
	)
	count, err = l.svcCtx.UserRoleModel.CountByRoleID(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if count > 0 {
		return errorx.NewDefaultError("该角色下存在用户，不能删除")
	}
	err = l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.RoleModel.WithTrans(ctx).Delete(l.ctx, req.ID); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.RolePermissionModel.WithTrans(ctx).DeleteByRoleID(l.ctx, req.ID); err != nil {
			return errorx.NewDataBaseError(err)
		}
		return nil
	})
	return err
}
