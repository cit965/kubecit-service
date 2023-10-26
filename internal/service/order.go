package service

import (
	"context"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/pkg/common"
)

// CreateOrder 创建订单
func (s *KubecitService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {

	order, err := s.orderCase.Create(ctx, req.GetCourseIds())
	if err != nil {
		return nil, err
	}
	var orderResponse pb.CreateOrderReply
	err = common.DeepCopyConvertType(&orderResponse, order)
	if err != nil {
		return nil, err
	}
	return &orderResponse, nil
}

func (s KubecitService) MyOrderList(ctx context.Context, req *pb.MyOrderRequest) (*pb.OrderListReply, error) {
	orders, err := s.orderCase.MyOrder(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	orderList := make([]*pb.CreateOrderReply, 0)
	err = common.DeepCopyConvertType(&orderList, orders)
	if err != nil {
		return nil, err
	}
	return &pb.OrderListReply{
		OrderList: orderList,
	}, nil
}
