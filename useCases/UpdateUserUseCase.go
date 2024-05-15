package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/user"
	"gorm.io/gorm"
)

type UpdateUserUseCase struct {
	userRepository gateways.UserRepository
}

type UpdateVerificationStatusCommand struct {
	isVerified bool
}

func NewUpdateUserUseCase(db *gorm.DB) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: *gateways.NewUserRepository(db),
	}
}

func (g *UpdateUserUseCase) UpdateVerificationStatus(userId uint32, isVerified bool) (user.User, error) {
	return g.userRepository.UpdateVerificationStatus(userId, isVerified)
}
