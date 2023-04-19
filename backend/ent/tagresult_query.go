// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/tagresult"
)

// TagResultQuery is the builder for querying TagResult entities.
type TagResultQuery struct {
	config
	ctx        *QueryContext
	order      []tagresult.OrderOption
	inters     []Interceptor
	predicates []predicate.TagResult
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TagResultQuery builder.
func (trq *TagResultQuery) Where(ps ...predicate.TagResult) *TagResultQuery {
	trq.predicates = append(trq.predicates, ps...)
	return trq
}

// Limit the number of records to be returned by this query.
func (trq *TagResultQuery) Limit(limit int) *TagResultQuery {
	trq.ctx.Limit = &limit
	return trq
}

// Offset to start from.
func (trq *TagResultQuery) Offset(offset int) *TagResultQuery {
	trq.ctx.Offset = &offset
	return trq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trq *TagResultQuery) Unique(unique bool) *TagResultQuery {
	trq.ctx.Unique = &unique
	return trq
}

// Order specifies how the records should be ordered.
func (trq *TagResultQuery) Order(o ...tagresult.OrderOption) *TagResultQuery {
	trq.order = append(trq.order, o...)
	return trq
}

// First returns the first TagResult entity from the query.
// Returns a *NotFoundError when no TagResult was found.
func (trq *TagResultQuery) First(ctx context.Context) (*TagResult, error) {
	nodes, err := trq.Limit(1).All(setContextOp(ctx, trq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tagresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trq *TagResultQuery) FirstX(ctx context.Context) *TagResult {
	node, err := trq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TagResult ID from the query.
// Returns a *NotFoundError when no TagResult ID was found.
func (trq *TagResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(1).IDs(setContextOp(ctx, trq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tagresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trq *TagResultQuery) FirstIDX(ctx context.Context) int {
	id, err := trq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TagResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TagResult entity is found.
// Returns a *NotFoundError when no TagResult entities are found.
func (trq *TagResultQuery) Only(ctx context.Context) (*TagResult, error) {
	nodes, err := trq.Limit(2).All(setContextOp(ctx, trq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tagresult.Label}
	default:
		return nil, &NotSingularError{tagresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trq *TagResultQuery) OnlyX(ctx context.Context) *TagResult {
	node, err := trq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TagResult ID in the query.
// Returns a *NotSingularError when more than one TagResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (trq *TagResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(2).IDs(setContextOp(ctx, trq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tagresult.Label}
	default:
		err = &NotSingularError{tagresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trq *TagResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := trq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TagResults.
func (trq *TagResultQuery) All(ctx context.Context) ([]*TagResult, error) {
	ctx = setContextOp(ctx, trq.ctx, "All")
	if err := trq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TagResult, *TagResultQuery]()
	return withInterceptors[[]*TagResult](ctx, trq, qr, trq.inters)
}

// AllX is like All, but panics if an error occurs.
func (trq *TagResultQuery) AllX(ctx context.Context) []*TagResult {
	nodes, err := trq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TagResult IDs.
func (trq *TagResultQuery) IDs(ctx context.Context) (ids []int, err error) {
	if trq.ctx.Unique == nil && trq.path != nil {
		trq.Unique(true)
	}
	ctx = setContextOp(ctx, trq.ctx, "IDs")
	if err = trq.Select(tagresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trq *TagResultQuery) IDsX(ctx context.Context) []int {
	ids, err := trq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trq *TagResultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, trq.ctx, "Count")
	if err := trq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, trq, querierCount[*TagResultQuery](), trq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (trq *TagResultQuery) CountX(ctx context.Context) int {
	count, err := trq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trq *TagResultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, trq.ctx, "Exist")
	switch _, err := trq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (trq *TagResultQuery) ExistX(ctx context.Context) bool {
	exist, err := trq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TagResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trq *TagResultQuery) Clone() *TagResultQuery {
	if trq == nil {
		return nil
	}
	return &TagResultQuery{
		config:     trq.config,
		ctx:        trq.ctx.Clone(),
		order:      append([]tagresult.OrderOption{}, trq.order...),
		inters:     append([]Interceptor{}, trq.inters...),
		predicates: append([]predicate.TagResult{}, trq.predicates...),
		// clone intermediate query.
		sql:  trq.sql.Clone(),
		path: trq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TagResult.Query().
//		GroupBy(tagresult.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (trq *TagResultQuery) GroupBy(field string, fields ...string) *TagResultGroupBy {
	trq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TagResultGroupBy{build: trq}
	grbuild.flds = &trq.ctx.Fields
	grbuild.label = tagresult.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.TagResult.Query().
//		Select(tagresult.FieldName).
//		Scan(ctx, &v)
func (trq *TagResultQuery) Select(fields ...string) *TagResultSelect {
	trq.ctx.Fields = append(trq.ctx.Fields, fields...)
	sbuild := &TagResultSelect{TagResultQuery: trq}
	sbuild.label = tagresult.Label
	sbuild.flds, sbuild.scan = &trq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TagResultSelect configured with the given aggregations.
func (trq *TagResultQuery) Aggregate(fns ...AggregateFunc) *TagResultSelect {
	return trq.Select().Aggregate(fns...)
}

func (trq *TagResultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range trq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, trq); err != nil {
				return err
			}
		}
	}
	for _, f := range trq.ctx.Fields {
		if !tagresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if trq.path != nil {
		prev, err := trq.path(ctx)
		if err != nil {
			return err
		}
		trq.sql = prev
	}
	return nil
}

func (trq *TagResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TagResult, error) {
	var (
		nodes   = []*TagResult{}
		withFKs = trq.withFKs
		_spec   = trq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, tagresult.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TagResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TagResult{config: trq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, trq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (trq *TagResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trq.querySpec()
	_spec.Node.Columns = trq.ctx.Fields
	if len(trq.ctx.Fields) > 0 {
		_spec.Unique = trq.ctx.Unique != nil && *trq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, trq.driver, _spec)
}

func (trq *TagResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tagresult.Table, tagresult.Columns, sqlgraph.NewFieldSpec(tagresult.FieldID, field.TypeInt))
	_spec.From = trq.sql
	if unique := trq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if trq.path != nil {
		_spec.Unique = true
	}
	if fields := trq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tagresult.FieldID)
		for i := range fields {
			if fields[i] != tagresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := trq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := trq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := trq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := trq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (trq *TagResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trq.driver.Dialect())
	t1 := builder.Table(tagresult.Table)
	columns := trq.ctx.Fields
	if len(columns) == 0 {
		columns = tagresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if trq.sql != nil {
		selector = trq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if trq.ctx.Unique != nil && *trq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range trq.predicates {
		p(selector)
	}
	for _, p := range trq.order {
		p(selector)
	}
	if offset := trq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := trq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TagResultGroupBy is the group-by builder for TagResult entities.
type TagResultGroupBy struct {
	selector
	build *TagResultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trgb *TagResultGroupBy) Aggregate(fns ...AggregateFunc) *TagResultGroupBy {
	trgb.fns = append(trgb.fns, fns...)
	return trgb
}

// Scan applies the selector query and scans the result into the given value.
func (trgb *TagResultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trgb.build.ctx, "GroupBy")
	if err := trgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TagResultQuery, *TagResultGroupBy](ctx, trgb.build, trgb, trgb.build.inters, v)
}

func (trgb *TagResultGroupBy) sqlScan(ctx context.Context, root *TagResultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(trgb.fns))
	for _, fn := range trgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*trgb.flds)+len(trgb.fns))
		for _, f := range *trgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*trgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TagResultSelect is the builder for selecting fields of TagResult entities.
type TagResultSelect struct {
	*TagResultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (trs *TagResultSelect) Aggregate(fns ...AggregateFunc) *TagResultSelect {
	trs.fns = append(trs.fns, fns...)
	return trs
}

// Scan applies the selector query and scans the result into the given value.
func (trs *TagResultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trs.ctx, "Select")
	if err := trs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TagResultQuery, *TagResultSelect](ctx, trs.TagResultQuery, trs, trs.inters, v)
}

func (trs *TagResultSelect) sqlScan(ctx context.Context, root *TagResultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(trs.fns))
	for _, fn := range trs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*trs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}