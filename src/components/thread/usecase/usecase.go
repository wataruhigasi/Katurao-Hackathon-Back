package usecase

import (
	"database/sql"

	infrac "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"
	transaction_infra "github.com/wataruhigasi/Katurao-Hackathon-Back/components/transaction/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ThreadUsecase interface {
	FindAll() ([]*model.Thread, error)
	Create(t *model.Thread, c *model.Comment) error
}

type threadUsecaseImpl struct {
	tr  infra.ThreadRepository
	cr  infrac.CommentRepository
	txr transaction_infra.Repo
}

func New(tr infra.ThreadRepository, cr infrac.CommentRepository, txr transaction_infra.Repo) *threadUsecaseImpl {
	return &threadUsecaseImpl{
		tr: tr,
		cr: cr,
		txr: txr,
	}
}

func (tu *threadUsecaseImpl) FindAll() ([]*model.Thread, error) {
	return tu.tr.FindAll()
}

func (tu *threadUsecaseImpl) Create(t *model.Thread, c *model.Comment) error {
	err := tu.txr.Transaction(
		func(tx *sql.Tx) error {
			res, err := tu.tr.CreateTx(tx, t)
			if err != nil {
				return err
			}

			threadID, err := res.LastInsertId()
			if err != nil {
				return err
			}

			if err := tu.cr.CreateTx(tx, c, threadID); err != nil {
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
