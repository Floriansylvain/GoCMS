package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/user"
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
