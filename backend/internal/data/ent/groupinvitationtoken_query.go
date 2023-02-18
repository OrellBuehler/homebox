// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/group"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/groupinvitationtoken"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/predicate"
)

// GroupInvitationTokenQuery is the builder for querying GroupInvitationToken entities.
type GroupInvitationTokenQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.GroupInvitationToken
	withGroup  *GroupQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupInvitationTokenQuery builder.
func (gitq *GroupInvitationTokenQuery) Where(ps ...predicate.GroupInvitationToken) *GroupInvitationTokenQuery {
	gitq.predicates = append(gitq.predicates, ps...)
	return gitq
}

// Limit the number of records to be returned by this query.
func (gitq *GroupInvitationTokenQuery) Limit(limit int) *GroupInvitationTokenQuery {
	gitq.ctx.Limit = &limit
	return gitq
}

// Offset to start from.
func (gitq *GroupInvitationTokenQuery) Offset(offset int) *GroupInvitationTokenQuery {
	gitq.ctx.Offset = &offset
	return gitq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gitq *GroupInvitationTokenQuery) Unique(unique bool) *GroupInvitationTokenQuery {
	gitq.ctx.Unique = &unique
	return gitq
}

// Order specifies how the records should be ordered.
func (gitq *GroupInvitationTokenQuery) Order(o ...OrderFunc) *GroupInvitationTokenQuery {
	gitq.order = append(gitq.order, o...)
	return gitq
}

// QueryGroup chains the current query on the "group" edge.
func (gitq *GroupInvitationTokenQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gitq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gitq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gitq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupinvitationtoken.Table, groupinvitationtoken.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, groupinvitationtoken.GroupTable, groupinvitationtoken.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(gitq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupInvitationToken entity from the query.
// Returns a *NotFoundError when no GroupInvitationToken was found.
func (gitq *GroupInvitationTokenQuery) First(ctx context.Context) (*GroupInvitationToken, error) {
	nodes, err := gitq.Limit(1).All(setContextOp(ctx, gitq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupinvitationtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) FirstX(ctx context.Context) *GroupInvitationToken {
	node, err := gitq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupInvitationToken ID from the query.
// Returns a *NotFoundError when no GroupInvitationToken ID was found.
func (gitq *GroupInvitationTokenQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gitq.Limit(1).IDs(setContextOp(ctx, gitq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupinvitationtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gitq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupInvitationToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupInvitationToken entity is found.
// Returns a *NotFoundError when no GroupInvitationToken entities are found.
func (gitq *GroupInvitationTokenQuery) Only(ctx context.Context) (*GroupInvitationToken, error) {
	nodes, err := gitq.Limit(2).All(setContextOp(ctx, gitq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupinvitationtoken.Label}
	default:
		return nil, &NotSingularError{groupinvitationtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) OnlyX(ctx context.Context) *GroupInvitationToken {
	node, err := gitq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupInvitationToken ID in the query.
// Returns a *NotSingularError when more than one GroupInvitationToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (gitq *GroupInvitationTokenQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gitq.Limit(2).IDs(setContextOp(ctx, gitq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupinvitationtoken.Label}
	default:
		err = &NotSingularError{groupinvitationtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gitq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupInvitationTokens.
func (gitq *GroupInvitationTokenQuery) All(ctx context.Context) ([]*GroupInvitationToken, error) {
	ctx = setContextOp(ctx, gitq.ctx, "All")
	if err := gitq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupInvitationToken, *GroupInvitationTokenQuery]()
	return withInterceptors[[]*GroupInvitationToken](ctx, gitq, qr, gitq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) AllX(ctx context.Context) []*GroupInvitationToken {
	nodes, err := gitq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupInvitationToken IDs.
func (gitq *GroupInvitationTokenQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if gitq.ctx.Unique == nil && gitq.path != nil {
		gitq.Unique(true)
	}
	ctx = setContextOp(ctx, gitq.ctx, "IDs")
	if err = gitq.Select(groupinvitationtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gitq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gitq *GroupInvitationTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gitq.ctx, "Count")
	if err := gitq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gitq, querierCount[*GroupInvitationTokenQuery](), gitq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) CountX(ctx context.Context) int {
	count, err := gitq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gitq *GroupInvitationTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gitq.ctx, "Exist")
	switch _, err := gitq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gitq *GroupInvitationTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := gitq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupInvitationTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gitq *GroupInvitationTokenQuery) Clone() *GroupInvitationTokenQuery {
	if gitq == nil {
		return nil
	}
	return &GroupInvitationTokenQuery{
		config:     gitq.config,
		ctx:        gitq.ctx.Clone(),
		order:      append([]OrderFunc{}, gitq.order...),
		inters:     append([]Interceptor{}, gitq.inters...),
		predicates: append([]predicate.GroupInvitationToken{}, gitq.predicates...),
		withGroup:  gitq.withGroup.Clone(),
		// clone intermediate query.
		sql:  gitq.sql.Clone(),
		path: gitq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gitq *GroupInvitationTokenQuery) WithGroup(opts ...func(*GroupQuery)) *GroupInvitationTokenQuery {
	query := (&GroupClient{config: gitq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gitq.withGroup = query
	return gitq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupInvitationToken.Query().
//		GroupBy(groupinvitationtoken.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gitq *GroupInvitationTokenQuery) GroupBy(field string, fields ...string) *GroupInvitationTokenGroupBy {
	gitq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupInvitationTokenGroupBy{build: gitq}
	grbuild.flds = &gitq.ctx.Fields
	grbuild.label = groupinvitationtoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.GroupInvitationToken.Query().
//		Select(groupinvitationtoken.FieldCreatedAt).
//		Scan(ctx, &v)
func (gitq *GroupInvitationTokenQuery) Select(fields ...string) *GroupInvitationTokenSelect {
	gitq.ctx.Fields = append(gitq.ctx.Fields, fields...)
	sbuild := &GroupInvitationTokenSelect{GroupInvitationTokenQuery: gitq}
	sbuild.label = groupinvitationtoken.Label
	sbuild.flds, sbuild.scan = &gitq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupInvitationTokenSelect configured with the given aggregations.
func (gitq *GroupInvitationTokenQuery) Aggregate(fns ...AggregateFunc) *GroupInvitationTokenSelect {
	return gitq.Select().Aggregate(fns...)
}

func (gitq *GroupInvitationTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gitq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gitq); err != nil {
				return err
			}
		}
	}
	for _, f := range gitq.ctx.Fields {
		if !groupinvitationtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gitq.path != nil {
		prev, err := gitq.path(ctx)
		if err != nil {
			return err
		}
		gitq.sql = prev
	}
	return nil
}

func (gitq *GroupInvitationTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupInvitationToken, error) {
	var (
		nodes       = []*GroupInvitationToken{}
		withFKs     = gitq.withFKs
		_spec       = gitq.querySpec()
		loadedTypes = [1]bool{
			gitq.withGroup != nil,
		}
	)
	if gitq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, groupinvitationtoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupInvitationToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupInvitationToken{config: gitq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gitq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gitq.withGroup; query != nil {
		if err := gitq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupInvitationToken, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gitq *GroupInvitationTokenQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupInvitationToken, init func(*GroupInvitationToken), assign func(*GroupInvitationToken, *Group)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GroupInvitationToken)
	for i := range nodes {
		if nodes[i].group_invitation_tokens == nil {
			continue
		}
		fk := *nodes[i].group_invitation_tokens
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
			return fmt.Errorf(`unexpected foreign-key "group_invitation_tokens" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gitq *GroupInvitationTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gitq.querySpec()
	_spec.Node.Columns = gitq.ctx.Fields
	if len(gitq.ctx.Fields) > 0 {
		_spec.Unique = gitq.ctx.Unique != nil && *gitq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gitq.driver, _spec)
}

func (gitq *GroupInvitationTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupinvitationtoken.Table, groupinvitationtoken.Columns, sqlgraph.NewFieldSpec(groupinvitationtoken.FieldID, field.TypeUUID))
	_spec.From = gitq.sql
	if unique := gitq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gitq.path != nil {
		_spec.Unique = true
	}
	if fields := gitq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupinvitationtoken.FieldID)
		for i := range fields {
			if fields[i] != groupinvitationtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gitq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gitq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gitq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gitq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gitq *GroupInvitationTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gitq.driver.Dialect())
	t1 := builder.Table(groupinvitationtoken.Table)
	columns := gitq.ctx.Fields
	if len(columns) == 0 {
		columns = groupinvitationtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gitq.sql != nil {
		selector = gitq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gitq.ctx.Unique != nil && *gitq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gitq.predicates {
		p(selector)
	}
	for _, p := range gitq.order {
		p(selector)
	}
	if offset := gitq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gitq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupInvitationTokenGroupBy is the group-by builder for GroupInvitationToken entities.
type GroupInvitationTokenGroupBy struct {
	selector
	build *GroupInvitationTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gitgb *GroupInvitationTokenGroupBy) Aggregate(fns ...AggregateFunc) *GroupInvitationTokenGroupBy {
	gitgb.fns = append(gitgb.fns, fns...)
	return gitgb
}

// Scan applies the selector query and scans the result into the given value.
func (gitgb *GroupInvitationTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gitgb.build.ctx, "GroupBy")
	if err := gitgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupInvitationTokenQuery, *GroupInvitationTokenGroupBy](ctx, gitgb.build, gitgb, gitgb.build.inters, v)
}

func (gitgb *GroupInvitationTokenGroupBy) sqlScan(ctx context.Context, root *GroupInvitationTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gitgb.fns))
	for _, fn := range gitgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gitgb.flds)+len(gitgb.fns))
		for _, f := range *gitgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gitgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gitgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupInvitationTokenSelect is the builder for selecting fields of GroupInvitationToken entities.
type GroupInvitationTokenSelect struct {
	*GroupInvitationTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gits *GroupInvitationTokenSelect) Aggregate(fns ...AggregateFunc) *GroupInvitationTokenSelect {
	gits.fns = append(gits.fns, fns...)
	return gits
}

// Scan applies the selector query and scans the result into the given value.
func (gits *GroupInvitationTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gits.ctx, "Select")
	if err := gits.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupInvitationTokenQuery, *GroupInvitationTokenSelect](ctx, gits.GroupInvitationTokenQuery, gits, gits.inters, v)
}

func (gits *GroupInvitationTokenSelect) sqlScan(ctx context.Context, root *GroupInvitationTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gits.fns))
	for _, fn := range gits.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gits.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gits.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
