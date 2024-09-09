package dictionary

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
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
	var (
		dic     *model.Dictionary
		isExist bool
		err     error
	)
	dic, err = l.svcCtx.DictionaryModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if !dic.IsEdit {
		return errorx.NewDefaultError("该字典不可编辑")
	}

	isExist, err = l.svcCtx.DictionaryModel.ExistByDicCode(l.ctx, &req.ID, req.DicCode)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if isExist {
		return errorx.NewDefaultError("字典编码已存在")
	}

	err = l.svcCtx.DictionaryModel.Update(l.ctx, req.ID, map[string]interface{}{
		"pid":        req.PID,
		"dic_name":   req.DicName,
		"dic_code":   req.DicCode,
		"dic_desc":   req.DicDesc,
		"dic_value":  req.DicValue,
		"dic_value2": req.DicValue2,
		"dic_value3": req.DicValue3,
		"sort":       req.Sort,
	})
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
