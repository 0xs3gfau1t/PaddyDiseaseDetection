package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("name"),
		field.String("email").Unique(),
		field.String("location"),
		field.String("coord").Optional(),
		field.String("password"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("diseases_identified", DiseaseIdentified.Type).Ref("uploaded_by"),
	}
}
