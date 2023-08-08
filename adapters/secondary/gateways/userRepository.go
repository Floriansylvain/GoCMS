package gateways

import (
	entity "GohCMS2/adapters/secondary/gateways/models"
	. "GohCMS2/domain/gateways"
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
	return domain.FromDb(user.ID, user.Username, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
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
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	creationResult := u.db.Create(&entity.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	})
	if creationResult.Error != nil {
		return domain.User{}, creationResult.Error
	}

	var createdUser entity.User
	creationResult.Scan(&createdUser)

	return domain.FromDb(
			createdUser.ID,
			createdUser.Username,
			createdUser.Password,
			createdUser.Email,
			createdUser.CreatedAt,
			createdUser.UpdatedAt),
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

var _ IUserRepository = &UserRepository{}
