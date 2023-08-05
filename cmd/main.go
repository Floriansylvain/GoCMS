package main

import (
	"GohCMS2/api"
	"fmt"
)

func main() {
	api.InitContainer()

	fmt.Println(api.CreateArticle("Title", "Content"))
	fmt.Println(api.GetArticle(1))
}
