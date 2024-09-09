package role

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/role"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 角色分页列表
func RolePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RolePageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRolePageLogic(r.Context(), svcCtx)
		resp, err := l.RolePage(&req)
		response.Response(w, resp, err)
	}
}
