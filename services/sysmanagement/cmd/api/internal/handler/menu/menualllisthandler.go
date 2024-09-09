package menu

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/menu"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 获取所有菜单
func MenuAllListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := menu.NewMenuAllListLogic(r.Context(), svcCtx)
		resp, err := l.MenuAllList()
		response.Response(w, resp, err)
	}
}
