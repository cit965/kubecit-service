// Code generated by ent, DO NOT EDIT.

package setting

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the setting type in the database.
	Label = "setting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldCover holds the string denoting the cover field in the database.
	FieldCover = "cover"
	// Table holds the table name of the setting in the database.
	Table = "settings"
)

// Columns holds all SQL columns for setting fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDetail,
	FieldCover,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Setting queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDetail orders the results by the detail field.
func ByDetail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDetail, opts...).ToFunc()
}

// ByCover orders the results by the cover field.
func ByCover(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCover, opts...).ToFunc()
}
