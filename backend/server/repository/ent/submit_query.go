// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/taskresult"
)

// SubmitQuery is the builder for querying Submit entities.
type SubmitQuery struct {
	config
	ctx             *QueryContext
	order           []submit.OrderOption
	inters          []Interceptor
	predicates      []predicate.Submit
	withTaskResults *TaskResultQuery
	withGroups      *GroupQuery
	withContests    *ContestQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubmitQuery builder.
func (sq *SubmitQuery) Where(ps ...predicate.Submit) *SubmitQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SubmitQuery) Limit(limit int) *SubmitQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SubmitQuery) Offset(offset int) *SubmitQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubmitQuery) Unique(unique bool) *SubmitQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SubmitQuery) Order(o ...submit.OrderOption) *SubmitQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryTaskResults chains the current query on the "taskResults" edge.
func (sq *SubmitQuery) QueryTaskResults() *TaskResultQuery {
	query := (&TaskResultClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, selector),
			sqlgraph.To(taskresult.Table, taskresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, submit.TaskResultsTable, submit.TaskResultsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (sq *SubmitQuery) QueryGroups() *GroupQuery {
	query := (&GroupClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, submit.GroupsTable, submit.GroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContests chains the current query on the "contests" edge.
func (sq *SubmitQuery) QueryContests() *ContestQuery {
	query := (&ContestClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, selector),
			sqlgraph.To(contest.Table, contest.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, submit.ContestsTable, submit.ContestsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Submit entity from the query.
// Returns a *NotFoundError when no Submit was found.
func (sq *SubmitQuery) First(ctx context.Context) (*Submit, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{submit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubmitQuery) FirstX(ctx context.Context) *Submit {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Submit ID from the query.
// Returns a *NotFoundError when no Submit ID was found.
func (sq *SubmitQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{submit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubmitQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Submit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Submit entity is found.
// Returns a *NotFoundError when no Submit entities are found.
func (sq *SubmitQuery) Only(ctx context.Context) (*Submit, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{submit.Label}
	default:
		return nil, &NotSingularError{submit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubmitQuery) OnlyX(ctx context.Context) *Submit {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Submit ID in the query.
// Returns a *NotSingularError when more than one Submit ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubmitQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{submit.Label}
	default:
		err = &NotSingularError{submit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubmitQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Submits.
func (sq *SubmitQuery) All(ctx context.Context) ([]*Submit, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Submit, *SubmitQuery]()
	return withInterceptors[[]*Submit](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubmitQuery) AllX(ctx context.Context) []*Submit {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Submit IDs.
func (sq *SubmitQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(submit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubmitQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubmitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SubmitQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubmitQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubmitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubmitQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubmitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubmitQuery) Clone() *SubmitQuery {
	if sq == nil {
		return nil
	}
	return &SubmitQuery{
		config:          sq.config,
		ctx:             sq.ctx.Clone(),
		order:           append([]submit.OrderOption{}, sq.order...),
		inters:          append([]Interceptor{}, sq.inters...),
		predicates:      append([]predicate.Submit{}, sq.predicates...),
		withTaskResults: sq.withTaskResults.Clone(),
		withGroups:      sq.withGroups.Clone(),
		withContests:    sq.withContests.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithTaskResults tells the query-builder to eager-load the nodes that are connected to
// the "taskResults" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubmitQuery) WithTaskResults(opts ...func(*TaskResultQuery)) *SubmitQuery {
	query := (&TaskResultClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withTaskResults = query
	return sq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubmitQuery) WithGroups(opts ...func(*GroupQuery)) *SubmitQuery {
	query := (&GroupClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withGroups = query
	return sq
}

// WithContests tells the query-builder to eager-load the nodes that are connected to
// the "contests" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubmitQuery) WithContests(opts ...func(*ContestQuery)) *SubmitQuery {
	query := (&ContestClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withContests = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Submit.Query().
//		GroupBy(submit.FieldURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubmitQuery) GroupBy(field string, fields ...string) *SubmitGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubmitGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = submit.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//	}
//
//	client.Submit.Query().
//		Select(submit.FieldURL).
//		Scan(ctx, &v)
func (sq *SubmitQuery) Select(fields ...string) *SubmitSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SubmitSelect{SubmitQuery: sq}
	sbuild.label = submit.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubmitSelect configured with the given aggregations.
func (sq *SubmitQuery) Aggregate(fns ...AggregateFunc) *SubmitSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SubmitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !submit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubmitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Submit, error) {
	var (
		nodes       = []*Submit{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [3]bool{
			sq.withTaskResults != nil,
			sq.withGroups != nil,
			sq.withContests != nil,
		}
	)
	if sq.withGroups != nil || sq.withContests != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, submit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Submit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Submit{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withTaskResults; query != nil {
		if err := sq.loadTaskResults(ctx, query, nodes,
			func(n *Submit) { n.Edges.TaskResults = []*TaskResult{} },
			func(n *Submit, e *TaskResult) { n.Edges.TaskResults = append(n.Edges.TaskResults, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withGroups; query != nil {
		if err := sq.loadGroups(ctx, query, nodes, nil,
			func(n *Submit, e *Group) { n.Edges.Groups = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withContests; query != nil {
		if err := sq.loadContests(ctx, query, nodes, nil,
			func(n *Submit, e *Contest) { n.Edges.Contests = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubmitQuery) loadTaskResults(ctx context.Context, query *TaskResultQuery, nodes []*Submit, init func(*Submit), assign func(*Submit, *TaskResult)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Submit)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TaskResult(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(submit.TaskResultsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.submit_task_results
		if fk == nil {
			return fmt.Errorf(`foreign-key "submit_task_results" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "submit_task_results" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SubmitQuery) loadGroups(ctx context.Context, query *GroupQuery, nodes []*Submit, init func(*Submit), assign func(*Submit, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Submit)
	for i := range nodes {
		if nodes[i].group_submits == nil {
			continue
		}
		fk := *nodes[i].group_submits
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_submits" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SubmitQuery) loadContests(ctx context.Context, query *ContestQuery, nodes []*Submit, init func(*Submit), assign func(*Submit, *Contest)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Submit)
	for i := range nodes {
		if nodes[i].contest_submits == nil {
			continue
		}
		fk := *nodes[i].contest_submits
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(contest.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "contest_submits" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SubmitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubmitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submit.FieldID)
		for i := range fields {
			if fields[i] != submit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubmitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(submit.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = submit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubmitGroupBy is the group-by builder for Submit entities.
type SubmitGroupBy struct {
	selector
	build *SubmitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubmitGroupBy) Aggregate(fns ...AggregateFunc) *SubmitGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SubmitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubmitQuery, *SubmitGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SubmitGroupBy) sqlScan(ctx context.Context, root *SubmitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubmitSelect is the builder for selecting fields of Submit entities.
type SubmitSelect struct {
	*SubmitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SubmitSelect) Aggregate(fns ...AggregateFunc) *SubmitSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubmitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubmitQuery, *SubmitSelect](ctx, ss.SubmitQuery, ss, ss.inters, v)
}

func (ss *SubmitSelect) sqlScan(ctx context.Context, root *SubmitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
