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
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/taskresult"
)

// TaskResultUpdate is the builder for updating TaskResult entities.
type TaskResultUpdate struct {
	config
	hooks    []Hook
	mutation *TaskResultMutation
}

// Where appends a list predicates to the TaskResultUpdate builder.
func (tru *TaskResultUpdate) Where(ps ...predicate.TaskResult) *TaskResultUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetRequestPerSec sets the "request_per_sec" field.
func (tru *TaskResultUpdate) SetRequestPerSec(i int) *TaskResultUpdate {
	tru.mutation.ResetRequestPerSec()
	tru.mutation.SetRequestPerSec(i)
	return tru
}

// AddRequestPerSec adds i to the "request_per_sec" field.
func (tru *TaskResultUpdate) AddRequestPerSec(i int) *TaskResultUpdate {
	tru.mutation.AddRequestPerSec(i)
	return tru
}

// SetErrorMessage sets the "error_message" field.
func (tru *TaskResultUpdate) SetErrorMessage(s string) *TaskResultUpdate {
	tru.mutation.SetErrorMessage(s)
	return tru
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (tru *TaskResultUpdate) SetNillableErrorMessage(s *string) *TaskResultUpdate {
	if s != nil {
		tru.SetErrorMessage(*s)
	}
	return tru
}

// ClearErrorMessage clears the value of the "error_message" field.
func (tru *TaskResultUpdate) ClearErrorMessage() *TaskResultUpdate {
	tru.mutation.ClearErrorMessage()
	return tru
}

// SetURL sets the "url" field.
func (tru *TaskResultUpdate) SetURL(s string) *TaskResultUpdate {
	tru.mutation.SetURL(s)
	return tru
}

// SetMethod sets the "method" field.
func (tru *TaskResultUpdate) SetMethod(s string) *TaskResultUpdate {
	tru.mutation.SetMethod(s)
	return tru
}

// SetRequestContentType sets the "request_content_type" field.
func (tru *TaskResultUpdate) SetRequestContentType(s string) *TaskResultUpdate {
	tru.mutation.SetRequestContentType(s)
	return tru
}

// SetRequestBody sets the "request_body" field.
func (tru *TaskResultUpdate) SetRequestBody(s string) *TaskResultUpdate {
	tru.mutation.SetRequestBody(s)
	return tru
}

// SetNillableRequestBody sets the "request_body" field if the given value is not nil.
func (tru *TaskResultUpdate) SetNillableRequestBody(s *string) *TaskResultUpdate {
	if s != nil {
		tru.SetRequestBody(*s)
	}
	return tru
}

// ClearRequestBody clears the value of the "request_body" field.
func (tru *TaskResultUpdate) ClearRequestBody() *TaskResultUpdate {
	tru.mutation.ClearRequestBody()
	return tru
}

// SetThreadNum sets the "thread_num" field.
func (tru *TaskResultUpdate) SetThreadNum(i int) *TaskResultUpdate {
	tru.mutation.ResetThreadNum()
	tru.mutation.SetThreadNum(i)
	return tru
}

// AddThreadNum adds i to the "thread_num" field.
func (tru *TaskResultUpdate) AddThreadNum(i int) *TaskResultUpdate {
	tru.mutation.AddThreadNum(i)
	return tru
}

// SetAttemptCount sets the "attempt_count" field.
func (tru *TaskResultUpdate) SetAttemptCount(i int) *TaskResultUpdate {
	tru.mutation.ResetAttemptCount()
	tru.mutation.SetAttemptCount(i)
	return tru
}

// AddAttemptCount adds i to the "attempt_count" field.
func (tru *TaskResultUpdate) AddAttemptCount(i int) *TaskResultUpdate {
	tru.mutation.AddAttemptCount(i)
	return tru
}

// SetCreatedAt sets the "created_at" field.
func (tru *TaskResultUpdate) SetCreatedAt(t time.Time) *TaskResultUpdate {
	tru.mutation.SetCreatedAt(t)
	return tru
}

// SetDeletedAt sets the "deleted_at" field.
func (tru *TaskResultUpdate) SetDeletedAt(t time.Time) *TaskResultUpdate {
	tru.mutation.SetDeletedAt(t)
	return tru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tru *TaskResultUpdate) SetNillableDeletedAt(t *time.Time) *TaskResultUpdate {
	if t != nil {
		tru.SetDeletedAt(*t)
	}
	return tru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tru *TaskResultUpdate) ClearDeletedAt() *TaskResultUpdate {
	tru.mutation.ClearDeletedAt()
	return tru
}

// Mutation returns the TaskResultMutation object of the builder.
func (tru *TaskResultUpdate) Mutation() *TaskResultMutation {
	return tru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TaskResultUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TaskResultMutation](ctx, tru.sqlSave, tru.mutation, tru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TaskResultUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TaskResultUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TaskResultUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tru *TaskResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(taskresult.Table, taskresult.Columns, sqlgraph.NewFieldSpec(taskresult.FieldID, field.TypeInt))
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.RequestPerSec(); ok {
		_spec.SetField(taskresult.FieldRequestPerSec, field.TypeInt, value)
	}
	if value, ok := tru.mutation.AddedRequestPerSec(); ok {
		_spec.AddField(taskresult.FieldRequestPerSec, field.TypeInt, value)
	}
	if value, ok := tru.mutation.ErrorMessage(); ok {
		_spec.SetField(taskresult.FieldErrorMessage, field.TypeString, value)
	}
	if tru.mutation.ErrorMessageCleared() {
		_spec.ClearField(taskresult.FieldErrorMessage, field.TypeString)
	}
	if value, ok := tru.mutation.URL(); ok {
		_spec.SetField(taskresult.FieldURL, field.TypeString, value)
	}
	if value, ok := tru.mutation.Method(); ok {
		_spec.SetField(taskresult.FieldMethod, field.TypeString, value)
	}
	if value, ok := tru.mutation.RequestContentType(); ok {
		_spec.SetField(taskresult.FieldRequestContentType, field.TypeString, value)
	}
	if value, ok := tru.mutation.RequestBody(); ok {
		_spec.SetField(taskresult.FieldRequestBody, field.TypeString, value)
	}
	if tru.mutation.RequestBodyCleared() {
		_spec.ClearField(taskresult.FieldRequestBody, field.TypeString)
	}
	if value, ok := tru.mutation.ThreadNum(); ok {
		_spec.SetField(taskresult.FieldThreadNum, field.TypeInt, value)
	}
	if value, ok := tru.mutation.AddedThreadNum(); ok {
		_spec.AddField(taskresult.FieldThreadNum, field.TypeInt, value)
	}
	if value, ok := tru.mutation.AttemptCount(); ok {
		_spec.SetField(taskresult.FieldAttemptCount, field.TypeInt, value)
	}
	if value, ok := tru.mutation.AddedAttemptCount(); ok {
		_spec.AddField(taskresult.FieldAttemptCount, field.TypeInt, value)
	}
	if value, ok := tru.mutation.CreatedAt(); ok {
		_spec.SetField(taskresult.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tru.mutation.DeletedAt(); ok {
		_spec.SetField(taskresult.FieldDeletedAt, field.TypeTime, value)
	}
	if tru.mutation.DeletedAtCleared() {
		_spec.ClearField(taskresult.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tru.mutation.done = true
	return n, nil
}

// TaskResultUpdateOne is the builder for updating a single TaskResult entity.
type TaskResultUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskResultMutation
}

// SetRequestPerSec sets the "request_per_sec" field.
func (truo *TaskResultUpdateOne) SetRequestPerSec(i int) *TaskResultUpdateOne {
	truo.mutation.ResetRequestPerSec()
	truo.mutation.SetRequestPerSec(i)
	return truo
}

// AddRequestPerSec adds i to the "request_per_sec" field.
func (truo *TaskResultUpdateOne) AddRequestPerSec(i int) *TaskResultUpdateOne {
	truo.mutation.AddRequestPerSec(i)
	return truo
}

// SetErrorMessage sets the "error_message" field.
func (truo *TaskResultUpdateOne) SetErrorMessage(s string) *TaskResultUpdateOne {
	truo.mutation.SetErrorMessage(s)
	return truo
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (truo *TaskResultUpdateOne) SetNillableErrorMessage(s *string) *TaskResultUpdateOne {
	if s != nil {
		truo.SetErrorMessage(*s)
	}
	return truo
}

// ClearErrorMessage clears the value of the "error_message" field.
func (truo *TaskResultUpdateOne) ClearErrorMessage() *TaskResultUpdateOne {
	truo.mutation.ClearErrorMessage()
	return truo
}

// SetURL sets the "url" field.
func (truo *TaskResultUpdateOne) SetURL(s string) *TaskResultUpdateOne {
	truo.mutation.SetURL(s)
	return truo
}

// SetMethod sets the "method" field.
func (truo *TaskResultUpdateOne) SetMethod(s string) *TaskResultUpdateOne {
	truo.mutation.SetMethod(s)
	return truo
}

// SetRequestContentType sets the "request_content_type" field.
func (truo *TaskResultUpdateOne) SetRequestContentType(s string) *TaskResultUpdateOne {
	truo.mutation.SetRequestContentType(s)
	return truo
}

// SetRequestBody sets the "request_body" field.
func (truo *TaskResultUpdateOne) SetRequestBody(s string) *TaskResultUpdateOne {
	truo.mutation.SetRequestBody(s)
	return truo
}

// SetNillableRequestBody sets the "request_body" field if the given value is not nil.
func (truo *TaskResultUpdateOne) SetNillableRequestBody(s *string) *TaskResultUpdateOne {
	if s != nil {
		truo.SetRequestBody(*s)
	}
	return truo
}

// ClearRequestBody clears the value of the "request_body" field.
func (truo *TaskResultUpdateOne) ClearRequestBody() *TaskResultUpdateOne {
	truo.mutation.ClearRequestBody()
	return truo
}

// SetThreadNum sets the "thread_num" field.
func (truo *TaskResultUpdateOne) SetThreadNum(i int) *TaskResultUpdateOne {
	truo.mutation.ResetThreadNum()
	truo.mutation.SetThreadNum(i)
	return truo
}

// AddThreadNum adds i to the "thread_num" field.
func (truo *TaskResultUpdateOne) AddThreadNum(i int) *TaskResultUpdateOne {
	truo.mutation.AddThreadNum(i)
	return truo
}

// SetAttemptCount sets the "attempt_count" field.
func (truo *TaskResultUpdateOne) SetAttemptCount(i int) *TaskResultUpdateOne {
	truo.mutation.ResetAttemptCount()
	truo.mutation.SetAttemptCount(i)
	return truo
}

// AddAttemptCount adds i to the "attempt_count" field.
func (truo *TaskResultUpdateOne) AddAttemptCount(i int) *TaskResultUpdateOne {
	truo.mutation.AddAttemptCount(i)
	return truo
}

// SetCreatedAt sets the "created_at" field.
func (truo *TaskResultUpdateOne) SetCreatedAt(t time.Time) *TaskResultUpdateOne {
	truo.mutation.SetCreatedAt(t)
	return truo
}

// SetDeletedAt sets the "deleted_at" field.
func (truo *TaskResultUpdateOne) SetDeletedAt(t time.Time) *TaskResultUpdateOne {
	truo.mutation.SetDeletedAt(t)
	return truo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (truo *TaskResultUpdateOne) SetNillableDeletedAt(t *time.Time) *TaskResultUpdateOne {
	if t != nil {
		truo.SetDeletedAt(*t)
	}
	return truo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (truo *TaskResultUpdateOne) ClearDeletedAt() *TaskResultUpdateOne {
	truo.mutation.ClearDeletedAt()
	return truo
}

// Mutation returns the TaskResultMutation object of the builder.
func (truo *TaskResultUpdateOne) Mutation() *TaskResultMutation {
	return truo.mutation
}

// Where appends a list predicates to the TaskResultUpdate builder.
func (truo *TaskResultUpdateOne) Where(ps ...predicate.TaskResult) *TaskResultUpdateOne {
	truo.mutation.Where(ps...)
	return truo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TaskResultUpdateOne) Select(field string, fields ...string) *TaskResultUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TaskResult entity.
func (truo *TaskResultUpdateOne) Save(ctx context.Context) (*TaskResult, error) {
	return withHooks[*TaskResult, TaskResultMutation](ctx, truo.sqlSave, truo.mutation, truo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TaskResultUpdateOne) SaveX(ctx context.Context) *TaskResult {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TaskResultUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TaskResultUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (truo *TaskResultUpdateOne) sqlSave(ctx context.Context) (_node *TaskResult, err error) {
	_spec := sqlgraph.NewUpdateSpec(taskresult.Table, taskresult.Columns, sqlgraph.NewFieldSpec(taskresult.FieldID, field.TypeInt))
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskresult.FieldID)
		for _, f := range fields {
			if !taskresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskresult.FieldID {
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
	if value, ok := truo.mutation.RequestPerSec(); ok {
		_spec.SetField(taskresult.FieldRequestPerSec, field.TypeInt, value)
	}
	if value, ok := truo.mutation.AddedRequestPerSec(); ok {
		_spec.AddField(taskresult.FieldRequestPerSec, field.TypeInt, value)
	}
	if value, ok := truo.mutation.ErrorMessage(); ok {
		_spec.SetField(taskresult.FieldErrorMessage, field.TypeString, value)
	}
	if truo.mutation.ErrorMessageCleared() {
		_spec.ClearField(taskresult.FieldErrorMessage, field.TypeString)
	}
	if value, ok := truo.mutation.URL(); ok {
		_spec.SetField(taskresult.FieldURL, field.TypeString, value)
	}
	if value, ok := truo.mutation.Method(); ok {
		_spec.SetField(taskresult.FieldMethod, field.TypeString, value)
	}
	if value, ok := truo.mutation.RequestContentType(); ok {
		_spec.SetField(taskresult.FieldRequestContentType, field.TypeString, value)
	}
	if value, ok := truo.mutation.RequestBody(); ok {
		_spec.SetField(taskresult.FieldRequestBody, field.TypeString, value)
	}
	if truo.mutation.RequestBodyCleared() {
		_spec.ClearField(taskresult.FieldRequestBody, field.TypeString)
	}
	if value, ok := truo.mutation.ThreadNum(); ok {
		_spec.SetField(taskresult.FieldThreadNum, field.TypeInt, value)
	}
	if value, ok := truo.mutation.AddedThreadNum(); ok {
		_spec.AddField(taskresult.FieldThreadNum, field.TypeInt, value)
	}
	if value, ok := truo.mutation.AttemptCount(); ok {
		_spec.SetField(taskresult.FieldAttemptCount, field.TypeInt, value)
	}
	if value, ok := truo.mutation.AddedAttemptCount(); ok {
		_spec.AddField(taskresult.FieldAttemptCount, field.TypeInt, value)
	}
	if value, ok := truo.mutation.CreatedAt(); ok {
		_spec.SetField(taskresult.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := truo.mutation.DeletedAt(); ok {
		_spec.SetField(taskresult.FieldDeletedAt, field.TypeTime, value)
	}
	if truo.mutation.DeletedAtCleared() {
		_spec.ClearField(taskresult.FieldDeletedAt, field.TypeTime)
	}
	_node = &TaskResult{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	truo.mutation.done = true
	return _node, nil
}
