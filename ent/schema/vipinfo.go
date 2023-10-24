package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// VipInfo holds the schema definition for the VipInfo entity.
type VipInfo struct {
	ent.Schema
}

// Fields of the VipInfo.
func (VipInfo) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Comment("会员id"),
		field.Int8("vip_type").Comment("会员类型"),
		field.Time("start_at").Default(time.Now).Comment("会员权益生效时间"),
		field.Time("expire_at").Comment("会员权益失效时间"),
		field.Int("user_id").Comment("用户id"),
	}
}

// Edges of the VipInfo.
func (VipInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_info", User.Type).Unique().Ref("vip_info").Field("user_id").Required(),
	}
}
