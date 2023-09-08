package model

import "time"

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Position  Position  `json:"position"`
}

func NewArticle(id int, title string, createdAt time.Time, body string, position Position) *Article {
	return &Article{
		ID:        id,
		Title:     title,
		CreatedAt: createdAt,
		Body:      body,
		Position:  position,
	}
}
