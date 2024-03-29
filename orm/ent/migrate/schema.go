// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoresColumns holds the columns for the "cores" table.
	CoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "count", Type: field.TypeInt, Default: 0},
		{Name: "created", Type: field.TypeString},
		{Name: "updated", Type: field.TypeString},
	}
	// CoresTable holds the schema information for the "cores" table.
	CoresTable = &schema.Table{
		Name:       "cores",
		Columns:    CoresColumns,
		PrimaryKey: []*schema.Column{CoresColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoresTable,
	}
)

func init() {
}
