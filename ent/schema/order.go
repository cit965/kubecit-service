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
		field.String("order_sn").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Comment("平台自己生成的订单号"),

		field.Int32("pay_type").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0).Comment("支付类型 1(支付宝)， 2(微信)"),

		field.Int32("pay_status").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(0).Comment("支付状态 0(待支付), 1(成功)， 2(失败)，3（关闭）4(订单取消) 5(退款中)6（退款取消）"),

		field.Float("trade_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal", // Override MySQL.
		}).Comment("订单价格"),

		field.String("trade_no").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional().Comment("交易号 就是微信、支付宝的订单号 查账"),

		field.Time("pay_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Comment("支付时间"),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(time.Now()).Comment("创建时间"),

		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(time.Now()).UpdateDefault(time.Now).Comment("更新时间"),
	}

}

// Edges of the Orders.
func (Orders) Edges() []ent.Edge {
	return nil
}
