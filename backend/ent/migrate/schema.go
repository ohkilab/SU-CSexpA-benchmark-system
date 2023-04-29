// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContestsColumns holds the columns for the "contests" table.
	ContestsColumns = []*schema.Column{
		{Name: "year", Type: field.TypeInt, Increment: true},
		{Name: "qualifier_start_at", Type: field.TypeTime},
		{Name: "qualifier_end_at", Type: field.TypeTime},
		{Name: "qualifier_submit_limit", Type: field.TypeInt},
		{Name: "final_start_at", Type: field.TypeTime},
		{Name: "final_end_at", Type: field.TypeTime},
		{Name: "final_submit_limit", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// ContestsTable holds the schema information for the "contests" table.
	ContestsTable = &schema.Table{
		Name:       "contests",
		Columns:    ContestsColumns,
		PrimaryKey: []*schema.Column{ContestsColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "year", Type: field.TypeInt},
		{Name: "score", Type: field.TypeInt},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"contestant", "guest"}},
		{Name: "encrypted_password", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "group_score",
				Unique:  false,
				Columns: []*schema.Column{GroupsColumns[2]},
			},
		},
	}
	// SubmitsColumns holds the columns for the "submits" table.
	SubmitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "year", Type: field.TypeInt},
		{Name: "score", Type: field.TypeInt},
		{Name: "language", Type: field.TypeEnum, Enums: []string{"php", "go", "rust", "javascript", "csharp", "cpp", "ruby", "python"}},
		{Name: "submited_at", Type: field.TypeTime},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// SubmitsTable holds the schema information for the "submits" table.
	SubmitsTable = &schema.Table{
		Name:       "submits",
		Columns:    SubmitsColumns,
		PrimaryKey: []*schema.Column{SubmitsColumns[0]},
	}
	// TagResultsColumns holds the columns for the "tag_results" table.
	TagResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "score", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "submit_tag_results", Type: field.TypeString, Nullable: true},
	}
	// TagResultsTable holds the schema information for the "tag_results" table.
	TagResultsTable = &schema.Table{
		Name:       "tag_results",
		Columns:    TagResultsColumns,
		PrimaryKey: []*schema.Column{TagResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_results_submits_tagResults",
				Columns:    []*schema.Column{TagResultsColumns[5]},
				RefColumns: []*schema.Column{SubmitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupSubmitsColumns holds the columns for the "group_submits" table.
	GroupSubmitsColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeString},
		{Name: "submit_id", Type: field.TypeString},
	}
	// GroupSubmitsTable holds the schema information for the "group_submits" table.
	GroupSubmitsTable = &schema.Table{
		Name:       "group_submits",
		Columns:    GroupSubmitsColumns,
		PrimaryKey: []*schema.Column{GroupSubmitsColumns[0], GroupSubmitsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_submits_group_id",
				Columns:    []*schema.Column{GroupSubmitsColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_submits_submit_id",
				Columns:    []*schema.Column{GroupSubmitsColumns[1]},
				RefColumns: []*schema.Column{SubmitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContestsTable,
		GroupsTable,
		SubmitsTable,
		TagResultsTable,
		GroupSubmitsTable,
	}
)

func init() {
	TagResultsTable.ForeignKeys[0].RefTable = SubmitsTable
	GroupSubmitsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupSubmitsTable.ForeignKeys[1].RefTable = SubmitsTable
}
