package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) LoginByJson(ctx context.Context, req *pb.LoginByJsonRequest) (*pb.LoginByJsonReply, error) {
	return &pb.LoginByJsonReply{}, nil
}

func (s *KubecitService) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	return &pb.CreateTokenReply{}, nil
}
