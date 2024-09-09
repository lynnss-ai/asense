package dictionary

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
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
	var (
		dic   *model.Dictionary
		count int64
		err   error
	)
	dic, err = l.svcCtx.DictionaryModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if !dic.IsEdit {
		return errorx.NewDefaultError("该字典不可删除")
	}
	count, err = l.svcCtx.DictionaryModel.CountByPid(l.ctx, &req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if count > 0 {
		return errorx.NewDefaultError("该字典存在子节点，不能删除")
	}

	err = l.svcCtx.DictionaryModel.Delete(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
