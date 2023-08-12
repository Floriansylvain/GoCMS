package api

import (
	"GohCMS2/adapters/secondary/gateways/models"
	. "GohCMS2/useCases"
	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type LocalContainer struct {
	CreateArticleUseCase *CreateArticleUseCase
	GetArticleUseCase    *GetArticleUseCase
	ListArticlesUseCase  *ListArticlesUseCase
	GetUserUseCase       *GetUserUseCase
	CreateUserUseCase    *CreateUserUseCase
	ListUsersUseCase     *ListUsersUseCase
	GetPageUseCase       *GetPageUseCase
}

var Container *LocalContainer

func setContainer(
	createArticle *CreateArticleUseCase,
	getArticle *GetArticleUseCase,
	listArticle *ListArticlesUseCase,
	getUser *GetUserUseCase,
	createUser *CreateUserUseCase,
	listUsers *ListUsersUseCase,
	getPage *GetPageUseCase,
) *LocalContainer {
	Container = &LocalContainer{
		CreateArticleUseCase: createArticle,
		GetArticleUseCase:    getArticle,
		ListArticlesUseCase:  listArticle,
		GetUserUseCase:       getUser,
		CreateUserUseCase:    createUser,
		ListUsersUseCase:     listUsers,
		GetPageUseCase:       getPage,
	}
	return Container
}

func InitContainer() {
	if Container != nil {
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
	_ = digContainer.Provide(func(db *gorm.DB) *CreateUserUseCase { return NewCreateUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *ListUsersUseCase { return NewListUsersUseCase(db) })
	_ = digContainer.Provide(NewGetPageUseCase)

	_ = digContainer.Invoke(setContainer)
}
