package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// PaymentStrategy 定义支付策略接口
type PaymentStrategy interface {
	Pay(ctx context.Context, orderId, payType int32) (*Order, error)
}

type PaymentStrategyUseCase struct {
	paymentStrategy PaymentStrategy
	log             *log.Helper
}

func NewPaymentStrategyCase(paymentStrategy PaymentStrategy, logger log.Logger) *PaymentStrategyUseCase {
	return &PaymentStrategyUseCase{
		paymentStrategy: paymentStrategy,
		log:             log.NewHelper(logger),
	}
}

func (p *PaymentStrategyUseCase) PayOrder(ctx context.Context, OrderId, payType int32) (*Order, error) {
	return p.paymentStrategy.Pay(ctx, OrderId, payType)
}
