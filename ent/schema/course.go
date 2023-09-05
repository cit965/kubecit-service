package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Int32("level"),
		field.Time("updated_at"),
		field.String("name"),
		field.String("detail"),
		field.String("cover"),
		field.Float32("price"),
		field.String("tags"),
		field.String("created_at"),
		field.Int32("status"),
		field.Int("category_id").Optional(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Category.Type).
			Ref("courses").
			Unique().
			Field("category_id"),
	}
}
