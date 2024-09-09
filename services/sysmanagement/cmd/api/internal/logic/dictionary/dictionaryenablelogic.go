package dictionary

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryEnableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 启用|禁用字典
func NewDictionaryEnableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryEnableLogic {
	return &DictionaryEnableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryEnableLogic) DictionaryEnable(req *types.ComIDPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
