package schema

import "entgo.io/ent"

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return nil
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return nil
}
