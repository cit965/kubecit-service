package service

import (
	"context"

	v1 "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) Category(ctx context.Context, in *v1.Empty) (*v1.CategoryResp, error) {
	categories, err := s.uc.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	var cs []*v1.Category
	for _, v := range categories {
		cs = append(cs, &v1.Category{
			CategoryName: v.CategoryName,
			Id:           v.Id,
		})
	}
	return &v1.CategoryResp{Categories: cs}, nil
}
