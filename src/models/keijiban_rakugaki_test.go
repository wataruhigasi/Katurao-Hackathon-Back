// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testKeijibanRakugakis(t *testing.T) {
	t.Parallel()

	query := KeijibanRakugakis()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testKeijibanRakugakisDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
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

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKeijibanRakugakisQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := KeijibanRakugakis().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKeijibanRakugakisSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := KeijibanRakugakiSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKeijibanRakugakisExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := KeijibanRakugakiExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if KeijibanRakugaki exists: %s", err)
	}
	if !e {
		t.Errorf("Expected KeijibanRakugakiExists to return true, but got false.")
	}
}

func testKeijibanRakugakisFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	keijibanRakugakiFound, err := FindKeijibanRakugaki(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if keijibanRakugakiFound == nil {
		t.Error("want a record, got nil")
	}
}

func testKeijibanRakugakisBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = KeijibanRakugakis().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testKeijibanRakugakisOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := KeijibanRakugakis().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testKeijibanRakugakisAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	keijibanRakugakiOne := &KeijibanRakugaki{}
	keijibanRakugakiTwo := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, keijibanRakugakiOne, keijibanRakugakiDBTypes, false, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}
	if err = randomize.Struct(seed, keijibanRakugakiTwo, keijibanRakugakiDBTypes, false, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = keijibanRakugakiOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = keijibanRakugakiTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := KeijibanRakugakis().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testKeijibanRakugakisCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	keijibanRakugakiOne := &KeijibanRakugaki{}
	keijibanRakugakiTwo := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, keijibanRakugakiOne, keijibanRakugakiDBTypes, false, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}
	if err = randomize.Struct(seed, keijibanRakugakiTwo, keijibanRakugakiDBTypes, false, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = keijibanRakugakiOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = keijibanRakugakiTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func keijibanRakugakiBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func keijibanRakugakiAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *KeijibanRakugaki) error {
	*o = KeijibanRakugaki{}
	return nil
}

func testKeijibanRakugakisHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &KeijibanRakugaki{}
	o := &KeijibanRakugaki{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, false); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki object: %s", err)
	}

	AddKeijibanRakugakiHook(boil.BeforeInsertHook, keijibanRakugakiBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiBeforeInsertHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.AfterInsertHook, keijibanRakugakiAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiAfterInsertHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.AfterSelectHook, keijibanRakugakiAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiAfterSelectHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.BeforeUpdateHook, keijibanRakugakiBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiBeforeUpdateHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.AfterUpdateHook, keijibanRakugakiAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiAfterUpdateHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.BeforeDeleteHook, keijibanRakugakiBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiBeforeDeleteHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.AfterDeleteHook, keijibanRakugakiAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiAfterDeleteHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.BeforeUpsertHook, keijibanRakugakiBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiBeforeUpsertHooks = []KeijibanRakugakiHook{}

	AddKeijibanRakugakiHook(boil.AfterUpsertHook, keijibanRakugakiAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	keijibanRakugakiAfterUpsertHooks = []KeijibanRakugakiHook{}
}

func testKeijibanRakugakisInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testKeijibanRakugakisInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(keijibanRakugakiColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testKeijibanRakugakisReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
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

func testKeijibanRakugakisReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := KeijibanRakugakiSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testKeijibanRakugakisSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := KeijibanRakugakis().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	keijibanRakugakiDBTypes = map[string]string{`ID`: `bigint`, `Body`: `varchar`, `Position`: `json`, `CreatedAt`: `datetime`}
	_                       = bytes.MinRead
)

func testKeijibanRakugakisUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(keijibanRakugakiPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(keijibanRakugakiAllColumns) == len(keijibanRakugakiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testKeijibanRakugakisSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(keijibanRakugakiAllColumns) == len(keijibanRakugakiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &KeijibanRakugaki{}
	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, keijibanRakugakiDBTypes, true, keijibanRakugakiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(keijibanRakugakiAllColumns, keijibanRakugakiPrimaryKeyColumns) {
		fields = keijibanRakugakiAllColumns
	} else {
		fields = strmangle.SetComplement(
			keijibanRakugakiAllColumns,
			keijibanRakugakiPrimaryKeyColumns,
		)
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

	slice := KeijibanRakugakiSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testKeijibanRakugakisUpsert(t *testing.T) {
	t.Parallel()

	if len(keijibanRakugakiAllColumns) == len(keijibanRakugakiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLKeijibanRakugakiUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := KeijibanRakugaki{}
	if err = randomize.Struct(seed, &o, keijibanRakugakiDBTypes, false); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert KeijibanRakugaki: %s", err)
	}

	count, err := KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, keijibanRakugakiDBTypes, false, keijibanRakugakiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KeijibanRakugaki struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert KeijibanRakugaki: %s", err)
	}

	count, err = KeijibanRakugakis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}