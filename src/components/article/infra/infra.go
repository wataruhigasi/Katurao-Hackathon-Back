package infra

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/models"
)

type Repo interface {
	FindAll() ([]*model.Article, error)
	Create(*model.Article) error
	ChangePosition(int64, *model.Position) error
}

type repoImpl struct {
	conn *sql.DB
}

func NewRepo(conn *sql.DB) *repoImpl {
	return &repoImpl{
		conn: conn,
	}
}

func (r *repoImpl) FindAll() ([]*model.Article, error) {
	ctx := context.Background()

	dto, err := models.Articles().All(ctx, r.conn)
	if err != nil {
		return nil, err
	}

	as := make([]*model.Article, 0, len(dto))
	for _, v := range dto {
		a, err := toArticle(v)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}

	return as, nil
}

func toArticle(a *models.Article) (*model.Article, error) {
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

func (r *repoImpl) Create(a *model.Article) error {
	ctx := context.Background()

	json := &types.JSON{}
	if err := json.Marshal(a.Position); err != nil {
		return err
	}

	// INSERT INTO articles (title, created_at, body, position) VALUES (?, ?, ?, ?)
	dto := models.Article{
		ID:       int64(a.ID),
		Title:    a.Title,
		Body:     a.Body,
		Position: *json,
	}

	return dto.Insert(ctx, r.conn, boil.Infer())
}

func (r *repoImpl) ChangePosition(id int64, p *model.Position) error {
	ctx := context.Background()

	a, err :=models.FindArticle(ctx, r.conn, id)
	if err != nil {
		return err
	}

	json := &types.JSON{}
	if err := json.Marshal(p); err != nil {
		return err
	}

	a.Position = *json

	if _, err := a.Update(ctx, r.conn, boil.Infer()); err != nil {
		return err
	}

	return nil
}
