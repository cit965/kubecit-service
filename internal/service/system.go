package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

// SystemSettings 提供给官网一些系统设置，比如微信二维码，logo
func (s *KubecitService) SystemSettings(ctx context.Context, req *pb.Empty) (*pb.SystemSettingsReply, error) {
	return &pb.SystemSettingsReply{
		Logourl:     "",
		Wechaturl:   "",
		Computerurl: "",
	}, nil
}
