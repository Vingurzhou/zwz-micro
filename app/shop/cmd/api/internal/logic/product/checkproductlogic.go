package product

import (
	"context"
	"fmt"
	"sync"

	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var sum int64 = 0
var productNum int64 = 10000
var mutex sync.Mutex

type CheckProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductLogic {
	return &CheckProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckProductLogic) CheckProduct(req *types.CheckProductReq) (resp *types.CheckProductResp, err error) {
	mutex.Lock()
	defer mutex.Unlock()
	var result bool
	if sum < productNum {
		sum += 1
		result = true
	}
	result = false
	fmt.Println(result)
	return
}
