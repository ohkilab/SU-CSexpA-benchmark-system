// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContestsColumns holds the columns for the "contests" table.
	ContestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "submit_limit", Type: field.TypeInt},
		{Name: "year", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
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
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
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
				Columns: []*schema.Column{GroupsColumns[3]},
			},
		},
	}
	// SubmitsColumns holds the columns for the "submits" table.
	SubmitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ip_addr", Type: field.TypeString},
		{Name: "year", Type: field.TypeInt},
		{Name: "score", Type: field.TypeInt, Nullable: true},
		{Name: "language", Type: field.TypeEnum, Nullable: true, Enums: []string{"php", "go", "rust", "javascript", "csharp", "cpp", "ruby", "python"}},
		{Name: "submited_at", Type: field.TypeTime},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "contest_submits", Type: field.TypeInt, Nullable: true},
		{Name: "group_submits", Type: field.TypeInt, Nullable: true},
	}
	// SubmitsTable holds the schema information for the "submits" table.
	SubmitsTable = &schema.Table{
		Name:       "submits",
		Columns:    SubmitsColumns,
		PrimaryKey: []*schema.Column{SubmitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "submits_contests_submits",
				Columns:    []*schema.Column{SubmitsColumns[8]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "submits_groups_submits",
				Columns:    []*schema.Column{SubmitsColumns[9]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TaskResultsColumns holds the columns for the "task_results" table.
	TaskResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "request_per_sec", Type: field.TypeInt},
		{Name: "url", Type: field.TypeString},
		{Name: "method", Type: field.TypeString},
		{Name: "request_content_type", Type: field.TypeString},
		{Name: "request_body", Type: field.TypeString, Nullable: true},
		{Name: "thread_num", Type: field.TypeInt},
		{Name: "attempt_count", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "submit_task_results", Type: field.TypeInt, Nullable: true},
	}
	// TaskResultsTable holds the schema information for the "task_results" table.
	TaskResultsTable = &schema.Table{
		Name:       "task_results",
		Columns:    TaskResultsColumns,
		PrimaryKey: []*schema.Column{TaskResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_results_submits_taskResults",
				Columns:    []*schema.Column{TaskResultsColumns[10]},
				RefColumns: []*schema.Column{SubmitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContestsTable,
		GroupsTable,
		SubmitsTable,
		TaskResultsTable,
	}
)

func init() {
	SubmitsTable.ForeignKeys[0].RefTable = ContestsTable
	SubmitsTable.ForeignKeys[1].RefTable = GroupsTable
	TaskResultsTable.ForeignKeys[0].RefTable = SubmitsTable
}
