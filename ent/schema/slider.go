package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Slider holds the schema definition for the Slider entity.
type Slider struct {
	ent.Schema
}

// Fields of the Slider.
func (Slider) Fields() []ent.Field {
	return []ent.Field{
		field.String("createBy"),
		field.String("imageName"),
		field.Time("createTime"),
		field.String("updateBy"),
		field.String("imageRemark"),
		field.String("imageUrl"),
		field.String("pcHref"),
		field.Time("updateTime"),
		field.String("id"),
		field.String("appHref"),
	}
}

// Edges of the Slider.
func (Slider) Edges() []ent.Edge {
	return nil
}
