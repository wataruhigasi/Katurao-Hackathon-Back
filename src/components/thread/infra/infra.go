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
	FindAll() ([]*model.Thread, error)
	Create(*model.Thread) error
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

func (r *repoImpl) FindAll() ([]*model.Thread, error) {
	ctx := context.Background()

	dto, err := models.Threads().All(ctx, r.conn)
	if err != nil {
		return nil, err
	}

	ts := make([]*model.Thread, 0, len(dto))
	for _, v := range dto {
		t, err := toThread(v)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	return ts, nil
}

func toThread(t *models.Thread) (*model.Thread, error) {
	var p model.Position
	if err := t.Position.Unmarshal(&p); err != nil {
		return nil, err
	}

	return &model.Thread{
		ID:              int(t.ID),
		Title:           t.Title,
		CreatedAt:       t.CreatedAt,
		Position:        p,
		ThreadRakugakis: make([]*model.ThreadRakugaki, 0),
	}, nil
}

func (r *repoImpl) Create(t *model.Thread) error {
	ctx := context.Background()

	pos := &types.JSON{}
	if err := pos.Marshal(t.Position); err != nil {
		return err
	}

	dto := models.Thread{
		Title:    t.Title,
		Position: *pos,
	}

	return dto.Insert(ctx, r.conn, boil.Infer())
}

func (r *repoImpl) ChangePosition(id int64, p *model.Position) error {
	ctx := context.Background()

	t, err := models.FindThread(ctx, r.conn, id)
	if err != nil {
		return err
	}

	json := &types.JSON{}
	if err := json.Marshal(p); err != nil {
		return err
	}

	t.Position = *json

	if _, err := t.Update(ctx, r.conn, boil.Infer()); err != nil {
		return err
	}

	return nil
}
