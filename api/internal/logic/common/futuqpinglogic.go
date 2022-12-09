package common

import (
	"context"

	"futuq/api/internal/svc"
	"futuq/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FutuqPingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFutuqPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FutuqPingLogic {
	return &FutuqPingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FutuqPingLogic) FutuqPing() (resp *types.FutuqPingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
