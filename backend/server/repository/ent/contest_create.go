// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
)

// ContestCreate is the builder for creating a Contest entity.
type ContestCreate struct {
	config
	mutation *ContestMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (cc *ContestCreate) SetTitle(s string) *ContestCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetStartAt sets the "start_at" field.
func (cc *ContestCreate) SetStartAt(t time.Time) *ContestCreate {
	cc.mutation.SetStartAt(t)
	return cc
}

// SetEndAt sets the "end_at" field.
func (cc *ContestCreate) SetEndAt(t time.Time) *ContestCreate {
	cc.mutation.SetEndAt(t)
	return cc
}

// SetSubmitLimit sets the "submit_limit" field.
func (cc *ContestCreate) SetSubmitLimit(i int) *ContestCreate {
	cc.mutation.SetSubmitLimit(i)
	return cc
}

// SetSlug sets the "slug" field.
func (cc *ContestCreate) SetSlug(s string) *ContestCreate {
	cc.mutation.SetSlug(s)
	return cc
}

// SetTagSelectionLogic sets the "tag_selection_logic" field.
func (cc *ContestCreate) SetTagSelectionLogic(csl contest.TagSelectionLogic) *ContestCreate {
	cc.mutation.SetTagSelectionLogic(csl)
	return cc
}

// SetValidator sets the "validator" field.
func (cc *ContestCreate) SetValidator(s string) *ContestCreate {
	cc.mutation.SetValidator(s)
	return cc
}

// SetTimeLimitPerTask sets the "time_limit_per_task" field.
func (cc *ContestCreate) SetTimeLimitPerTask(i int64) *ContestCreate {
	cc.mutation.SetTimeLimitPerTask(i)
	return cc
}

// SetNillableTimeLimitPerTask sets the "time_limit_per_task" field if the given value is not nil.
func (cc *ContestCreate) SetNillableTimeLimitPerTask(i *int64) *ContestCreate {
	if i != nil {
		cc.SetTimeLimitPerTask(*i)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContestCreate) SetCreatedAt(t time.Time) *ContestCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContestCreate) SetUpdatedAt(t time.Time) *ContestCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableUpdatedAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ContestCreate) SetID(i int) *ContestCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (cc *ContestCreate) AddSubmitIDs(ids ...int) *ContestCreate {
	cc.mutation.AddSubmitIDs(ids...)
	return cc
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (cc *ContestCreate) AddSubmits(s ...*Submit) *ContestCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSubmitIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cc *ContestCreate) Mutation() *ContestMutation {
	return cc.mutation
}

// Save creates the Contest in the database.
func (cc *ContestCreate) Save(ctx context.Context) (*Contest, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContestCreate) SaveX(ctx context.Context) *Contest {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContestCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContestCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContestCreate) defaults() {
	if _, ok := cc.mutation.TimeLimitPerTask(); !ok {
		v := contest.DefaultTimeLimitPerTask
		cc.mutation.SetTimeLimitPerTask(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContestCreate) check() error {
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Contest.title"`)}
	}
	if _, ok := cc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "Contest.start_at"`)}
	}
	if _, ok := cc.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New(`ent: missing required field "Contest.end_at"`)}
	}
	if _, ok := cc.mutation.SubmitLimit(); !ok {
		return &ValidationError{Name: "submit_limit", err: errors.New(`ent: missing required field "Contest.submit_limit"`)}
	}
	if _, ok := cc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Contest.slug"`)}
	}
	if _, ok := cc.mutation.TagSelectionLogic(); !ok {
		return &ValidationError{Name: "tag_selection_logic", err: errors.New(`ent: missing required field "Contest.tag_selection_logic"`)}
	}
	if v, ok := cc.mutation.TagSelectionLogic(); ok {
		if err := contest.TagSelectionLogicValidator(v); err != nil {
			return &ValidationError{Name: "tag_selection_logic", err: fmt.Errorf(`ent: validator failed for field "Contest.tag_selection_logic": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Validator(); !ok {
		return &ValidationError{Name: "validator", err: errors.New(`ent: missing required field "Contest.validator"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Contest.created_at"`)}
	}
	return nil
}

func (cc *ContestCreate) sqlSave(ctx context.Context) (*Contest, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ContestCreate) createSpec() (*Contest, *sqlgraph.CreateSpec) {
	var (
		_node = &Contest{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(contest.Table, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(contest.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := cc.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := cc.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := cc.mutation.SubmitLimit(); ok {
		_spec.SetField(contest.FieldSubmitLimit, field.TypeInt, value)
		_node.SubmitLimit = value
	}
	if value, ok := cc.mutation.Slug(); ok {
		_spec.SetField(contest.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := cc.mutation.TagSelectionLogic(); ok {
		_spec.SetField(contest.FieldTagSelectionLogic, field.TypeEnum, value)
		_node.TagSelectionLogic = value
	}
	if value, ok := cc.mutation.Validator(); ok {
		_spec.SetField(contest.FieldValidator, field.TypeString, value)
		_node.Validator = value
	}
	if value, ok := cc.mutation.TimeLimitPerTask(); ok {
		_spec.SetField(contest.FieldTimeLimitPerTask, field.TypeInt64, value)
		_node.TimeLimitPerTask = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(contest.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.SubmitsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContestCreateBulk is the builder for creating many Contest entities in bulk.
type ContestCreateBulk struct {
	config
	builders []*ContestCreate
}

// Save creates the Contest entities in the database.
func (ccb *ContestCreateBulk) Save(ctx context.Context) ([]*Contest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contest, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContestCreateBulk) SaveX(ctx context.Context) []*Contest {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContestCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContestCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
