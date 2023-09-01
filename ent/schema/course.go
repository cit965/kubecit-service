package schema

import "entgo.io/ent"

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return nil
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return nil
}
