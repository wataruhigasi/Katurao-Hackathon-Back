package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
}

func NewComment(id int, title string, createdAt time.Time, body string, author string) *Comment {
	return &Comment{
		ID:        id,
		CreatedAt: createdAt,
		Body:      body,
		Author:    author,
	}
}
