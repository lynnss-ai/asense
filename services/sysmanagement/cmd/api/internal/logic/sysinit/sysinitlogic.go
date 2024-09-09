package sysinit

import (
	"context"

	"asense/services/sysmanagement/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SysinitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 系统初始化
func NewSysinitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysinitLogic {
	return &SysinitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysinitLogic) Sysinit() error {
	// todo: add your logic here and delete this line

	return nil
}
