package user

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/user"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 刷新用户Token
func UserRefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUserRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.UserRefreshToken()
		response.Response(w, resp, err)
	}
}
