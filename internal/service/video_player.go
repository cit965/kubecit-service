package service

import (
	"context"
	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) VideoPlayerGetPlayerParam(ctx context.Context, req *pb.VideoPlayerGetParamReq) (*pb.VideoPlayerGetParamReply, error) {
	return s.videoPlayerUseCase.GetPlayerParam(ctx, req)
}
