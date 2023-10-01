// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"hospital/internal/modules/db/ent/disease"
	"hospital/internal/modules/db/ent/doctor"
	"hospital/internal/modules/db/ent/patient"
	"hospital/internal/modules/db/ent/predicate"
	"hospital/internal/modules/db/ent/room"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PatientQuery is the builder for querying Patient entities.
type PatientQuery struct {
	config
	ctx        *QueryContext
	order      []patient.Order
	inters     []Interceptor
	predicates []predicate.Patient
	withRepo   *RoomQuery
	withDoctor *DoctorQuery
	withIlls   *DiseaseQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PatientQuery builder.
func (pq *PatientQuery) Where(ps ...predicate.Patient) *PatientQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PatientQuery) Limit(limit int) *PatientQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PatientQuery) Offset(offset int) *PatientQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PatientQuery) Unique(unique bool) *PatientQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PatientQuery) Order(o ...patient.Order) *PatientQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryRepo chains the current query on the "repo" edge.
func (pq *PatientQuery) QueryRepo() *RoomQuery {
	query := (&RoomClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patient.RepoTable, patient.RepoColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctor chains the current query on the "doctor" edge.
func (pq *PatientQuery) QueryDoctor() *DoctorQuery {
	query := (&DoctorClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, selector),
			sqlgraph.To(doctor.Table, doctor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, patient.DoctorTable, patient.DoctorPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIlls chains the current query on the "ills" edge.
func (pq *PatientQuery) QueryIlls() *DiseaseQuery {
	query := (&DiseaseClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, selector),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patient.IllsTable, patient.IllsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Patient entity from the query.
// Returns a *NotFoundError when no Patient was found.
func (pq *PatientQuery) First(ctx context.Context) (*Patient, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{patient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PatientQuery) FirstX(ctx context.Context) *Patient {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Patient ID from the query.
// Returns a *NotFoundError when no Patient ID was found.
func (pq *PatientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{patient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PatientQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Patient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Patient entity is found.
// Returns a *NotFoundError when no Patient entities are found.
func (pq *PatientQuery) Only(ctx context.Context) (*Patient, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{patient.Label}
	default:
		return nil, &NotSingularError{patient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PatientQuery) OnlyX(ctx context.Context) *Patient {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Patient ID in the query.
// Returns a *NotSingularError when more than one Patient ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PatientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{patient.Label}
	default:
		err = &NotSingularError{patient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PatientQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Patients.
func (pq *PatientQuery) All(ctx context.Context) ([]*Patient, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Patient, *PatientQuery]()
	return withInterceptors[[]*Patient](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PatientQuery) AllX(ctx context.Context) []*Patient {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Patient IDs.
func (pq *PatientQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(patient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PatientQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PatientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PatientQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PatientQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PatientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PatientQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PatientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PatientQuery) Clone() *PatientQuery {
	if pq == nil {
		return nil
	}
	return &PatientQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]patient.Order{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Patient{}, pq.predicates...),
		withRepo:   pq.withRepo.Clone(),
		withDoctor: pq.withDoctor.Clone(),
		withIlls:   pq.withIlls.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithRepo tells the query-builder to eager-load the nodes that are connected to
// the "repo" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PatientQuery) WithRepo(opts ...func(*RoomQuery)) *PatientQuery {
	query := (&RoomClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withRepo = query
	return pq
}

// WithDoctor tells the query-builder to eager-load the nodes that are connected to
// the "doctor" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PatientQuery) WithDoctor(opts ...func(*DoctorQuery)) *PatientQuery {
	query := (&DoctorClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withDoctor = query
	return pq
}

// WithIlls tells the query-builder to eager-load the nodes that are connected to
// the "ills" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PatientQuery) WithIlls(opts ...func(*DiseaseQuery)) *PatientQuery {
	query := (&DiseaseClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withIlls = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Surname string `json:"surname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Patient.Query().
//		GroupBy(patient.FieldSurname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PatientQuery) GroupBy(field string, fields ...string) *PatientGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PatientGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = patient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Surname string `json:"surname,omitempty"`
//	}
//
//	client.Patient.Query().
//		Select(patient.FieldSurname).
//		Scan(ctx, &v)
func (pq *PatientQuery) Select(fields ...string) *PatientSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PatientSelect{PatientQuery: pq}
	sbuild.label = patient.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PatientSelect configured with the given aggregations.
func (pq *PatientQuery) Aggregate(fns ...AggregateFunc) *PatientSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PatientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !patient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PatientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Patient, error) {
	var (
		nodes       = []*Patient{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [3]bool{
			pq.withRepo != nil,
			pq.withDoctor != nil,
			pq.withIlls != nil,
		}
	)
	if pq.withIlls != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, patient.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Patient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Patient{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withRepo; query != nil {
		if err := pq.loadRepo(ctx, query, nodes, nil,
			func(n *Patient, e *Room) { n.Edges.Repo = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withDoctor; query != nil {
		if err := pq.loadDoctor(ctx, query, nodes,
			func(n *Patient) { n.Edges.Doctor = []*Doctor{} },
			func(n *Patient, e *Doctor) { n.Edges.Doctor = append(n.Edges.Doctor, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withIlls; query != nil {
		if err := pq.loadIlls(ctx, query, nodes, nil,
			func(n *Patient, e *Disease) { n.Edges.Ills = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PatientQuery) loadRepo(ctx context.Context, query *RoomQuery, nodes []*Patient, init func(*Patient), assign func(*Patient, *Room)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Patient)
	for i := range nodes {
		fk := nodes[i].RoomNumber
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(room.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "roomNumber" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PatientQuery) loadDoctor(ctx context.Context, query *DoctorQuery, nodes []*Patient, init func(*Patient), assign func(*Patient, *Doctor)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Patient)
	nids := make(map[int]map[*Patient]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(patient.DoctorTable)
		s.Join(joinT).On(s.C(doctor.FieldID), joinT.C(patient.DoctorPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(patient.DoctorPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(patient.DoctorPrimaryKey[1]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Patient]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Doctor](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "doctor" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PatientQuery) loadIlls(ctx context.Context, query *DiseaseQuery, nodes []*Patient, init func(*Patient), assign func(*Patient, *Disease)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Patient)
	for i := range nodes {
		if nodes[i].disease_has == nil {
			continue
		}
		fk := *nodes[i].disease_has
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(disease.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "disease_has" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PatientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PatientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(patient.Table, patient.Columns, sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, patient.FieldID)
		for i := range fields {
			if fields[i] != patient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withRepo != nil {
			_spec.Node.AddColumnOnce(patient.FieldRoomNumber)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PatientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(patient.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = patient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PatientGroupBy is the group-by builder for Patient entities.
type PatientGroupBy struct {
	selector
	build *PatientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PatientGroupBy) Aggregate(fns ...AggregateFunc) *PatientGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PatientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PatientQuery, *PatientGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PatientGroupBy) sqlScan(ctx context.Context, root *PatientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PatientSelect is the builder for selecting fields of Patient entities.
type PatientSelect struct {
	*PatientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PatientSelect) Aggregate(fns ...AggregateFunc) *PatientSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PatientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PatientQuery, *PatientSelect](ctx, ps.PatientQuery, ps, ps.inters, v)
}

func (ps *PatientSelect) sqlScan(ctx context.Context, root *PatientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
