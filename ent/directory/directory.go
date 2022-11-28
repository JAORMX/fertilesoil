// Code generated by ent, DO NOT EDIT.

package directory

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the directory type in the database.
	Label = "directory"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldIsRoot holds the string denoting the is_root field in the database.
	FieldIsRoot = "is_root"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the directory in the database.
	Table = "directories"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "directories"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "directory_children"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "directories"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "directory_children"
)

// Columns holds all SQL columns for directory fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMetadata,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsRoot,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "directories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"directory_children",
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/JAORMX/fertilesoil/ent/runtime"
var (
	Hooks [4]ent.Hook
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsRoot holds the default value on creation for the "is_root" field.
	DefaultIsRoot bool
)
