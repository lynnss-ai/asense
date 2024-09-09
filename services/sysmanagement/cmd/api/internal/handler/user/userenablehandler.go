package user

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/user"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 启用|禁用用户
func UserEnableHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComIDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserEnableLogic(r.Context(), svcCtx)
		err := l.UserEnable(&req)
		response.Response(w, nil, err)
	}
}
