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
	Create(*model.Thread) (sql.Result, error)
	CreateTx(*sql.Tx, *model.Thread) (sql.Result, error)
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
		ID:        int(t.ID),
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		Position:  p,
		Comments:  nil,
	}, nil
}

func (r *repoImpl) Create(t *model.Thread) (sql.Result, error) {
	pos := &types.JSON{}
	if err := pos.Marshal(t.Position); err != nil {
		return nil, err
	}

	dto := models.Thread{
		Title:    t.Title,
		Position: *pos,
	}

	// ここで、ORMを使っていないのはsql.Resultを返して、そのメソッドのLastInsertID()を使うため
	// Exec()によって、sql.Resultが返るがsqlboilerのInsertでは返さない

	p, err := r.conn.Prepare("INSERT INTO threads (title, position) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := p.Exec(dto.Title, *pos)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repoImpl) CreateTx(tx *sql.Tx, t *model.Thread) (sql.Result, error) {
	pos := &types.JSON{}
	if err := pos.Marshal(t.Position); err != nil {
		return nil, err
	}

	dto := models.Thread{
		Title:    t.Title,
		Position: *pos,
	}

	// ここで、ORMを使っていないのはsql.Resultを返して、そのメソッドのLastInsertID()を使うため
	// Exec()によって、sql.Resultが返るがsqlboilerのInsertでは返さない

	p, err := tx.Prepare("INSERT INTO threads (title, position) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := p.Exec(dto.Title, *pos)
	if err != nil {
		return nil, err
	}

	return res, nil
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
