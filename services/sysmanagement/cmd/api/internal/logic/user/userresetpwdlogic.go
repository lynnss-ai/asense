package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserResetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 重置密码
func NewUserResetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserResetPwdLogic {
	return &UserResetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserResetPwdLogic) UserResetPwd(req *types.ComUserResetPwdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
