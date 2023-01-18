package articles

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
)

func getApiFullUrl() string {
	return fmt.Sprintf("%v%v", os.Getenv("APP_API_ADDRESS"), os.Getenv("APP_BASE_API_PATH"))
}

func getArticleSkipTakeFullUrl(skip uint64, take uint64) string {
	return fmt.Sprintf("%v/articles?skip=%v&take=%v", getApiFullUrl(), skip+take, take)
}

func getBuiltGetResponse(articles []Article, skip uint64, take uint64) map[string]any {
	slicedArticles := articles[skip : skip+take]
	return map[string]any{
		"content": slicedArticles,
		"total":   len(slicedArticles),
		"pagination": map[string]any{
			"skip": skip,
			"take": take,
			"links": map[string]any{
				"next":     getArticleSkipTakeFullUrl(skip+take, take),
				"previous": getArticleSkipTakeFullUrl(skip-take, take),
			},
		},
	}
}

func parseUintQueryParam(c *gin.Context, param string) (uint64, error) {
	value, err := strconv.ParseUint(c.Query(param), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%s query parameter must be a positive number.", param))
	}
	return value, err
}

func getArticleIdFilter(titleID string) map[string]any {
	filter := map[string]any{}
	if titleID != "" {
		filter["titleID"] = titleID
	}
	return filter
}

func Get(c *gin.Context) {
	titleID := c.Params.ByName("id")

	skip, skipErr := parseUintQueryParam(c, "skip")
	take, takeErr := parseUintQueryParam(c, "take")
	if skipErr != nil || takeErr != nil {
		return
	}

	articles, err := database.GetDocuments(articlesLocation, getArticleIdFilter(titleID))
	if err != nil {
		api.SendBadRequest(c, fmt.Sprintf("The ID '%v' doesn't match any article.", titleID))
		return
	}

	articlesArray := ParseArticlesFromBytesToArray(articles)
	articlesArrayLength := uint64(len(articlesArray))
	if skip >= articlesArrayLength || take == 0 {
		c.JSON(http.StatusOK, map[string]any{
			"content": []string{},
			"total":   "0",
		})
		return
	}

	c.JSON(http.StatusOK, getBuiltGetResponse(articlesArray, skip, take))
}
