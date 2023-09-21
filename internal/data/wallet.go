package data

import (
	"context"
	"fmt"
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

// 定义一个查询并锁定的函数
func (w walletRepo) QueryAndLock(ctx context.Context, client *ent.Client, userId, goldLeafAmount int32) (*ent.Wallet, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// 使用 `ForUpdate()` 锁定查询操作
	walletObj, err := tx.Wallet.Query().Where(wallet.UserID(userId)).ForUpdate().First(ctx)
	//fmt.Println(err)
	// 执行业务逻辑操作...
	var wl *ent.Wallet
	if ent.IsNotFound(err) {
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		_, err := w.data.db.User.Query().Where(user.IDEQ(int(userId))).First(ctx)
		if err != nil {
			return nil, err
		}
		wl, err = w.data.db.Wallet.Create().SetUserID(userId).SetGoldLeaf(goldLeafAmount).Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		amount := walletObj.GoldLeaf + goldLeafAmount
		wl = walletObj.Update().SetGoldLeaf(amount).SaveX(ctx)
		// 提交事务
		if err := tx.Commit(); err != nil {
			return nil, err
		}
	}

	return wl, nil
}

func (w walletRepo) RechargeGoldLeaf(ctx context.Context, userId, goldLeafAmount int32) (*biz.Wallet, error) {
	wl, err := w.QueryAndLock(ctx, w.data.db, userId, goldLeafAmount)
	fmt.Println(wl)
	if err != nil {
		return nil, err
	}

	return &biz.Wallet{
		Id:               int32(wl.ID),
		GoldLeaf:         wl.GoldLeaf,
		SilverLeaf:       wl.SilverLeaf,
		FrozenGoldLeaf:   wl.FrozenGoldLeaf,
		FrozenSilverLeaf: wl.FrozenSilverLeaf,
		UserId:           wl.UserID,
		CreateAt:         wl.CreatedAt,
		UpdateAt:         wl.UpdatedAt,
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
		CreateAt:         walletObj.CreatedAt,
		UpdateAt:         walletObj.UpdatedAt,
	}, nil
}
