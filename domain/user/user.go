package user

import "time"

type User struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromApi(
	username string,
	password string,
	email string,
) User {
	return User{
		Username: username,
		Password: password,
		Email:    email,
	}
}

func FromDb(
	id uint32,
	username string,
	password string,
	email string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
