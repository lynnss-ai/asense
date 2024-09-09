package user

import (
	"asense/common/errorx"
	"asense/common/utils/timeutil"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户详情
func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.ComIDPathReq) (resp *types.UserResp, err error) {
	var (
		user        *model.User
		userRoleIds []*string
		roles       []*model.Role
		roleIds     []*string
		roleKvList  []*types.ComKvResp
	)

	user, err = l.svcCtx.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	userRoleIds, err = l.svcCtx.UserRoleModel.ListByUserID(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	roles, err = l.svcCtx.RoleModel.ListByIds(l.ctx, userRoleIds)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}

	for _, role := range roles {
		roleIds = append(roleIds, &role.ID)
		roleKvList = append(roleKvList, &types.ComKvResp{
			Key:   role.ID,
			Value: role.RoleName,
		})
	}

	resp = &types.UserResp{
		ID:            user.ID,
		Name:          user.Name,
		UserName:      user.UserName,
		Email:         user.Email,
		Phone:         user.Phone,
		Avatar:        user.Avatar,
		RoleIds:       roleIds,
		Roles:         roleKvList,
		LastLoginTime: timeutil.TimeFormat(user.LastLoginTime),
		IsEnable:      user.IsEnable,
		IsAdmin:       user.IsAdmin,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	return resp, nil
}
