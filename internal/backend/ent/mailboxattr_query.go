// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/internal/backend/ent/mailboxattr"
	"github.com/ProtonMail/gluon/internal/backend/ent/predicate"
)

// MailboxAttrQuery is the builder for querying MailboxAttr entities.
type MailboxAttrQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MailboxAttr
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MailboxAttrQuery builder.
func (maq *MailboxAttrQuery) Where(ps ...predicate.MailboxAttr) *MailboxAttrQuery {
	maq.predicates = append(maq.predicates, ps...)
	return maq
}

// Limit adds a limit step to the query.
func (maq *MailboxAttrQuery) Limit(limit int) *MailboxAttrQuery {
	maq.limit = &limit
	return maq
}

// Offset adds an offset step to the query.
func (maq *MailboxAttrQuery) Offset(offset int) *MailboxAttrQuery {
	maq.offset = &offset
	return maq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (maq *MailboxAttrQuery) Unique(unique bool) *MailboxAttrQuery {
	maq.unique = &unique
	return maq
}

// Order adds an order step to the query.
func (maq *MailboxAttrQuery) Order(o ...OrderFunc) *MailboxAttrQuery {
	maq.order = append(maq.order, o...)
	return maq
}

// First returns the first MailboxAttr entity from the query.
// Returns a *NotFoundError when no MailboxAttr was found.
func (maq *MailboxAttrQuery) First(ctx context.Context) (*MailboxAttr, error) {
	nodes, err := maq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mailboxattr.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (maq *MailboxAttrQuery) FirstX(ctx context.Context) *MailboxAttr {
	node, err := maq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MailboxAttr ID from the query.
// Returns a *NotFoundError when no MailboxAttr ID was found.
func (maq *MailboxAttrQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = maq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mailboxattr.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (maq *MailboxAttrQuery) FirstIDX(ctx context.Context) int {
	id, err := maq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MailboxAttr entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MailboxAttr entity is found.
// Returns a *NotFoundError when no MailboxAttr entities are found.
func (maq *MailboxAttrQuery) Only(ctx context.Context) (*MailboxAttr, error) {
	nodes, err := maq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mailboxattr.Label}
	default:
		return nil, &NotSingularError{mailboxattr.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (maq *MailboxAttrQuery) OnlyX(ctx context.Context) *MailboxAttr {
	node, err := maq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MailboxAttr ID in the query.
// Returns a *NotSingularError when more than one MailboxAttr ID is found.
// Returns a *NotFoundError when no entities are found.
func (maq *MailboxAttrQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = maq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = &NotSingularError{mailboxattr.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (maq *MailboxAttrQuery) OnlyIDX(ctx context.Context) int {
	id, err := maq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MailboxAttrs.
func (maq *MailboxAttrQuery) All(ctx context.Context) ([]*MailboxAttr, error) {
	if err := maq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return maq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (maq *MailboxAttrQuery) AllX(ctx context.Context) []*MailboxAttr {
	nodes, err := maq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MailboxAttr IDs.
func (maq *MailboxAttrQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := maq.Select(mailboxattr.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (maq *MailboxAttrQuery) IDsX(ctx context.Context) []int {
	ids, err := maq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (maq *MailboxAttrQuery) Count(ctx context.Context) (int, error) {
	if err := maq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return maq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (maq *MailboxAttrQuery) CountX(ctx context.Context) int {
	count, err := maq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (maq *MailboxAttrQuery) Exist(ctx context.Context) (bool, error) {
	if err := maq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return maq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (maq *MailboxAttrQuery) ExistX(ctx context.Context) bool {
	exist, err := maq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MailboxAttrQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (maq *MailboxAttrQuery) Clone() *MailboxAttrQuery {
	if maq == nil {
		return nil
	}
	return &MailboxAttrQuery{
		config:     maq.config,
		limit:      maq.limit,
		offset:     maq.offset,
		order:      append([]OrderFunc{}, maq.order...),
		predicates: append([]predicate.MailboxAttr{}, maq.predicates...),
		// clone intermediate query.
		sql:    maq.sql.Clone(),
		path:   maq.path,
		unique: maq.unique,
	}
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
//	client.MailboxAttr.Query().
//		GroupBy(mailboxattr.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (maq *MailboxAttrQuery) GroupBy(field string, fields ...string) *MailboxAttrGroupBy {
	group := &MailboxAttrGroupBy{config: maq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := maq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return maq.sqlQuery(ctx), nil
	}
	return group
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
//	client.MailboxAttr.Query().
//		Select(mailboxattr.FieldValue).
//		Scan(ctx, &v)
//
func (maq *MailboxAttrQuery) Select(fields ...string) *MailboxAttrSelect {
	maq.fields = append(maq.fields, fields...)
	return &MailboxAttrSelect{MailboxAttrQuery: maq}
}

func (maq *MailboxAttrQuery) prepareQuery(ctx context.Context) error {
	for _, f := range maq.fields {
		if !mailboxattr.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if maq.path != nil {
		prev, err := maq.path(ctx)
		if err != nil {
			return err
		}
		maq.sql = prev
	}
	return nil
}

func (maq *MailboxAttrQuery) sqlAll(ctx context.Context) ([]*MailboxAttr, error) {
	var (
		nodes   = []*MailboxAttr{}
		withFKs = maq.withFKs
		_spec   = maq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mailboxattr.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MailboxAttr{config: maq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, maq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (maq *MailboxAttrQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := maq.querySpec()
	_spec.Node.Columns = maq.fields
	if len(maq.fields) > 0 {
		_spec.Unique = maq.unique != nil && *maq.unique
	}
	return sqlgraph.CountNodes(ctx, maq.driver, _spec)
}

func (maq *MailboxAttrQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := maq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (maq *MailboxAttrQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mailboxattr.Table,
			Columns: mailboxattr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mailboxattr.FieldID,
			},
		},
		From:   maq.sql,
		Unique: true,
	}
	if unique := maq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := maq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mailboxattr.FieldID)
		for i := range fields {
			if fields[i] != mailboxattr.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := maq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := maq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := maq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := maq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (maq *MailboxAttrQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(maq.driver.Dialect())
	t1 := builder.Table(mailboxattr.Table)
	columns := maq.fields
	if len(columns) == 0 {
		columns = mailboxattr.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if maq.sql != nil {
		selector = maq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if maq.unique != nil && *maq.unique {
		selector.Distinct()
	}
	for _, p := range maq.predicates {
		p(selector)
	}
	for _, p := range maq.order {
		p(selector)
	}
	if offset := maq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := maq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MailboxAttrGroupBy is the group-by builder for MailboxAttr entities.
type MailboxAttrGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (magb *MailboxAttrGroupBy) Aggregate(fns ...AggregateFunc) *MailboxAttrGroupBy {
	magb.fns = append(magb.fns, fns...)
	return magb
}

// Scan applies the group-by query and scans the result into the given value.
func (magb *MailboxAttrGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := magb.path(ctx)
	if err != nil {
		return err
	}
	magb.sql = query
	return magb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := magb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(magb.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := magb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) StringsX(ctx context.Context) []string {
	v, err := magb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = magb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) StringX(ctx context.Context) string {
	v, err := magb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(magb.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := magb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) IntsX(ctx context.Context) []int {
	v, err := magb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = magb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) IntX(ctx context.Context) int {
	v, err := magb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(magb.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := magb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := magb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = magb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) Float64X(ctx context.Context) float64 {
	v, err := magb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(magb.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := magb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := magb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (magb *MailboxAttrGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = magb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (magb *MailboxAttrGroupBy) BoolX(ctx context.Context) bool {
	v, err := magb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (magb *MailboxAttrGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range magb.fields {
		if !mailboxattr.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := magb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := magb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (magb *MailboxAttrGroupBy) sqlQuery() *sql.Selector {
	selector := magb.sql.Select()
	aggregation := make([]string, 0, len(magb.fns))
	for _, fn := range magb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(magb.fields)+len(magb.fns))
		for _, f := range magb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(magb.fields...)...)
}

// MailboxAttrSelect is the builder for selecting fields of MailboxAttr entities.
type MailboxAttrSelect struct {
	*MailboxAttrQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mas *MailboxAttrSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mas.prepareQuery(ctx); err != nil {
		return err
	}
	mas.sql = mas.MailboxAttrQuery.sqlQuery(ctx)
	return mas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mas *MailboxAttrSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mas.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mas *MailboxAttrSelect) StringsX(ctx context.Context) []string {
	v, err := mas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mas *MailboxAttrSelect) StringX(ctx context.Context) string {
	v, err := mas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mas.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mas *MailboxAttrSelect) IntsX(ctx context.Context) []int {
	v, err := mas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mas *MailboxAttrSelect) IntX(ctx context.Context) int {
	v, err := mas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mas.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mas *MailboxAttrSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mas *MailboxAttrSelect) Float64X(ctx context.Context) float64 {
	v, err := mas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mas.fields) > 1 {
		return nil, errors.New("ent: MailboxAttrSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mas *MailboxAttrSelect) BoolsX(ctx context.Context) []bool {
	v, err := mas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mas *MailboxAttrSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailboxattr.Label}
	default:
		err = fmt.Errorf("ent: MailboxAttrSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mas *MailboxAttrSelect) BoolX(ctx context.Context) bool {
	v, err := mas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mas *MailboxAttrSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mas.sql.Query()
	if err := mas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
