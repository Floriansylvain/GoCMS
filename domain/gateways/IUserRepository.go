package gateways

import . "GohCMS2/domain/user"

type IUserRepository interface {
	Get(id uint32) (User, error)
	GetByUsername(username string) (User, error)
	GetAll() []User
	Create(user User) (User, error)
}
