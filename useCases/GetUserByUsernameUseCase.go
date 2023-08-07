package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/user"
	"gorm.io/gorm"
)

type GetUserByUsernameUseCase struct {
	userRepository UserRepository
}

func NewGetUserByUsernameUseCase(db *gorm.DB) *GetUserByUsernameUseCase {
	return &GetUserByUsernameUseCase{
		userRepository: *NewUserRepository(db),
	}
}

func (g *GetUserByUsernameUseCase) GetUserByUsername(username string) (User, error) {
	return g.userRepository.GetByUsername(username)
}
