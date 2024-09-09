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
	var (
		total      int64
		list       []*model.User
		resultList []*types.UserResp
	)
	total, list, err = l.svcCtx.UserModel.FindPage(l.ctx, req.Page, req.PageSize, req.Filter, req.IsEnable)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}

	for _, item := range list {
		var (
			roleIds    []*string
			roles      []*model.Role
			roleKvList []*types.ComKvResp
		)
		roleIds, err = l.svcCtx.UserRoleModel.ListByUserID(l.ctx, item.ID)
		if err != nil {
			return nil, errorx.NewDataBaseError(err)
		}
		roles, err = l.svcCtx.RoleModel.ListByIds(l.ctx, roleIds)
		if err != nil {
			return nil, errorx.NewDataBaseError(err)
		}

		for _, role := range roles {
			roleKvList = append(roleKvList, &types.ComKvResp{
				Key:   role.ID,
				Value: role.RoleName,
			})
		}

		resultList = append(resultList, &types.UserResp{
			ID:            item.ID,
			Name:          item.Name,
			Avatar:        item.Avatar,
			UserName:      item.UserName,
			Email:         item.Email,
			Phone:         item.Phone,
			LastLoginTime: timeutil.TimeFormat(item.LastLoginTime),
			IsEnable:      item.IsEnable,
			IsAdmin:       item.IsAdmin,
			RoleIds:       roleIds,
			Roles:         roleKvList,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
		})
	}
	return &types.UserPageResp{
		Total: total,
		Items: resultList,
	}, nil
}
