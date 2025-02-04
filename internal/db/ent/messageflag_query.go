// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/db/ent/message"
	"github.com/ProtonMail/gluon/internal/db/ent/messageflag"
	"github.com/ProtonMail/gluon/internal/db/ent/predicate"
)

// MessageFlagQuery is the builder for querying MessageFlag entities.
type MessageFlagQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.MessageFlag
	withMessages *MessageQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageFlagQuery builder.
func (mfq *MessageFlagQuery) Where(ps ...predicate.MessageFlag) *MessageFlagQuery {
	mfq.predicates = append(mfq.predicates, ps...)
	return mfq
}

// Limit adds a limit step to the query.
func (mfq *MessageFlagQuery) Limit(limit int) *MessageFlagQuery {
	mfq.limit = &limit
	return mfq
}

// Offset adds an offset step to the query.
func (mfq *MessageFlagQuery) Offset(offset int) *MessageFlagQuery {
	mfq.offset = &offset
	return mfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mfq *MessageFlagQuery) Unique(unique bool) *MessageFlagQuery {
	mfq.unique = &unique
	return mfq
}

// Order adds an order step to the query.
func (mfq *MessageFlagQuery) Order(o ...OrderFunc) *MessageFlagQuery {
	mfq.order = append(mfq.order, o...)
	return mfq
}

// QueryMessages chains the current query on the "messages" edge.
func (mfq *MessageFlagQuery) QueryMessages() *MessageQuery {
	query := &MessageQuery{config: mfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(messageflag.Table, messageflag.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, messageflag.MessagesTable, messageflag.MessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MessageFlag entity from the query.
// Returns a *NotFoundError when no MessageFlag was found.
func (mfq *MessageFlagQuery) First(ctx context.Context) (*MessageFlag, error) {
	nodes, err := mfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messageflag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mfq *MessageFlagQuery) FirstX(ctx context.Context) *MessageFlag {
	node, err := mfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageFlag ID from the query.
// Returns a *NotFoundError when no MessageFlag ID was found.
func (mfq *MessageFlagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messageflag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mfq *MessageFlagQuery) FirstIDX(ctx context.Context) int {
	id, err := mfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageFlag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageFlag entity is found.
// Returns a *NotFoundError when no MessageFlag entities are found.
func (mfq *MessageFlagQuery) Only(ctx context.Context) (*MessageFlag, error) {
	nodes, err := mfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messageflag.Label}
	default:
		return nil, &NotSingularError{messageflag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mfq *MessageFlagQuery) OnlyX(ctx context.Context) *MessageFlag {
	node, err := mfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageFlag ID in the query.
// Returns a *NotSingularError when more than one MessageFlag ID is found.
// Returns a *NotFoundError when no entities are found.
func (mfq *MessageFlagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messageflag.Label}
	default:
		err = &NotSingularError{messageflag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mfq *MessageFlagQuery) OnlyIDX(ctx context.Context) int {
	id, err := mfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageFlags.
func (mfq *MessageFlagQuery) All(ctx context.Context) ([]*MessageFlag, error) {
	if err := mfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mfq *MessageFlagQuery) AllX(ctx context.Context) []*MessageFlag {
	nodes, err := mfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageFlag IDs.
func (mfq *MessageFlagQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mfq.Select(messageflag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mfq *MessageFlagQuery) IDsX(ctx context.Context) []int {
	ids, err := mfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mfq *MessageFlagQuery) Count(ctx context.Context) (int, error) {
	if err := mfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mfq *MessageFlagQuery) CountX(ctx context.Context) int {
	count, err := mfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mfq *MessageFlagQuery) Exist(ctx context.Context) (bool, error) {
	if err := mfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mfq *MessageFlagQuery) ExistX(ctx context.Context) bool {
	exist, err := mfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageFlagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mfq *MessageFlagQuery) Clone() *MessageFlagQuery {
	if mfq == nil {
		return nil
	}
	return &MessageFlagQuery{
		config:       mfq.config,
		limit:        mfq.limit,
		offset:       mfq.offset,
		order:        append([]OrderFunc{}, mfq.order...),
		predicates:   append([]predicate.MessageFlag{}, mfq.predicates...),
		withMessages: mfq.withMessages.Clone(),
		// clone intermediate query.
		sql:    mfq.sql.Clone(),
		path:   mfq.path,
		unique: mfq.unique,
	}
}

// WithMessages tells the query-builder to eager-load the nodes that are connected to
// the "messages" edge. The optional arguments are used to configure the query builder of the edge.
func (mfq *MessageFlagQuery) WithMessages(opts ...func(*MessageQuery)) *MessageFlagQuery {
	query := &MessageQuery{config: mfq.config}
	for _, opt := range opts {
		opt(query)
	}
	mfq.withMessages = query
	return mfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"Value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageFlag.Query().
//		GroupBy(messageflag.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mfq *MessageFlagQuery) GroupBy(field string, fields ...string) *MessageFlagGroupBy {
	grbuild := &MessageFlagGroupBy{config: mfq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mfq.sqlQuery(ctx), nil
	}
	grbuild.label = messageflag.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value string `json:"Value,omitempty"`
//	}
//
//	client.MessageFlag.Query().
//		Select(messageflag.FieldValue).
//		Scan(ctx, &v)
func (mfq *MessageFlagQuery) Select(fields ...string) *MessageFlagSelect {
	mfq.fields = append(mfq.fields, fields...)
	selbuild := &MessageFlagSelect{MessageFlagQuery: mfq}
	selbuild.label = messageflag.Label
	selbuild.flds, selbuild.scan = &mfq.fields, selbuild.Scan
	return selbuild
}

func (mfq *MessageFlagQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mfq.fields {
		if !messageflag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mfq.path != nil {
		prev, err := mfq.path(ctx)
		if err != nil {
			return err
		}
		mfq.sql = prev
	}
	return nil
}

func (mfq *MessageFlagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MessageFlag, error) {
	var (
		nodes       = []*MessageFlag{}
		withFKs     = mfq.withFKs
		_spec       = mfq.querySpec()
		loadedTypes = [1]bool{
			mfq.withMessages != nil,
		}
	)
	if mfq.withMessages != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, messageflag.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*MessageFlag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &MessageFlag{config: mfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mfq.withMessages; query != nil {
		if err := mfq.loadMessages(ctx, query, nodes, nil,
			func(n *MessageFlag, e *Message) { n.Edges.Messages = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mfq *MessageFlagQuery) loadMessages(ctx context.Context, query *MessageQuery, nodes []*MessageFlag, init func(*MessageFlag), assign func(*MessageFlag, *Message)) error {
	ids := make([]imap.InternalMessageID, 0, len(nodes))
	nodeids := make(map[imap.InternalMessageID][]*MessageFlag)
	for i := range nodes {
		if nodes[i].message_flags == nil {
			continue
		}
		fk := *nodes[i].message_flags
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(message.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "message_flags" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mfq *MessageFlagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mfq.querySpec()
	_spec.Node.Columns = mfq.fields
	if len(mfq.fields) > 0 {
		_spec.Unique = mfq.unique != nil && *mfq.unique
	}
	return sqlgraph.CountNodes(ctx, mfq.driver, _spec)
}

func (mfq *MessageFlagQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mfq *MessageFlagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageflag.Table,
			Columns: messageflag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageflag.FieldID,
			},
		},
		From:   mfq.sql,
		Unique: true,
	}
	if unique := mfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageflag.FieldID)
		for i := range fields {
			if fields[i] != messageflag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mfq *MessageFlagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mfq.driver.Dialect())
	t1 := builder.Table(messageflag.Table)
	columns := mfq.fields
	if len(columns) == 0 {
		columns = messageflag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mfq.sql != nil {
		selector = mfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mfq.unique != nil && *mfq.unique {
		selector.Distinct()
	}
	for _, p := range mfq.predicates {
		p(selector)
	}
	for _, p := range mfq.order {
		p(selector)
	}
	if offset := mfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageFlagGroupBy is the group-by builder for MessageFlag entities.
type MessageFlagGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mfgb *MessageFlagGroupBy) Aggregate(fns ...AggregateFunc) *MessageFlagGroupBy {
	mfgb.fns = append(mfgb.fns, fns...)
	return mfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mfgb *MessageFlagGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mfgb.path(ctx)
	if err != nil {
		return err
	}
	mfgb.sql = query
	return mfgb.sqlScan(ctx, v)
}

func (mfgb *MessageFlagGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mfgb.fields {
		if !messageflag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mfgb *MessageFlagGroupBy) sqlQuery() *sql.Selector {
	selector := mfgb.sql.Select()
	aggregation := make([]string, 0, len(mfgb.fns))
	for _, fn := range mfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mfgb.fields)+len(mfgb.fns))
		for _, f := range mfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mfgb.fields...)...)
}

// MessageFlagSelect is the builder for selecting fields of MessageFlag entities.
type MessageFlagSelect struct {
	*MessageFlagQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mfs *MessageFlagSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mfs.prepareQuery(ctx); err != nil {
		return err
	}
	mfs.sql = mfs.MessageFlagQuery.sqlQuery(ctx)
	return mfs.sqlScan(ctx, v)
}

func (mfs *MessageFlagSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mfs.sql.Query()
	if err := mfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
