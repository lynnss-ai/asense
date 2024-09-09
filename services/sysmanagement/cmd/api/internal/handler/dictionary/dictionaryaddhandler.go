package dictionary

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/dictionary"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 新增字典
func DictionaryAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewDictionaryAddLogic(r.Context(), svcCtx)
		err := l.DictionaryAdd(&req)
		response.Response(w, nil, err)
	}
}
