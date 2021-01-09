// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/mediafile"
	"github.com/dreamvo/gilfoyle/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MediaFileQuery is the builder for querying MediaFile entities.
type MediaFileQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.MediaFile
	// eager-loading edges.
	withMedia *MediaQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mfq *MediaFileQuery) Where(ps ...predicate.MediaFile) *MediaFileQuery {
	mfq.predicates = append(mfq.predicates, ps...)
	return mfq
}

// Limit adds a limit step to the query.
func (mfq *MediaFileQuery) Limit(limit int) *MediaFileQuery {
	mfq.limit = &limit
	return mfq
}

// Offset adds an offset step to the query.
func (mfq *MediaFileQuery) Offset(offset int) *MediaFileQuery {
	mfq.offset = &offset
	return mfq
}

// Order adds an order step to the query.
func (mfq *MediaFileQuery) Order(o ...OrderFunc) *MediaFileQuery {
	mfq.order = append(mfq.order, o...)
	return mfq
}

// QueryMedia chains the current query on the media edge.
func (mfq *MediaFileQuery) QueryMedia() *MediaQuery {
	query := &MediaQuery{config: mfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mfq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mediafile.Table, mediafile.FieldID, selector),
			sqlgraph.To(media.Table, media.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mediafile.MediaTable, mediafile.MediaColumn),
		)
		fromU = sqlgraph.SetNeighbors(mfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MediaFile entity in the query. Returns *NotFoundError when no mediafile was found.
func (mfq *MediaFileQuery) First(ctx context.Context) (*MediaFile, error) {
	nodes, err := mfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mediafile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mfq *MediaFileQuery) FirstX(ctx context.Context) *MediaFile {
	node, err := mfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MediaFile id in the query. Returns *NotFoundError when no id was found.
func (mfq *MediaFileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mediafile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mfq *MediaFileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MediaFile entity in the query, returns an error if not exactly one entity was returned.
func (mfq *MediaFileQuery) Only(ctx context.Context) (*MediaFile, error) {
	nodes, err := mfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mediafile.Label}
	default:
		return nil, &NotSingularError{mediafile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mfq *MediaFileQuery) OnlyX(ctx context.Context) *MediaFile {
	node, err := mfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only MediaFile id in the query, returns an error if not exactly one id was returned.
func (mfq *MediaFileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = &NotSingularError{mediafile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mfq *MediaFileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MediaFiles.
func (mfq *MediaFileQuery) All(ctx context.Context) ([]*MediaFile, error) {
	if err := mfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mfq *MediaFileQuery) AllX(ctx context.Context) []*MediaFile {
	nodes, err := mfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MediaFile ids.
func (mfq *MediaFileQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := mfq.Select(mediafile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mfq *MediaFileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mfq *MediaFileQuery) Count(ctx context.Context) (int, error) {
	if err := mfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mfq *MediaFileQuery) CountX(ctx context.Context) int {
	count, err := mfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mfq *MediaFileQuery) Exist(ctx context.Context) (bool, error) {
	if err := mfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mfq *MediaFileQuery) ExistX(ctx context.Context) bool {
	exist, err := mfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mfq *MediaFileQuery) Clone() *MediaFileQuery {
	if mfq == nil {
		return nil
	}
	return &MediaFileQuery{
		config:     mfq.config,
		limit:      mfq.limit,
		offset:     mfq.offset,
		order:      append([]OrderFunc{}, mfq.order...),
		predicates: append([]predicate.MediaFile{}, mfq.predicates...),
		withMedia:  mfq.withMedia.Clone(),
		// clone intermediate query.
		sql:  mfq.sql.Clone(),
		path: mfq.path,
	}
}

//  WithMedia tells the query-builder to eager-loads the nodes that are connected to
// the "media" edge. The optional arguments used to configure the query builder of the edge.
func (mfq *MediaFileQuery) WithMedia(opts ...func(*MediaQuery)) *MediaFileQuery {
	query := &MediaQuery{config: mfq.config}
	for _, opt := range opts {
		opt(query)
	}
	mfq.withMedia = query
	return mfq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Format string `json:"format,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MediaFile.Query().
//		GroupBy(mediafile.FieldFormat).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mfq *MediaFileQuery) GroupBy(field string, fields ...string) *MediaFileGroupBy {
	group := &MediaFileGroupBy{config: mfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mfq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Format string `json:"format,omitempty"`
//	}
//
//	client.MediaFile.Query().
//		Select(mediafile.FieldFormat).
//		Scan(ctx, &v)
//
func (mfq *MediaFileQuery) Select(field string, fields ...string) *MediaFileSelect {
	selector := &MediaFileSelect{config: mfq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mfq.sqlQuery(), nil
	}
	return selector
}

func (mfq *MediaFileQuery) prepareQuery(ctx context.Context) error {
	if mfq.path != nil {
		prev, err := mfq.path(ctx)
		if err != nil {
			return err
		}
		mfq.sql = prev
	}
	return nil
}

func (mfq *MediaFileQuery) sqlAll(ctx context.Context) ([]*MediaFile, error) {
	var (
		nodes       = []*MediaFile{}
		withFKs     = mfq.withFKs
		_spec       = mfq.querySpec()
		loadedTypes = [1]bool{
			mfq.withMedia != nil,
		}
	)
	if mfq.withMedia != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mediafile.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &MediaFile{config: mfq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, mfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mfq.withMedia; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*MediaFile)
		for i := range nodes {
			if fk := nodes[i].media; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(media.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "media" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Media = n
			}
		}
	}

	return nodes, nil
}

func (mfq *MediaFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mfq.querySpec()
	return sqlgraph.CountNodes(ctx, mfq.driver, _spec)
}

func (mfq *MediaFileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mfq *MediaFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mediafile.Table,
			Columns: mediafile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mediafile.FieldID,
			},
		},
		From:   mfq.sql,
		Unique: true,
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
				ps[i](selector, mediafile.ValidColumn)
			}
		}
	}
	return _spec
}

func (mfq *MediaFileQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mfq.driver.Dialect())
	t1 := builder.Table(mediafile.Table)
	selector := builder.Select(t1.Columns(mediafile.Columns...)...).From(t1)
	if mfq.sql != nil {
		selector = mfq.sql
		selector.Select(selector.Columns(mediafile.Columns...)...)
	}
	for _, p := range mfq.predicates {
		p(selector)
	}
	for _, p := range mfq.order {
		p(selector, mediafile.ValidColumn)
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

// MediaFileGroupBy is the builder for group-by MediaFile entities.
type MediaFileGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mfgb *MediaFileGroupBy) Aggregate(fns ...AggregateFunc) *MediaFileGroupBy {
	mfgb.fns = append(mfgb.fns, fns...)
	return mfgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mfgb *MediaFileGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mfgb.path(ctx)
	if err != nil {
		return err
	}
	mfgb.sql = query
	return mfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mfgb.fields) > 1 {
		return nil, errors.New("ent: MediaFileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) StringsX(ctx context.Context) []string {
	v, err := mfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) StringX(ctx context.Context) string {
	v, err := mfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mfgb.fields) > 1 {
		return nil, errors.New("ent: MediaFileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) IntsX(ctx context.Context) []int {
	v, err := mfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) IntX(ctx context.Context) int {
	v, err := mfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mfgb.fields) > 1 {
		return nil, errors.New("ent: MediaFileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mfgb.fields) > 1 {
		return nil, errors.New("ent: MediaFileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mfgb *MediaFileGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mfgb *MediaFileGroupBy) BoolX(ctx context.Context) bool {
	v, err := mfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mfgb *MediaFileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mfgb.fields {
		if !mediafile.ValidColumn(f) {
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

func (mfgb *MediaFileGroupBy) sqlQuery() *sql.Selector {
	selector := mfgb.sql
	columns := make([]string, 0, len(mfgb.fields)+len(mfgb.fns))
	columns = append(columns, mfgb.fields...)
	for _, fn := range mfgb.fns {
		columns = append(columns, fn(selector, mediafile.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(mfgb.fields...)
}

// MediaFileSelect is the builder for select fields of MediaFile entities.
type MediaFileSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (mfs *MediaFileSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := mfs.path(ctx)
	if err != nil {
		return err
	}
	mfs.sql = query
	return mfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mfs *MediaFileSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mfs.fields) > 1 {
		return nil, errors.New("ent: MediaFileSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mfs *MediaFileSelect) StringsX(ctx context.Context) []string {
	v, err := mfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mfs *MediaFileSelect) StringX(ctx context.Context) string {
	v, err := mfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mfs.fields) > 1 {
		return nil, errors.New("ent: MediaFileSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mfs *MediaFileSelect) IntsX(ctx context.Context) []int {
	v, err := mfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mfs *MediaFileSelect) IntX(ctx context.Context) int {
	v, err := mfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mfs.fields) > 1 {
		return nil, errors.New("ent: MediaFileSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mfs *MediaFileSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mfs *MediaFileSelect) Float64X(ctx context.Context) float64 {
	v, err := mfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mfs.fields) > 1 {
		return nil, errors.New("ent: MediaFileSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mfs *MediaFileSelect) BoolsX(ctx context.Context) []bool {
	v, err := mfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mfs *MediaFileSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediafile.Label}
	default:
		err = fmt.Errorf("ent: MediaFileSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mfs *MediaFileSelect) BoolX(ctx context.Context) bool {
	v, err := mfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mfs *MediaFileSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mfs.fields {
		if !mediafile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := mfs.sqlQuery().Query()
	if err := mfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mfs *MediaFileSelect) sqlQuery() sql.Querier {
	selector := mfs.sql
	selector.Select(selector.Columns(mfs.fields...)...)
	return selector
}
