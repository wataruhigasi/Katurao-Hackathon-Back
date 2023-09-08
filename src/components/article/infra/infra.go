package infra

import (
	"context"
	"database/sql"

	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/models"
)

type ArticleRepository interface {
	FindArticles() ([]*model.Article, error)
}

type articleRepositoryImpl struct {
	conn *sql.DB
}

func NewArticleRepository(conn *sql.DB) *articleRepositoryImpl {
	return &articleRepositoryImpl{
		conn: conn,
	}
}

func (ar *articleRepositoryImpl) FindArticles() ([]*model.Article, error) {
	ctx := context.Background()

	dto, err := models.Articles().All(ctx, ar.conn)
	if err != nil {
		return nil, err
	}

	as := make([]*model.Article, len(dto))
	for _, v := range dto {
		a, err := ToArticle(v)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}

	return as, nil
}

func ToArticle(a *models.Article) (*model.Article, error) {
	var p model.Position
	if err := a.Position.Unmarshal(&p); err != nil {
		return nil, err
	}

	return &model.Article{
		ID:        int(a.ID),
		Title:     a.Title,
		CreatedAt: a.CreatedAt,
		Body:      a.Body,
		Position:  p,
	}, nil
}
