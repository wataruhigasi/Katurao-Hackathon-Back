package infra

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/models"
)

type CommentRepository interface {
	FindAll(int) ([]*model.Comment, error)
	Create(*model.Comment, int) error
}

type commentRepositoryImpl struct {
	conn *sql.DB
}

func NewCommentRepository(conn *sql.DB) *commentRepositoryImpl {
	return &commentRepositoryImpl{
		conn: conn,
	}
}

func (cr *commentRepositoryImpl) FindAll(id int) ([]*model.Comment, error) {
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

func (cr *commentRepositoryImpl) Create(c *model.Comment, id int) error {
	ctx := context.Background()

	dto := models.Comment{
		ID:       int64(c.ID),
		Body:     c.Body,
		Author:   c.Author,
		ThreadID: int64(id),
	}

	return dto.Insert(ctx, cr.conn, boil.Infer())
}
