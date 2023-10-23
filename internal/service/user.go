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

func (s *KubecitService) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.UserInfoReply, error) {
	return s.userUseCase.CurrentUserInfo(ctx)
}

func (s *KubecitService) BecomeVip(ctx context.Context, req *pb.BecomeVipRequest) (*pb.BecomeVipReply, error) {
	return s.userUseCase.Become(ctx, req)
}

func (s *KubecitService) TradeCallback(ctx context.Context, req *pb.TradeCallbackRequest) (*pb.TradeCallbackReply, error) {
	return s.userUseCase.Callback(ctx, req)
}
