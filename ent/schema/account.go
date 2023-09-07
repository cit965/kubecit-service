package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{

		field.Uint64("user_id"),
		field.String("openid").MaxLen(32),
		field.String("password").MaxLen(32),
		field.String("method").MaxLen(32),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
func (Account) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("openid", "method").
			Unique(),
	}
}
