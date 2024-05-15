package gateways

import "GoCMS/domain/user"

type IUserRepository interface {
	Get(id uint32) (user.User, error)
	GetByUsername(username string) (user.User, error)
	GetAll() []user.User
	Create(user user.User) (user.User, error)
	Delete(id uint32) error
	UpdateVerificationStatus(userId uint32, isVerified bool) (user.User, error)
}
