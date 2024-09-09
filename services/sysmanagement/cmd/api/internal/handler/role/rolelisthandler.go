package role

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/role"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 角色列表
func RoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComFilterFormReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleListLogic(r.Context(), svcCtx)
		resp, err := l.RoleList(&req)
		response.Response(w, resp, err)
	}
}
