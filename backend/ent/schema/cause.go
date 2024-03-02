package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Cause struct {
	ent.Schema
}

func (Cause) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("image"),
	}
}
func (Cause) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("disease", Disease.Type),
	}
}
