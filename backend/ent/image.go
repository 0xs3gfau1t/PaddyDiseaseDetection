// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"segFault/PaddyDiseaseDetection/ent/image"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier string `json:"identifier,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldIdentifier:
			values[i] = new(sql.NullString)
		case image.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case image.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case image.FieldIdentifier:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[j])
			} else if value.Valid {
				i.Identifier = value.String
			}
		case image.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("identifier=")
	builder.WriteString(i.Identifier)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image