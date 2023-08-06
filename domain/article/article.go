package article

import "time"

type Article struct {
	ID        uint32    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromApi(
	title string,
	body string,
) Article {
	return Article{
		Title: title,
		Body:  body,
	}
}

func FromDb(
	id uint32,
	title string,
	body string,
	createdAt time.Time,
	updatedAt time.Time,
) Article {
	return Article{
		ID:        id,
		Title:     title,
		Body:      body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
