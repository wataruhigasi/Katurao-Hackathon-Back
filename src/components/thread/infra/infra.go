package infra

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/models"
)

type ThreadRepository interface {
	FindAll() ([]*model.Thread, error)
	Create(*model.Thread) (sql.Result, error)
	GetLastInsertID() (int64, error)
}

type threadRepositoryImpl struct {
	conn *sql.DB
}

func NewThreadRepository(conn *sql.DB) *threadRepositoryImpl {
	return &threadRepositoryImpl{
		conn: conn,
	}
}

func (tr *threadRepositoryImpl) GetLastInsertID() (int64, error) {
	ctx := context.Background()

	dto, err := models.Threads().All(ctx, tr.conn)
	if err != nil {
		return 0, err
	}

	return dto[len(dto)-1].ID, nil
}

func (tr *threadRepositoryImpl) FindAll() ([]*model.Thread, error) {
	ctx := context.Background()

	dto, err := models.Threads().All(ctx, tr.conn)
	if err != nil {
		return nil, err
	}

	ts := make([]*model.Thread, 0, len(dto))
	for _, v := range dto {
		t, err := ToThread(v)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	return ts, nil
}

func ToThread(t *models.Thread) (*model.Thread, error) {
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

func (tr *threadRepositoryImpl) Create(t *model.Thread) (sql.Result, error) {
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

	p, err := tr.conn.Prepare("INSERT INTO threads (title, position) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := p.Exec(dto.Title, *pos)
	if err != nil {
		return nil, err
	}

	return res, nil

}
