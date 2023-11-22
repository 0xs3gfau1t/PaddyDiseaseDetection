// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"segFault/PaddyDiseaseDetection/ent/disease"
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/ent/predicate"
	"segFault/PaddyDiseaseDetection/ent/solution"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DiseaseQuery is the builder for querying Disease entities.
type DiseaseQuery struct {
	config
	ctx                   *QueryContext
	order                 []disease.OrderOption
	inters                []Interceptor
	predicates            []predicate.Disease
	withSolutions         *SolutionQuery
	withDiseaseIdentified *DiseaseIdentifiedQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiseaseQuery builder.
func (dq *DiseaseQuery) Where(ps ...predicate.Disease) *DiseaseQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DiseaseQuery) Limit(limit int) *DiseaseQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DiseaseQuery) Offset(offset int) *DiseaseQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DiseaseQuery) Unique(unique bool) *DiseaseQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DiseaseQuery) Order(o ...disease.OrderOption) *DiseaseQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QuerySolutions chains the current query on the "solutions" edge.
func (dq *DiseaseQuery) QuerySolutions() *SolutionQuery {
	query := (&SolutionClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(disease.Table, disease.FieldID, selector),
			sqlgraph.To(solution.Table, solution.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, disease.SolutionsTable, disease.SolutionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDiseaseIdentified chains the current query on the "disease_identified" edge.
func (dq *DiseaseQuery) QueryDiseaseIdentified() *DiseaseIdentifiedQuery {
	query := (&DiseaseIdentifiedClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(disease.Table, disease.FieldID, selector),
			sqlgraph.To(diseaseidentified.Table, diseaseidentified.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, disease.DiseaseIdentifiedTable, disease.DiseaseIdentifiedPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Disease entity from the query.
// Returns a *NotFoundError when no Disease was found.
func (dq *DiseaseQuery) First(ctx context.Context) (*Disease, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{disease.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DiseaseQuery) FirstX(ctx context.Context) *Disease {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Disease ID from the query.
// Returns a *NotFoundError when no Disease ID was found.
func (dq *DiseaseQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{disease.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DiseaseQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Disease entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Disease entity is found.
// Returns a *NotFoundError when no Disease entities are found.
func (dq *DiseaseQuery) Only(ctx context.Context) (*Disease, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{disease.Label}
	default:
		return nil, &NotSingularError{disease.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DiseaseQuery) OnlyX(ctx context.Context) *Disease {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Disease ID in the query.
// Returns a *NotSingularError when more than one Disease ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DiseaseQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{disease.Label}
	default:
		err = &NotSingularError{disease.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DiseaseQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Diseases.
func (dq *DiseaseQuery) All(ctx context.Context) ([]*Disease, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Disease, *DiseaseQuery]()
	return withInterceptors[[]*Disease](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DiseaseQuery) AllX(ctx context.Context) []*Disease {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Disease IDs.
func (dq *DiseaseQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(disease.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DiseaseQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DiseaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DiseaseQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DiseaseQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DiseaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DiseaseQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiseaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DiseaseQuery) Clone() *DiseaseQuery {
	if dq == nil {
		return nil
	}
	return &DiseaseQuery{
		config:                dq.config,
		ctx:                   dq.ctx.Clone(),
		order:                 append([]disease.OrderOption{}, dq.order...),
		inters:                append([]Interceptor{}, dq.inters...),
		predicates:            append([]predicate.Disease{}, dq.predicates...),
		withSolutions:         dq.withSolutions.Clone(),
		withDiseaseIdentified: dq.withDiseaseIdentified.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithSolutions tells the query-builder to eager-load the nodes that are connected to
// the "solutions" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DiseaseQuery) WithSolutions(opts ...func(*SolutionQuery)) *DiseaseQuery {
	query := (&SolutionClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withSolutions = query
	return dq
}

// WithDiseaseIdentified tells the query-builder to eager-load the nodes that are connected to
// the "disease_identified" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DiseaseQuery) WithDiseaseIdentified(opts ...func(*DiseaseIdentifiedQuery)) *DiseaseQuery {
	query := (&DiseaseIdentifiedClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDiseaseIdentified = query
	return dq
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
//	client.Disease.Query().
//		GroupBy(disease.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DiseaseQuery) GroupBy(field string, fields ...string) *DiseaseGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DiseaseGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = disease.Label
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
//	client.Disease.Query().
//		Select(disease.FieldName).
//		Scan(ctx, &v)
func (dq *DiseaseQuery) Select(fields ...string) *DiseaseSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DiseaseSelect{DiseaseQuery: dq}
	sbuild.label = disease.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DiseaseSelect configured with the given aggregations.
func (dq *DiseaseQuery) Aggregate(fns ...AggregateFunc) *DiseaseSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DiseaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !disease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DiseaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Disease, error) {
	var (
		nodes       = []*Disease{}
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withSolutions != nil,
			dq.withDiseaseIdentified != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Disease).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Disease{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withSolutions; query != nil {
		if err := dq.loadSolutions(ctx, query, nodes,
			func(n *Disease) { n.Edges.Solutions = []*Solution{} },
			func(n *Disease, e *Solution) { n.Edges.Solutions = append(n.Edges.Solutions, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withDiseaseIdentified; query != nil {
		if err := dq.loadDiseaseIdentified(ctx, query, nodes,
			func(n *Disease) { n.Edges.DiseaseIdentified = []*DiseaseIdentified{} },
			func(n *Disease, e *DiseaseIdentified) {
				n.Edges.DiseaseIdentified = append(n.Edges.DiseaseIdentified, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DiseaseQuery) loadSolutions(ctx context.Context, query *SolutionQuery, nodes []*Disease, init func(*Disease), assign func(*Disease, *Solution)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Disease)
	nids := make(map[uuid.UUID]map[*Disease]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(disease.SolutionsTable)
		s.Join(joinT).On(s.C(solution.FieldID), joinT.C(disease.SolutionsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(disease.SolutionsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(disease.SolutionsPrimaryKey[0]))
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
					nids[inValue] = map[*Disease]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Solution](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "solutions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DiseaseQuery) loadDiseaseIdentified(ctx context.Context, query *DiseaseIdentifiedQuery, nodes []*Disease, init func(*Disease), assign func(*Disease, *DiseaseIdentified)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Disease)
	nids := make(map[uuid.UUID]map[*Disease]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(disease.DiseaseIdentifiedTable)
		s.Join(joinT).On(s.C(diseaseidentified.FieldID), joinT.C(disease.DiseaseIdentifiedPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(disease.DiseaseIdentifiedPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(disease.DiseaseIdentifiedPrimaryKey[1]))
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
					nids[inValue] = map[*Disease]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*DiseaseIdentified](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "disease_identified" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (dq *DiseaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DiseaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(disease.Table, disease.Columns, sqlgraph.NewFieldSpec(disease.FieldID, field.TypeUUID))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, disease.FieldID)
		for i := range fields {
			if fields[i] != disease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DiseaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(disease.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = disease.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DiseaseGroupBy is the group-by builder for Disease entities.
type DiseaseGroupBy struct {
	selector
	build *DiseaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DiseaseGroupBy) Aggregate(fns ...AggregateFunc) *DiseaseGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DiseaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiseaseQuery, *DiseaseGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DiseaseGroupBy) sqlScan(ctx context.Context, root *DiseaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DiseaseSelect is the builder for selecting fields of Disease entities.
type DiseaseSelect struct {
	*DiseaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DiseaseSelect) Aggregate(fns ...AggregateFunc) *DiseaseSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DiseaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiseaseQuery, *DiseaseSelect](ctx, ds.DiseaseQuery, ds, ds.inters, v)
}

func (ds *DiseaseSelect) sqlScan(ctx context.Context, root *DiseaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
