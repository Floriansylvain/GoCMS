package api

import (
	. "GohCMS2/domain/article"
	. "GohCMS2/useCases"
)

func GetArticle(id int) Article {
	return container.GetArticleUseCase.GetArticle(id)
}

func CreateArticle(title string, body string) Article {
	return container.CreateArticleUseCase.CreateAarticle(CreateArticleCommand{Title: title, Body: body})
}
