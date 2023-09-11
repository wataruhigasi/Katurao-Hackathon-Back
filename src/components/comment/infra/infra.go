package infra

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/models"
)

type CommentRepo interface {
	FindAll(int) ([]*model.Comment, error)
	Create(*model.Comment, int64) error
	CreateTx(*sql.Tx, *model.Comment, int64) error
}

type commentRepoImpl struct {
	conn *sql.DB
}

func NewRepo(conn *sql.DB) *commentRepoImpl {
	return &commentRepoImpl{
		conn: conn,
	}
}

func (cr *commentRepoImpl) FindAll(id int) ([]*model.Comment, error) {
	ctx := context.Background()

	dto, err := models.Comments(qm.Where("thread_id = ?", id)).All(ctx, cr.conn)

	if err != nil {
		return nil, err
	}

	cs := make([]*model.Comment, 0, len(dto))
	for _, v := range dto {
		c, err := ToComment(v)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}

	return cs, nil
}

func ToComment(c *models.Comment) (*model.Comment, error) {

	return &model.Comment{
		ID:        int(c.ID),
		CreatedAt: c.CreatedAt,
		Body:      c.Body,
		Author:    c.Author,
	}, nil
}

func (cr *commentRepoImpl) Create(c *model.Comment, threadID int64) error {
	ctx := context.Background()

	dto := models.Comment{
		ID:       int64(c.ID),
		Body:     c.Body,
		Author:   c.Author,
		ThreadID: threadID,
	}

	return dto.Insert(ctx, cr.conn, boil.Infer())
}

func (cr *commentRepoImpl) CreateTx(tx *sql.Tx, c *model.Comment, threadID int64) error {
	ctx := context.Background()

	dto := models.Comment{
		ID:       int64(c.ID),
		Body:     c.Body,
		Author:   c.Author,
		ThreadID: threadID,
	}

	return dto.Insert(ctx, tx, boil.Infer())
}
