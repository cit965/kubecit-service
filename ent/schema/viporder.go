package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// VipOrder holds the schema definition for the VipOrder entity.
type VipOrder struct {
	ent.Schema
}

// Fields of the VipOrder.
func (VipOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("biz_id").Unique().Comment("业务订单号"),
		field.Int8("vip_type").Comment("会员类型"),
		field.Int8("pay_type").Comment("支付类型"),
		field.Int8("pay_status").Comment("支付状态").Optional(),
		field.Time("create_at").Default(time.Now).Comment("创建时间"),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.Int("user_id").Optional().Unique().Comment("用户id"),
		field.Float("price").Comment("订单价格"),
	}
}

// Edges of the VipOrder.
func (VipOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_order", User.Type).Ref("vip_order").Field("user_id").Unique(),
	}
}
