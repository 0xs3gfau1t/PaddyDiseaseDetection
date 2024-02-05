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
		field.Int("severity").Range(1, 10).Default(1),
		field.Time("created_at").Default(time.Now),
		field.Enum("status").Values("processing", "processed", "queued", "failed").Default("queued"),
	}
}
func (DiseaseIdentified) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("uploaded_by", User.Type).Unique().Required(),
		edge.To("disease", Disease.Type).Unique(),
		edge.From("image", Image.Type).Ref("disease_identified").Unique(),
	}
}
