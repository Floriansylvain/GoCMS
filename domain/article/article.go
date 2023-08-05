package article

type Article struct {
	Id        int
	Title     string
	Body      string
	createdAt string
	updatedAt string
}

func FromApi(
	title string,
	body string,
) Article {
	return Article{
		Title: title,
		Body:  body,
	}
}

func FromDb(
	id int,
	title string,
	body string,
	createdAt string,
	updatedAt string,
) Article {
	return Article{
		Id:        id,
		Title:     title,
		Body:      body,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
