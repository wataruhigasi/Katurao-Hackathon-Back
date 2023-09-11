package usecase

import (
	"database/sql"

	comment_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"
	transaction_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/transaction/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type Usecase interface {
	FindAll() ([]*model.Thread, error)
	Create(t *model.Thread, c *model.Comment) error
}

type usecaseImpl struct {
	r  infra.Repo
	cr  comment_infra.Repo
	txr transaction_infra.Repo
}

func New(r infra.Repo, cr comment_infra.Repo) *usecaseImpl {
	return &usecaseImpl{
		r: r,
		cr: cr,
	}
}

func (u *usecaseImpl) FindAll() ([]*model.Thread, error) {
	return u.r.FindAll()
}

func (u *usecaseImpl) Create(t *model.Thread, c *model.Comment) error {
	err := u.txr.Transaction(
		func(tx *sql.Tx) error {
			res, err := u.r.CreateTx(tx, t)
			if err != nil {
				return err
			}

			threadID, err := res.LastInsertId()
			if err != nil {
				return err
			}

			if err := u.cr.CreateTx(tx, c, threadID); err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return err
	}

	return nil
}
