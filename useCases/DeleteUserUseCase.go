package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"gorm.io/gorm"
)

type DeleteUserUseCase struct {
	userRepository gateways.UserRepository
}

func NewDeleteUserUseCase(db *gorm.DB) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: *gateways.NewUserRepository(db),
	}
}

func (g *DeleteUserUseCase) DeleteUser(userId uint32) error {
	return g.userRepository.Delete(userId)
}
