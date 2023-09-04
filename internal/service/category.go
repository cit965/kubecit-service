package service

import (
	"context"
	"kubecit-service/internal/biz"

	pb "kubecit-service/api/helloworld/v1"
)

type CategoryService struct {
	pb.UnimplementedCategoryServer

	cs *biz.CategoryUsecase
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) Category(ctx context.Context, req *pb.Empty) (*pb.CategoryResp, error) {
	categories, err := s.cs.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	var cs []*pb.CategoryInfo
	for _, v := range categories {
		cs = append(cs, &pb.CategoryInfo{
			CategoryName: v.CategoryName,
			Id:           v.Id,
			ParentId:     v.ParentId,
			Level:        v.Level,
			Status:       v.Status,
		})
	}
	return &pb.CategoryResp{Categories: cs}, nil
}
