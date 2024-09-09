package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户信息
func NewUserCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurrentLogic {
	return &UserCurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCurrentLogic) UserCurrent() (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
