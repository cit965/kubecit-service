package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

type SliderService struct {
	pb.UnimplementedSliderServer
}

func NewSliderService() *SliderService {
	return &SliderService{}
}

func (s *SliderService) GetSliders(ctx context.Context, req *pb.GetSlidersRequest) (*pb.GetSlidersReply, error) {
	return &pb.GetSlidersReply{}, nil
}
