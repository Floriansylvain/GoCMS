package post

import (
	entity "GoCMS/adapters/secondary/gateways/models"
	domain "GoCMS/domain/image"
	"time"
)

type Post struct {
	ID        uint32          `json:"id"`
	Title     string          `json:"title"`
	Body      string          `json:"body"`
	Images    []*domain.Image `json:"images"`
	IsOnline  bool            `json:"is_online"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
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
	images []*entity.Image,
	isOnline bool,
	createdAt time.Time,
	updatedAt time.Time,
) Post {
	domainImages := make([]*domain.Image, len(images))
	for i, img := range images {
		domainImage := domain.FromDB(
			img.ID,
			img.Path,
			img.PostID,
			img.CreatedAt,
			img.UpdatedAt,
		)
		domainImages[i] = &domainImage
	}
	return Post{
		ID:        id,
		Title:     title,
		Body:      body,
		Images:    domainImages,
		IsOnline:  isOnline,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
