package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
  ent.Schema
}

func (User) Fields() []ent.Field{
  return []ent.Field{
    field.UUID("id", uuid.New()).Unique(),
    field.String("name"),
    field.String("email"),
    field.String("location"),
    field.String("coord"),
  }
}
