// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Host is an object representing the database table.
type Host struct {
	ID       null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Hostname string     `boil:"hostname" json:"hostname" toml:"hostname" yaml:"hostname"`
	Ipaddr   string     `boil:"ipaddr" json:"ipaddr" toml:"ipaddr" yaml:"ipaddr"`
	Desc     string     `boil:"desc" json:"desc" toml:"desc" yaml:"desc"`

	R *hostR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hostL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HostColumns = struct {
	ID       string
	Hostname string
	Ipaddr   string
	Desc     string
}{
	ID:       "id",
	Hostname: "hostname",
	Ipaddr:   "ipaddr",
	Desc:     "desc",
}

var HostTableColumns = struct {
	ID       string
	Hostname string
	Ipaddr   string
	Desc     string
}{
	ID:       "host.id",
	Hostname: "host.hostname",
	Ipaddr:   "host.ipaddr",
	Desc:     "host.desc",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var HostWhere = struct {
	ID       whereHelpernull_Int64
	Hostname whereHelperstring
	Ipaddr   whereHelperstring
	Desc     whereHelperstring
}{
	ID:       whereHelpernull_Int64{field: "\"host\".\"id\""},
	Hostname: whereHelperstring{field: "\"host\".\"hostname\""},
	Ipaddr:   whereHelperstring{field: "\"host\".\"ipaddr\""},
	Desc:     whereHelperstring{field: "\"host\".\"desc\""},
}

// HostRels is where relationship names are stored.
var HostRels = struct {
	Sgroups string
}{
	Sgroups: "Sgroups",
}

// hostR is where relationships are stored.
type hostR struct {
	Sgroups SgroupSlice `boil:"Sgroups" json:"Sgroups" toml:"Sgroups" yaml:"Sgroups"`
}

// NewStruct creates a new relationship struct
func (*hostR) NewStruct() *hostR {
	return &hostR{}
}

func (r *hostR) GetSgroups() SgroupSlice {
	if r == nil {
		return nil
	}
	return r.Sgroups
}

// hostL is where Load methods for each relationship are stored.
type hostL struct{}

var (
	hostAllColumns            = []string{"id", "hostname", "ipaddr", "desc"}
	hostColumnsWithoutDefault = []string{"hostname", "ipaddr"}
	hostColumnsWithDefault    = []string{"id", "desc"}
	hostPrimaryKeyColumns     = []string{"id"}
	hostGeneratedColumns      = []string{"id"}
)

type (
	// HostSlice is an alias for a slice of pointers to Host.
	// This should almost always be used instead of []Host.
	HostSlice []*Host
	// HostHook is the signature for custom Host hook methods
	HostHook func(context.Context, boil.ContextExecutor, *Host) error

	hostQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hostType                 = reflect.TypeOf(&Host{})
	hostMapping              = queries.MakeStructMapping(hostType)
	hostPrimaryKeyMapping, _ = queries.BindMapping(hostType, hostMapping, hostPrimaryKeyColumns)
	hostInsertCacheMut       sync.RWMutex
	hostInsertCache          = make(map[string]insertCache)
	hostUpdateCacheMut       sync.RWMutex
	hostUpdateCache          = make(map[string]updateCache)
	hostUpsertCacheMut       sync.RWMutex
	hostUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hostAfterSelectHooks []HostHook

var hostBeforeInsertHooks []HostHook
var hostAfterInsertHooks []HostHook

var hostBeforeUpdateHooks []HostHook
var hostAfterUpdateHooks []HostHook

var hostBeforeDeleteHooks []HostHook
var hostAfterDeleteHooks []HostHook

var hostBeforeUpsertHooks []HostHook
var hostAfterUpsertHooks []HostHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Host) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Host) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Host) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Host) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Host) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Host) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Host) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Host) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Host) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHostHook registers your hook function for all future operations.
func AddHostHook(hookPoint boil.HookPoint, hostHook HostHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		hostAfterSelectHooks = append(hostAfterSelectHooks, hostHook)
	case boil.BeforeInsertHook:
		hostBeforeInsertHooks = append(hostBeforeInsertHooks, hostHook)
	case boil.AfterInsertHook:
		hostAfterInsertHooks = append(hostAfterInsertHooks, hostHook)
	case boil.BeforeUpdateHook:
		hostBeforeUpdateHooks = append(hostBeforeUpdateHooks, hostHook)
	case boil.AfterUpdateHook:
		hostAfterUpdateHooks = append(hostAfterUpdateHooks, hostHook)
	case boil.BeforeDeleteHook:
		hostBeforeDeleteHooks = append(hostBeforeDeleteHooks, hostHook)
	case boil.AfterDeleteHook:
		hostAfterDeleteHooks = append(hostAfterDeleteHooks, hostHook)
	case boil.BeforeUpsertHook:
		hostBeforeUpsertHooks = append(hostBeforeUpsertHooks, hostHook)
	case boil.AfterUpsertHook:
		hostAfterUpsertHooks = append(hostAfterUpsertHooks, hostHook)
	}
}

// One returns a single host record from the query.
func (q hostQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Host, error) {
	o := &Host{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for host")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Host records from the query.
func (q hostQuery) All(ctx context.Context, exec boil.ContextExecutor) (HostSlice, error) {
	var o []*Host

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Host slice")
	}

	if len(hostAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Host records in the query.
func (q hostQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count host rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q hostQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if host exists")
	}

	return count > 0, nil
}

// Sgroups retrieves all the sgroup's Sgroups with an executor.
func (o *Host) Sgroups(mods ...qm.QueryMod) sgroupQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"host_sgroup\" on \"sgroup\".\"id\" = \"host_sgroup\".\"sgroup_id\""),
		qm.Where("\"host_sgroup\".\"host_id\"=?", o.ID),
	)

	return Sgroups(queryMods...)
}

// LoadSgroups allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (hostL) LoadSgroups(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHost interface{}, mods queries.Applicator) error {
	var slice []*Host
	var object *Host

	if singular {
		var ok bool
		object, ok = maybeHost.(*Host)
		if !ok {
			object = new(Host)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeHost)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeHost))
			}
		}
	} else {
		s, ok := maybeHost.(*[]*Host)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeHost)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeHost))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &hostR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hostR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"sgroup\".\"id\", \"sgroup\".\"name\", \"sgroup\".\"desc\", \"a\".\"host_id\""),
		qm.From("\"sgroup\""),
		qm.InnerJoin("\"host_sgroup\" as \"a\" on \"sgroup\".\"id\" = \"a\".\"sgroup_id\""),
		qm.WhereIn("\"a\".\"host_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sgroup")
	}

	var resultSlice []*Sgroup

	var localJoinCols []int64
	for results.Next() {
		one := new(Sgroup)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.Name, &one.Desc, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for sgroup")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice sgroup")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sgroup")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sgroup")
	}

	if len(sgroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Sgroups = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sgroupR{}
			}
			foreign.R.Hosts = append(foreign.R.Hosts, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if queries.Equal(local.ID, localJoinCol) {
				local.R.Sgroups = append(local.R.Sgroups, foreign)
				if foreign.R == nil {
					foreign.R = &sgroupR{}
				}
				foreign.R.Hosts = append(foreign.R.Hosts, local)
				break
			}
		}
	}

	return nil
}

// AddSgroups adds the given related objects to the existing relationships
// of the host, optionally inserting them as new records.
// Appends related to o.R.Sgroups.
// Sets related.R.Hosts appropriately.
func (o *Host) AddSgroups(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Sgroup) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"host_sgroup\" (\"host_id\", \"sgroup_id\") values (?, ?)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &hostR{
			Sgroups: related,
		}
	} else {
		o.R.Sgroups = append(o.R.Sgroups, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sgroupR{
				Hosts: HostSlice{o},
			}
		} else {
			rel.R.Hosts = append(rel.R.Hosts, o)
		}
	}
	return nil
}

// SetSgroups removes all previously related items of the
// host replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Hosts's Sgroups accordingly.
// Replaces o.R.Sgroups with related.
// Sets related.R.Hosts's Sgroups accordingly.
func (o *Host) SetSgroups(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Sgroup) error {
	query := "delete from \"host_sgroup\" where \"host_id\" = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeSgroupsFromHostsSlice(o, related)
	if o.R != nil {
		o.R.Sgroups = nil
	}

	return o.AddSgroups(ctx, exec, insert, related...)
}

// RemoveSgroups relationships from objects passed in.
// Removes related items from R.Sgroups (uses pointer comparison, removal does not keep order)
// Sets related.R.Hosts.
func (o *Host) RemoveSgroups(ctx context.Context, exec boil.ContextExecutor, related ...*Sgroup) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"host_sgroup\" where \"host_id\" = ? and \"sgroup_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeSgroupsFromHostsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Sgroups {
			if rel != ri {
				continue
			}

			ln := len(o.R.Sgroups)
			if ln > 1 && i < ln-1 {
				o.R.Sgroups[i] = o.R.Sgroups[ln-1]
			}
			o.R.Sgroups = o.R.Sgroups[:ln-1]
			break
		}
	}

	return nil
}

func removeSgroupsFromHostsSlice(o *Host, related []*Sgroup) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Hosts {
			if !queries.Equal(o.ID, ri.ID) {
				continue
			}

			ln := len(rel.R.Hosts)
			if ln > 1 && i < ln-1 {
				rel.R.Hosts[i] = rel.R.Hosts[ln-1]
			}
			rel.R.Hosts = rel.R.Hosts[:ln-1]
			break
		}
	}
}

// Hosts retrieves all the records using an executor.
func Hosts(mods ...qm.QueryMod) hostQuery {
	mods = append(mods, qm.From("\"host\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"host\".*"})
	}

	return hostQuery{q}
}

// FindHost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHost(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Host, error) {
	hostObj := &Host{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"host\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, hostObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from host")
	}

	if err = hostObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hostObj, err
	}

	return hostObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Host) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no host provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hostInsertCacheMut.RLock()
	cache, cached := hostInsertCache[key]
	hostInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hostAllColumns,
			hostColumnsWithDefault,
			hostColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, hostGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(hostType, hostMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hostType, hostMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"host\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"host\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into host")
	}

	if !cached {
		hostInsertCacheMut.Lock()
		hostInsertCache[key] = cache
		hostInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Host.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Host) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hostUpdateCacheMut.RLock()
	cache, cached := hostUpdateCache[key]
	hostUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hostAllColumns,
			hostPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, hostGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update host, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"host\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, hostPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hostType, hostMapping, append(wl, hostPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update host row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for host")
	}

	if !cached {
		hostUpdateCacheMut.Lock()
		hostUpdateCache[key] = cache
		hostUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q hostQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for host")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for host")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"host\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, hostPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in host slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all host")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Host) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no host provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	hostUpsertCacheMut.RLock()
	cache, cached := hostUpsertCache[key]
	hostUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			hostAllColumns,
			hostColumnsWithDefault,
			hostColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			hostAllColumns,
			hostPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert host, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(hostPrimaryKeyColumns))
			copy(conflict, hostPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"host\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(hostType, hostMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hostType, hostMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert host")
	}

	if !cached {
		hostUpsertCacheMut.Lock()
		hostUpsertCache[key] = cache
		hostUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Host record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Host) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Host provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hostPrimaryKeyMapping)
	sql := "DELETE FROM \"host\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from host")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for host")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q hostQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hostQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from host")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for host")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hostBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"host\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, hostPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from host slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for host")
	}

	if len(hostAfterDeleteHooks) != 0 {
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
func (o *Host) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"host\".* FROM \"host\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, hostPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HostSlice")
	}

	*o = slice

	return nil
}

// HostExists checks if the Host row exists.
func HostExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"host\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if host exists")
	}

	return exists, nil
}

// Exists checks if the Host row exists.
func (o *Host) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HostExists(ctx, exec, o.ID)
}