package role

import (
	"asense/common/components"
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/common/utils/characterutil"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增角色
func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req *types.RoleReq) error {
	var (
		user    *model.User
		isExist bool
		err     error
	)
	userID := components.GetAuthKeyJwtUserID(l.ctx)

	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if !user.IsAdmin {
		return errorx.NewDefaultError("非管理员不能添加角色")
	}
	isExist, err = l.svcCtx.RoleModel.ExistByRoleCode(l.ctx, req.RoleCode)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if isExist {
		return errorx.NewDefaultError("角色编码已存在")
	}

	role := model.Role{
		ID:              dbcore.NewId(),
		RoleName:        req.RoleName,
		RoleCode:        characterutil.StitchingBuilderStr(req.RoleCode),
		RoleDesc:        req.RoleDesc,
		SelectedMenuIds: "",
		IsAdmin:         false,
		IsSetPermission: false,
		IsEnable:        true,
	}

	err = l.svcCtx.RoleModel.Insert(l.ctx, &role)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
