package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/user"
	"gorm.io/gorm"
)

type UpdateUserUseCase struct {
	userRepository gateways.UserRepository
}

func NewUpdateUserUseCase(db *gorm.DB) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: *gateways.NewUserRepository(db),
	}
}

func (g *UpdateUserUseCase) UpdateVerificationStatus(userId uint32, isVerified bool) (user.User, error) {
	return g.userRepository.UpdateVerificationStatus(userId, isVerified)
}

func (g *UpdateUserUseCase) UpdatePasswordResetCode(userId uint32, code string) (user.User, error) {
	return g.userRepository.UpdatePasswordResetCode(userId, code)
}

func (g *UpdateUserUseCase) UpdatePassword(userId uint32, password string) (user.User, error) {
	return g.userRepository.UpdatePassword(userId, password)
}
