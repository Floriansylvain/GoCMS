package article

type Article struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
