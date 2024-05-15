package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/user"
	"gorm.io/gorm"
)

type CreateUserUseCase struct {
	userRepository gateways.UserRepository
}

type CreateUserCommand struct {
	Username         string
	Password         string
	Email            string
	VerificationCode string
}

func NewCreateUserUseCase(db *gorm.DB) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: *gateways.NewUserRepository(db),
	}
}

func (g *CreateUserUseCase) CreateUser(createUser CreateUserCommand) (user.User, error) {
	return g.userRepository.Create(user.FromApi(
		createUser.Username,
		createUser.Password,
		createUser.Email,
		createUser.VerificationCode,
	))
}
