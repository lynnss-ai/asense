package user

import (
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"errors"
	"gorm.io/gorm"

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
	var (
		user      *model.User
		userRoles []*model.UserRole
	)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorx.NewDefaultError("该用户不存在")
		}
		return errorx.NewDataBaseError(err)
	}
	for _, roleId := range req.RoleIds {
		userRoles = append(userRoles, &model.UserRole{
			ID:     dbcore.NewId(),
			UserID: user.ID,
			RoleID: roleId,
		})
	}

	err = l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.UserModel.WithTrans(ctx).Update(l.ctx, req.ID, map[string]interface{}{
			"name":   req.Name,
			"avatar": req.Avatar,
			"email":  req.Email,
			"phone":  req.Phone,
		}); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.UserRoleModel.WithTrans(ctx).DeleteByUserID(l.ctx, req.ID); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.UserRoleModel.WithTrans(ctx).BatchInsert(l.ctx, userRoles); err != nil {
			return errorx.NewDataBaseError(err)
		}
		return nil
	})
	return err
}
