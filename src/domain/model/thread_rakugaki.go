package model

import "time"

type ThreadRakugaki struct {
	ID        int       `json:"id"`
	ThreadID  int       `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Position  Position  `json:"position"`
}

func NewThreadRakugaki(id, threadID int, title string, createdAt time.Time, body string, position Position) *ThreadRakugaki {
	return &ThreadRakugaki{
		ID:        id,
		ThreadID:  threadID,
		CreatedAt: createdAt,
		Body:      body,
		Position:  position,
	}
}
