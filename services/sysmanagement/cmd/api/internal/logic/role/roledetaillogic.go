package role

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"strings"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 角色详情
func NewRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDetailLogic {
	return &RoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDetailLogic) RoleDetail(req *types.ComIDPathReq) (resp *types.RoleResp, err error) {
	var (
		role            *model.Role
		menuIds         []*string
		selectedMenuIds []string
	)
	role, err = l.svcCtx.RoleModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	menuIds, err = l.svcCtx.RolePermissionModel.ListByRoleID(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}

	if role.SelectedMenuIds != "" {
		selectedMenuIds = strings.Split(role.SelectedMenuIds, ",")
	} else {
		selectedMenuIds = []string{}
	}

	resp = &types.RoleResp{
		ID:              role.ID,
		RoleName:        role.RoleName,
		RoleCode:        role.RoleCode,
		RoleDesc:        role.RoleDesc,
		IsSetPermission: role.IsSetPermission,
		IsEnable:        role.IsEnable,
		IsAdmin:         role.IsAdmin,
		MenuIds:         menuIds,
		SelectedMenuIds: selectedMenuIds,
		CreatedAt:       role.CreatedAt,
		UpdatedAt:       role.UpdatedAt,
	}
	return resp, nil
}
