package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Orders holds the schema definition for the Orders entity.
type Orders struct {
	ent.Schema
}

// Fields of the Orders.
func (Orders) Fields() []ent.Field {

	return []ent.Field{
		field.Int32("user_id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Comment("用户id"),

		field.String("order_sn").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Comment("平台自己生成的订单号"),

		field.Int32("pay_type").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(1).Comment("支付类型 NO_PAY = 1 ; // 还未支付 ALIPAY = 2; // 支付宝 WECHAT = 3; // 微信支付 GOLDEN_LEAF = 4;  // 金叶子 SILVER_LEAF = 5;  //银叶子"),

		field.Int32("pay_status").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(1).Comment("支付状态 UNPAID = 1; // 未支付 PAID = 2; // 已支付 FAILED = 3; //支付失败 CLOSED = 4; //关闭 CANCELED = 5; //取消 REFUNDING = 6; //退款中 REFUNDED = 7; //退款成功 REFUND_FAILED = 8;//退款失败 "),

		field.Int32("trade_price").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Comment("订单价格(单位分)"),

		field.String("trade_no").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional().Comment("交易号 就是微信、支付宝的订单号 查账"),

		field.Time("pay_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Comment("支付时间"),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(time.Now).Comment("创建时间"),

		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}

}

// Edges of the Orders.
func (Orders) Edges() []ent.Edge {
	return nil
}
