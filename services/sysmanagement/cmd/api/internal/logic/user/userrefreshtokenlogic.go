package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新用户Token
func NewUserRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRefreshTokenLogic {
	return &UserRefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRefreshTokenLogic) UserRefreshToken() (resp *types.ComLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
