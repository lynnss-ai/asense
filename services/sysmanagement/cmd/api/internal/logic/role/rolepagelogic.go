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

type RolePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 角色分页列表
func NewRolePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePageLogic {
	return &RolePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolePageLogic) RolePage(req *types.RolePageReq) (resp *types.RolePageResp, err error) {
	var (
		total      int64
		list       []*model.Role
		resultList []*types.RoleResp
	)
	total, list, err = l.svcCtx.RoleModel.FindPage(l.ctx, req.Page, req.PageSize, req.Filter, req.IsEnable)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	resultList = make([]*types.RoleResp, 0, len(list))
	for _, item := range list {
		var (
			menuIds         []*string
			selectedMenuIds []string
		)
		menuIds, err = l.svcCtx.RolePermissionModel.ListByRoleID(l.ctx, item.ID)
		if err != nil {
			return nil, errorx.NewDataBaseError(err)
		}

		if item.SelectedMenuIds != "" {
			selectedMenuIds = strings.Split(item.SelectedMenuIds, ",")
		} else {
			selectedMenuIds = []string{}
		}

		resultList = append(resultList, &types.RoleResp{
			ID:              item.ID,
			RoleName:        item.RoleName,
			RoleCode:        item.RoleCode,
			RoleDesc:        item.RoleDesc,
			IsSetPermission: item.IsSetPermission,
			IsEnable:        item.IsEnable,
			IsAdmin:         item.IsAdmin,
			SelectedMenuIds: selectedMenuIds,
			MenuIds:         menuIds,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}

	return &types.RolePageResp{Total: total, Items: resultList}, nil
}
