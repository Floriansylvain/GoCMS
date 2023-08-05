package api

import (
	. "GohCMS2/useCases"
	"errors"
	"go.uber.org/dig"
)

type Container struct {
	CreateArticleUseCase *CreateArticleUseCase
	GetArticleUseCase    *GetArticleUseCase
}

var container *Container

func InitContainer() *Container {
	if container != nil {
		return container
	}

	digContainer := dig.New()

	err1 := digContainer.Provide(NewCreateArticleUseCase)
	err2 := digContainer.Provide(NewGetArticleUseCase)

	if err := errors.Join(err1, err2); err != nil {
		panic(err)
	}

	err := digContainer.Invoke(func(createArticle *CreateArticleUseCase, getArticle *GetArticleUseCase) {
		container = &Container{
			CreateArticleUseCase: createArticle,
			GetArticleUseCase:    getArticle,
		}
	})
	if err != nil {
		panic(err)
	}

	return container
}
