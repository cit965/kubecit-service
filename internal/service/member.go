package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

type MemberService struct {
	pb.UnimplementedMemberServer
}

func NewMemberService() *MemberService {
	return &MemberService{}
}

func (s *MemberService) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoReply, error) {
	return &pb.GetInfoReply{}, nil
}
