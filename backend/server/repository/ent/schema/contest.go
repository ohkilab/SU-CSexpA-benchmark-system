package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("slug").Unique(),
		// if auto, {contest_slug}/random.txt is used
		// if manual, {contest_slug}/\d.txt is used (number means attempt count)
		field.Enum("tag_selection_logic").Values("auto", "manual"),
		field.String("validator"),
		field.Int64("time_limit_per_task").Default(int64(30 * time.Second)).Optional(),
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

func (Contest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Unique(),
	}
}
