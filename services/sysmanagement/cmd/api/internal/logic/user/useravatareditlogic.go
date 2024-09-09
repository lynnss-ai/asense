package user

import (
	"asense/common/components"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"errors"
	"gorm.io/gorm"

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
	var (
		user *model.User
		err  error
	)
	userID := components.GetAuthKeyJwtUserID(l.ctx)
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorx.NewDefaultError("用户不存在")
		}
		return errorx.NewDataBaseError(err)
	}

	if err = l.svcCtx.UserModel.Update(l.ctx, user.ID, map[string]interface{}{
		"avatar": req.Avatar,
	}); err != nil {
		return errorx.NewDataBaseError(err)
	}

	return nil
}
