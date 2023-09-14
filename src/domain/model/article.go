package model

import "time"

type Article struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Position  Position  `json:"position"`
}

func NewArticle(id int, createdAt time.Time, body string, position Position) *Article {
	return &Article{
		ID:        id,
		CreatedAt: createdAt,
		Body:      body,
		Position:  position,
	}
}
