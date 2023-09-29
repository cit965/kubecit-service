package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
)

func (s KubecitService) OrderPay(ctx context.Context, req *pb.OrderPayRequest) (*pb.CreateOrderReply, error) {

	order, err := s.paymentStrategyUseCase.PayOrder(ctx, req.OrderId, int32(req.PayType))
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderReply{
		UserId:     order.UserId,
		OrderSn:    order.OrderSn,
		PayType:    pb.PayType(order.PayType),
		PayStatus:  pb.PayStatus(order.PayStatus),
		TradePrice: order.TradePrice,
		TradeNo:    order.TradeNo,
		PayTime:    timestamppb.New(order.PayTime),
		CreateTime: timestamppb.New(order.CreateTime),
		UpdateTime: timestamppb.New(order.UpdateTime),
	}, nil
}
