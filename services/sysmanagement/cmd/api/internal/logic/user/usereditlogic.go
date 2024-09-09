package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑用户
func NewUserEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEditLogic {
	return &UserEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserEditLogic) UserEdit(req *types.UserEditReq) error {
	// todo: add your logic here and delete this line

	return nil
}
