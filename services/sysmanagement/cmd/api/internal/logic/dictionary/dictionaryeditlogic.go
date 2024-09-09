package dictionary

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"asense/services/sysmanagement/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑字典
func NewDictionaryEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryEditLogic {
	return &DictionaryEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryEditLogic) DictionaryEdit(req *types.DictionaryEditReq) error {
	// todo: add your logic here and delete this line

	return nil
}
