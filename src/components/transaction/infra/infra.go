package infra

import (
	"database/sql"
)

type Repo interface {
	Transaction(f func(tx *sql.Tx) error) error
}

var _ Repo = (*repoImpl)(nil)

type repoImpl struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *repoImpl {
	return &repoImpl{
		db: db,
	}
}

func (r *repoImpl) Transaction(f func(tx *sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err := f(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
