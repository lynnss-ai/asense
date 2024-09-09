package role

import (
	"asense/common/components"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑角色
func NewRoleEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleEditLogic {
	return &RoleEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleEditLogic) RoleEdit(req *types.RoleEditReq) error {
	var (
		user *model.User
		err  error
	)
	userID := components.GetAuthKeyJwtUserID(l.ctx)
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if !user.IsAdmin {
		return errorx.NewDefaultError("非管理员不能编辑角色")
	}

	err = l.svcCtx.RoleModel.Update(l.ctx, req.ID,
		map[string]interface{}{
			"role_name": req.RoleName,
			"role_desc": req.RoleDesc,
		})
	if err != nil {
		return errorx.NewDataBaseError(err)
	}

	return nil
}
