package post

import "time"

type Post struct {
	ID        uint32    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromApi(
	title string,
	body string,
) Post {
	return Post{
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
) Post {
	return Post{
		ID:        id,
		Title:     title,
		Body:      body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
