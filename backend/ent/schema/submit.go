package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Submit holds the schema definition for the Submit entity.
type Submit struct {
	ent.Schema
}

var languages = []string{"php", "go", "rust", "javascript", "csharp", "cpp", "ruby", "python"}

// Fields of the Submit.
func (Submit) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("ip_addr"),
		field.Int("year").Positive(),
		field.Int("score").Positive(),
		field.Enum("language").Values(languages...),
		field.Time("submited_at"),
		field.Time("completed_at").Optional(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Submit.
func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tagResults", TagResult.Type),
		edge.From("group", Group.Type).Ref("submits"),
	}
}
