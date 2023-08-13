package api

import (
	"GohCMS2/adapters/secondary/gateways/models"
	"GohCMS2/useCases"
	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type LocalContainer struct {
	CreateArticleUseCase *useCases.CreateArticleUseCase
	GetArticleUseCase    *useCases.GetArticleUseCase
	ListArticlesUseCase  *useCases.ListArticlesUseCase
	GetUserUseCase       *useCases.GetUserUseCase
	CreateUserUseCase    *useCases.CreateUserUseCase
	ListUsersUseCase     *useCases.ListUsersUseCase
	GetPageUseCase       *useCases.GetPageUseCase
}

var Container *LocalContainer

func setContainer(
	createArticle *useCases.CreateArticleUseCase,
	getArticle *useCases.GetArticleUseCase,
	listArticle *useCases.ListArticlesUseCase,
	getUser *useCases.GetUserUseCase,
	createUser *useCases.CreateUserUseCase,
	listUsers *useCases.ListUsersUseCase,
	getPage *useCases.GetPageUseCase,
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

	dbName := os.Getenv("DB_FILE")
	if err := os.MkdirAll(filepath.Dir(dbName), os.ModePerm); err != nil {
		panic("Unable to create necessary subdirectories: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Unable to open the database: " + err.Error())
	}

	_ = db.AutoMigrate(&models.Article{}, &models.User{})

	_ = digContainer.Provide(func() *gorm.DB { return db })

	_ = digContainer.Provide(func(db *gorm.DB) *useCases.CreateArticleUseCase { return useCases.NewCreateArticleUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.GetArticleUseCase { return useCases.NewGetArticleUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.ListArticlesUseCase { return useCases.NewListArticlesUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.GetUserUseCase { return useCases.NewGetUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.CreateUserUseCase { return useCases.NewCreateUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.ListUsersUseCase { return useCases.NewListUsersUseCase(db) })
	_ = digContainer.Provide(useCases.NewGetPageUseCase)

	_ = digContainer.Invoke(setContainer)
}
