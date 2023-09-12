package model

import "time"

type KeijibanRakugaki struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Position  Position  `json:"position"`
}

func NewKeijibanRakugaki(id int, title string, createdAt time.Time, body string, position Position) *KeijibanRakugaki {
	return &KeijibanRakugaki{
		ID:        id,
		CreatedAt: createdAt,
		Body:      body,
		Position:  position,
	}
}
