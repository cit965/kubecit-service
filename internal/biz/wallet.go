package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Wallet struct {
	Id               int32
	GoldLeaf         int32
	SilverLeaf       int32
	FrozenGoldLeaf   int32
	FrozenSilverLeaf int32
	UserId           int32
	CreateAt         time.Time
	UpdateAt         time.Time
}

type WalletRepo interface {
	RechargeGoldLeaf(ctx context.Context, userId, goldLeafAmount int32) (*Wallet, error)
	Balance(ctx context.Context, userId int32) (*Wallet, error)
}

type WalletUseCase struct {
	walletRepo WalletRepo
	log        *log.Helper
}

func NewWalletUseCase(walletRepo WalletRepo, logger log.Logger) *WalletUseCase {
	return &WalletUseCase{
		walletRepo: walletRepo,
		log:        log.NewHelper(logger),
	}
}

func (wac *WalletUseCase) RechargeGoldLeafs(ctx context.Context, userId, goldLeafAmount int32) (*Wallet, error) {
	return wac.walletRepo.RechargeGoldLeaf(ctx, userId, goldLeafAmount)
}

func (wac WalletUseCase) Balances(ctx context.Context, userId int32) (*Wallet, error) {
	return wac.walletRepo.Balance(ctx, userId)

}
