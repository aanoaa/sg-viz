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

func testPolicies(t *testing.T) {
	t.Parallel()

	query := Policies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPoliciesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
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

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPoliciesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Policies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPoliciesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PolicySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPoliciesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PolicyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Policy exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PolicyExists to return true, but got false.")
	}
}

func testPoliciesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	policyFound, err := FindPolicy(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if policyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPoliciesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Policies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPoliciesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Policies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPoliciesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	policyOne := &Policy{}
	policyTwo := &Policy{}
	if err = randomize.Struct(seed, policyOne, policyDBTypes, false, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}
	if err = randomize.Struct(seed, policyTwo, policyDBTypes, false, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = policyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = policyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Policies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPoliciesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	policyOne := &Policy{}
	policyTwo := &Policy{}
	if err = randomize.Struct(seed, policyOne, policyDBTypes, false, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}
	if err = randomize.Struct(seed, policyTwo, policyDBTypes, false, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = policyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = policyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func policyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func policyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Policy) error {
	*o = Policy{}
	return nil
}

func testPoliciesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Policy{}
	o := &Policy{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, policyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Policy object: %s", err)
	}

	AddPolicyHook(boil.BeforeInsertHook, policyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	policyBeforeInsertHooks = []PolicyHook{}

	AddPolicyHook(boil.AfterInsertHook, policyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	policyAfterInsertHooks = []PolicyHook{}

	AddPolicyHook(boil.AfterSelectHook, policyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	policyAfterSelectHooks = []PolicyHook{}

	AddPolicyHook(boil.BeforeUpdateHook, policyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	policyBeforeUpdateHooks = []PolicyHook{}

	AddPolicyHook(boil.AfterUpdateHook, policyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	policyAfterUpdateHooks = []PolicyHook{}

	AddPolicyHook(boil.BeforeDeleteHook, policyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	policyBeforeDeleteHooks = []PolicyHook{}

	AddPolicyHook(boil.AfterDeleteHook, policyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	policyAfterDeleteHooks = []PolicyHook{}

	AddPolicyHook(boil.BeforeUpsertHook, policyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	policyBeforeUpsertHooks = []PolicyHook{}

	AddPolicyHook(boil.AfterUpsertHook, policyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	policyAfterUpsertHooks = []PolicyHook{}
}

func testPoliciesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPoliciesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(policyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPolicyToOneSgroupUsingToSgroup(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Policy
	var foreign Sgroup

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, policyDBTypes, false, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.To, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ToSgroup().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddSgroupHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := PolicySlice{&local}
	if err = local.L.LoadToSgroup(ctx, tx, false, (*[]*Policy)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ToSgroup == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ToSgroup = nil
	if err = local.L.LoadToSgroup(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ToSgroup == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testPolicyToOneSgroupUsingFromSgroup(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Policy
	var foreign Sgroup

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, policyDBTypes, false, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.From, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FromSgroup().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddSgroupHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := PolicySlice{&local}
	if err = local.L.LoadFromSgroup(ctx, tx, false, (*[]*Policy)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromSgroup == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FromSgroup = nil
	if err = local.L.LoadFromSgroup(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromSgroup == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testPolicyToOneSetOpSgroupUsingToSgroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Policy
	var b, c Sgroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, policyDBTypes, false, strmangle.SetComplement(policyPrimaryKeyColumns, policyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Sgroup{&b, &c} {
		err = a.SetToSgroup(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ToSgroup != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ToPolicies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.To, x.ID) {
			t.Error("foreign key was wrong value", a.To)
		}

		zero := reflect.Zero(reflect.TypeOf(a.To))
		reflect.Indirect(reflect.ValueOf(&a.To)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.To, x.ID) {
			t.Error("foreign key was wrong value", a.To, x.ID)
		}
	}
}
func testPolicyToOneSetOpSgroupUsingFromSgroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Policy
	var b, c Sgroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, policyDBTypes, false, strmangle.SetComplement(policyPrimaryKeyColumns, policyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sgroupDBTypes, false, strmangle.SetComplement(sgroupPrimaryKeyColumns, sgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Sgroup{&b, &c} {
		err = a.SetFromSgroup(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FromSgroup != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FromPolicies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.From, x.ID) {
			t.Error("foreign key was wrong value", a.From)
		}

		zero := reflect.Zero(reflect.TypeOf(a.From))
		reflect.Indirect(reflect.ValueOf(&a.From)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.From, x.ID) {
			t.Error("foreign key was wrong value", a.From, x.ID)
		}
	}
}

func testPoliciesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
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

func testPoliciesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PolicySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPoliciesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Policies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	policyDBTypes = map[string]string{`ID`: `INTEGER`, `From`: `INTEGER`, `To`: `INTEGER`, `Port`: `INTEGER`, `Protocol`: `TEXT`, `Desc`: `TEXT`}
	_             = bytes.MinRead
)

func testPoliciesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(policyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(policyAllColumns) == len(policyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, policyDBTypes, true, policyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPoliciesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(policyAllColumns) == len(policyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Policy{}
	if err = randomize.Struct(seed, o, policyDBTypes, true, policyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, policyDBTypes, true, policyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(policyAllColumns, policyPrimaryKeyColumns) {
		fields = policyAllColumns
	} else {
		fields = strmangle.SetComplement(
			policyAllColumns,
			policyPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, policyGeneratedColumns)
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

	slice := PolicySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPoliciesUpsert(t *testing.T) {
	t.Parallel()
	if len(policyAllColumns) == len(policyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Policy{}
	if err = randomize.Struct(seed, &o, policyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Policy: %s", err)
	}

	count, err := Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, policyDBTypes, false, policyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Policy struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Policy: %s", err)
	}

	count, err = Policies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
