package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Image struct {
	ent.Schema
}

func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier"),
		field.UUID("id", uuid.New()).Unique(),
		field.Time("created_at").Default(time.Now),
	}
}
