package user

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/user"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 用户分页列表
func UserPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserPageLogic(r.Context(), svcCtx)
		resp, err := l.UserPage(&req)
		response.Response(w, resp, err)
	}
}
