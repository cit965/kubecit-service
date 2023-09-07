package service

import (
	"context"
	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) LoginByJson(ctx context.Context, req *pb.LoginByJsonRequest) (*pb.LoginByJsonReply, error) {
	return s.userUseCase.LoginByJson(ctx, req)
}

func (s *KubecitService) RegisterUsername(ctx context.Context, req *pb.RegisterUsernameRequest) (*pb.RegisterUsernameReply, error) {
	return s.userUseCase.RegisterUsername(ctx, req)
}

func (s *KubecitService) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	return &pb.CreateTokenReply{}, nil
}
