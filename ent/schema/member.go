package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("orderNumber"),
		field.String("vipName"),
		field.String("vipId"),
		field.String("vipDesc"),
		field.Time("startTime"),
		field.Time("endTime"),
		field.String("id"),
		field.Bool("isExpired"),
		field.String("memberId"),
		field.String("vipIcon"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("vipMember").Unique(),
	}
}
