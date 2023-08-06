package api

import (
	. "GohCMS2/useCases"
	"go.uber.org/dig"
)

type Container struct {
	CreateArticleUseCase *CreateArticleUseCase
	GetArticleUseCase    *GetArticleUseCase
	ListArticlesUseCase  *ListArticlesUseCase
}

var container *Container

func createContainer(
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

	_ = digContainer.Provide(NewCreateArticleUseCase)
	_ = digContainer.Provide(NewGetArticleUseCase)
	_ = digContainer.Provide(NewListArticlesUseCase)

	_ = digContainer.Invoke(createContainer)
}
