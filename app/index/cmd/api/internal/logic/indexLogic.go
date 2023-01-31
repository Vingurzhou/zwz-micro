package logic

import (
	"context"

	"looklook/app/index/cmd/api/internal/svc"
	"looklook/app/index/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	return &types.Response{
		Code: 200,
		Msg:  "Welcome zhouwenzhe's Microservices",
	}, nil
}
