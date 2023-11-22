package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type DiseaseIdentified struct {
	ent.Schema
}

func (DiseaseIdentified) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("location"),
		field.Int("severity").Range(1, 10),
		field.Time("created_at").Default(time.Now),
		field.Strings("photos"),
		field.Enum("status").Values("processing", "processed", "queued", "failed"),
	}
}
func (DiseaseIdentified) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("uploded_by", User.Type).Ref("disease_identified"),
		edge.To("disease", Disease.Type),
	}
}
