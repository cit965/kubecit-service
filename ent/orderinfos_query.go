// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kubecit-service/ent/orderinfos"
	"kubecit-service/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderInfosQuery is the builder for querying OrderInfos entities.
type OrderInfosQuery struct {
	config
	ctx        *QueryContext
	order      []orderinfos.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderInfos
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderInfosQuery builder.
func (oiq *OrderInfosQuery) Where(ps ...predicate.OrderInfos) *OrderInfosQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit the number of records to be returned by this query.
func (oiq *OrderInfosQuery) Limit(limit int) *OrderInfosQuery {
	oiq.ctx.Limit = &limit
	return oiq
}

// Offset to start from.
func (oiq *OrderInfosQuery) Offset(offset int) *OrderInfosQuery {
	oiq.ctx.Offset = &offset
	return oiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oiq *OrderInfosQuery) Unique(unique bool) *OrderInfosQuery {
	oiq.ctx.Unique = &unique
	return oiq
}

// Order specifies how the records should be ordered.
func (oiq *OrderInfosQuery) Order(o ...orderinfos.OrderOption) *OrderInfosQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// First returns the first OrderInfos entity from the query.
// Returns a *NotFoundError when no OrderInfos was found.
func (oiq *OrderInfosQuery) First(ctx context.Context) (*OrderInfos, error) {
	nodes, err := oiq.Limit(1).All(setContextOp(ctx, oiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderinfos.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrderInfosQuery) FirstX(ctx context.Context) *OrderInfos {
	node, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderInfos ID from the query.
// Returns a *NotFoundError when no OrderInfos ID was found.
func (oiq *OrderInfosQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(1).IDs(setContextOp(ctx, oiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderinfos.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oiq *OrderInfosQuery) FirstIDX(ctx context.Context) int {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderInfos entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderInfos entity is found.
// Returns a *NotFoundError when no OrderInfos entities are found.
func (oiq *OrderInfosQuery) Only(ctx context.Context) (*OrderInfos, error) {
	nodes, err := oiq.Limit(2).All(setContextOp(ctx, oiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderinfos.Label}
	default:
		return nil, &NotSingularError{orderinfos.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrderInfosQuery) OnlyX(ctx context.Context) *OrderInfos {
	node, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderInfos ID in the query.
// Returns a *NotSingularError when more than one OrderInfos ID is found.
// Returns a *NotFoundError when no entities are found.
func (oiq *OrderInfosQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(2).IDs(setContextOp(ctx, oiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderinfos.Label}
	default:
		err = &NotSingularError{orderinfos.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oiq *OrderInfosQuery) OnlyIDX(ctx context.Context) int {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderInfosSlice.
func (oiq *OrderInfosQuery) All(ctx context.Context) ([]*OrderInfos, error) {
	ctx = setContextOp(ctx, oiq.ctx, "All")
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderInfos, *OrderInfosQuery]()
	return withInterceptors[[]*OrderInfos](ctx, oiq, qr, oiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrderInfosQuery) AllX(ctx context.Context) []*OrderInfos {
	nodes, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderInfos IDs.
func (oiq *OrderInfosQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oiq.ctx.Unique == nil && oiq.path != nil {
		oiq.Unique(true)
	}
	ctx = setContextOp(ctx, oiq.ctx, "IDs")
	if err = oiq.Select(orderinfos.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrderInfosQuery) IDsX(ctx context.Context) []int {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrderInfosQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oiq.ctx, "Count")
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oiq, querierCount[*OrderInfosQuery](), oiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oiq *OrderInfosQuery) CountX(ctx context.Context) int {
	count, err := oiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oiq *OrderInfosQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oiq.ctx, "Exist")
	switch _, err := oiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrderInfosQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderInfosQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrderInfosQuery) Clone() *OrderInfosQuery {
	if oiq == nil {
		return nil
	}
	return &OrderInfosQuery{
		config:     oiq.config,
		ctx:        oiq.ctx.Clone(),
		order:      append([]orderinfos.OrderOption{}, oiq.order...),
		inters:     append([]Interceptor{}, oiq.inters...),
		predicates: append([]predicate.OrderInfos{}, oiq.predicates...),
		// clone intermediate query.
		sql:  oiq.sql.Clone(),
		path: oiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID int32 `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderInfos.Query().
//		GroupBy(orderinfos.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oiq *OrderInfosQuery) GroupBy(field string, fields ...string) *OrderInfosGroupBy {
	oiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderInfosGroupBy{build: oiq}
	grbuild.flds = &oiq.ctx.Fields
	grbuild.label = orderinfos.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID int32 `json:"order_id,omitempty"`
//	}
//
//	client.OrderInfos.Query().
//		Select(orderinfos.FieldOrderID).
//		Scan(ctx, &v)
func (oiq *OrderInfosQuery) Select(fields ...string) *OrderInfosSelect {
	oiq.ctx.Fields = append(oiq.ctx.Fields, fields...)
	sbuild := &OrderInfosSelect{OrderInfosQuery: oiq}
	sbuild.label = orderinfos.Label
	sbuild.flds, sbuild.scan = &oiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderInfosSelect configured with the given aggregations.
func (oiq *OrderInfosQuery) Aggregate(fns ...AggregateFunc) *OrderInfosSelect {
	return oiq.Select().Aggregate(fns...)
}

func (oiq *OrderInfosQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oiq); err != nil {
				return err
			}
		}
	}
	for _, f := range oiq.ctx.Fields {
		if !orderinfos.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrderInfosQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderInfos, error) {
	var (
		nodes = []*OrderInfos{}
		_spec = oiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderInfos).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderInfos{config: oiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oiq *OrderInfosQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	_spec.Node.Columns = oiq.ctx.Fields
	if len(oiq.ctx.Fields) > 0 {
		_spec.Unique = oiq.ctx.Unique != nil && *oiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrderInfosQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderinfos.Table, orderinfos.Columns, sqlgraph.NewFieldSpec(orderinfos.FieldID, field.TypeInt))
	_spec.From = oiq.sql
	if unique := oiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oiq.path != nil {
		_spec.Unique = true
	}
	if fields := oiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderinfos.FieldID)
		for i := range fields {
			if fields[i] != orderinfos.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oiq *OrderInfosQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orderinfos.Table)
	columns := oiq.ctx.Fields
	if len(columns) == 0 {
		columns = orderinfos.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oiq.ctx.Unique != nil && *oiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderInfosGroupBy is the group-by builder for OrderInfos entities.
type OrderInfosGroupBy struct {
	selector
	build *OrderInfosQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrderInfosGroupBy) Aggregate(fns ...AggregateFunc) *OrderInfosGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the selector query and scans the result into the given value.
func (oigb *OrderInfosGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oigb.build.ctx, "GroupBy")
	if err := oigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderInfosQuery, *OrderInfosGroupBy](ctx, oigb.build, oigb, oigb.build.inters, v)
}

func (oigb *OrderInfosGroupBy) sqlScan(ctx context.Context, root *OrderInfosQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oigb.fns))
	for _, fn := range oigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oigb.flds)+len(oigb.fns))
		for _, f := range *oigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderInfosSelect is the builder for selecting fields of OrderInfos entities.
type OrderInfosSelect struct {
	*OrderInfosQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ois *OrderInfosSelect) Aggregate(fns ...AggregateFunc) *OrderInfosSelect {
	ois.fns = append(ois.fns, fns...)
	return ois
}

// Scan applies the selector query and scans the result into the given value.
func (ois *OrderInfosSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ois.ctx, "Select")
	if err := ois.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderInfosQuery, *OrderInfosSelect](ctx, ois.OrderInfosQuery, ois, ois.inters, v)
}

func (ois *OrderInfosSelect) sqlScan(ctx context.Context, root *OrderInfosQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ois.fns))
	for _, fn := range ois.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ois.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}