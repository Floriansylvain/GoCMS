package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/user"
	"gorm.io/gorm"
)

type CreateUserUseCase struct {
	userRepository UserRepository
}

type CreateUserCommand struct {
	Username string
	Password string
	Email    string
}

func NewCreateUserUseCase(db *gorm.DB) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: *NewUserRepository(db),
	}
}

func (g *CreateUserUseCase) CreateUser(user CreateUserCommand) (User, error) {
	return g.userRepository.Create(FromApi(user.Username, user.Password, user.Email))
}
