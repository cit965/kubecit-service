package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// VipProduct holds the schema definition for the VipProduct entity.
type VipProduct struct {
	ent.Schema
}

// Fields of the VipProduct.
func (VipProduct) Fields() []ent.Field {
	return []ent.Field{
		//field.Int8("").Comment(""),
		field.Float("price").Comment("会员价格"),
		field.String("name").Comment("会员名称"),
		field.Text("description").Comment("会员描述"),
	}
}

// Edges of the VipProduct.
func (VipProduct) Edges() []ent.Edge {
	return nil
}
