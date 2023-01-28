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
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/attachment"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/group"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/item"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/itemfield"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/label"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/location"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/maintenanceentry"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/predicate"
)

// ItemQuery is the builder for querying Item entities.
type ItemQuery struct {
	config
	ctx                    *QueryContext
	order                  []OrderFunc
	inters                 []Interceptor
	predicates             []predicate.Item
	withParent             *ItemQuery
	withChildren           *ItemQuery
	withGroup              *GroupQuery
	withLabel              *LabelQuery
	withLocation           *LocationQuery
	withFields             *ItemFieldQuery
	withMaintenanceEntries *MaintenanceEntryQuery
	withAttachments        *AttachmentQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemQuery builder.
func (iq *ItemQuery) Where(ps ...predicate.Item) *ItemQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *ItemQuery) Limit(limit int) *ItemQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *ItemQuery) Offset(offset int) *ItemQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ItemQuery) Unique(unique bool) *ItemQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *ItemQuery) Order(o ...OrderFunc) *ItemQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryParent chains the current query on the "parent" edge.
func (iq *ItemQuery) QueryParent() *ItemQuery {
	query := (&ItemClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.ParentTable, item.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (iq *ItemQuery) QueryChildren() *ItemQuery {
	query := (&ItemClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.ChildrenTable, item.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (iq *ItemQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.GroupTable, item.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLabel chains the current query on the "label" edge.
func (iq *ItemQuery) QueryLabel() *LabelQuery {
	query := (&LabelClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(label.Table, label.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, item.LabelTable, item.LabelPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLocation chains the current query on the "location" edge.
func (iq *ItemQuery) QueryLocation() *LocationQuery {
	query := (&LocationClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.LocationTable, item.LocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFields chains the current query on the "fields" edge.
func (iq *ItemQuery) QueryFields() *ItemFieldQuery {
	query := (&ItemFieldClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(itemfield.Table, itemfield.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.FieldsTable, item.FieldsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMaintenanceEntries chains the current query on the "maintenance_entries" edge.
func (iq *ItemQuery) QueryMaintenanceEntries() *MaintenanceEntryQuery {
	query := (&MaintenanceEntryClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(maintenanceentry.Table, maintenanceentry.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.MaintenanceEntriesTable, item.MaintenanceEntriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttachments chains the current query on the "attachments" edge.
func (iq *ItemQuery) QueryAttachments() *AttachmentQuery {
	query := (&AttachmentClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.AttachmentsTable, item.AttachmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Item entity from the query.
// Returns a *NotFoundError when no Item was found.
func (iq *ItemQuery) First(ctx context.Context) (*Item, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{item.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ItemQuery) FirstX(ctx context.Context) *Item {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Item ID from the query.
// Returns a *NotFoundError when no Item ID was found.
func (iq *ItemQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{item.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ItemQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Item entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Item entity is found.
// Returns a *NotFoundError when no Item entities are found.
func (iq *ItemQuery) Only(ctx context.Context) (*Item, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{item.Label}
	default:
		return nil, &NotSingularError{item.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ItemQuery) OnlyX(ctx context.Context) *Item {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Item ID in the query.
// Returns a *NotSingularError when more than one Item ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ItemQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{item.Label}
	default:
		err = &NotSingularError{item.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ItemQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Items.
func (iq *ItemQuery) All(ctx context.Context) ([]*Item, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Item, *ItemQuery]()
	return withInterceptors[[]*Item](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *ItemQuery) AllX(ctx context.Context) []*Item {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Item IDs.
func (iq *ItemQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err := iq.Select(item.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ItemQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*ItemQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ItemQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ItemQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ItemQuery) Clone() *ItemQuery {
	if iq == nil {
		return nil
	}
	return &ItemQuery{
		config:                 iq.config,
		ctx:                    iq.ctx.Clone(),
		order:                  append([]OrderFunc{}, iq.order...),
		inters:                 append([]Interceptor{}, iq.inters...),
		predicates:             append([]predicate.Item{}, iq.predicates...),
		withParent:             iq.withParent.Clone(),
		withChildren:           iq.withChildren.Clone(),
		withGroup:              iq.withGroup.Clone(),
		withLabel:              iq.withLabel.Clone(),
		withLocation:           iq.withLocation.Clone(),
		withFields:             iq.withFields.Clone(),
		withMaintenanceEntries: iq.withMaintenanceEntries.Clone(),
		withAttachments:        iq.withAttachments.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithParent(opts ...func(*ItemQuery)) *ItemQuery {
	query := (&ItemClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withParent = query
	return iq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithChildren(opts ...func(*ItemQuery)) *ItemQuery {
	query := (&ItemClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withChildren = query
	return iq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithGroup(opts ...func(*GroupQuery)) *ItemQuery {
	query := (&GroupClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withGroup = query
	return iq
}

// WithLabel tells the query-builder to eager-load the nodes that are connected to
// the "label" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithLabel(opts ...func(*LabelQuery)) *ItemQuery {
	query := (&LabelClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withLabel = query
	return iq
}

// WithLocation tells the query-builder to eager-load the nodes that are connected to
// the "location" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithLocation(opts ...func(*LocationQuery)) *ItemQuery {
	query := (&LocationClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withLocation = query
	return iq
}

// WithFields tells the query-builder to eager-load the nodes that are connected to
// the "fields" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithFields(opts ...func(*ItemFieldQuery)) *ItemQuery {
	query := (&ItemFieldClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withFields = query
	return iq
}

// WithMaintenanceEntries tells the query-builder to eager-load the nodes that are connected to
// the "maintenance_entries" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithMaintenanceEntries(opts ...func(*MaintenanceEntryQuery)) *ItemQuery {
	query := (&MaintenanceEntryClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withMaintenanceEntries = query
	return iq
}

// WithAttachments tells the query-builder to eager-load the nodes that are connected to
// the "attachments" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithAttachments(opts ...func(*AttachmentQuery)) *ItemQuery {
	query := (&AttachmentClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withAttachments = query
	return iq
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
//	client.Item.Query().
//		GroupBy(item.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *ItemQuery) GroupBy(field string, fields ...string) *ItemGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = item.Label
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
//	client.Item.Query().
//		Select(item.FieldCreatedAt).
//		Scan(ctx, &v)
func (iq *ItemQuery) Select(fields ...string) *ItemSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &ItemSelect{ItemQuery: iq}
	sbuild.label = item.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemSelect configured with the given aggregations.
func (iq *ItemQuery) Aggregate(fns ...AggregateFunc) *ItemSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *ItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !item.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *ItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Item, error) {
	var (
		nodes       = []*Item{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [8]bool{
			iq.withParent != nil,
			iq.withChildren != nil,
			iq.withGroup != nil,
			iq.withLabel != nil,
			iq.withLocation != nil,
			iq.withFields != nil,
			iq.withMaintenanceEntries != nil,
			iq.withAttachments != nil,
		}
	)
	if iq.withParent != nil || iq.withGroup != nil || iq.withLocation != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, item.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Item).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Item{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withParent; query != nil {
		if err := iq.loadParent(ctx, query, nodes, nil,
			func(n *Item, e *Item) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withChildren; query != nil {
		if err := iq.loadChildren(ctx, query, nodes,
			func(n *Item) { n.Edges.Children = []*Item{} },
			func(n *Item, e *Item) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withGroup; query != nil {
		if err := iq.loadGroup(ctx, query, nodes, nil,
			func(n *Item, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withLabel; query != nil {
		if err := iq.loadLabel(ctx, query, nodes,
			func(n *Item) { n.Edges.Label = []*Label{} },
			func(n *Item, e *Label) { n.Edges.Label = append(n.Edges.Label, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withLocation; query != nil {
		if err := iq.loadLocation(ctx, query, nodes, nil,
			func(n *Item, e *Location) { n.Edges.Location = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withFields; query != nil {
		if err := iq.loadFields(ctx, query, nodes,
			func(n *Item) { n.Edges.Fields = []*ItemField{} },
			func(n *Item, e *ItemField) { n.Edges.Fields = append(n.Edges.Fields, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withMaintenanceEntries; query != nil {
		if err := iq.loadMaintenanceEntries(ctx, query, nodes,
			func(n *Item) { n.Edges.MaintenanceEntries = []*MaintenanceEntry{} },
			func(n *Item, e *MaintenanceEntry) { n.Edges.MaintenanceEntries = append(n.Edges.MaintenanceEntries, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withAttachments; query != nil {
		if err := iq.loadAttachments(ctx, query, nodes,
			func(n *Item) { n.Edges.Attachments = []*Attachment{} },
			func(n *Item, e *Attachment) { n.Edges.Attachments = append(n.Edges.Attachments, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *ItemQuery) loadParent(ctx context.Context, query *ItemQuery, nodes []*Item, init func(*Item), assign func(*Item, *Item)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Item)
	for i := range nodes {
		if nodes[i].item_children == nil {
			continue
		}
		fk := *nodes[i].item_children
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(item.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_children" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *ItemQuery) loadChildren(ctx context.Context, query *ItemQuery, nodes []*Item, init func(*Item), assign func(*Item, *Item)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Item)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Item(func(s *sql.Selector) {
		s.Where(sql.InValues(item.ChildrenColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.item_children
		if fk == nil {
			return fmt.Errorf(`foreign-key "item_children" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_children" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *ItemQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*Item, init func(*Item), assign func(*Item, *Group)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Item)
	for i := range nodes {
		if nodes[i].group_items == nil {
			continue
		}
		fk := *nodes[i].group_items
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
			return fmt.Errorf(`unexpected foreign-key "group_items" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *ItemQuery) loadLabel(ctx context.Context, query *LabelQuery, nodes []*Item, init func(*Item), assign func(*Item, *Label)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Item)
	nids := make(map[uuid.UUID]map[*Item]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(item.LabelTable)
		s.Join(joinT).On(s.C(label.FieldID), joinT.C(item.LabelPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(item.LabelPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(item.LabelPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Item]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Label](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "label" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (iq *ItemQuery) loadLocation(ctx context.Context, query *LocationQuery, nodes []*Item, init func(*Item), assign func(*Item, *Location)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Item)
	for i := range nodes {
		if nodes[i].location_items == nil {
			continue
		}
		fk := *nodes[i].location_items
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_items" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *ItemQuery) loadFields(ctx context.Context, query *ItemFieldQuery, nodes []*Item, init func(*Item), assign func(*Item, *ItemField)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Item)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ItemField(func(s *sql.Selector) {
		s.Where(sql.InValues(item.FieldsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.item_fields
		if fk == nil {
			return fmt.Errorf(`foreign-key "item_fields" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_fields" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *ItemQuery) loadMaintenanceEntries(ctx context.Context, query *MaintenanceEntryQuery, nodes []*Item, init func(*Item), assign func(*Item, *MaintenanceEntry)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Item)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.MaintenanceEntry(func(s *sql.Selector) {
		s.Where(sql.InValues(item.MaintenanceEntriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ItemID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *ItemQuery) loadAttachments(ctx context.Context, query *AttachmentQuery, nodes []*Item, init func(*Item), assign func(*Item, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Item)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(item.AttachmentsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.item_attachments
		if fk == nil {
			return fmt.Errorf(`foreign-key "item_attachments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_attachments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *ItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *ItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for i := range fields {
			if fields[i] != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *ItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(item.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = item.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemGroupBy is the group-by builder for Item entities.
type ItemGroupBy struct {
	selector
	build *ItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ItemGroupBy) Aggregate(fns ...AggregateFunc) *ItemGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *ItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemQuery, *ItemGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *ItemGroupBy) sqlScan(ctx context.Context, root *ItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ItemSelect is the builder for selecting fields of Item entities.
type ItemSelect struct {
	*ItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *ItemSelect) Aggregate(fns ...AggregateFunc) *ItemSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *ItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemQuery, *ItemSelect](ctx, is.ItemQuery, is, is.inters, v)
}

func (is *ItemSelect) sqlScan(ctx context.Context, root *ItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
