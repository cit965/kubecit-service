package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) RechargeWallet(ctx context.Context, req *pb.RechargeWalletRequest) (*pb.WalletInfo, error) {
	wallet, err := s.walletCase.RechargeGoldLeafs(ctx, req.UserId, req.GoldLeafAmount)
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

func (s KubecitService) WalletBalance(ctx context.Context, req *pb.Empty) (*pb.WalletInfo, error) {
	userRow := ctx.Value("user_id")
	userId, ok := userRow.(int32)
	if !ok {
		return nil, errors.BadRequest("解析不出用户ID", "用户未登录")
	}
	//userId := int32(1)
	wallet, err := s.walletCase.Balances(ctx, userId)
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
