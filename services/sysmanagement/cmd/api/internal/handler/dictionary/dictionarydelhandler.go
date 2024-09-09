package dictionary

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/dictionary"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 删除字典
func DictionaryDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComIDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewDictionaryDelLogic(r.Context(), svcCtx)
		err := l.DictionaryDel(&req)
		response.Response(w, nil, err)
	}
}
