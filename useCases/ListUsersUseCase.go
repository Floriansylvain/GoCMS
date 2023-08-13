package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/user"
	"gorm.io/gorm"
)

type ListUsersUseCase struct {
	userRepository gateways.UserRepository
}

func NewListUsersUseCase(db *gorm.DB) *ListUsersUseCase {
	return &ListUsersUseCase{
		userRepository: *gateways.NewUserRepository(db),
	}
}

func (g *ListUsersUseCase) ListUsers() []user.User {
	return g.userRepository.GetAll()
}
