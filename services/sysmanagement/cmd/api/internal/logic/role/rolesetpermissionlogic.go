package role

import (
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"
	"strings"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleSetPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置角色权限
func NewRoleSetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleSetPermissionLogic {
	return &RoleSetPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleSetPermissionLogic) RoleSetPermission(req *types.RoleSetPermissionReq) error {
	var rolePermissions []*model.RolePermission

	for _, menuId := range req.MenuIds {
		rolePermissions = append(rolePermissions, &model.RolePermission{
			ID:     dbcore.NewId(),
			RoleID: req.RoleID,
			MenuID: menuId,
		})
	}
	err := l.svcCtx.Tx.ExecTx(l.ctx, func(ctx context.Context) error {
		if err := l.svcCtx.RolePermissionModel.WithTrans(ctx).DeleteByRoleID(l.ctx, req.RoleID); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.RoleModel.WithTrans(ctx).Update(l.ctx, req.RoleID, map[string]interface{}{
			"is_set_permission": true,
			"selected_menu_ids": strings.Join(req.SelectedMenuIds, ","),
		}); err != nil {
			return errorx.NewDataBaseError(err)
		}
		if err := l.svcCtx.RolePermissionModel.WithTrans(ctx).BatchInsert(l.ctx, rolePermissions); err != nil {
			return errorx.NewDataBaseError(err)
		}
		return nil
	})
	return err
}
