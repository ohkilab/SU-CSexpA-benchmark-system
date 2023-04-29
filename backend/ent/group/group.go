// Code generated by ent, DO NOT EDIT.

package group

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldEncryptedPassword holds the string denoting the encrypted_password field in the database.
	FieldEncryptedPassword = "encrypted_password"
	// EdgeSubmits holds the string denoting the submits edge name in mutations.
	EdgeSubmits = "submits"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// SubmitsTable is the table that holds the submits relation/edge. The primary key declared below.
	SubmitsTable = "group_submits"
	// SubmitsInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmitsInverseTable = "submits"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldYear,
	FieldScore,
	FieldRole,
	FieldEncryptedPassword,
}

var (
	// SubmitsPrimaryKey and SubmitsColumn2 are the table columns denoting the
	// primary key for the submits relation (M2M).
	SubmitsPrimaryKey = []string{"group_id", "submit_id"}
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

var (
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int) error
)

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleContestant Role = "contestant"
	RoleGuest      Role = "guest"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleContestant, RoleGuest:
		return nil
	default:
		return fmt.Errorf("group: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the Group queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByEncryptedPassword orders the results by the encrypted_password field.
func ByEncryptedPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEncryptedPassword, opts...).ToFunc()
}

// BySubmitsCount orders the results by submits count.
func BySubmitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmitsStep(), opts...)
	}
}

// BySubmits orders the results by submits terms.
func BySubmits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmitsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubmitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubmitsTable, SubmitsPrimaryKey...),
	)
}
