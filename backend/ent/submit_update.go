// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/tagresult"
)

// SubmitUpdate is the builder for updating Submit entities.
type SubmitUpdate struct {
	config
	hooks    []Hook
	mutation *SubmitMutation
}

// Where appends a list predicates to the SubmitUpdate builder.
func (su *SubmitUpdate) Where(ps ...predicate.Submit) *SubmitUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetIPAddr sets the "ip_addr" field.
func (su *SubmitUpdate) SetIPAddr(s string) *SubmitUpdate {
	su.mutation.SetIPAddr(s)
	return su
}

// SetYear sets the "year" field.
func (su *SubmitUpdate) SetYear(i int) *SubmitUpdate {
	su.mutation.ResetYear()
	su.mutation.SetYear(i)
	return su
}

// AddYear adds i to the "year" field.
func (su *SubmitUpdate) AddYear(i int) *SubmitUpdate {
	su.mutation.AddYear(i)
	return su
}

// SetScore sets the "score" field.
func (su *SubmitUpdate) SetScore(i int) *SubmitUpdate {
	su.mutation.ResetScore()
	su.mutation.SetScore(i)
	return su
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableScore(i *int) *SubmitUpdate {
	if i != nil {
		su.SetScore(*i)
	}
	return su
}

// AddScore adds i to the "score" field.
func (su *SubmitUpdate) AddScore(i int) *SubmitUpdate {
	su.mutation.AddScore(i)
	return su
}

// ClearScore clears the value of the "score" field.
func (su *SubmitUpdate) ClearScore() *SubmitUpdate {
	su.mutation.ClearScore()
	return su
}

// SetLanguage sets the "language" field.
func (su *SubmitUpdate) SetLanguage(s submit.Language) *SubmitUpdate {
	su.mutation.SetLanguage(s)
	return su
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableLanguage(s *submit.Language) *SubmitUpdate {
	if s != nil {
		su.SetLanguage(*s)
	}
	return su
}

// ClearLanguage clears the value of the "language" field.
func (su *SubmitUpdate) ClearLanguage() *SubmitUpdate {
	su.mutation.ClearLanguage()
	return su
}

// SetSubmitedAt sets the "submited_at" field.
func (su *SubmitUpdate) SetSubmitedAt(t time.Time) *SubmitUpdate {
	su.mutation.SetSubmitedAt(t)
	return su
}

// SetCompletedAt sets the "completed_at" field.
func (su *SubmitUpdate) SetCompletedAt(t time.Time) *SubmitUpdate {
	su.mutation.SetCompletedAt(t)
	return su
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableCompletedAt(t *time.Time) *SubmitUpdate {
	if t != nil {
		su.SetCompletedAt(*t)
	}
	return su
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (su *SubmitUpdate) ClearCompletedAt() *SubmitUpdate {
	su.mutation.ClearCompletedAt()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SubmitUpdate) SetUpdatedAt(t time.Time) *SubmitUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableUpdatedAt(t *time.Time) *SubmitUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (su *SubmitUpdate) ClearUpdatedAt() *SubmitUpdate {
	su.mutation.ClearUpdatedAt()
	return su
}

// AddTagResultIDs adds the "tagResults" edge to the TagResult entity by IDs.
func (su *SubmitUpdate) AddTagResultIDs(ids ...int) *SubmitUpdate {
	su.mutation.AddTagResultIDs(ids...)
	return su
}

// AddTagResults adds the "tagResults" edges to the TagResult entity.
func (su *SubmitUpdate) AddTagResults(t ...*TagResult) *SubmitUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTagResultIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (su *SubmitUpdate) AddGroupIDs(ids ...string) *SubmitUpdate {
	su.mutation.AddGroupIDs(ids...)
	return su
}

// AddGroups adds the "groups" edges to the Group entity.
func (su *SubmitUpdate) AddGroups(g ...*Group) *SubmitUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return su.AddGroupIDs(ids...)
}

// AddContestIDs adds the "contests" edge to the Contest entity by IDs.
func (su *SubmitUpdate) AddContestIDs(ids ...int) *SubmitUpdate {
	su.mutation.AddContestIDs(ids...)
	return su
}

// AddContests adds the "contests" edges to the Contest entity.
func (su *SubmitUpdate) AddContests(c ...*Contest) *SubmitUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddContestIDs(ids...)
}

// Mutation returns the SubmitMutation object of the builder.
func (su *SubmitUpdate) Mutation() *SubmitMutation {
	return su.mutation
}

// ClearTagResults clears all "tagResults" edges to the TagResult entity.
func (su *SubmitUpdate) ClearTagResults() *SubmitUpdate {
	su.mutation.ClearTagResults()
	return su
}

// RemoveTagResultIDs removes the "tagResults" edge to TagResult entities by IDs.
func (su *SubmitUpdate) RemoveTagResultIDs(ids ...int) *SubmitUpdate {
	su.mutation.RemoveTagResultIDs(ids...)
	return su
}

// RemoveTagResults removes "tagResults" edges to TagResult entities.
func (su *SubmitUpdate) RemoveTagResults(t ...*TagResult) *SubmitUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTagResultIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (su *SubmitUpdate) ClearGroups() *SubmitUpdate {
	su.mutation.ClearGroups()
	return su
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (su *SubmitUpdate) RemoveGroupIDs(ids ...string) *SubmitUpdate {
	su.mutation.RemoveGroupIDs(ids...)
	return su
}

// RemoveGroups removes "groups" edges to Group entities.
func (su *SubmitUpdate) RemoveGroups(g ...*Group) *SubmitUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return su.RemoveGroupIDs(ids...)
}

// ClearContests clears all "contests" edges to the Contest entity.
func (su *SubmitUpdate) ClearContests() *SubmitUpdate {
	su.mutation.ClearContests()
	return su
}

// RemoveContestIDs removes the "contests" edge to Contest entities by IDs.
func (su *SubmitUpdate) RemoveContestIDs(ids ...int) *SubmitUpdate {
	su.mutation.RemoveContestIDs(ids...)
	return su
}

// RemoveContests removes "contests" edges to Contest entities.
func (su *SubmitUpdate) RemoveContests(c ...*Contest) *SubmitUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveContestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubmitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SubmitMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubmitUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubmitUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubmitUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubmitUpdate) check() error {
	if v, ok := su.mutation.Year(); ok {
		if err := submit.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Submit.year": %w`, err)}
		}
	}
	if v, ok := su.mutation.Score(); ok {
		if err := submit.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Submit.score": %w`, err)}
		}
	}
	if v, ok := su.mutation.Language(); ok {
		if err := submit.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Submit.language": %w`, err)}
		}
	}
	return nil
}

func (su *SubmitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.IPAddr(); ok {
		_spec.SetField(submit.FieldIPAddr, field.TypeString, value)
	}
	if value, ok := su.mutation.Year(); ok {
		_spec.SetField(submit.FieldYear, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedYear(); ok {
		_spec.AddField(submit.FieldYear, field.TypeInt, value)
	}
	if value, ok := su.mutation.Score(); ok {
		_spec.SetField(submit.FieldScore, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedScore(); ok {
		_spec.AddField(submit.FieldScore, field.TypeInt, value)
	}
	if su.mutation.ScoreCleared() {
		_spec.ClearField(submit.FieldScore, field.TypeInt)
	}
	if value, ok := su.mutation.Language(); ok {
		_spec.SetField(submit.FieldLanguage, field.TypeEnum, value)
	}
	if su.mutation.LanguageCleared() {
		_spec.ClearField(submit.FieldLanguage, field.TypeEnum)
	}
	if value, ok := su.mutation.SubmitedAt(); ok {
		_spec.SetField(submit.FieldSubmitedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.CompletedAt(); ok {
		_spec.SetField(submit.FieldCompletedAt, field.TypeTime, value)
	}
	if su.mutation.CompletedAtCleared() {
		_spec.ClearField(submit.FieldCompletedAt, field.TypeTime)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(submit.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UpdatedAtCleared() {
		_spec.ClearField(submit.FieldUpdatedAt, field.TypeTime)
	}
	if su.mutation.TagResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TagResultsTable,
			Columns: []string{submit.TagResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTagResultsIDs(); len(nodes) > 0 && !su.mutation.TagResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TagResultsTable,
			Columns: []string{submit.TagResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TagResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TagResultsTable,
			Columns: []string{submit.TagResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: submit.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !su.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: submit.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: submit.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ContestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: submit.ContestsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedContestsIDs(); len(nodes) > 0 && !su.mutation.ContestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: submit.ContestsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ContestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: submit.ContestsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubmitUpdateOne is the builder for updating a single Submit entity.
type SubmitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubmitMutation
}

// SetIPAddr sets the "ip_addr" field.
func (suo *SubmitUpdateOne) SetIPAddr(s string) *SubmitUpdateOne {
	suo.mutation.SetIPAddr(s)
	return suo
}

// SetYear sets the "year" field.
func (suo *SubmitUpdateOne) SetYear(i int) *SubmitUpdateOne {
	suo.mutation.ResetYear()
	suo.mutation.SetYear(i)
	return suo
}

// AddYear adds i to the "year" field.
func (suo *SubmitUpdateOne) AddYear(i int) *SubmitUpdateOne {
	suo.mutation.AddYear(i)
	return suo
}

// SetScore sets the "score" field.
func (suo *SubmitUpdateOne) SetScore(i int) *SubmitUpdateOne {
	suo.mutation.ResetScore()
	suo.mutation.SetScore(i)
	return suo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableScore(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetScore(*i)
	}
	return suo
}

// AddScore adds i to the "score" field.
func (suo *SubmitUpdateOne) AddScore(i int) *SubmitUpdateOne {
	suo.mutation.AddScore(i)
	return suo
}

// ClearScore clears the value of the "score" field.
func (suo *SubmitUpdateOne) ClearScore() *SubmitUpdateOne {
	suo.mutation.ClearScore()
	return suo
}

// SetLanguage sets the "language" field.
func (suo *SubmitUpdateOne) SetLanguage(s submit.Language) *SubmitUpdateOne {
	suo.mutation.SetLanguage(s)
	return suo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableLanguage(s *submit.Language) *SubmitUpdateOne {
	if s != nil {
		suo.SetLanguage(*s)
	}
	return suo
}

// ClearLanguage clears the value of the "language" field.
func (suo *SubmitUpdateOne) ClearLanguage() *SubmitUpdateOne {
	suo.mutation.ClearLanguage()
	return suo
}

// SetSubmitedAt sets the "submited_at" field.
func (suo *SubmitUpdateOne) SetSubmitedAt(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetSubmitedAt(t)
	return suo
}

// SetCompletedAt sets the "completed_at" field.
func (suo *SubmitUpdateOne) SetCompletedAt(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetCompletedAt(t)
	return suo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableCompletedAt(t *time.Time) *SubmitUpdateOne {
	if t != nil {
		suo.SetCompletedAt(*t)
	}
	return suo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (suo *SubmitUpdateOne) ClearCompletedAt() *SubmitUpdateOne {
	suo.mutation.ClearCompletedAt()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SubmitUpdateOne) SetUpdatedAt(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableUpdatedAt(t *time.Time) *SubmitUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suo *SubmitUpdateOne) ClearUpdatedAt() *SubmitUpdateOne {
	suo.mutation.ClearUpdatedAt()
	return suo
}

// AddTagResultIDs adds the "tagResults" edge to the TagResult entity by IDs.
func (suo *SubmitUpdateOne) AddTagResultIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.AddTagResultIDs(ids...)
	return suo
}

// AddTagResults adds the "tagResults" edges to the TagResult entity.
func (suo *SubmitUpdateOne) AddTagResults(t ...*TagResult) *SubmitUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTagResultIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (suo *SubmitUpdateOne) AddGroupIDs(ids ...string) *SubmitUpdateOne {
	suo.mutation.AddGroupIDs(ids...)
	return suo
}

// AddGroups adds the "groups" edges to the Group entity.
func (suo *SubmitUpdateOne) AddGroups(g ...*Group) *SubmitUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return suo.AddGroupIDs(ids...)
}

// AddContestIDs adds the "contests" edge to the Contest entity by IDs.
func (suo *SubmitUpdateOne) AddContestIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.AddContestIDs(ids...)
	return suo
}

// AddContests adds the "contests" edges to the Contest entity.
func (suo *SubmitUpdateOne) AddContests(c ...*Contest) *SubmitUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddContestIDs(ids...)
}

// Mutation returns the SubmitMutation object of the builder.
func (suo *SubmitUpdateOne) Mutation() *SubmitMutation {
	return suo.mutation
}

// ClearTagResults clears all "tagResults" edges to the TagResult entity.
func (suo *SubmitUpdateOne) ClearTagResults() *SubmitUpdateOne {
	suo.mutation.ClearTagResults()
	return suo
}

// RemoveTagResultIDs removes the "tagResults" edge to TagResult entities by IDs.
func (suo *SubmitUpdateOne) RemoveTagResultIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.RemoveTagResultIDs(ids...)
	return suo
}

// RemoveTagResults removes "tagResults" edges to TagResult entities.
func (suo *SubmitUpdateOne) RemoveTagResults(t ...*TagResult) *SubmitUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTagResultIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (suo *SubmitUpdateOne) ClearGroups() *SubmitUpdateOne {
	suo.mutation.ClearGroups()
	return suo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (suo *SubmitUpdateOne) RemoveGroupIDs(ids ...string) *SubmitUpdateOne {
	suo.mutation.RemoveGroupIDs(ids...)
	return suo
}

// RemoveGroups removes "groups" edges to Group entities.
func (suo *SubmitUpdateOne) RemoveGroups(g ...*Group) *SubmitUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return suo.RemoveGroupIDs(ids...)
}

// ClearContests clears all "contests" edges to the Contest entity.
func (suo *SubmitUpdateOne) ClearContests() *SubmitUpdateOne {
	suo.mutation.ClearContests()
	return suo
}

// RemoveContestIDs removes the "contests" edge to Contest entities by IDs.
func (suo *SubmitUpdateOne) RemoveContestIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.RemoveContestIDs(ids...)
	return suo
}

// RemoveContests removes "contests" edges to Contest entities.
func (suo *SubmitUpdateOne) RemoveContests(c ...*Contest) *SubmitUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveContestIDs(ids...)
}

// Where appends a list predicates to the SubmitUpdate builder.
func (suo *SubmitUpdateOne) Where(ps ...predicate.Submit) *SubmitUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubmitUpdateOne) Select(field string, fields ...string) *SubmitUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Submit entity.
func (suo *SubmitUpdateOne) Save(ctx context.Context) (*Submit, error) {
	return withHooks[*Submit, SubmitMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubmitUpdateOne) SaveX(ctx context.Context) *Submit {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubmitUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubmitUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubmitUpdateOne) check() error {
	if v, ok := suo.mutation.Year(); ok {
		if err := submit.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Submit.year": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Score(); ok {
		if err := submit.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Submit.score": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Language(); ok {
		if err := submit.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Submit.language": %w`, err)}
		}
	}
	return nil
}

func (suo *SubmitUpdateOne) sqlSave(ctx context.Context) (_node *Submit, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Submit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submit.FieldID)
		for _, f := range fields {
			if !submit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != submit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.IPAddr(); ok {
		_spec.SetField(submit.FieldIPAddr, field.TypeString, value)
	}
	if value, ok := suo.mutation.Year(); ok {
		_spec.SetField(submit.FieldYear, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedYear(); ok {
		_spec.AddField(submit.FieldYear, field.TypeInt, value)
	}
	if value, ok := suo.mutation.Score(); ok {
		_spec.SetField(submit.FieldScore, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedScore(); ok {
		_spec.AddField(submit.FieldScore, field.TypeInt, value)
	}
	if suo.mutation.ScoreCleared() {
		_spec.ClearField(submit.FieldScore, field.TypeInt)
	}
	if value, ok := suo.mutation.Language(); ok {
		_spec.SetField(submit.FieldLanguage, field.TypeEnum, value)
	}
	if suo.mutation.LanguageCleared() {
		_spec.ClearField(submit.FieldLanguage, field.TypeEnum)
	}
	if value, ok := suo.mutation.SubmitedAt(); ok {
		_spec.SetField(submit.FieldSubmitedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CompletedAt(); ok {
		_spec.SetField(submit.FieldCompletedAt, field.TypeTime, value)
	}
	if suo.mutation.CompletedAtCleared() {
		_spec.ClearField(submit.FieldCompletedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(submit.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UpdatedAtCleared() {
		_spec.ClearField(submit.FieldUpdatedAt, field.TypeTime)
	}
	if suo.mutation.TagResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TagResultsTable,
			Columns: []string{submit.TagResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTagResultsIDs(); len(nodes) > 0 && !suo.mutation.TagResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TagResultsTable,
			Columns: []string{submit.TagResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TagResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TagResultsTable,
			Columns: []string{submit.TagResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: submit.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !suo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: submit.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: submit.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ContestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: submit.ContestsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedContestsIDs(); len(nodes) > 0 && !suo.mutation.ContestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: submit.ContestsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ContestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: submit.ContestsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Submit{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
