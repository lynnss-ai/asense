package dictionary

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/dictionary"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 字典列表
func DictionaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewDictionaryListLogic(r.Context(), svcCtx)
		resp, err := l.DictionaryList(&req)
		response.Response(w, resp, err)
	}
}
