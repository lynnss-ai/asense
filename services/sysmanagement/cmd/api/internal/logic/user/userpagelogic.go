package user

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户分页列表
func NewUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageLogic {
	return &UserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPageLogic) UserPage(req *types.UserPageReq) (resp *types.UserPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
