package user

import (
	"asense/common/components"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"time"

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
	var (
		user  *model.User
		token string
	)
	now := time.Now().Unix()

	userId := components.GetAuthKeyJwtUserID(l.ctx)
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	if !user.IsEnable {
		return nil, errorx.NewDefaultError("该账户已被禁用")
	}

	token, err = components.GeneratorJwtToken(components.CtxAuthKeyByJWTUserID, l.svcCtx.Config.JetAuth.AccessSecret, now, l.svcCtx.Config.JetAuth.AccessExpire, user.ID)
	if err != nil {
		return nil, errorx.NewDefaultError("生成token失败")
	}

	return &types.ComLoginResp{
		AccessToken:  token,
		AccessExpire: now + l.svcCtx.Config.JetAuth.AccessExpire,
		RefreshAfter: now + l.svcCtx.Config.JetAuth.AccessExpire/2,
	}, nil
}
