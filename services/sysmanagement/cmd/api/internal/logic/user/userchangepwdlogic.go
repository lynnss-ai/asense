package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改密码
func NewUserChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChangePwdLogic {
	return &UserChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChangePwdLogic) UserChangePwd(req *types.ComUserChangePwdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
