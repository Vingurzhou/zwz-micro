package product

import (
	"context"

	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductInfoLogic {
	return &GetProductInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductInfoLogic) GetProductInfo(req *types.GetProductInfoReq) (resp *types.GetProductInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
