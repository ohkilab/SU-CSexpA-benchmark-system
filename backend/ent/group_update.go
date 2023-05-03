// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetYear sets the "year" field.
func (gu *GroupUpdate) SetYear(i int) *GroupUpdate {
	gu.mutation.ResetYear()
	gu.mutation.SetYear(i)
	return gu
}

// AddYear adds i to the "year" field.
func (gu *GroupUpdate) AddYear(i int) *GroupUpdate {
	gu.mutation.AddYear(i)
	return gu
}

// SetScore sets the "score" field.
func (gu *GroupUpdate) SetScore(i int) *GroupUpdate {
	gu.mutation.ResetScore()
	gu.mutation.SetScore(i)
	return gu
}

// AddScore adds i to the "score" field.
func (gu *GroupUpdate) AddScore(i int) *GroupUpdate {
	gu.mutation.AddScore(i)
	return gu
}

// SetRole sets the "role" field.
func (gu *GroupUpdate) SetRole(gr group.Role) *GroupUpdate {
	gu.mutation.SetRole(gr)
	return gu
}

// SetEncryptedPassword sets the "encrypted_password" field.
func (gu *GroupUpdate) SetEncryptedPassword(s string) *GroupUpdate {
	gu.mutation.SetEncryptedPassword(s)
	return gu
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (gu *GroupUpdate) AddSubmitIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddSubmitIDs(ids...)
	return gu
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (gu *GroupUpdate) AddSubmits(s ...*Submit) *GroupUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gu.AddSubmitIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearSubmits clears all "submits" edges to the Submit entity.
func (gu *GroupUpdate) ClearSubmits() *GroupUpdate {
	gu.mutation.ClearSubmits()
	return gu
}

// RemoveSubmitIDs removes the "submits" edge to Submit entities by IDs.
func (gu *GroupUpdate) RemoveSubmitIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveSubmitIDs(ids...)
	return gu
}

// RemoveSubmits removes "submits" edges to Submit entities.
func (gu *GroupUpdate) RemoveSubmits(s ...*Submit) *GroupUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gu.RemoveSubmitIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GroupMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.Year(); ok {
		if err := group.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Group.year": %w`, err)}
		}
	}
	if v, ok := gu.mutation.Score(); ok {
		if err := group.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Group.score": %w`, err)}
		}
	}
	if v, ok := gu.mutation.Role(); ok {
		if err := group.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Group.role": %w`, err)}
		}
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.Year(); ok {
		_spec.SetField(group.FieldYear, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedYear(); ok {
		_spec.AddField(group.FieldYear, field.TypeInt, value)
	}
	if value, ok := gu.mutation.Score(); ok {
		_spec.SetField(group.FieldScore, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedScore(); ok {
		_spec.AddField(group.FieldScore, field.TypeInt, value)
	}
	if value, ok := gu.mutation.Role(); ok {
		_spec.SetField(group.FieldRole, field.TypeEnum, value)
	}
	if value, ok := gu.mutation.EncryptedPassword(); ok {
		_spec.SetField(group.FieldEncryptedPassword, field.TypeString, value)
	}
	if gu.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.SubmitsTable,
			Columns: group.SubmitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedSubmitsIDs(); len(nodes) > 0 && !gu.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.SubmitsTable,
			Columns: group.SubmitsPrimaryKey,
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
	if nodes := gu.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.SubmitsTable,
			Columns: group.SubmitsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetYear sets the "year" field.
func (guo *GroupUpdateOne) SetYear(i int) *GroupUpdateOne {
	guo.mutation.ResetYear()
	guo.mutation.SetYear(i)
	return guo
}

// AddYear adds i to the "year" field.
func (guo *GroupUpdateOne) AddYear(i int) *GroupUpdateOne {
	guo.mutation.AddYear(i)
	return guo
}

// SetScore sets the "score" field.
func (guo *GroupUpdateOne) SetScore(i int) *GroupUpdateOne {
	guo.mutation.ResetScore()
	guo.mutation.SetScore(i)
	return guo
}

// AddScore adds i to the "score" field.
func (guo *GroupUpdateOne) AddScore(i int) *GroupUpdateOne {
	guo.mutation.AddScore(i)
	return guo
}

// SetRole sets the "role" field.
func (guo *GroupUpdateOne) SetRole(gr group.Role) *GroupUpdateOne {
	guo.mutation.SetRole(gr)
	return guo
}

// SetEncryptedPassword sets the "encrypted_password" field.
func (guo *GroupUpdateOne) SetEncryptedPassword(s string) *GroupUpdateOne {
	guo.mutation.SetEncryptedPassword(s)
	return guo
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (guo *GroupUpdateOne) AddSubmitIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddSubmitIDs(ids...)
	return guo
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (guo *GroupUpdateOne) AddSubmits(s ...*Submit) *GroupUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return guo.AddSubmitIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearSubmits clears all "submits" edges to the Submit entity.
func (guo *GroupUpdateOne) ClearSubmits() *GroupUpdateOne {
	guo.mutation.ClearSubmits()
	return guo
}

// RemoveSubmitIDs removes the "submits" edge to Submit entities by IDs.
func (guo *GroupUpdateOne) RemoveSubmitIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveSubmitIDs(ids...)
	return guo
}

// RemoveSubmits removes "submits" edges to Submit entities.
func (guo *GroupUpdateOne) RemoveSubmits(s ...*Submit) *GroupUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return guo.RemoveSubmitIDs(ids...)
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return withHooks[*Group, GroupMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.Year(); ok {
		if err := group.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Group.year": %w`, err)}
		}
	}
	if v, ok := guo.mutation.Score(); ok {
		if err := group.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Group.score": %w`, err)}
		}
	}
	if v, ok := guo.mutation.Role(); ok {
		if err := group.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Group.role": %w`, err)}
		}
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.Year(); ok {
		_spec.SetField(group.FieldYear, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedYear(); ok {
		_spec.AddField(group.FieldYear, field.TypeInt, value)
	}
	if value, ok := guo.mutation.Score(); ok {
		_spec.SetField(group.FieldScore, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedScore(); ok {
		_spec.AddField(group.FieldScore, field.TypeInt, value)
	}
	if value, ok := guo.mutation.Role(); ok {
		_spec.SetField(group.FieldRole, field.TypeEnum, value)
	}
	if value, ok := guo.mutation.EncryptedPassword(); ok {
		_spec.SetField(group.FieldEncryptedPassword, field.TypeString, value)
	}
	if guo.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.SubmitsTable,
			Columns: group.SubmitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedSubmitsIDs(); len(nodes) > 0 && !guo.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.SubmitsTable,
			Columns: group.SubmitsPrimaryKey,
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
	if nodes := guo.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.SubmitsTable,
			Columns: group.SubmitsPrimaryKey,
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
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
