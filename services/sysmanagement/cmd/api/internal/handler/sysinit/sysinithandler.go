package sysinit

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/sysinit"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 系统初始化
func SysinitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := sysinit.NewSysinitLogic(r.Context(), svcCtx)
		err := l.Sysinit()
		response.Response(w, nil, err)
	}
}
