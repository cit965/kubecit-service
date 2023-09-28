package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/ent"
	"kubecit-service/ent/orders"
	"kubecit-service/ent/wallet"
	"kubecit-service/internal/biz"
	"kubecit-service/internal/pkg/common"
	"kubecit-service/internal/pkg/errs"
	"time"
)

type Strategy struct {
	data *Data
	log  *log.Helper
}

func NewStrategy(data *Data, logger log.Logger) biz.PaymentStrategy {
	return &Strategy{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *Strategy) Pay(ctx context.Context, orderId, payType int32) (*biz.Order, error) {
	// 实现金叶子支付逻辑
	fmt.Println("@@@@@")
	result, err := g.data.WithResultTx(ctx, func(tx *ent.Tx) (interface{}, error) {
		var orderTx *biz.Order
		var err error
		if payType == int32(pb.PayType_GOLDEN_LEAF) {
			orderTx, err = g.GoldLeafDeductionTX(ctx, orderId)
			if err != nil {
				return nil, err
			}
		} else if payType == int32(pb.PayType_SILVER_LEAF) {
			orderTx, err = g.SilverLeafDeductionTX(ctx, orderId)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("目前仅支持金叶子和银叶子支付")
		}
		return orderTx, nil
	})
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	if order, ok := result.(*biz.Order); ok {
		return order, nil
	}
	return nil, err
}

func (g *Strategy) GoldLeafDeductionTX(ctx context.Context, orderId int32) (*biz.Order, error) {
	//userId, err := common.GetUserFromCtx(ctx)
	//if err != nil {
	//	return nil, err
	//}
	userId := int32(1)
	order, err := g.data.db.Orders.Query().Where(orders.IDEQ(int(orderId))).First(ctx)
	if ent.IsNotFound(err) {
		return nil, errs.OrderNotFound(err)
	} else {
		walletObj, err := g.data.db.Wallet.Query().Where(wallet.UserID(int32(userId))).First(ctx)
		if ent.IsNotFound(err) {
			return nil, errs.NotEnoughMoney(err)
		} else {
			fmt.Println(order.TradePrice, walletObj.GoldLeaf)
			if order.TradePrice > walletObj.GoldLeaf {
				return nil, errs.NotEnoughMoney(errors.New("余额不足，请及时充值！"))
			} else {
				goldLeafAmount := walletObj.GoldLeaf - order.TradePrice
				_, err = walletObj.Update().SetGoldLeaf(goldLeafAmount).Save(ctx)
				if err != nil {
					_, err := order.Update().SetPayTime(time.Now()).SetPayStatus(int32(pb.PayStatus_FAILED)).SetPayType(int32(pb.PayType_GOLDEN_LEAF)).Save(ctx)
					return nil, errs.BadError(err, "订单支付失败！")
				}
				ord, err := order.Update().SetPayTime(time.Now()).SetPayStatus(int32(pb.PayStatus_PAID)).SetPayType(int32(pb.PayType_GOLDEN_LEAF)).Save(ctx)
				if err != nil {
					return nil, errs.BadError(err, "订单支付失败！")
				}
				return &biz.Order{
					Id:         ord.ID,
					UserId:     ord.UserID,
					OrderSn:    ord.OrderSn,
					PayType:    ord.PayType,
					PayStatus:  ord.PayStatus,
					TradePrice: ord.TradePrice,
					TradeNo:    ord.TradeNo,
					PayTime:    ord.PayTime,
					CreateTime: ord.CreateTime,
					UpdateTime: ord.UpdateTime,
				}, nil
			}
		}
	}
}

func (g *Strategy) SilverLeafDeductionTX(ctx context.Context, orderId int32) (*biz.Order, error) {
	userId, err := common.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	order, err := g.data.db.Orders.Query().Where(orders.IDEQ(int(orderId))).First(ctx)
	if ent.IsNotFound(err) {
		return nil, errs.OrderNotFound(err)
	} else {
		walletObj, err := g.data.db.Wallet.Query().Where(wallet.UserID(int32(userId))).First(ctx)
		if ent.IsNotFound(err) {
			return nil, errs.NotEnoughMoney(err)
		} else {
			if order.TradePrice > walletObj.SilverLeaf {
				return nil, errs.NotEnoughMoney(errors.New("余额不足，请及时充值！"))
			} else {
				silverLeafAmount := walletObj.SilverLeaf - order.TradePrice
				_, err = walletObj.Update().SetGoldLeaf(silverLeafAmount).Save(ctx)
				if err != nil {
					_, err := order.Update().SetPayTime(time.Now()).SetPayStatus(int32(pb.PayStatus_FAILED)).SetPayType(int32(pb.PayType_SILVER_LEAF)).Save(ctx)
					return nil, errs.BadError(err, "订单支付失败！")
				}
				ord, err := order.Update().SetPayTime(time.Now()).SetPayStatus(int32(pb.PayStatus_PAID)).SetPayType(int32(pb.PayType_SILVER_LEAF)).Save(ctx)
				if err != nil {
					return nil, errs.BadError(err, "订单支付失败！")
				}
				return &biz.Order{
					Id:         ord.ID,
					UserId:     ord.UserID,
					OrderSn:    ord.OrderSn,
					PayType:    ord.PayType,
					PayStatus:  ord.PayStatus,
					TradePrice: ord.TradePrice,
					TradeNo:    ord.TradeNo,
					PayTime:    ord.PayTime,
					CreateTime: ord.CreateTime,
					UpdateTime: ord.UpdateTime,
				}, nil
			}
		}
	}

}
