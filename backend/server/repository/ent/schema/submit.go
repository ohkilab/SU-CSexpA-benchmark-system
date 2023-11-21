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
		field.Int("id"),
		field.String("url"),
		field.Int("score").Optional(),
		field.Enum("language").Values(languages...).Optional(),
		field.String("message").Optional(),
		field.String("status"), // enum は grpc で定義済み
		field.Int("task_num"),
		field.Time("submited_at"),
		field.Time("completed_at").Optional(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Submit.
func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("taskResults", TaskResult.Type),
		edge.From("groups", Group.Type).Ref("submits").Unique(),
		edge.From("contests", Contest.Type).Ref("submits").Unique(),
	}
}
