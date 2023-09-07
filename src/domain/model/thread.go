package model

import "time"

type Thread struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Position  Position  `json:"coordinate"`
	Comments  []*Comment  `json:"comments"`
}

func NewThread(id int, title string, createdAt time.Time, position Position) *Thread {
	return &Thread{
		ID: id,
		Title: title,
		CreatedAt: createdAt,
		Position: position,
		Comments: []*Comment{},
	}
}