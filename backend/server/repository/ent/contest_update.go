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
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
)

// ContestUpdate is the builder for updating Contest entities.
type ContestUpdate struct {
	config
	hooks    []Hook
	mutation *ContestMutation
}

// Where appends a list predicates to the ContestUpdate builder.
func (cu *ContestUpdate) Where(ps ...predicate.Contest) *ContestUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ContestUpdate) SetTitle(s string) *ContestUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetStartAt sets the "start_at" field.
func (cu *ContestUpdate) SetStartAt(t time.Time) *ContestUpdate {
	cu.mutation.SetStartAt(t)
	return cu
}

// SetEndAt sets the "end_at" field.
func (cu *ContestUpdate) SetEndAt(t time.Time) *ContestUpdate {
	cu.mutation.SetEndAt(t)
	return cu
}

// SetSubmitLimit sets the "submit_limit" field.
func (cu *ContestUpdate) SetSubmitLimit(i int) *ContestUpdate {
	cu.mutation.ResetSubmitLimit()
	cu.mutation.SetSubmitLimit(i)
	return cu
}

// AddSubmitLimit adds i to the "submit_limit" field.
func (cu *ContestUpdate) AddSubmitLimit(i int) *ContestUpdate {
	cu.mutation.AddSubmitLimit(i)
	return cu
}

// SetSlug sets the "slug" field.
func (cu *ContestUpdate) SetSlug(s string) *ContestUpdate {
	cu.mutation.SetSlug(s)
	return cu
}

// SetTagSelectionLogic sets the "tag_selection_logic" field.
func (cu *ContestUpdate) SetTagSelectionLogic(csl contest.TagSelectionLogic) *ContestUpdate {
	cu.mutation.SetTagSelectionLogic(csl)
	return cu
}

// SetValidator sets the "validator" field.
func (cu *ContestUpdate) SetValidator(s string) *ContestUpdate {
	cu.mutation.SetValidator(s)
	return cu
}

// SetTimeLimitPerTask sets the "time_limit_per_task" field.
func (cu *ContestUpdate) SetTimeLimitPerTask(i int64) *ContestUpdate {
	cu.mutation.ResetTimeLimitPerTask()
	cu.mutation.SetTimeLimitPerTask(i)
	return cu
}

// SetNillableTimeLimitPerTask sets the "time_limit_per_task" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableTimeLimitPerTask(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetTimeLimitPerTask(*i)
	}
	return cu
}

// AddTimeLimitPerTask adds i to the "time_limit_per_task" field.
func (cu *ContestUpdate) AddTimeLimitPerTask(i int64) *ContestUpdate {
	cu.mutation.AddTimeLimitPerTask(i)
	return cu
}

// ClearTimeLimitPerTask clears the value of the "time_limit_per_task" field.
func (cu *ContestUpdate) ClearTimeLimitPerTask() *ContestUpdate {
	cu.mutation.ClearTimeLimitPerTask()
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *ContestUpdate) SetCreatedAt(t time.Time) *ContestUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ContestUpdate) SetUpdatedAt(t time.Time) *ContestUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableUpdatedAt(t *time.Time) *ContestUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cu *ContestUpdate) ClearUpdatedAt() *ContestUpdate {
	cu.mutation.ClearUpdatedAt()
	return cu
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (cu *ContestUpdate) AddSubmitIDs(ids ...int) *ContestUpdate {
	cu.mutation.AddSubmitIDs(ids...)
	return cu
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (cu *ContestUpdate) AddSubmits(s ...*Submit) *ContestUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSubmitIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cu *ContestUpdate) Mutation() *ContestMutation {
	return cu.mutation
}

// ClearSubmits clears all "submits" edges to the Submit entity.
func (cu *ContestUpdate) ClearSubmits() *ContestUpdate {
	cu.mutation.ClearSubmits()
	return cu
}

// RemoveSubmitIDs removes the "submits" edge to Submit entities by IDs.
func (cu *ContestUpdate) RemoveSubmitIDs(ids ...int) *ContestUpdate {
	cu.mutation.RemoveSubmitIDs(ids...)
	return cu
}

// RemoveSubmits removes "submits" edges to Submit entities.
func (cu *ContestUpdate) RemoveSubmits(s ...*Submit) *ContestUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSubmitIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContestUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContestUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContestUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ContestUpdate) check() error {
	if v, ok := cu.mutation.TagSelectionLogic(); ok {
		if err := contest.TagSelectionLogicValidator(v); err != nil {
			return &ValidationError{Name: "tag_selection_logic", err: fmt.Errorf(`ent: validator failed for field "Contest.tag_selection_logic": %w`, err)}
		}
	}
	return nil
}

func (cu *ContestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(contest.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.SubmitLimit(); ok {
		_spec.SetField(contest.FieldSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedSubmitLimit(); ok {
		_spec.AddField(contest.FieldSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Slug(); ok {
		_spec.SetField(contest.FieldSlug, field.TypeString, value)
	}
	if value, ok := cu.mutation.TagSelectionLogic(); ok {
		_spec.SetField(contest.FieldTagSelectionLogic, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Validator(); ok {
		_spec.SetField(contest.FieldValidator, field.TypeString, value)
	}
	if value, ok := cu.mutation.TimeLimitPerTask(); ok {
		_spec.SetField(contest.FieldTimeLimitPerTask, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedTimeLimitPerTask(); ok {
		_spec.AddField(contest.FieldTimeLimitPerTask, field.TypeInt64, value)
	}
	if cu.mutation.TimeLimitPerTaskCleared() {
		_spec.ClearField(contest.FieldTimeLimitPerTask, field.TypeInt64)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(contest.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.ClearField(contest.FieldUpdatedAt, field.TypeTime)
	}
	if cu.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSubmitsIDs(); len(nodes) > 0 && !cu.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContestUpdateOne is the builder for updating a single Contest entity.
type ContestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContestMutation
}

// SetTitle sets the "title" field.
func (cuo *ContestUpdateOne) SetTitle(s string) *ContestUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetStartAt sets the "start_at" field.
func (cuo *ContestUpdateOne) SetStartAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetStartAt(t)
	return cuo
}

// SetEndAt sets the "end_at" field.
func (cuo *ContestUpdateOne) SetEndAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetEndAt(t)
	return cuo
}

// SetSubmitLimit sets the "submit_limit" field.
func (cuo *ContestUpdateOne) SetSubmitLimit(i int) *ContestUpdateOne {
	cuo.mutation.ResetSubmitLimit()
	cuo.mutation.SetSubmitLimit(i)
	return cuo
}

// AddSubmitLimit adds i to the "submit_limit" field.
func (cuo *ContestUpdateOne) AddSubmitLimit(i int) *ContestUpdateOne {
	cuo.mutation.AddSubmitLimit(i)
	return cuo
}

// SetSlug sets the "slug" field.
func (cuo *ContestUpdateOne) SetSlug(s string) *ContestUpdateOne {
	cuo.mutation.SetSlug(s)
	return cuo
}

// SetTagSelectionLogic sets the "tag_selection_logic" field.
func (cuo *ContestUpdateOne) SetTagSelectionLogic(csl contest.TagSelectionLogic) *ContestUpdateOne {
	cuo.mutation.SetTagSelectionLogic(csl)
	return cuo
}

// SetValidator sets the "validator" field.
func (cuo *ContestUpdateOne) SetValidator(s string) *ContestUpdateOne {
	cuo.mutation.SetValidator(s)
	return cuo
}

// SetTimeLimitPerTask sets the "time_limit_per_task" field.
func (cuo *ContestUpdateOne) SetTimeLimitPerTask(i int64) *ContestUpdateOne {
	cuo.mutation.ResetTimeLimitPerTask()
	cuo.mutation.SetTimeLimitPerTask(i)
	return cuo
}

// SetNillableTimeLimitPerTask sets the "time_limit_per_task" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableTimeLimitPerTask(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetTimeLimitPerTask(*i)
	}
	return cuo
}

// AddTimeLimitPerTask adds i to the "time_limit_per_task" field.
func (cuo *ContestUpdateOne) AddTimeLimitPerTask(i int64) *ContestUpdateOne {
	cuo.mutation.AddTimeLimitPerTask(i)
	return cuo
}

// ClearTimeLimitPerTask clears the value of the "time_limit_per_task" field.
func (cuo *ContestUpdateOne) ClearTimeLimitPerTask() *ContestUpdateOne {
	cuo.mutation.ClearTimeLimitPerTask()
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *ContestUpdateOne) SetCreatedAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ContestUpdateOne) SetUpdatedAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableUpdatedAt(t *time.Time) *ContestUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cuo *ContestUpdateOne) ClearUpdatedAt() *ContestUpdateOne {
	cuo.mutation.ClearUpdatedAt()
	return cuo
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (cuo *ContestUpdateOne) AddSubmitIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.AddSubmitIDs(ids...)
	return cuo
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (cuo *ContestUpdateOne) AddSubmits(s ...*Submit) *ContestUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSubmitIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cuo *ContestUpdateOne) Mutation() *ContestMutation {
	return cuo.mutation
}

// ClearSubmits clears all "submits" edges to the Submit entity.
func (cuo *ContestUpdateOne) ClearSubmits() *ContestUpdateOne {
	cuo.mutation.ClearSubmits()
	return cuo
}

// RemoveSubmitIDs removes the "submits" edge to Submit entities by IDs.
func (cuo *ContestUpdateOne) RemoveSubmitIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.RemoveSubmitIDs(ids...)
	return cuo
}

// RemoveSubmits removes "submits" edges to Submit entities.
func (cuo *ContestUpdateOne) RemoveSubmits(s ...*Submit) *ContestUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSubmitIDs(ids...)
}

// Where appends a list predicates to the ContestUpdate builder.
func (cuo *ContestUpdateOne) Where(ps ...predicate.Contest) *ContestUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContestUpdateOne) Select(field string, fields ...string) *ContestUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contest entity.
func (cuo *ContestUpdateOne) Save(ctx context.Context) (*Contest, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContestUpdateOne) SaveX(ctx context.Context) *Contest {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContestUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContestUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ContestUpdateOne) check() error {
	if v, ok := cuo.mutation.TagSelectionLogic(); ok {
		if err := contest.TagSelectionLogicValidator(v); err != nil {
			return &ValidationError{Name: "tag_selection_logic", err: fmt.Errorf(`ent: validator failed for field "Contest.tag_selection_logic": %w`, err)}
		}
	}
	return nil
}

func (cuo *ContestUpdateOne) sqlSave(ctx context.Context) (_node *Contest, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contest.FieldID)
		for _, f := range fields {
			if !contest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(contest.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.SubmitLimit(); ok {
		_spec.SetField(contest.FieldSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedSubmitLimit(); ok {
		_spec.AddField(contest.FieldSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Slug(); ok {
		_spec.SetField(contest.FieldSlug, field.TypeString, value)
	}
	if value, ok := cuo.mutation.TagSelectionLogic(); ok {
		_spec.SetField(contest.FieldTagSelectionLogic, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Validator(); ok {
		_spec.SetField(contest.FieldValidator, field.TypeString, value)
	}
	if value, ok := cuo.mutation.TimeLimitPerTask(); ok {
		_spec.SetField(contest.FieldTimeLimitPerTask, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedTimeLimitPerTask(); ok {
		_spec.AddField(contest.FieldTimeLimitPerTask, field.TypeInt64, value)
	}
	if cuo.mutation.TimeLimitPerTaskCleared() {
		_spec.ClearField(contest.FieldTimeLimitPerTask, field.TypeInt64)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(contest.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(contest.FieldUpdatedAt, field.TypeTime)
	}
	if cuo.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSubmitsIDs(); len(nodes) > 0 && !cuo.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Contest{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
