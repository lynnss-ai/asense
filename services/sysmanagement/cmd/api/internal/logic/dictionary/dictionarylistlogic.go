package dictionary

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 字典列表
func NewDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryListLogic {
	return &DictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryListLogic) DictionaryList(req *types.DictionaryListReq) (resp *types.DictionaryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
