package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("level"),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now).Comment("修改时间"),
		field.String("name"),
		field.String("detail"),
		field.String("cover"),
		field.Float32("price"),
		field.String("tags"),
		field.Time("created_at").Default(time.Now()).Comment("创建时间"),
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
		edge.To("chapters", Chapter.Type),
	}
}
