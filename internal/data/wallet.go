package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/ent"
	"kubecit-service/ent/user"
	"kubecit-service/ent/wallet"
	"kubecit-service/internal/biz"
)

type walletRepo struct {
	data *Data
	log  *log.Helper
}

func NewWalletRepo(data *Data, logger log.Logger) biz.WalletRepo {
	return &walletRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (w walletRepo) RechargeGoldLeaf(ctx context.Context, userId, goldLeafAmount int32) (*biz.Wallet, error) {
	walletObj, err := w.data.db.Wallet.Query().Where(wallet.UserID(userId)).First(ctx)
	var wl *ent.Wallet
	if ent.IsNotFound(err) {
		u, err := w.data.db.User.Query().Where(user.IDEQ(int(userId))).First(ctx)
		if err != nil {
			return nil, err
		}
		wl = w.data.db.Wallet.Create().SetUserID(userId).SetGoldLeaf(goldLeafAmount).SetUsername(u.Username).SaveX(ctx)
	} else {
		amount := walletObj.GoldLeaf + goldLeafAmount
		wl = walletObj.Update().SetGoldLeaf(amount).SaveX(ctx)
	}

	return &biz.Wallet{
		Id:               int32(wl.ID),
		GoldLeaf:         wl.GoldLeaf,
		SilverLeaf:       wl.SilverLeaf,
		FrozenGoldLeaf:   wl.FrozenGoldLeaf,
		FrozenSilverLeaf: wl.FrozenSilverLeaf,
		UserId:           wl.UserID,
		UserName:         wl.Username,
		CreateAt:         wl.CreateAt,
		UpdateAt:         wl.UpdateAt,
	}, nil
}

func (w walletRepo) Balance(ctx context.Context, userId int32) (*biz.Wallet, error) {
	walletObj, err := w.data.db.Wallet.Query().Where(wallet.UserID(userId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Wallet{
		Id:               int32(walletObj.ID),
		GoldLeaf:         walletObj.GoldLeaf,
		SilverLeaf:       walletObj.SilverLeaf,
		FrozenGoldLeaf:   walletObj.FrozenGoldLeaf,
		FrozenSilverLeaf: walletObj.FrozenSilverLeaf,
		UserId:           walletObj.UserID,
		UserName:         walletObj.Username,
		CreateAt:         walletObj.CreateAt,
		UpdateAt:         walletObj.UpdateAt,
	}, nil
}
