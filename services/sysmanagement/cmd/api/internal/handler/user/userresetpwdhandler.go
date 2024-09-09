package user

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/user"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 重置密码
func UserResetPwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComUserResetPwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserResetPwdLogic(r.Context(), svcCtx)
		err := l.UserResetPwd(&req)
		response.Response(w, nil, err)
	}
}
