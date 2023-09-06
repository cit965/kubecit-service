package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
)

func (s *KubecitService) CreateSlider(ctx context.Context, req *pb.CreateSliderRequest) (*pb.CreateSliderReply, error) {
	ins := &biz.Slider{
		Title:     req.Title,
		Content:   req.Content,
		ImageLink: req.ImageLink,
		IsValid:   req.IsValid,
		Priority:  int(req.Priority),
	}
	res, err := s.su.CreateSlider(ctx, ins)
	if err != nil {
		return nil, err
	}
	return &pb.CreateSliderReply{
		Data: &pb.SliderInfo{
			Id:        int32(res.Id),
			Title:     res.Title,
			Content:   res.Content,
			ImageLink: res.ImageLink,
			CreateAt:  timestamppb.New(res.CreateAt),
			UpdateAt:  timestamppb.New(res.UpdateAt),
			IsValid:   res.IsValid,
			Priority:  int32(res.Priority),
		},
	}, nil
}

func (s *KubecitService) GetSlider(ctx context.Context, req *pb.GetSliderRequest) (*pb.GetSliderReply, error) {
	res, err := s.su.GetSliderById(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetSliderReply{
		Data: &pb.SliderInfo{
			Id:        int32(res.Id),
			Title:     res.Title,
			Content:   res.Content,
			ImageLink: res.ImageLink,
			CreateAt:  timestamppb.New(res.CreateAt),
			UpdateAt:  timestamppb.New(res.UpdateAt),
			IsValid:   res.IsValid,
			Priority:  int32(res.Priority),
		},
	}, nil
}

func (s *KubecitService) DeleteSlider(ctx context.Context, req *pb.DeleteSliderRequest) (*pb.DeleteSliderReply, error) {
	res, err := s.su.DeleteById(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSliderReply{
		Count: int32(res),
	}, nil
}

func (s *KubecitService) UpdateSlider(ctx context.Context, req *pb.UpdateSliderRequest) (*pb.UpdateSliderReply, error) {
	ins := &biz.Slider{
		Title:     req.Title,
		Content:   req.Content,
		ImageLink: req.ImageLink,
		IsValid:   req.IsValid,
		Priority:  int(req.Priority),
	}
	res, err := s.su.UpdateById(ctx, int(req.Id), ins)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSliderReply{
		Data: &pb.SliderInfo{
			Id:        int32(res.Id),
			Title:     res.Title,
			Content:   res.Content,
			ImageLink: res.ImageLink,
			CreateAt:  timestamppb.New(res.CreateAt),
			UpdateAt:  timestamppb.New(res.UpdateAt),
			IsValid:   res.IsValid,
			Priority:  int32(res.Priority),
		},
	}, nil
}

func (s *KubecitService) ListSlidersByPriority(ctx context.Context, req *pb.ListSlidersByPriorityRequest) (*pb.ListSlidersByPriorityReply, error) {
	res, err := s.su.ListByPriority(ctx, int(req.GetPriority()), int(req.GetCount()))
	if err != nil {
		return nil, err
	}
	dataSet := make([]*pb.SliderInfo, 0)
	for _, value := range res {
		data := &pb.SliderInfo{
			Id:        int32(value.Id),
			Title:     value.Title,
			Content:   value.Content,
			ImageLink: value.ImageLink,
			CreateAt:  timestamppb.New(value.CreateAt),
			UpdateAt:  timestamppb.New(value.UpdateAt),
			IsValid:   value.IsValid,
			Priority:  int32(value.Priority),
		}
		dataSet = append(dataSet, data)
	}
	return &pb.ListSlidersByPriorityReply{
		Data: dataSet,
	}, nil
}
