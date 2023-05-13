// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/taskresult"
)

// SubmitCreate is the builder for creating a Submit entity.
type SubmitCreate struct {
	config
	mutation *SubmitMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (sc *SubmitCreate) SetURL(s string) *SubmitCreate {
	sc.mutation.SetURL(s)
	return sc
}

// SetYear sets the "year" field.
func (sc *SubmitCreate) SetYear(i int) *SubmitCreate {
	sc.mutation.SetYear(i)
	return sc
}

// SetScore sets the "score" field.
func (sc *SubmitCreate) SetScore(i int) *SubmitCreate {
	sc.mutation.SetScore(i)
	return sc
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (sc *SubmitCreate) SetNillableScore(i *int) *SubmitCreate {
	if i != nil {
		sc.SetScore(*i)
	}
	return sc
}

// SetLanguage sets the "language" field.
func (sc *SubmitCreate) SetLanguage(s submit.Language) *SubmitCreate {
	sc.mutation.SetLanguage(s)
	return sc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (sc *SubmitCreate) SetNillableLanguage(s *submit.Language) *SubmitCreate {
	if s != nil {
		sc.SetLanguage(*s)
	}
	return sc
}

// SetMessage sets the "message" field.
func (sc *SubmitCreate) SetMessage(s string) *SubmitCreate {
	sc.mutation.SetMessage(s)
	return sc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (sc *SubmitCreate) SetNillableMessage(s *string) *SubmitCreate {
	if s != nil {
		sc.SetMessage(*s)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *SubmitCreate) SetStatus(s string) *SubmitCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetSubmitedAt sets the "submited_at" field.
func (sc *SubmitCreate) SetSubmitedAt(t time.Time) *SubmitCreate {
	sc.mutation.SetSubmitedAt(t)
	return sc
}

// SetCompletedAt sets the "completed_at" field.
func (sc *SubmitCreate) SetCompletedAt(t time.Time) *SubmitCreate {
	sc.mutation.SetCompletedAt(t)
	return sc
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (sc *SubmitCreate) SetNillableCompletedAt(t *time.Time) *SubmitCreate {
	if t != nil {
		sc.SetCompletedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubmitCreate) SetUpdatedAt(t time.Time) *SubmitCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubmitCreate) SetNillableUpdatedAt(t *time.Time) *SubmitCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SubmitCreate) SetID(i int) *SubmitCreate {
	sc.mutation.SetID(i)
	return sc
}

// AddTaskResultIDs adds the "taskResults" edge to the TaskResult entity by IDs.
func (sc *SubmitCreate) AddTaskResultIDs(ids ...int) *SubmitCreate {
	sc.mutation.AddTaskResultIDs(ids...)
	return sc
}

// AddTaskResults adds the "taskResults" edges to the TaskResult entity.
func (sc *SubmitCreate) AddTaskResults(t ...*TaskResult) *SubmitCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTaskResultIDs(ids...)
}

// SetGroupsID sets the "groups" edge to the Group entity by ID.
func (sc *SubmitCreate) SetGroupsID(id int) *SubmitCreate {
	sc.mutation.SetGroupsID(id)
	return sc
}

// SetNillableGroupsID sets the "groups" edge to the Group entity by ID if the given value is not nil.
func (sc *SubmitCreate) SetNillableGroupsID(id *int) *SubmitCreate {
	if id != nil {
		sc = sc.SetGroupsID(*id)
	}
	return sc
}

// SetGroups sets the "groups" edge to the Group entity.
func (sc *SubmitCreate) SetGroups(g *Group) *SubmitCreate {
	return sc.SetGroupsID(g.ID)
}

// SetContestsID sets the "contests" edge to the Contest entity by ID.
func (sc *SubmitCreate) SetContestsID(id int) *SubmitCreate {
	sc.mutation.SetContestsID(id)
	return sc
}

// SetNillableContestsID sets the "contests" edge to the Contest entity by ID if the given value is not nil.
func (sc *SubmitCreate) SetNillableContestsID(id *int) *SubmitCreate {
	if id != nil {
		sc = sc.SetContestsID(*id)
	}
	return sc
}

// SetContests sets the "contests" edge to the Contest entity.
func (sc *SubmitCreate) SetContests(c *Contest) *SubmitCreate {
	return sc.SetContestsID(c.ID)
}

// Mutation returns the SubmitMutation object of the builder.
func (sc *SubmitCreate) Mutation() *SubmitMutation {
	return sc.mutation
}

// Save creates the Submit in the database.
func (sc *SubmitCreate) Save(ctx context.Context) (*Submit, error) {
	return withHooks[*Submit, SubmitMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubmitCreate) SaveX(ctx context.Context) *Submit {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubmitCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubmitCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubmitCreate) check() error {
	if _, ok := sc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Submit.url"`)}
	}
	if _, ok := sc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Submit.year"`)}
	}
	if v, ok := sc.mutation.Year(); ok {
		if err := submit.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Submit.year": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Language(); ok {
		if err := submit.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Submit.language": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Submit.status"`)}
	}
	if _, ok := sc.mutation.SubmitedAt(); !ok {
		return &ValidationError{Name: "submited_at", err: errors.New(`ent: missing required field "Submit.submited_at"`)}
	}
	return nil
}

func (sc *SubmitCreate) sqlSave(ctx context.Context) (*Submit, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubmitCreate) createSpec() (*Submit, *sqlgraph.CreateSpec) {
	var (
		_node = &Submit{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(submit.Table, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.URL(); ok {
		_spec.SetField(submit.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := sc.mutation.Year(); ok {
		_spec.SetField(submit.FieldYear, field.TypeInt, value)
		_node.Year = value
	}
	if value, ok := sc.mutation.Score(); ok {
		_spec.SetField(submit.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if value, ok := sc.mutation.Language(); ok {
		_spec.SetField(submit.FieldLanguage, field.TypeEnum, value)
		_node.Language = value
	}
	if value, ok := sc.mutation.Message(); ok {
		_spec.SetField(submit.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(submit.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.SubmitedAt(); ok {
		_spec.SetField(submit.FieldSubmitedAt, field.TypeTime, value)
		_node.SubmitedAt = value
	}
	if value, ok := sc.mutation.CompletedAt(); ok {
		_spec.SetField(submit.FieldCompletedAt, field.TypeTime, value)
		_node.CompletedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(submit.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := sc.mutation.TaskResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TaskResultsTable,
			Columns: []string{submit.TaskResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.GroupsTable,
			Columns: []string{submit.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_submits = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ContestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.ContestsTable,
			Columns: []string{submit.ContestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.contest_submits = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubmitCreateBulk is the builder for creating many Submit entities in bulk.
type SubmitCreateBulk struct {
	config
	builders []*SubmitCreate
}

// Save creates the Submit entities in the database.
func (scb *SubmitCreateBulk) Save(ctx context.Context) ([]*Submit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Submit, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubmitMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubmitCreateBulk) SaveX(ctx context.Context) []*Submit {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubmitCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubmitCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
