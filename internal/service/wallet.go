package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) RechargeWallet(ctx context.Context, req *pb.RechargeWalletRequest) (*pb.WalletInfo, error) {
	wallet, err := s.walletCase.RechargeGoldLeaf(ctx, req.UserId, req.GoldLeafAmount)
	if err != nil {
		return nil, err
	}
	return &pb.WalletInfo{
		GoldLeaf:         wallet.GoldLeaf,
		SilverLeaf:       wallet.SilverLeaf,
		FrozenGoldLeaf:   wallet.FrozenGoldLeaf,
		FrozenSilverLeaf: wallet.FrozenSilverLeaf,
		UserId:           wallet.UserId,
		UserName:         wallet.UserName,
		CreateAt:         timestamppb.New(wallet.CreateAt),
		UpdateAt:         timestamppb.New(wallet.UpdateAt),
		Id:               wallet.Id,
	}, nil
}
