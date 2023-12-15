package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Solution struct {
	ent.Schema
}

func (Solution) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("name"),
		field.Strings("photos"),
		field.String("description"),
		field.Strings("ingredients"),
	}
}

func (Solution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("disease", Disease.Type).Ref("solutions"),
	}
}
