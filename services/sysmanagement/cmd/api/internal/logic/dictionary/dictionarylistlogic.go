package dictionary

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
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
	var (
		list       []*model.Dictionary
		resultList []*types.DictionaryResp
	)
	isHide := false

	list, err = l.svcCtx.DictionaryModel.List(l.ctx, nil, req.IsEnable, nil, &isHide, req.Filter)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}

	for _, item := range list {
		resultList = append(resultList, &types.DictionaryResp{
			ID:        item.ID,
			PID:       item.PID,
			DicName:   item.DicName,
			DicCode:   item.DicCode,
			DicDesc:   item.DicDesc,
			DicValue:  item.DicValue,
			DicValue2: item.DicValue2,
			DicValue3: item.DicValue3,
			Sort:      item.Sort,
			IsEdit:    item.IsEdit,
			IsEnable:  item.IsEnable,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return &types.DictionaryListResp{Items: resultList}, nil
}
