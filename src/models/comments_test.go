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

func testComments(t *testing.T) {
	t.Parallel()

	query := Comments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCommentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
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

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Comments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CommentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CommentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Comment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CommentExists to return true, but got false.")
	}
}

func testCommentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	commentFound, err := FindComment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if commentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCommentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Comments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCommentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Comments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCommentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	commentOne := &Comment{}
	commentTwo := &Comment{}
	if err = randomize.Struct(seed, commentOne, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err = randomize.Struct(seed, commentTwo, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = commentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = commentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Comments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCommentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	commentOne := &Comment{}
	commentTwo := &Comment{}
	if err = randomize.Struct(seed, commentOne, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err = randomize.Struct(seed, commentTwo, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = commentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = commentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func commentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func testCommentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Comment{}
	o := &Comment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, commentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Comment object: %s", err)
	}

	AddCommentHook(boil.BeforeInsertHook, commentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	commentBeforeInsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterInsertHook, commentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	commentAfterInsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterSelectHook, commentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	commentAfterSelectHooks = []CommentHook{}

	AddCommentHook(boil.BeforeUpdateHook, commentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	commentBeforeUpdateHooks = []CommentHook{}

	AddCommentHook(boil.AfterUpdateHook, commentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	commentAfterUpdateHooks = []CommentHook{}

	AddCommentHook(boil.BeforeDeleteHook, commentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	commentBeforeDeleteHooks = []CommentHook{}

	AddCommentHook(boil.AfterDeleteHook, commentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	commentAfterDeleteHooks = []CommentHook{}

	AddCommentHook(boil.BeforeUpsertHook, commentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	commentBeforeUpsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterUpsertHook, commentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	commentAfterUpsertHooks = []CommentHook{}
}

func testCommentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(commentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
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

func testCommentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CommentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCommentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Comments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	commentDBTypes = map[string]string{`ID`: `bigint`, `ArticleID`: `bigint`, `Body`: `varchar`, `Author`: `varchar`, `CreatedAt`: `datetime`}
	_              = bytes.MinRead
)

func testCommentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(commentAllColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, commentDBTypes, true, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCommentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(commentAllColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, commentDBTypes, true, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(commentAllColumns, commentPrimaryKeyColumns) {
		fields = commentAllColumns
	} else {
		fields = strmangle.SetComplement(
			commentAllColumns,
			commentPrimaryKeyColumns,
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

	slice := CommentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCommentsUpsert(t *testing.T) {
	t.Parallel()

	if len(commentAllColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCommentUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Comment{}
	if err = randomize.Struct(seed, &o, commentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Comment: %s", err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, commentDBTypes, false, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Comment: %s", err)
	}

	count, err = Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
