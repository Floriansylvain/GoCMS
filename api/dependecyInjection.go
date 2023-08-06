package api

import (
	"GohCMS2/adapters/secondary/gateways/models"
	. "GohCMS2/useCases"
	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Container struct {
	CreateArticleUseCase *CreateArticleUseCase
	GetArticleUseCase    *GetArticleUseCase
	ListArticlesUseCase  *ListArticlesUseCase
}

var container *Container

func setContainer(
	createArticle *CreateArticleUseCase,
	getArticle *GetArticleUseCase,
	listArticle *ListArticlesUseCase,
) *Container {
	container = &Container{
		CreateArticleUseCase: createArticle,
		GetArticleUseCase:    getArticle,
		ListArticlesUseCase:  listArticle,
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
	_ = db.AutoMigrate(&models.Article{})

	_ = digContainer.Provide(func() *gorm.DB { return db })

	_ = digContainer.Provide(func(db *gorm.DB) *CreateArticleUseCase { return NewCreateArticleUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *GetArticleUseCase { return NewGetArticleUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *ListArticlesUseCase { return NewListArticlesUseCase(db) })

	_ = digContainer.Invoke(setContainer)
}
