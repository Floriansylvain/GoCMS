package models

import (
	"gorm.io/gorm"
	"time"
)

type Image struct {
	gorm.Model
	ID        uint32 `gorm:"primary_key;auto_increment;not_null"`
	Path      string
	PostID    uint32
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
