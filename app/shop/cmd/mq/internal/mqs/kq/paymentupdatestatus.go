package kq

import (
	"context"
	"encoding/json"
	"fmt"
	"looklook/app/shop/cmd/mq/internal/svc"
	"looklook/common/kqueue"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type PaymentUpdateStatusMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentUpdateStatusMq {
	return &PaymentUpdateStatusMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentUpdateStatusMq) Consume(_, val string) error {

	fmt.Printf(" PaymentUpdateStatusMq Consume val : %s \n", val)
	//解析数据
	var message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	//执行业务..
	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *PaymentUpdateStatusMq) execService(message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage) error {

	orderTradeState := l.getOrderTradeStateByPaymentTradeState(message.PayStatus)
	if orderTradeState != -99 {
		//update homestay order state
		//_, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(l.ctx, &order.UpdateHomestayOrderTradeStateReq{
		//	Sn:         message.OrderSn,
		//	TradeState: orderTradeState,
		//})
		//if err != nil {
		//	return errors.Wrapf(xerr.NewErrMsg("update homestay order state fail"), "update homestay order state fail err : %v ,message:%+v", err, message)
		//}
	}

	return nil
}

// Get order status based on payment status.
func (l *PaymentUpdateStatusMq) getOrderTradeStateByPaymentTradeState(thirdPaymentPayStatus int64) int64 {

	switch thirdPaymentPayStatus {
	//case paymentModel.ThirdPaymentPayTradeStateSuccess:
	//	return model.HomestayOrderTradeStateWaitUse
	//case paymentModel.ThirdPaymentPayTradeStateRefund:
	//	return model.HomestayOrderTradeStateRefund
	default:
		return -99
	}

}
