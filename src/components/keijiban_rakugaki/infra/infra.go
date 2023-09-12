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
	FindAll() ([]*model.KeijibanRakugaki, error)
	Create(*model.KeijibanRakugaki) error
}

type repoImpl struct {
	conn *sql.DB
}

func NewRepo(conn *sql.DB) *repoImpl {
	return &repoImpl{
		conn: conn,
	}
}

func (r *repoImpl) FindAll() ([]*model.KeijibanRakugaki, error) {
	ctx := context.Background()

	dto, err := models.KeijibanRakugakis().All(ctx, r.conn)
	if err != nil {
		return nil, err
	}

	as := make([]*model.KeijibanRakugaki, 0, len(dto))
	for _, v := range dto {
		a, err := toKeijibanRakugakis(v)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}

	return as, nil
}

func toKeijibanRakugakis(a *models.KeijibanRakugaki) (*model.KeijibanRakugaki, error) {
	var p model.Position
	if err := a.Position.Unmarshal(&p); err != nil {
		return nil, err
	}

	return &model.KeijibanRakugaki{
		ID:        int(a.ID),
		CreatedAt: a.CreatedAt,
		Body:      a.Body,
		Position:  p,
	}, nil
}

func (r *repoImpl) Create(a *model.KeijibanRakugaki) error {
	ctx := context.Background()

	json := &types.JSON{}
	if err := json.Marshal(a.Position); err != nil {
		return err
	}

	dto := models.KeijibanRakugaki{
		ID:       int64(a.ID),
		Body:     a.Body,
		Position: *json,
	}

	return dto.Insert(ctx, r.conn, boil.Infer())
}
