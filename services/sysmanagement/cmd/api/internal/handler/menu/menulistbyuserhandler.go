package menu

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/menu"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 根据当前用户获取菜单列表
func MenuListByUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := menu.NewMenuListByUserLogic(r.Context(), svcCtx)
		resp, err := l.MenuListByUser()
		response.Response(w, resp, err)
	}
}
