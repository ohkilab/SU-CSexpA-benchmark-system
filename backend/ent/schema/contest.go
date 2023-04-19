package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Contest holds the schema definition for the Contest entity.
type Contest struct {
	ent.Schema
}

// Fields of the Contest.
func (Contest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("year").Positive(),
		field.Time("qualifier_start_at"),
		field.Time("qualifier_end_at"),
		field.Int("qualifier_submit_limit"),
		field.Time("final_start_at"),
		field.Time("final_end_at"),
		field.Int("final_submit_limit"),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return nil
}
