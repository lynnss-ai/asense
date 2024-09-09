package role

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/role"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 设置角色权限
func RoleSetPermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleSetPermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleSetPermissionLogic(r.Context(), svcCtx)
		err := l.RoleSetPermission(&req)
		response.Response(w, nil, err)
	}
}
