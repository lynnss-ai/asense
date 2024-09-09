package dictionary

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 字典详情
func NewDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryDetailLogic {
	return &DictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryDetailLogic) DictionaryDetail(req *types.ComIDPathReq) (resp *types.DictionaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
