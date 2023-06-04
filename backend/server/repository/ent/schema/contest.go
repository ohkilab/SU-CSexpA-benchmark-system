package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contest holds the schema definition for the Contest entity.
type Contest struct {
	ent.Schema
}

// Fields of the Contest.
func (Contest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.Time("start_at"),
		field.Time("end_at"),
		field.Int("submit_limit"),
		field.Int("year").Positive(),
		// if auto, {contest_id}/random.txt is used
		// if manual, {contest_id}/\d.txt is used (number means attempt count)
		field.Enum("tag_selection_logic").Values("auto", "manual"),
		field.Time("created_at"),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submits", Submit.Type),
	}
}
