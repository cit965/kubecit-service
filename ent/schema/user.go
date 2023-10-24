package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("channel"),
		field.Uint8("role_id"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teacher", Teacher.Type).Unique(),
		edge.To("apply_record", ApplyRecord.Type),
		edge.To("vip_info", VipInfo.Type).Unique(),
		edge.To("vip_order", VipOrder.Type),
	}
}
