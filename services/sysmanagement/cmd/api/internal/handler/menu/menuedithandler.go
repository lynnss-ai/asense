package menu

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/menu"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 编辑菜单
func MenuEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewMenuEditLogic(r.Context(), svcCtx)
		err := l.MenuEdit(&req)
		response.Response(w, nil, err)
	}
}
