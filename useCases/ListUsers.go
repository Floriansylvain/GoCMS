package useCases

import (
	"GoCMS/domain/gateways"
	"GoCMS/domain/user"
)

type ListUsersUseCase struct {
	userRepository gateways.IUserRepository
}

func NewListUsersUseCase(userRepository gateways.IUserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{userRepository}
}

func (g *ListUsersUseCase) ListUsers() []user.User {
	return g.userRepository.GetAll()
}
