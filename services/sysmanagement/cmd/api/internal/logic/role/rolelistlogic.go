package role

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 角色列表
func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.ComFilterFormReq) (resp *types.ComKvListResp, err error) {
	var (
		list     []*model.Role
		roleList []*types.ComKvResp
	)
	isEnable := true
	list, err = l.svcCtx.RoleModel.ListBySetPermission(l.ctx, req.Filter, &isEnable)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	for _, item := range list {
		roleList = append(roleList, &types.ComKvResp{
			Key:   item.ID,
			Value: item.RoleName,
		})
	}
	return &types.ComKvListResp{Items: roleList}, nil
}
