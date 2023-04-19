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
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
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

// SetQualifierStartAt sets the "qualifier_start_at" field.
func (cu *ContestUpdate) SetQualifierStartAt(t time.Time) *ContestUpdate {
	cu.mutation.SetQualifierStartAt(t)
	return cu
}

// SetQualifierEndAt sets the "qualifier_end_at" field.
func (cu *ContestUpdate) SetQualifierEndAt(t time.Time) *ContestUpdate {
	cu.mutation.SetQualifierEndAt(t)
	return cu
}

// SetQualifierSubmitLimit sets the "qualifier_submit_limit" field.
func (cu *ContestUpdate) SetQualifierSubmitLimit(i int) *ContestUpdate {
	cu.mutation.ResetQualifierSubmitLimit()
	cu.mutation.SetQualifierSubmitLimit(i)
	return cu
}

// AddQualifierSubmitLimit adds i to the "qualifier_submit_limit" field.
func (cu *ContestUpdate) AddQualifierSubmitLimit(i int) *ContestUpdate {
	cu.mutation.AddQualifierSubmitLimit(i)
	return cu
}

// SetFinalStartAt sets the "final_start_at" field.
func (cu *ContestUpdate) SetFinalStartAt(t time.Time) *ContestUpdate {
	cu.mutation.SetFinalStartAt(t)
	return cu
}

// SetFinalEndAt sets the "final_end_at" field.
func (cu *ContestUpdate) SetFinalEndAt(t time.Time) *ContestUpdate {
	cu.mutation.SetFinalEndAt(t)
	return cu
}

// SetFinalSubmitLimit sets the "final_submit_limit" field.
func (cu *ContestUpdate) SetFinalSubmitLimit(i int) *ContestUpdate {
	cu.mutation.ResetFinalSubmitLimit()
	cu.mutation.SetFinalSubmitLimit(i)
	return cu
}

// AddFinalSubmitLimit adds i to the "final_submit_limit" field.
func (cu *ContestUpdate) AddFinalSubmitLimit(i int) *ContestUpdate {
	cu.mutation.AddFinalSubmitLimit(i)
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

// Mutation returns the ContestMutation object of the builder.
func (cu *ContestUpdate) Mutation() *ContestMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ContestMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
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

func (cu *ContestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.QualifierStartAt(); ok {
		_spec.SetField(contest.FieldQualifierStartAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.QualifierEndAt(); ok {
		_spec.SetField(contest.FieldQualifierEndAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.QualifierSubmitLimit(); ok {
		_spec.SetField(contest.FieldQualifierSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedQualifierSubmitLimit(); ok {
		_spec.AddField(contest.FieldQualifierSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cu.mutation.FinalStartAt(); ok {
		_spec.SetField(contest.FieldFinalStartAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.FinalEndAt(); ok {
		_spec.SetField(contest.FieldFinalEndAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.FinalSubmitLimit(); ok {
		_spec.SetField(contest.FieldFinalSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedFinalSubmitLimit(); ok {
		_spec.AddField(contest.FieldFinalSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.ClearField(contest.FieldUpdatedAt, field.TypeTime)
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

// SetQualifierStartAt sets the "qualifier_start_at" field.
func (cuo *ContestUpdateOne) SetQualifierStartAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetQualifierStartAt(t)
	return cuo
}

// SetQualifierEndAt sets the "qualifier_end_at" field.
func (cuo *ContestUpdateOne) SetQualifierEndAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetQualifierEndAt(t)
	return cuo
}

// SetQualifierSubmitLimit sets the "qualifier_submit_limit" field.
func (cuo *ContestUpdateOne) SetQualifierSubmitLimit(i int) *ContestUpdateOne {
	cuo.mutation.ResetQualifierSubmitLimit()
	cuo.mutation.SetQualifierSubmitLimit(i)
	return cuo
}

// AddQualifierSubmitLimit adds i to the "qualifier_submit_limit" field.
func (cuo *ContestUpdateOne) AddQualifierSubmitLimit(i int) *ContestUpdateOne {
	cuo.mutation.AddQualifierSubmitLimit(i)
	return cuo
}

// SetFinalStartAt sets the "final_start_at" field.
func (cuo *ContestUpdateOne) SetFinalStartAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetFinalStartAt(t)
	return cuo
}

// SetFinalEndAt sets the "final_end_at" field.
func (cuo *ContestUpdateOne) SetFinalEndAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetFinalEndAt(t)
	return cuo
}

// SetFinalSubmitLimit sets the "final_submit_limit" field.
func (cuo *ContestUpdateOne) SetFinalSubmitLimit(i int) *ContestUpdateOne {
	cuo.mutation.ResetFinalSubmitLimit()
	cuo.mutation.SetFinalSubmitLimit(i)
	return cuo
}

// AddFinalSubmitLimit adds i to the "final_submit_limit" field.
func (cuo *ContestUpdateOne) AddFinalSubmitLimit(i int) *ContestUpdateOne {
	cuo.mutation.AddFinalSubmitLimit(i)
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

// Mutation returns the ContestMutation object of the builder.
func (cuo *ContestUpdateOne) Mutation() *ContestMutation {
	return cuo.mutation
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
	return withHooks[*Contest, ContestMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
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

func (cuo *ContestUpdateOne) sqlSave(ctx context.Context) (_node *Contest, err error) {
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
	if value, ok := cuo.mutation.QualifierStartAt(); ok {
		_spec.SetField(contest.FieldQualifierStartAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.QualifierEndAt(); ok {
		_spec.SetField(contest.FieldQualifierEndAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.QualifierSubmitLimit(); ok {
		_spec.SetField(contest.FieldQualifierSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedQualifierSubmitLimit(); ok {
		_spec.AddField(contest.FieldQualifierSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.FinalStartAt(); ok {
		_spec.SetField(contest.FieldFinalStartAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.FinalEndAt(); ok {
		_spec.SetField(contest.FieldFinalEndAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.FinalSubmitLimit(); ok {
		_spec.SetField(contest.FieldFinalSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedFinalSubmitLimit(); ok {
		_spec.AddField(contest.FieldFinalSubmitLimit, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(contest.FieldUpdatedAt, field.TypeTime)
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
