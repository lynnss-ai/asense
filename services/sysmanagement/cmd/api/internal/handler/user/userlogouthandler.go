package user

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/user"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 退出登录
func UserLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUserLogoutLogic(r.Context(), svcCtx)
		err := l.UserLogout()
		response.Response(w, nil, err)
	}
}
