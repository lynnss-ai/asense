package user

import (
	"asense/common/components"
	"asense/common/errorx"
	"asense/common/utils/encryptutil"
	"asense/common/utils/randomutil"
	"asense/services/sysmanagement/model"
	"context"
	"errors"
	"gorm.io/gorm"

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
	var (
		user *model.User
		err  error
	)
	userId := components.GetAuthKeyJwtUserID(l.ctx)
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorx.NewDefaultError("该用户不存在")
		}
		return errorx.NewDataBaseError(err)
	}
	if !user.IsAdmin {
		return errorx.NewDefaultError("请联系管理员重置密码")
	}
	salt := randomutil.GetRandomNumStr(32)
	password, _ := encryptutil.GeneratePassword(req.Password, salt)
	userMap := map[string]interface{}{
		"password": password,
		"salt":     salt,
	}
	err = l.svcCtx.UserModel.Update(l.ctx, req.UserID, userMap)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
