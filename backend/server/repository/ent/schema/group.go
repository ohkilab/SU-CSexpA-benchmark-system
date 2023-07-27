package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.Enum("role").Values("contestant", "guest", "admin"),
		field.String("encrypted_password"),
		field.Time("created_at"),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submits", Submit.Type),
	}
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{}
}
