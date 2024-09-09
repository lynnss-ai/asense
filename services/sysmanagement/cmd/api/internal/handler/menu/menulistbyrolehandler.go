package menu

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/menu"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 根据角色获取菜单列表
func MenuListByRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComIDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewMenuListByRoleLogic(r.Context(), svcCtx)
		resp, err := l.MenuListByRole(&req)
		response.Response(w, resp, err)
	}
}
