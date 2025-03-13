package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/user"
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
