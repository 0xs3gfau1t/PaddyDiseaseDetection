// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldCoord holds the string denoting the coord field in the database.
	FieldCoord = "coord"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeDiseaseIdentified holds the string denoting the disease_identified edge name in mutations.
	EdgeDiseaseIdentified = "disease_identified"
	// Table holds the table name of the user in the database.
	Table = "users"
	// DiseaseIdentifiedTable is the table that holds the disease_identified relation/edge. The primary key declared below.
	DiseaseIdentifiedTable = "user_disease_identified"
	// DiseaseIdentifiedInverseTable is the table name for the DiseaseIdentified entity.
	// It exists in this package in order to avoid circular dependency with the "diseaseidentified" package.
	DiseaseIdentifiedInverseTable = "disease_identifieds"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldLocation,
	FieldCoord,
	FieldPassword,
}

var (
	// DiseaseIdentifiedPrimaryKey and DiseaseIdentifiedColumn2 are the table columns denoting the
	// primary key for the disease_identified relation (M2M).
	DiseaseIdentifiedPrimaryKey = []string{"user_id", "disease_identified_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByCoord orders the results by the coord field.
func ByCoord(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoord, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByDiseaseIdentifiedCount orders the results by disease_identified count.
func ByDiseaseIdentifiedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDiseaseIdentifiedStep(), opts...)
	}
}

// ByDiseaseIdentified orders the results by disease_identified terms.
func ByDiseaseIdentified(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDiseaseIdentifiedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDiseaseIdentifiedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DiseaseIdentifiedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DiseaseIdentifiedTable, DiseaseIdentifiedPrimaryKey...),
	)
}
