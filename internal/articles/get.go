package articles

import (
	"fmt"
	"math"
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
	return fmt.Sprintf("%varticles?skip=%v&take=%v", getApiFullUrl(), skip+take, take)
}

func getBuiltGetResponse(articles []Article, skip uint64, take uint64) map[string]any {
	articlesCap := uint64(len(articles))

	var normTake uint64
	if skip+take > articlesCap {
		normTake = articlesCap
	} else {
		normTake = skip + take
	}

	slicedArticles := articles[skip:normTake]
	total := len(slicedArticles)
	return map[string]any{
		"content": slicedArticles,
		"total":   total,
		"pagination": map[string]any{
			"skip": skip,
			"take": take,
			"links": map[string]any{
				"next":     getArticleSkipTakeFullUrl(skip+take, take),
				"previous": getArticleSkipTakeFullUrl(skip-take, take),
			},
		},
		"last_page": math.Ceil(float64(articlesCap) / float64(take)),
	}
}

func parseUintQueryParam(c *gin.Context, param string, defaultValue uint64) (uint64, error) {
	if c.Query(param) == "" {
		return defaultValue, nil
	}
	value, err := strconv.ParseUint(c.Query(param), 10, 0)
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

	skip, skipErr := parseUintQueryParam(c, "skip", 0)
	take, takeErr := parseUintQueryParam(c, "take", 10)
	if skipErr != nil || takeErr != nil {
		api.SendBadRequest(c, "Take and Skip query parameters must be positive numbers.")
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
