package model

import "time"

type Thread struct {
	ID              int               `json:"id"`
	Title           string            `json:"title"`
	CreatedAt       time.Time         `json:"created_at"`
	Position        Position          `json:"position"`
	ThreadRakugakis []*ThreadRakugaki `json:"thread_rakugakis"`
}

func NewThread(id int, title string, createdAt time.Time, position Position) *Thread {
	return &Thread{
		ID:              id,
		Title:           title,
		CreatedAt:       createdAt,
		Position:        position,
		ThreadRakugakis: []*ThreadRakugaki{},
	}
}
