package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("version").StorageKey("vearion"),
		field.String("title").MinLen(3).MaxLen(100),
		field.String("description").MinLen(3).MaxLen(1000),
		field.Bool("completed").Default(false),
		field.Time("created_at").Default(time.Now).
			Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.UUID("author_id", uuid.UUID{}),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("todos").Field("author_id").Unique().Required(),
	}
}
