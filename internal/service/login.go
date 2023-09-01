package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

type LoginService struct {
	pb.UnimplementedLoginServer
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) LoginByJson(ctx context.Context, req *pb.LoginByJsonRequest) (*pb.LoginByJsonReply, error) {
	return &pb.LoginByJsonReply{}, nil
}
