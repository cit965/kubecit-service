package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Order struct {
	Id         int
	UserId     int32
	OrderSn    string
	PayType    int32
	PayStatus  int32
	TradePrice int32
	TradeNo    string
	PayTime    time.Time
	CreateTime time.Time
	UpdateTime time.Time
	Details    []*OrderInfo
}

type OrderInfo struct {
	Id              int
	OrderId         int32
	ProductId       int32
	ProductName     string
	ProductPrice    int32
	ProductDescribe string
	CreateTime      time.Time
	UpdateTime      time.Time
}

type OrderRepo interface {
	Create(ctx context.Context, courseIds []int32) (*Order, error)
	MyOrder(ctx context.Context, pageNum, pageSize *int32) ([]*Order, error)
}

type OrderUseCase struct {
	orderRepo OrderRepo
	log       *log.Helper
}

func NewOrderUseCase(orderRepo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
		log:       log.NewHelper(logger),
	}
}

func (oc *OrderUseCase) Create(ctx context.Context, courseIds []int32) (*Order, error) {
	return oc.orderRepo.Create(ctx, courseIds)
}

func (oc *OrderUseCase) MyOrder(ctx context.Context, pageNum, pageSize *int32) ([]*Order, error) {
	return oc.orderRepo.MyOrder(ctx, pageNum, pageSize)
}
