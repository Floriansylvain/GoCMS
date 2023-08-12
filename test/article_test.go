package test

import (
	"GohCMS2/domain/article"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestArticle(t *testing.T) {
	StartServerIfNotAlready()
	WaitForServer()

	t.Run("Create", func(t *testing.T) {
		t.Run("Should return an article with the given title and body", func(t *testing.T) {
			jsonBody, err := json.Marshal(map[string]string{
				"title": "Test Title",
				"body":  "Test Body",
			})
			if err != nil {
				t.Fatal(err)
			}

			r, err := ApiRequest("POST", "/article", bytes.NewBuffer(jsonBody))

			var response article.Article
			bd, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = json.Unmarshal(bd, &response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusOK, r.StatusCode)
			assert.Equal(t, "Test Title", response.Title)
			assert.Equal(t, "Test Body", response.Body)
		})

		t.Run("Should return an error if the title is missing", func(t *testing.T) {
			jsonBody, err := json.Marshal(map[string]string{
				"body": "Test Body",
			})
			if err != nil {
				t.Fatal(err)
			}

			r, err := ApiRequest("POST", "/article", bytes.NewBuffer(jsonBody))

			assert.Equal(t, http.StatusBadRequest, r.StatusCode)
		})

		t.Run("Should return an error if the title is too short", func(t *testing.T) {
			jsonBody, err := json.Marshal(map[string]string{
				"title": "Te",
				"body":  "Test Body",
			})
			if err != nil {
				t.Fatal(err)
			}

			r, err := ApiRequest("POST", "/article", bytes.NewBuffer(jsonBody))

			assert.Equal(t, http.StatusBadRequest, r.StatusCode)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Should return an article with the given id", func(t *testing.T) {
			var createdArticle article.Article
			var articleToCreate = article.Article{
				Title: "Test Title",
				Body:  "Test Body",
			}
			db := GetDb()
			db.Create(&articleToCreate).Scan(&createdArticle)

			r, err := ApiRequest("GET", "/article/"+strconv.Itoa(int(createdArticle.ID)), nil)

			var response article.Article
			bd, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = json.Unmarshal(bd, &response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusOK, r.StatusCode)
			assert.Equal(t, createdArticle.ID, response.ID)
			assert.Equal(t, createdArticle.Title, response.Title)
			assert.Equal(t, createdArticle.Body, response.Body)
		})
	})

	t.Run("GetAll", func(t *testing.T) {
		t.Run("Should return all articles", func(t *testing.T) {
			var createdArticle article.Article
			var articleToCreate = article.Article{
				Title: "Test Title",
				Body:  "Test Body",
			}
			db := GetDb()
			db.Create(&articleToCreate).Scan(&createdArticle)

			r, err := ApiRequest("GET", "/article", nil)

			var response []article.Article
			bd, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = json.Unmarshal(bd, &response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, http.StatusOK, r.StatusCode)
			assert.Equal(t, createdArticle.Title, response[0].Title)
			assert.Equal(t, createdArticle.Body, response[0].Body)
		})
	})
}
