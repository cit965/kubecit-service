package schema

import "entgo.io/ent"

// Slider holds the schema definition for the Slider entity.
type Slider struct {
	ent.Schema
}

// Fields of the Slider.
func (Slider) Fields() []ent.Field {
	return nil
}

// Edges of the Slider.
func (Slider) Edges() []ent.Edge {
	return nil
}
