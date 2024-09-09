package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAvatarEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑用户头像
func NewUserAvatarEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAvatarEditLogic {
	return &UserAvatarEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAvatarEditLogic) UserAvatarEdit(req *types.UserAvatarEditReq) error {
	// todo: add your logic here and delete this line

	return nil
}
