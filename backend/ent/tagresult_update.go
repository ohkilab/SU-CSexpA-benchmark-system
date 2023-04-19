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
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/tagresult"
)

// TagResultUpdate is the builder for updating TagResult entities.
type TagResultUpdate struct {
	config
	hooks    []Hook
	mutation *TagResultMutation
}

// Where appends a list predicates to the TagResultUpdate builder.
func (tru *TagResultUpdate) Where(ps ...predicate.TagResult) *TagResultUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetName sets the "name" field.
func (tru *TagResultUpdate) SetName(s string) *TagResultUpdate {
	tru.mutation.SetName(s)
	return tru
}

// SetScore sets the "score" field.
func (tru *TagResultUpdate) SetScore(i int) *TagResultUpdate {
	tru.mutation.ResetScore()
	tru.mutation.SetScore(i)
	return tru
}

// AddScore adds i to the "score" field.
func (tru *TagResultUpdate) AddScore(i int) *TagResultUpdate {
	tru.mutation.AddScore(i)
	return tru
}

// SetCreatedAt sets the "created_at" field.
func (tru *TagResultUpdate) SetCreatedAt(t time.Time) *TagResultUpdate {
	tru.mutation.SetCreatedAt(t)
	return tru
}

// SetDeletedAt sets the "deleted_at" field.
func (tru *TagResultUpdate) SetDeletedAt(t time.Time) *TagResultUpdate {
	tru.mutation.SetDeletedAt(t)
	return tru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tru *TagResultUpdate) SetNillableDeletedAt(t *time.Time) *TagResultUpdate {
	if t != nil {
		tru.SetDeletedAt(*t)
	}
	return tru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tru *TagResultUpdate) ClearDeletedAt() *TagResultUpdate {
	tru.mutation.ClearDeletedAt()
	return tru
}

// Mutation returns the TagResultMutation object of the builder.
func (tru *TagResultUpdate) Mutation() *TagResultMutation {
	return tru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TagResultUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TagResultMutation](ctx, tru.sqlSave, tru.mutation, tru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TagResultUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TagResultUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TagResultUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tru *TagResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tagresult.Table, tagresult.Columns, sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt))
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.Name(); ok {
		_spec.SetField(tagresult.FieldName, field.TypeString, value)
	}
	if value, ok := tru.mutation.Score(); ok {
		_spec.SetField(tagresult.FieldScore, field.TypeInt, value)
	}
	if value, ok := tru.mutation.AddedScore(); ok {
		_spec.AddField(tagresult.FieldScore, field.TypeInt, value)
	}
	if value, ok := tru.mutation.CreatedAt(); ok {
		_spec.SetField(tagresult.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tru.mutation.DeletedAt(); ok {
		_spec.SetField(tagresult.FieldDeletedAt, field.TypeTime, value)
	}
	if tru.mutation.DeletedAtCleared() {
		_spec.ClearField(tagresult.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tagresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tru.mutation.done = true
	return n, nil
}

// TagResultUpdateOne is the builder for updating a single TagResult entity.
type TagResultUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagResultMutation
}

// SetName sets the "name" field.
func (truo *TagResultUpdateOne) SetName(s string) *TagResultUpdateOne {
	truo.mutation.SetName(s)
	return truo
}

// SetScore sets the "score" field.
func (truo *TagResultUpdateOne) SetScore(i int) *TagResultUpdateOne {
	truo.mutation.ResetScore()
	truo.mutation.SetScore(i)
	return truo
}

// AddScore adds i to the "score" field.
func (truo *TagResultUpdateOne) AddScore(i int) *TagResultUpdateOne {
	truo.mutation.AddScore(i)
	return truo
}

// SetCreatedAt sets the "created_at" field.
func (truo *TagResultUpdateOne) SetCreatedAt(t time.Time) *TagResultUpdateOne {
	truo.mutation.SetCreatedAt(t)
	return truo
}

// SetDeletedAt sets the "deleted_at" field.
func (truo *TagResultUpdateOne) SetDeletedAt(t time.Time) *TagResultUpdateOne {
	truo.mutation.SetDeletedAt(t)
	return truo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (truo *TagResultUpdateOne) SetNillableDeletedAt(t *time.Time) *TagResultUpdateOne {
	if t != nil {
		truo.SetDeletedAt(*t)
	}
	return truo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (truo *TagResultUpdateOne) ClearDeletedAt() *TagResultUpdateOne {
	truo.mutation.ClearDeletedAt()
	return truo
}

// Mutation returns the TagResultMutation object of the builder.
func (truo *TagResultUpdateOne) Mutation() *TagResultMutation {
	return truo.mutation
}

// Where appends a list predicates to the TagResultUpdate builder.
func (truo *TagResultUpdateOne) Where(ps ...predicate.TagResult) *TagResultUpdateOne {
	truo.mutation.Where(ps...)
	return truo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TagResultUpdateOne) Select(field string, fields ...string) *TagResultUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TagResult entity.
func (truo *TagResultUpdateOne) Save(ctx context.Context) (*TagResult, error) {
	return withHooks[*TagResult, TagResultMutation](ctx, truo.sqlSave, truo.mutation, truo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TagResultUpdateOne) SaveX(ctx context.Context) *TagResult {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TagResultUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TagResultUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (truo *TagResultUpdateOne) sqlSave(ctx context.Context) (_node *TagResult, err error) {
	_spec := sqlgraph.NewUpdateSpec(tagresult.Table, tagresult.Columns, sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt))
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TagResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tagresult.FieldID)
		for _, f := range fields {
			if !tagresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tagresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := truo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := truo.mutation.Name(); ok {
		_spec.SetField(tagresult.FieldName, field.TypeString, value)
	}
	if value, ok := truo.mutation.Score(); ok {
		_spec.SetField(tagresult.FieldScore, field.TypeInt, value)
	}
	if value, ok := truo.mutation.AddedScore(); ok {
		_spec.AddField(tagresult.FieldScore, field.TypeInt, value)
	}
	if value, ok := truo.mutation.CreatedAt(); ok {
		_spec.SetField(tagresult.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := truo.mutation.DeletedAt(); ok {
		_spec.SetField(tagresult.FieldDeletedAt, field.TypeTime, value)
	}
	if truo.mutation.DeletedAtCleared() {
		_spec.ClearField(tagresult.FieldDeletedAt, field.TypeTime)
	}
	_node = &TagResult{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tagresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	truo.mutation.done = true
	return _node, nil
}