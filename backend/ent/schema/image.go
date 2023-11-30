package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Image struct {
	ent.Schema
}

func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier"),
		field.Time("created_at").Default(time.Now),
	}
}

func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("disease", Disease.Type),
		edge.To("disease_identified", DiseaseIdentified.Type).Unique().Annotations(entsql.OnDelete(entsql.Cascade)).Required(),
	}
}
