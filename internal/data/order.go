package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/ent"
	"kubecit-service/ent/course"
	"kubecit-service/ent/orderinfos"
	"kubecit-service/ent/orders"
	"kubecit-service/internal/biz"
	"kubecit-service/internal/pkg/common"
	"math/rand"
	"time"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (or *orderRepo) Create(ctx context.Context, courseIds []int32) (*biz.Order, error) {
	result, err := or.data.WithResultTx(ctx, func(tx *ent.Tx) (interface{}, error) {
		orderTx, err := or.createOrderTx(ctx, courseIds)
		if err != nil {
			return nil, err
		}
		return orderTx, nil
	})
	if err != nil {
		return nil, err
	}
	if order, ok := result.(*biz.Order); ok {
		return order, nil
	}
	return nil, err
}

func (or *orderRepo) createOrderTx(ctx context.Context, courseIds []int32) (*biz.Order, error) {
	var numsInt []int
	for _, num := range courseIds {
		numsInt = append(numsInt, int(num))
	}
	coursePrice, err := or.data.db.Course.Query().Where(course.IDIn(numsInt...)).Aggregate(ent.Sum(course.FieldPrice)).Int(ctx)
	if err != nil {
		return nil, errors.BadRequest("", err.Error())
	}

	userId, err := common.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	orderObj, err := or.data.db.Orders.Create().SetOrderSn(GenerateOrderSn(int32(userId))).SetTradePrice(int32(coursePrice)).SetUserID(int32(userId)).Save(ctx)
	if err != nil {
		return nil, errors.BadRequest("创建订单失败", err.Error())
	}

	infos := make([]*biz.OrderInfo, 0)
	for _, id := range courseIds {
		c, err := or.data.db.Course.Query().Where(course.IDEQ(int(id))).First(ctx)
		if err != nil {
			return nil, errors.BadRequest("根据课程ID查询课程失败", err.Error())
		}
		info, err := or.data.db.OrderInfos.Create().SetOrderID(int32(orderObj.ID)).SetProductID(id).SetProductDescribe(c.Detail).SetProductName(c.Name).SetProductPrice(int32(c.Price)).Save(ctx)
		if err != nil {
			return nil, errors.BadRequest("创建订单详情失败", err.Error())
		}

		infos = append(infos, &biz.OrderInfo{
			Id:              info.ID,
			ProductName:     info.ProductName,
			ProductDescribe: info.ProductDescribe,
			ProductPrice:    c.Price,
			CreateTime:      info.CreateTime,
			UpdateTime:      info.UpdateTime,
			OrderId:         int32(orderObj.ID),
			ProductId:       id,
		})
	}
	order := &biz.Order{
		Id:         orderObj.ID,
		UserId:     int32(userId),
		OrderSn:    orderObj.OrderSn,
		PayType:    orderObj.PayType,
		PayStatus:  orderObj.PayStatus,
		TradePrice: orderObj.TradePrice,
		TradeNo:    orderObj.TradeNo,
		PayTime:    orderObj.PayTime,
		CreateTime: orderObj.CreateTime,
		UpdateTime: orderObj.UpdateTime,
		Info:       infos,
	}
	return order, nil
}

func GenerateOrderSn(userId int32) string {
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(), userId, rand.Intn(90)+10)
	return orderSn
}

func (or *orderRepo) MyOrder(ctx context.Context, pageNum, pageSize *int32) ([]*biz.Order, error) {
	userID, err := common.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	//userID := int32(1)
	qr := or.data.db.Orders.Query().Where(
		orders.UserIDEQ(int32(userID)),
	).Order(ent.Desc(orders.FieldCreateTime))

	limit, offset := common.ConvertPageSize(pageNum, pageSize)
	orderAll, err := qr.Limit(limit).Offset(offset).All(ctx)

	if err != nil {
		return nil, errors.BadRequest(err.Error(), "获取订单列表失败！")
	}
	ordersResult := make([]*biz.Order, 0)
	for _, order := range orderAll {
		orderInfos, err := or.data.db.OrderInfos.Query().Where(
			orderinfos.OrderIDEQ(int32(order.ID)),
		).All(ctx)
		if err != nil {
			return nil, err
		}
		orderDetails := make([]*biz.OrderInfo, 0)

		for _, orderInfo := range orderInfos {
			orderDetails = append(orderDetails, &biz.OrderInfo{
				Id:              orderInfo.ID,
				ProductName:     orderInfo.ProductName,
				ProductDescribe: orderInfo.ProductDescribe,
				ProductPrice:    orderInfo.ProductPrice,
				CreateTime:      orderInfo.CreateTime,
				UpdateTime:      orderInfo.UpdateTime,
				OrderId:         int32(order.ID),
				ProductId:       orderInfo.ProductID,
			})
		}
		ordersResult = append(ordersResult, &biz.Order{
			Id:         order.ID,
			UserId:     order.UserID,
			OrderSn:    order.OrderSn,
			PayType:    order.PayType,
			PayStatus:  order.PayStatus,
			TradePrice: order.TradePrice,
			TradeNo:    order.TradeNo,
			PayTime:    order.PayTime,
			CreateTime: order.CreateTime,
			UpdateTime: order.UpdateTime,
			Info:       orderDetails,
		})
	}
	return ordersResult, nil
}
