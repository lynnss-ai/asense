package dictionary

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增字典
func NewDictionaryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryAddLogic {
	return &DictionaryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryAddLogic) DictionaryAdd(req *types.DictionaryReq) error {
	// todo: add your logic here and delete this line

	return nil
}
