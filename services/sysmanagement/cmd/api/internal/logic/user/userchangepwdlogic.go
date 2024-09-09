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
	isSuccessPwd := encryptutil.ComparePassword(req.NewPassword, req.OldPassword, user.Salt)
	if !isSuccessPwd {
		return errorx.NewDefaultError("您输入的旧密码错误")
	}
	salt := randomutil.GetRandomNumStr(32)
	password, _ := encryptutil.GeneratePassword(req.NewPassword, salt)
	userMap := map[string]interface{}{
		"password": password,
		"salt":     salt,
	}
	err = l.svcCtx.UserModel.Update(l.ctx, userId, userMap)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
