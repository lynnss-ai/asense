package user

import (
	"asense/common/errorx"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelLogic {
	return &UserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDelLogic) UserDel(req *types.ComIDPathReq) error {
	err := l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.UserRoleModel.WithTrans(ctx).DeleteByUserID(l.ctx, req.ID); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.UserModel.WithTrans(ctx).Delete(l.ctx, req.ID); err != nil {
			return errorx.NewDataBaseError(err)
		}

		return nil
	})

	return err
}
