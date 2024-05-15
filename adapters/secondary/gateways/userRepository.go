package gateways

import (
	entity "GohCMS2/adapters/secondary/gateways/models"
	"GohCMS2/domain/gateways"
	domain "GohCMS2/domain/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func mapUserToDomain(user entity.User) domain.User {
	return domain.FromDb(
		user.ID,
		user.Username,
		user.Password,
		user.Email,
		user.IsVerified,
		user.VerificationCode,
		user.VerificationExpiration,
		user.CreatedAt, user.UpdatedAt,
	)
}

func (u *UserRepository) Get(id uint32) (domain.User, error) {
	var user entity.User
	err := u.db.Model(&entity.User{}).First(&user, id).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(user), nil
}

func (u *UserRepository) Create(user domain.User) (domain.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	hashedVerificationCode, _ := bcrypt.GenerateFromPassword([]byte(user.VerificationCode), 12)

	creationResult := u.db.Create(&entity.User{
		Username:               user.Username,
		Password:               string(hashedPassword),
		Email:                  user.Email,
		VerificationCode:       string(hashedVerificationCode),
		VerificationExpiration: user.VerificationExpiration,
	})
	if creationResult.Error != nil {
		return domain.User{}, creationResult.Error
	}

	var createdUser entity.User
	creationResult.Scan(&createdUser)

	return mapUserToDomain(createdUser),
		nil
}

func (u *UserRepository) GetAll() []domain.User {
	var users []entity.User
	u.db.Model(&entity.User{}).Find(&users)

	var domainUsers []domain.User
	for _, user := range users {
		domainUsers = append(domainUsers, mapUserToDomain(user))
	}

	return domainUsers
}

func (u *UserRepository) GetByUsername(username string) (domain.User, error) {
	var user entity.User
	err := u.db.Model(&entity.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(user), nil
}

func (u *UserRepository) UpdateVerificationStatus(userId uint32, isVerified bool) (domain.User, error) {
	var user entity.User
	err := u.db.Model(&entity.User{}).First(&user, userId).Error
	if err != nil {
		return domain.User{}, err
	}

	user.IsVerified = isVerified
	err = u.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(user), nil
}

var _ gateways.IUserRepository = &UserRepository{}
