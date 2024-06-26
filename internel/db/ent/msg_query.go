// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/msg"
	"IM/internel/db/ent/predicate"
	"IM/internel/db/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MsgQuery is the builder for querying Msg entities.
type MsgQuery struct {
	config
	ctx             *QueryContext
	order           []msg.OrderOption
	inters          []Interceptor
	predicates      []predicate.Msg
	withSendUser    *UserQuery
	withReceiveUser *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MsgQuery builder.
func (mq *MsgQuery) Where(ps ...predicate.Msg) *MsgQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MsgQuery) Limit(limit int) *MsgQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MsgQuery) Offset(offset int) *MsgQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MsgQuery) Unique(unique bool) *MsgQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MsgQuery) Order(o ...msg.OrderOption) *MsgQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QuerySendUser chains the current query on the "send_user" edge.
func (mq *MsgQuery) QuerySendUser() *UserQuery {
	query := (&UserClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(msg.Table, msg.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, msg.SendUserTable, msg.SendUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReceiveUser chains the current query on the "receive_user" edge.
func (mq *MsgQuery) QueryReceiveUser() *UserQuery {
	query := (&UserClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(msg.Table, msg.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, msg.ReceiveUserTable, msg.ReceiveUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Msg entity from the query.
// Returns a *NotFoundError when no Msg was found.
func (mq *MsgQuery) First(ctx context.Context) (*Msg, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{msg.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MsgQuery) FirstX(ctx context.Context) *Msg {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Msg ID from the query.
// Returns a *NotFoundError when no Msg ID was found.
func (mq *MsgQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{msg.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MsgQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Msg entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Msg entity is found.
// Returns a *NotFoundError when no Msg entities are found.
func (mq *MsgQuery) Only(ctx context.Context) (*Msg, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{msg.Label}
	default:
		return nil, &NotSingularError{msg.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MsgQuery) OnlyX(ctx context.Context) *Msg {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Msg ID in the query.
// Returns a *NotSingularError when more than one Msg ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MsgQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{msg.Label}
	default:
		err = &NotSingularError{msg.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MsgQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Msgs.
func (mq *MsgQuery) All(ctx context.Context) ([]*Msg, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Msg, *MsgQuery]()
	return withInterceptors[[]*Msg](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MsgQuery) AllX(ctx context.Context) []*Msg {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Msg IDs.
func (mq *MsgQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(msg.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MsgQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MsgQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MsgQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MsgQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MsgQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MsgQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MsgQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MsgQuery) Clone() *MsgQuery {
	if mq == nil {
		return nil
	}
	return &MsgQuery{
		config:          mq.config,
		ctx:             mq.ctx.Clone(),
		order:           append([]msg.OrderOption{}, mq.order...),
		inters:          append([]Interceptor{}, mq.inters...),
		predicates:      append([]predicate.Msg{}, mq.predicates...),
		withSendUser:    mq.withSendUser.Clone(),
		withReceiveUser: mq.withReceiveUser.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithSendUser tells the query-builder to eager-load the nodes that are connected to
// the "send_user" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MsgQuery) WithSendUser(opts ...func(*UserQuery)) *MsgQuery {
	query := (&UserClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withSendUser = query
	return mq
}

// WithReceiveUser tells the query-builder to eager-load the nodes that are connected to
// the "receive_user" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MsgQuery) WithReceiveUser(opts ...func(*UserQuery)) *MsgQuery {
	query := (&UserClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withReceiveUser = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Msg.Query().
//		GroupBy(msg.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MsgQuery) GroupBy(field string, fields ...string) *MsgGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MsgGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = msg.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at"`
//	}
//
//	client.Msg.Query().
//		Select(msg.FieldCreatedAt).
//		Scan(ctx, &v)
func (mq *MsgQuery) Select(fields ...string) *MsgSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MsgSelect{MsgQuery: mq}
	sbuild.label = msg.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MsgSelect configured with the given aggregations.
func (mq *MsgQuery) Aggregate(fns ...AggregateFunc) *MsgSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MsgQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !msg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MsgQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Msg, error) {
	var (
		nodes       = []*Msg{}
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withSendUser != nil,
			mq.withReceiveUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Msg).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Msg{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withSendUser; query != nil {
		if err := mq.loadSendUser(ctx, query, nodes, nil,
			func(n *Msg, e *User) { n.Edges.SendUser = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withReceiveUser; query != nil {
		if err := mq.loadReceiveUser(ctx, query, nodes, nil,
			func(n *Msg, e *User) { n.Edges.ReceiveUser = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MsgQuery) loadSendUser(ctx context.Context, query *UserQuery, nodes []*Msg, init func(*Msg), assign func(*Msg, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Msg)
	for i := range nodes {
		fk := nodes[i].SendID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "send_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MsgQuery) loadReceiveUser(ctx context.Context, query *UserQuery, nodes []*Msg, init func(*Msg), assign func(*Msg, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Msg)
	for i := range nodes {
		fk := nodes[i].ReceiveID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "receive_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mq *MsgQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MsgQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(msg.Table, msg.Columns, sqlgraph.NewFieldSpec(msg.FieldID, field.TypeInt64))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, msg.FieldID)
		for i := range fields {
			if fields[i] != msg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mq.withSendUser != nil {
			_spec.Node.AddColumnOnce(msg.FieldSendID)
		}
		if mq.withReceiveUser != nil {
			_spec.Node.AddColumnOnce(msg.FieldReceiveID)
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MsgQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(msg.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = msg.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MsgGroupBy is the group-by builder for Msg entities.
type MsgGroupBy struct {
	selector
	build *MsgQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MsgGroupBy) Aggregate(fns ...AggregateFunc) *MsgGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MsgGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MsgQuery, *MsgGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MsgGroupBy) sqlScan(ctx context.Context, root *MsgQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MsgSelect is the builder for selecting fields of Msg entities.
type MsgSelect struct {
	*MsgQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MsgSelect) Aggregate(fns ...AggregateFunc) *MsgSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MsgSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MsgQuery, *MsgSelect](ctx, ms.MsgQuery, ms, ms.inters, v)
}

func (ms *MsgSelect) sqlScan(ctx context.Context, root *MsgQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
