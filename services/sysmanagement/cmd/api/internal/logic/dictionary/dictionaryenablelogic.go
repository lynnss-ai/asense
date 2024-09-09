package dictionary

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
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
	var (
		dic *model.Dictionary
		err error
	)
	dic, err = l.svcCtx.DictionaryModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if !dic.IsEdit {
		return errorx.NewDefaultError("该项数据不可进行启用或禁用操作")
	}
	if dic.IsHide {
		return errorx.NewDefaultError("该项数据不可进行启用或禁用操作")
	}
	err = l.svcCtx.DictionaryModel.Enable(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
