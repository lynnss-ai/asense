package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserEnableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 启用|禁用用户
func NewUserEnableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEnableLogic {
	return &UserEnableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserEnableLogic) UserEnable(req *types.ComIDPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
