package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/user"
	"gorm.io/gorm"
)

type ListUsersUseCase struct {
	userRepository UserRepository
}

func NewListUsersUseCase(db *gorm.DB) *ListUsersUseCase {
	return &ListUsersUseCase{
		userRepository: *NewUserRepository(db),
	}
}

func (g *ListUsersUseCase) ListUsers() []User {
	return g.userRepository.GetAll()
}
