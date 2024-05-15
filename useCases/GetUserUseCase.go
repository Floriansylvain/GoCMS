package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/user"
	"gorm.io/gorm"
)

type GetUserUseCase struct {
	userRepository gateways.UserRepository
}

func NewGetUserUseCase(db *gorm.DB) *GetUserUseCase {
	return &GetUserUseCase{
		userRepository: *gateways.NewUserRepository(db),
	}
}

func (g *GetUserUseCase) GetUser(id uint32) (user.User, error) {
	return g.userRepository.Get(id)
}

func (g *GetUserUseCase) GetUserByUsername(username string) (user.User, error) {
	return g.userRepository.GetByUsername(username)
}
