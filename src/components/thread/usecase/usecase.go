package usecase

import (
	infrac "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ThreadUsecase interface {
	Create(t *model.Thread, c *model.Comment) error
}

type threadUsecaseImpl struct {
	tr infra.ThreadRepository
	cr infrac.CommentRepository
}

func NewThreadUsecase(tr infra.ThreadRepository) *threadUsecaseImpl {
	return &threadUsecaseImpl{
		tr: tr,
	}
}

func (tu *threadUsecaseImpl) Create(t *model.Thread, c *model.Comment) error {
	res, err := tu.tr.Create(t)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if err := tu.cr.Create(c, int(id)); err != nil {
		return err
	}

	return nil
}
