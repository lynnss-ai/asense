package user

import (
	"asense/common/components"
	"asense/common/errorx"
	"asense/common/utils/encryptutil"
	"asense/services/sysmanagement/model"
	"context"
	"time"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.ComLoginReq) (resp *types.ComLoginResp, err error) {
	var user *model.User

	now := time.Now().Unix()
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		return nil, errorx.NewDefaultError("用户名或密码不能为空")
	}
	user, err = l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	isLoginSuccess := encryptutil.ComparePassword(user.Password, req.Password, user.Salt)
	if !isLoginSuccess {
		//登录密码错误
		return nil, errorx.NewDefaultError("用户名或密码错误")
	}

	if !user.IsEnable {
		return nil, errorx.NewDefaultError("您的账号已被禁用,请联系管理员进行恢复")
	}
	token, err := components.GeneratorJwtToken(components.CtxAuthKeyByJWTUserID, l.svcCtx.Config.JetAuth.AccessSecret, now, l.svcCtx.Config.JetAuth.AccessExpire, user.ID)
	if err != nil {
		return nil, errorx.NewDefaultError("生成token失败")
	}
	if err := l.svcCtx.UserModel.Update(l.ctx, user.ID, map[string]interface{}{
		"last_login_time": time.Now(),
	}); err != nil {
		return nil, errorx.NewDataBaseError(err)
	}

	return &types.ComLoginResp{
		AccessToken:  token,
		AccessExpire: now + l.svcCtx.Config.JetAuth.AccessExpire,
		RefreshAfter: now + l.svcCtx.Config.JetAuth.AccessExpire/2,
	}, nil
}
