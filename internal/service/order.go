package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
)

// CreateOrder 创建订单
func (s *KubecitService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	//fmt.Println(req.GetCourseIds())
	order, err := s.orderCase.Create(ctx, req.GetCourseIds())
	if err != nil {
		return nil, err
	}

	details := make([]*pb.OrderDetail, 0)
	for _, info := range order.Info {
		details = append(details, &pb.OrderDetail{
			OrderId:           info.OrderId,
			CourseId:          info.CourseId,
			CourseName:        info.CourseName,
			CoursePrice:       info.CoursePrice,
			CourseDescription: info.CourseDescribe,
		})
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
		Details:    details,
	}, nil
}
