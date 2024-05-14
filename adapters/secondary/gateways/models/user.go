package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID                     uint32    `gorm:"primaryKey;autoIncrement"`
	Username               string    `gorm:"unique;not null"`
	Password               string    `gorm:"not null"`
	Email                  string    `gorm:"unique;not null"`
	IsVerified             bool      `gorm:"default=false;not null"`
	VerificationCode       string    `gorm:"unique;not null"`
	VerificationExpiration time.Time `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime"`
}
