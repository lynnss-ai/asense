package dictionary

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除字典
func NewDictionaryDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryDelLogic {
	return &DictionaryDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryDelLogic) DictionaryDel(req *types.ComIDPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
