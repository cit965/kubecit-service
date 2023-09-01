package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

type TokenService struct {
	pb.UnimplementedTokenServer
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	return &pb.CreateTokenReply{}, nil
}
