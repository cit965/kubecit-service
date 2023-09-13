package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("detail"),
		field.String("cover"),
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}
