// Code generated by ent, DO NOT EDIT.

package project

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldDocURL holds the string denoting the doc_url field in the database.
	FieldDocURL = "doc_url"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_projects"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// TeamsTable is the table that holds the teams relation/edge. The primary key declared below.
	TeamsTable = "team_projects"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "teams"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldURL,
	FieldImageURL,
	FieldDocURL,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "project_id"}
	// TeamsPrimaryKey and TeamsColumn2 are the table columns denoting the
	// primary key for the teams relation (M2M).
	TeamsPrimaryKey = []string{"team_id", "project_id"}
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

// OrderOption defines the ordering options for the Project queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByDocURL orders the results by the doc_url field.
func ByDocURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDocURL, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTeamsCount orders the results by teams count.
func ByTeamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamsStep(), opts...)
	}
}

// ByTeams orders the results by teams terms.
func ByTeams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
func newTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
	)
}
