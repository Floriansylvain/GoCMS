package user

import (
	"time"
)

type User struct {
	ID                     uint32    `json:"id"`
	Username               string    `json:"username"`
	Password               string    `json:"password"`
	Email                  string    `json:"email"`
	IsVerified             bool      `gorm:"default=false;not null"`
	VerificationCode       string    `gorm:"unique;not null"`
	VerificationExpiration time.Time `gorm:"not null"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

func FromApi(
	username string,
	password string,
	email string,
	verificationCode string,
) User {
	return User{
		Username:               username,
		Password:               password,
		Email:                  email,
		IsVerified:             false,
		VerificationCode:       verificationCode,
		VerificationExpiration: time.Now().Add(2 * time.Hour),
	}
}

func FromDb(
	id uint32,
	username string,
	password string,
	email string,
	isVerified bool,
	verificationCode string,
	verificationExpiration time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:                     id,
		Username:               username,
		Password:               password,
		Email:                  email,
		IsVerified:             isVerified,
		VerificationCode:       verificationCode,
		VerificationExpiration: verificationExpiration,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
	}
}
