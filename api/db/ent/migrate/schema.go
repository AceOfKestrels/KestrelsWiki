// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "updated", Type: field.TypeTime},
		{Name: "author", Type: field.TypeString},
		{Name: "commit_hash", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// MirrorsColumns holds the columns for the "mirrors" table.
	MirrorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "origin_path", Type: field.TypeString, Unique: true},
		{Name: "target_path", Type: field.TypeString},
	}
	// MirrorsTable holds the schema information for the "mirrors" table.
	MirrorsTable = &schema.Table{
		Name:       "mirrors",
		Columns:    MirrorsColumns,
		PrimaryKey: []*schema.Column{MirrorsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FilesTable,
		MirrorsTable,
	}
)

func init() {
}