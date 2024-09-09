package dictionary

import (
	"asense/common/dbcore"
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
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
	var (
		isExist bool
		err     error
	)
	isExist, err = l.svcCtx.DictionaryModel.ExistByDicCode(l.ctx, nil, req.DicCode)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	if isExist {
		return errorx.NewDefaultError("字典编码已存在")
	}
	dic := model.Dictionary{
		ID:        dbcore.NewId(),
		PID:       req.PID,
		DicName:   req.DicName,
		DicCode:   req.DicCode,
		DicDesc:   req.DicDesc,
		DicValue:  req.DicValue,
		DicValue2: req.DicValue2,
		DicValue3: req.DicValue3,
		Sort:      req.Sort,
		IsEdit:    true,
		IsEnable:  true,
		IsHide:    false,
	}
	err = l.svcCtx.DictionaryModel.Insert(l.ctx, &dic)
	if err != nil {
		return errorx.NewDataBaseError(err)
	}
	return nil
}
