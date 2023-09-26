package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/pkg/common"
)

func (s *KubecitService) RechargeWallet(ctx context.Context, req *pb.RechargeWalletRequest) (*pb.WalletInfo, error) {
	wallet, err := s.walletCase.RechargeGoldLeafs(ctx, req.UserId, req.GoldLeafAmount, req.SilverLeafAmount)
	if err != nil {
		return nil, err
	}
	return &pb.WalletInfo{
		GoldLeaf:         &wallet.GoldLeaf,
		SilverLeaf:       &wallet.SilverLeaf,
		FrozenGoldLeaf:   &wallet.FrozenGoldLeaf,
		FrozenSilverLeaf: &wallet.FrozenSilverLeaf,
		UserId:           &wallet.UserId,
		CreatedAt:        timestamppb.New(wallet.CreateAt),
		UpdatedAt:        timestamppb.New(wallet.UpdateAt),
		Id:               &wallet.Id,
	}, nil
}

func (s KubecitService) WalletBalance(ctx context.Context, req *pb.Empty) (*pb.WalletInfo, error) {
	userId, err := common.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	//userId := int32(100)
	wallet, err := s.walletCase.Balances(ctx, int32(userId))
	if err != nil {
		return nil, err
	}

	return &pb.WalletInfo{
		GoldLeaf:         &wallet.GoldLeaf,
		SilverLeaf:       &wallet.SilverLeaf,
		FrozenGoldLeaf:   &wallet.FrozenGoldLeaf,
		FrozenSilverLeaf: &wallet.FrozenSilverLeaf,
		UserId:           &wallet.UserId,
		CreatedAt:        timestamppb.New(wallet.CreateAt),
		UpdatedAt:        timestamppb.New(wallet.UpdateAt),
		Id:               &wallet.Id,
	}, nil
}
