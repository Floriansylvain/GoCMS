package image

import (
	"time"
)

type Image struct {
	ID        uint32
	Path      string
	PostID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
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
