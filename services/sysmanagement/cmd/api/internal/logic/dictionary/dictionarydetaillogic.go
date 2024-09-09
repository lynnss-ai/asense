package dictionary

import (
	"asense/common/errorx"
	"asense/services/sysmanagement/model"
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
	var dic *model.Dictionary
	dic, err = l.svcCtx.DictionaryModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDataBaseError(err)
	}
	resp = &types.DictionaryResp{
		ID:        dic.ID,
		PID:       dic.PID,
		DicName:   dic.DicName,
		DicCode:   dic.DicCode,
		DicDesc:   dic.DicDesc,
		DicValue:  dic.DicValue,
		DicValue2: dic.DicValue2,
		DicValue3: dic.DicValue3,
		Sort:      dic.Sort,
		IsEdit:    dic.IsEdit,
		IsEnable:  dic.IsEnable,
		CreatedAt: dic.CreatedAt,
		UpdatedAt: dic.UpdatedAt,
	}
	return resp, nil
}
