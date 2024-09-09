package user

import (
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/common/utils/encryptutil"
	"asense/common/utils/randomutil"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增用户
func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserReq) error {
	var (
		isExist   bool
		userRoles []*model.UserRole
		err       error
	)
	isExist, err = l.svcCtx.UserModel.ExistByUserName(l.ctx, req.UserName)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if isExist {
		return errorx.NewDefaultError("用户名已存在")
	}

	salt := randomutil.GetRandomNumStr(32)
	password, _ := encryptutil.GeneratePassword(req.Password, salt)

	user := model.User{
		ID:            dbcore.NewId(),
		Name:          req.Name,
		UserName:      req.UserName,
		Phone:         req.Phone,
		Password:      password,
		Email:         req.Email,
		Salt:          salt,
		Avatar:        req.Avatar,
		IsEnable:      true,
		IsAdmin:       false,
		Remark:        "",
		LastLoginTime: nil,
	}

	for _, roleId := range req.RoleIds {
		userRoles = append(userRoles, &model.UserRole{
			ID:     dbcore.NewId(),
			UserID: user.ID,
			RoleID: roleId,
		})
	}

	err = l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.UserModel.WithTrans(ctx).Insert(l.ctx, &user); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if len(userRoles) > 0 {
			if err := l.svcCtx.UserRoleModel.WithTrans(ctx).BatchInsert(l.ctx, userRoles); err != nil {
				return errorx.NewDataBaseError(err)
			}
		}
		return nil
	})
	return err
}
