package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TagResult holds the schema definition for the TagResult entity.
type TaskResult struct {
	ent.Schema
}

// Fields of the TagResult.
func (TaskResult) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("request_per_sec"),
		field.String("status"), // enum は grpc で定義済み
		field.String("error_message").Optional(),
		// request
		field.String("url"),
		field.String("method"),
		field.String("request_content_type"),
		field.String("request_body").Optional(),
		// config
		field.Int("thread_num"),
		field.Int("attempt_count"),
		field.Time("created_at"),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the TagResult.
func (TaskResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submits", Submit.Type).Ref("taskResults").Unique(),
	}
}
