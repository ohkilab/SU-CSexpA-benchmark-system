package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TagResult holds the schema definition for the TagResult entity.
type TagResult struct {
	ent.Schema
}

// Fields of the TagResult.
func (TagResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("score"),
		field.Time("created_at"),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the TagResult.
func (TagResult) Edges() []ent.Edge {
	return nil
}
