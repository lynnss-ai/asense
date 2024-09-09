package role

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/role"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 删除角色
func RoleDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComIDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleDelLogic(r.Context(), svcCtx)
		err := l.RoleDel(&req)
		response.Response(w, nil, err)
	}
}
