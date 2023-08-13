// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testHosts(t *testing.T) {
	t.Parallel()

	query := Hosts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHostsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHostsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Hosts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHostsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HostSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHostsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HostExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Host exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HostExists to return true, but got false.")
	}
}

func testHostsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	hostFound, err := FindHost(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if hostFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHostsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Hosts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHostsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Hosts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHostsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	hostOne := &Host{}
	hostTwo := &Host{}
	if err = randomize.Struct(seed, hostOne, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}
	if err = randomize.Struct(seed, hostTwo, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hostOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hostTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hosts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHostsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	hostOne := &Host{}
	hostTwo := &Host{}
	if err = randomize.Struct(seed, hostOne, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}
	if err = randomize.Struct(seed, hostTwo, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hostOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hostTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func hostBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func testHostsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Host{}
	o := &Host{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, hostDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Host object: %s", err)
	}

	AddHostHook(boil.BeforeInsertHook, hostBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	hostBeforeInsertHooks = []HostHook{}

	AddHostHook(boil.AfterInsertHook, hostAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	hostAfterInsertHooks = []HostHook{}

	AddHostHook(boil.AfterSelectHook, hostAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	hostAfterSelectHooks = []HostHook{}

	AddHostHook(boil.BeforeUpdateHook, hostBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	hostBeforeUpdateHooks = []HostHook{}

	AddHostHook(boil.AfterUpdateHook, hostAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	hostAfterUpdateHooks = []HostHook{}

	AddHostHook(boil.BeforeDeleteHook, hostBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	hostBeforeDeleteHooks = []HostHook{}

	AddHostHook(boil.AfterDeleteHook, hostAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	hostAfterDeleteHooks = []HostHook{}

	AddHostHook(boil.BeforeUpsertHook, hostBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	hostBeforeUpsertHooks = []HostHook{}

	AddHostHook(boil.AfterUpsertHook, hostAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	hostAfterUpsertHooks = []HostHook{}
}

func testHostsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHostsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(hostColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHostToManySgroups(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c Sgroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, sgroupDBTypes, false, sgroupColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sgroupDBTypes, false, sgroupColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"host_sgroup\" (\"host_id\", \"sgroup_id\") values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"host_sgroup\" (\"host_id\", \"sgroup_id\") values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Sgroups().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ID, b.ID) {
			bFound = true
		}
		if queries.Equal(v.ID, c.ID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HostSlice{&a}
	if err = a.L.LoadSgroups(ctx, tx, false, (*[]*Host)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sgroups); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Sgroups = nil
	if err = a.L.LoadSgroups(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sgroups); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHostToManyAddOpSgroups(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c, d, e Sgroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, false, strmangle.SetComplement(hostPrimaryKeyColumns, hostColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Sgroup{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Sgroup{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSgroups(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Hosts[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Hosts[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Sgroups[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Sgroups[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Sgroups().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testHostToManySetOpSgroups(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c, d, e Sgroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, false, strmangle.SetComplement(hostPrimaryKeyColumns, hostColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Sgroup{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSgroups(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Sgroups().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSgroups(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Sgroups().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Hosts) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Hosts) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Hosts[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Hosts[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Sgroups[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Sgroups[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testHostToManyRemoveOpSgroups(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c, d, e Sgroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, false, strmangle.SetComplement(hostPrimaryKeyColumns, hostColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Sgroup{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSgroups(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Sgroups().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSgroups(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Sgroups().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Hosts) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Hosts) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Hosts[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Hosts[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Sgroups) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Sgroups[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Sgroups[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testHostsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHostsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HostSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHostsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hosts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	hostDBTypes = map[string]string{`ID`: `INTEGER`, `Hostname`: `TEXT`, `Ipaddr`: `TEXT`, `Desc`: `TEXT`}
	_           = bytes.MinRead
)

func testHostsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(hostAllColumns) == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hostDBTypes, true, hostPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHostsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(hostAllColumns) == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hostDBTypes, true, hostPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(hostAllColumns, hostPrimaryKeyColumns) {
		fields = hostAllColumns
	} else {
		fields = strmangle.SetComplement(
			hostAllColumns,
			hostPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, hostGeneratedColumns)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := HostSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHostsUpsert(t *testing.T) {
	t.Parallel()
	if len(hostAllColumns) == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Host{}
	if err = randomize.Struct(seed, &o, hostDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Host: %s", err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, hostDBTypes, false, hostPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Host: %s", err)
	}

	count, err = Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}