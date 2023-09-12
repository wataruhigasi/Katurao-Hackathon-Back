// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// ThreadRakugaki is an object representing the database table.
type ThreadRakugaki struct {
	ID        int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ThreadID  int64      `boil:"thread_id" json:"thread_id" toml:"thread_id" yaml:"thread_id"`
	Body      string     `boil:"body" json:"body" toml:"body" yaml:"body"`
	Position  types.JSON `boil:"position" json:"position" toml:"position" yaml:"position"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *threadRakugakiR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L threadRakugakiL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ThreadRakugakiColumns = struct {
	ID        string
	ThreadID  string
	Body      string
	Position  string
	CreatedAt string
}{
	ID:        "id",
	ThreadID:  "thread_id",
	Body:      "body",
	Position:  "position",
	CreatedAt: "created_at",
}

var ThreadRakugakiTableColumns = struct {
	ID        string
	ThreadID  string
	Body      string
	Position  string
	CreatedAt string
}{
	ID:        "thread_rakugaki.id",
	ThreadID:  "thread_rakugaki.thread_id",
	Body:      "thread_rakugaki.body",
	Position:  "thread_rakugaki.position",
	CreatedAt: "thread_rakugaki.created_at",
}

// Generated where

var ThreadRakugakiWhere = struct {
	ID        whereHelperint64
	ThreadID  whereHelperint64
	Body      whereHelperstring
	Position  whereHelpertypes_JSON
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`thread_rakugaki`.`id`"},
	ThreadID:  whereHelperint64{field: "`thread_rakugaki`.`thread_id`"},
	Body:      whereHelperstring{field: "`thread_rakugaki`.`body`"},
	Position:  whereHelpertypes_JSON{field: "`thread_rakugaki`.`position`"},
	CreatedAt: whereHelpertime_Time{field: "`thread_rakugaki`.`created_at`"},
}

// ThreadRakugakiRels is where relationship names are stored.
var ThreadRakugakiRels = struct {
}{}

// threadRakugakiR is where relationships are stored.
type threadRakugakiR struct {
}

// NewStruct creates a new relationship struct
func (*threadRakugakiR) NewStruct() *threadRakugakiR {
	return &threadRakugakiR{}
}

// threadRakugakiL is where Load methods for each relationship are stored.
type threadRakugakiL struct{}

var (
	threadRakugakiAllColumns            = []string{"id", "thread_id", "body", "position", "created_at"}
	threadRakugakiColumnsWithoutDefault = []string{"thread_id", "body", "position"}
	threadRakugakiColumnsWithDefault    = []string{"id", "created_at"}
	threadRakugakiPrimaryKeyColumns     = []string{"id"}
	threadRakugakiGeneratedColumns      = []string{}
)

type (
	// ThreadRakugakiSlice is an alias for a slice of pointers to ThreadRakugaki.
	// This should almost always be used instead of []ThreadRakugaki.
	ThreadRakugakiSlice []*ThreadRakugaki
	// ThreadRakugakiHook is the signature for custom ThreadRakugaki hook methods
	ThreadRakugakiHook func(context.Context, boil.ContextExecutor, *ThreadRakugaki) error

	threadRakugakiQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	threadRakugakiType                 = reflect.TypeOf(&ThreadRakugaki{})
	threadRakugakiMapping              = queries.MakeStructMapping(threadRakugakiType)
	threadRakugakiPrimaryKeyMapping, _ = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, threadRakugakiPrimaryKeyColumns)
	threadRakugakiInsertCacheMut       sync.RWMutex
	threadRakugakiInsertCache          = make(map[string]insertCache)
	threadRakugakiUpdateCacheMut       sync.RWMutex
	threadRakugakiUpdateCache          = make(map[string]updateCache)
	threadRakugakiUpsertCacheMut       sync.RWMutex
	threadRakugakiUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var threadRakugakiAfterSelectHooks []ThreadRakugakiHook

var threadRakugakiBeforeInsertHooks []ThreadRakugakiHook
var threadRakugakiAfterInsertHooks []ThreadRakugakiHook

var threadRakugakiBeforeUpdateHooks []ThreadRakugakiHook
var threadRakugakiAfterUpdateHooks []ThreadRakugakiHook

var threadRakugakiBeforeDeleteHooks []ThreadRakugakiHook
var threadRakugakiAfterDeleteHooks []ThreadRakugakiHook

var threadRakugakiBeforeUpsertHooks []ThreadRakugakiHook
var threadRakugakiAfterUpsertHooks []ThreadRakugakiHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ThreadRakugaki) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ThreadRakugaki) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ThreadRakugaki) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ThreadRakugaki) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ThreadRakugaki) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ThreadRakugaki) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ThreadRakugaki) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ThreadRakugaki) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ThreadRakugaki) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadRakugakiAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddThreadRakugakiHook registers your hook function for all future operations.
func AddThreadRakugakiHook(hookPoint boil.HookPoint, threadRakugakiHook ThreadRakugakiHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		threadRakugakiAfterSelectHooks = append(threadRakugakiAfterSelectHooks, threadRakugakiHook)
	case boil.BeforeInsertHook:
		threadRakugakiBeforeInsertHooks = append(threadRakugakiBeforeInsertHooks, threadRakugakiHook)
	case boil.AfterInsertHook:
		threadRakugakiAfterInsertHooks = append(threadRakugakiAfterInsertHooks, threadRakugakiHook)
	case boil.BeforeUpdateHook:
		threadRakugakiBeforeUpdateHooks = append(threadRakugakiBeforeUpdateHooks, threadRakugakiHook)
	case boil.AfterUpdateHook:
		threadRakugakiAfterUpdateHooks = append(threadRakugakiAfterUpdateHooks, threadRakugakiHook)
	case boil.BeforeDeleteHook:
		threadRakugakiBeforeDeleteHooks = append(threadRakugakiBeforeDeleteHooks, threadRakugakiHook)
	case boil.AfterDeleteHook:
		threadRakugakiAfterDeleteHooks = append(threadRakugakiAfterDeleteHooks, threadRakugakiHook)
	case boil.BeforeUpsertHook:
		threadRakugakiBeforeUpsertHooks = append(threadRakugakiBeforeUpsertHooks, threadRakugakiHook)
	case boil.AfterUpsertHook:
		threadRakugakiAfterUpsertHooks = append(threadRakugakiAfterUpsertHooks, threadRakugakiHook)
	}
}

// One returns a single threadRakugaki record from the query.
func (q threadRakugakiQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ThreadRakugaki, error) {
	o := &ThreadRakugaki{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for thread_rakugaki")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ThreadRakugaki records from the query.
func (q threadRakugakiQuery) All(ctx context.Context, exec boil.ContextExecutor) (ThreadRakugakiSlice, error) {
	var o []*ThreadRakugaki

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ThreadRakugaki slice")
	}

	if len(threadRakugakiAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ThreadRakugaki records in the query.
func (q threadRakugakiQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count thread_rakugaki rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q threadRakugakiQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if thread_rakugaki exists")
	}

	return count > 0, nil
}

// ThreadRakugakis retrieves all the records using an executor.
func ThreadRakugakis(mods ...qm.QueryMod) threadRakugakiQuery {
	mods = append(mods, qm.From("`thread_rakugaki`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`thread_rakugaki`.*"})
	}

	return threadRakugakiQuery{q}
}

// FindThreadRakugaki retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindThreadRakugaki(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ThreadRakugaki, error) {
	threadRakugakiObj := &ThreadRakugaki{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `thread_rakugaki` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, threadRakugakiObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from thread_rakugaki")
	}

	if err = threadRakugakiObj.doAfterSelectHooks(ctx, exec); err != nil {
		return threadRakugakiObj, err
	}

	return threadRakugakiObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ThreadRakugaki) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no thread_rakugaki provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(threadRakugakiColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	threadRakugakiInsertCacheMut.RLock()
	cache, cached := threadRakugakiInsertCache[key]
	threadRakugakiInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			threadRakugakiAllColumns,
			threadRakugakiColumnsWithDefault,
			threadRakugakiColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `thread_rakugaki` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `thread_rakugaki` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `thread_rakugaki` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, threadRakugakiPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into thread_rakugaki")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == threadRakugakiMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for thread_rakugaki")
	}

CacheNoHooks:
	if !cached {
		threadRakugakiInsertCacheMut.Lock()
		threadRakugakiInsertCache[key] = cache
		threadRakugakiInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ThreadRakugaki.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ThreadRakugaki) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	threadRakugakiUpdateCacheMut.RLock()
	cache, cached := threadRakugakiUpdateCache[key]
	threadRakugakiUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			threadRakugakiAllColumns,
			threadRakugakiPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update thread_rakugaki, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `thread_rakugaki` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, threadRakugakiPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, append(wl, threadRakugakiPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update thread_rakugaki row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for thread_rakugaki")
	}

	if !cached {
		threadRakugakiUpdateCacheMut.Lock()
		threadRakugakiUpdateCache[key] = cache
		threadRakugakiUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q threadRakugakiQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for thread_rakugaki")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for thread_rakugaki")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ThreadRakugakiSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), threadRakugakiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `thread_rakugaki` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, threadRakugakiPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in threadRakugaki slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all threadRakugaki")
	}
	return rowsAff, nil
}

var mySQLThreadRakugakiUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ThreadRakugaki) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no thread_rakugaki provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(threadRakugakiColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLThreadRakugakiUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	threadRakugakiUpsertCacheMut.RLock()
	cache, cached := threadRakugakiUpsertCache[key]
	threadRakugakiUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			threadRakugakiAllColumns,
			threadRakugakiColumnsWithDefault,
			threadRakugakiColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			threadRakugakiAllColumns,
			threadRakugakiPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert thread_rakugaki, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`thread_rakugaki`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `thread_rakugaki` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for thread_rakugaki")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == threadRakugakiMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(threadRakugakiType, threadRakugakiMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for thread_rakugaki")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for thread_rakugaki")
	}

CacheNoHooks:
	if !cached {
		threadRakugakiUpsertCacheMut.Lock()
		threadRakugakiUpsertCache[key] = cache
		threadRakugakiUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ThreadRakugaki record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ThreadRakugaki) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ThreadRakugaki provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), threadRakugakiPrimaryKeyMapping)
	sql := "DELETE FROM `thread_rakugaki` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from thread_rakugaki")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for thread_rakugaki")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q threadRakugakiQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no threadRakugakiQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thread_rakugaki")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for thread_rakugaki")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ThreadRakugakiSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(threadRakugakiBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), threadRakugakiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `thread_rakugaki` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, threadRakugakiPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from threadRakugaki slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for thread_rakugaki")
	}

	if len(threadRakugakiAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ThreadRakugaki) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindThreadRakugaki(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ThreadRakugakiSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ThreadRakugakiSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), threadRakugakiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `thread_rakugaki`.* FROM `thread_rakugaki` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, threadRakugakiPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ThreadRakugakiSlice")
	}

	*o = slice

	return nil
}

// ThreadRakugakiExists checks if the ThreadRakugaki row exists.
func ThreadRakugakiExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `thread_rakugaki` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if thread_rakugaki exists")
	}

	return exists, nil
}

// Exists checks if the ThreadRakugaki row exists.
func (o *ThreadRakugaki) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ThreadRakugakiExists(ctx, exec, o.ID)
}
