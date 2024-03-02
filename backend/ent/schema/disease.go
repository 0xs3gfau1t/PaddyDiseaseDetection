package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Disease struct {
	ent.Schema
}

func (Disease) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("name"),
	}
}
func (Disease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("solutions", Solution.Type),
		edge.To("causes", Cause.Type),
		edge.From("disease_identified", DiseaseIdentified.Type).Ref("disease"),
		edge.From("image", Image.Type).Ref("disease"),
	}
}
