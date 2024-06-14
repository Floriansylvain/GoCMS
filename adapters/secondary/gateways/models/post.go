package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID        uint32 `gorm:"primary_key;auto_increment;not_null"`
	Title     string
	Body      string
	Images    []*Image  `gorm:"many2many:post_images;"`
	IsOnline  bool      `gorm:"not_null;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
