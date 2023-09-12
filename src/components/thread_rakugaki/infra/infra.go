package infra

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/models"
)

type Repo interface {
	FindAll(int64) ([]*model.ThreadRakugaki, error)
	Create(*model.ThreadRakugaki, int64) error
}

type repoImpl struct {
	conn *sql.DB
}

func NewRepo(conn *sql.DB) *repoImpl {
	return &repoImpl{
		conn: conn,
	}
}

func (r *repoImpl) FindAll(id int64) ([]*model.ThreadRakugaki, error) {
	ctx := context.Background()

	dto, err := models.ThreadRakugakis(qm.Where("thread_id = ?", id)).All(ctx, r.conn)
	if err != nil {
		return nil, err
	}

	trs := make([]*model.ThreadRakugaki, 0, len(dto))
	for _, v := range dto {
		a, err := toThreadRakugakis(v)
		if err != nil {
			return nil, err
		}
		trs = append(trs, a)
	}

	return trs, nil
}

func toThreadRakugakis(a *models.ThreadRakugaki) (*model.ThreadRakugaki, error) {
	var p model.Position
	if err := a.Position.Unmarshal(&p); err != nil {
		return nil, err
	}

	return &model.ThreadRakugaki{
		ID:        int(a.ID),
		ThreadID:  int(a.ThreadID),
		CreatedAt: a.CreatedAt,
		Body:      a.Body,
		Position:  p,
	}, nil
}

func (r *repoImpl) Create(a *model.ThreadRakugaki, id int64) error {
	ctx := context.Background()

	json := &types.JSON{}
	if err := json.Marshal(a.Position); err != nil {
		return err
	}

	dto := models.ThreadRakugaki{
		ID:       int64(a.ID),
		ThreadID: id,
		Body:     a.Body,
		Position: *json,
	}

	return dto.Insert(ctx, r.conn, boil.Infer())
}
