package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/user"
	"gorm.io/gorm"
)

type GetUserUseCase struct {
	userRepository UserRepository
}

func NewGetUserUseCase(db *gorm.DB) *GetUserUseCase {
	return &GetUserUseCase{
		userRepository: *NewUserRepository(db),
	}
}

func (g *GetUserUseCase) GetUser(id uint32) (User, error) {
	return g.userRepository.Get(id)
}

func (g *GetUserUseCase) GetUserByUsername(username string) (User, error) {
	return g.userRepository.GetByUsername(username)
}
