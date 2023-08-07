package api

import (
	"GohCMS2/adapters/secondary/gateways/models"
	. "GohCMS2/useCases"
	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Container struct {
	CreateArticleUseCase     *CreateArticleUseCase
	GetArticleUseCase        *GetArticleUseCase
	ListArticlesUseCase      *ListArticlesUseCase
	GetUserUseCase           *GetUserUseCase
	GetUserByUsernameUseCase *GetUserByUsernameUseCase
	CreateUserUseCase        *CreateUserUseCase
	ListUsersUseCase         *ListUsersUseCase
}

var container *Container

func setContainer(
	createArticle *CreateArticleUseCase,
	getArticle *GetArticleUseCase,
	listArticle *ListArticlesUseCase,
	getUser *GetUserUseCase,
	getUserByUsername *GetUserByUsernameUseCase,
	createUser *CreateUserUseCase,
	listUsers *ListUsersUseCase,
) *Container {
	container = &Container{
		CreateArticleUseCase:     createArticle,
		GetArticleUseCase:        getArticle,
		ListArticlesUseCase:      listArticle,
		GetUserUseCase:           getUser,
		GetUserByUsernameUseCase: getUserByUsername,
		CreateUserUseCase:        createUser,
		ListUsersUseCase:         listUsers,
	}
	return container
}

func InitContainer() {
	if container != nil {
		return
	}

	digContainer := dig.New()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&models.Article{}, &models.User{})

	_ = digContainer.Provide(func() *gorm.DB { return db })

	_ = digContainer.Provide(func(db *gorm.DB) *CreateArticleUseCase { return NewCreateArticleUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *GetArticleUseCase { return NewGetArticleUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *ListArticlesUseCase { return NewListArticlesUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *GetUserUseCase { return NewGetUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *GetUserByUsernameUseCase { return NewGetUserByUsernameUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *CreateUserUseCase { return NewCreateUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *ListUsersUseCase { return NewListUsersUseCase(db) })

	_ = digContainer.Invoke(setContainer)
}
