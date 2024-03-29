// Code generated by ent, DO NOT EDIT.

package failure

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the failure type in the database.
	Label = "failure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldJiraKey holds the string denoting the jira_key field in the database.
	FieldJiraKey = "jira_key"
	// FieldJiraStatus holds the string denoting the jira_status field in the database.
	FieldJiraStatus = "jira_status"
	// FieldErrorMessage holds the string denoting the error_message field in the database.
	FieldErrorMessage = "error_message"
	// FieldTitleFromJira holds the string denoting the title_from_jira field in the database.
	FieldTitleFromJira = "title_from_jira"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// FieldClosedDate holds the string denoting the closed_date field in the database.
	FieldClosedDate = "closed_date"
	// FieldLabels holds the string denoting the labels field in the database.
	FieldLabels = "labels"
	// EdgeFailures holds the string denoting the failures edge name in mutations.
	EdgeFailures = "failures"
	// TeamsFieldID holds the string denoting the ID field of the Teams.
	TeamsFieldID = "team_id"
	// Table holds the table name of the failure in the database.
	Table = "failures"
	// FailuresTable is the table that holds the failures relation/edge.
	FailuresTable = "failures"
	// FailuresInverseTable is the table name for the Teams entity.
	// It exists in this package in order to avoid circular dependency with the "teams" package.
	FailuresInverseTable = "teams"
	// FailuresColumn is the table column denoting the failures relation/edge.
	FailuresColumn = "teams_failures"
)

// Columns holds all SQL columns for failure fields.
var Columns = []string{
	FieldID,
	FieldJiraKey,
	FieldJiraStatus,
	FieldErrorMessage,
	FieldTitleFromJira,
	FieldCreatedDate,
	FieldClosedDate,
	FieldLabels,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "failures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"teams_failures",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// JiraKeyValidator is a validator for the "jira_key" field. It is called by the builders before save.
	JiraKeyValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
