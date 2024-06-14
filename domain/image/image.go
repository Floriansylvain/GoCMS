package image

import (
	"time"
)

type Image struct {
	ID        uint32    `json:"id"`
	Path      string    `json:"path"`
	PostID    uint32    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDB(id uint32, path string, postId uint32, createdAt time.Time, updatedAt time.Time) Image {
	return Image{
		ID:        id,
		Path:      path,
		PostID:    postId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
