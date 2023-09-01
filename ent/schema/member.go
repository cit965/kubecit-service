package schema

import "entgo.io/ent"

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return nil
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
