package dictionary

import (
	"asense/common/response"
	"asense/services/sysmanagement/cmd/api/internal/logic/dictionary"
	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 启用|禁用字典
func DictionaryEnableHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComIDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewDictionaryEnableLogic(r.Context(), svcCtx)
		err := l.DictionaryEnable(&req)
		response.Response(w, nil, err)
	}
}
