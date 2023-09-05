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
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
