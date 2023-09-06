package test

import (
	"GohCMS2/domain/post"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

var TestCreatePostSuccess = func(t *testing.T) {
	jsonBody, err := json.Marshal(map[string]string{
		"title": "Test Title",
		"body":  "Test Body",
	})
	if err != nil {
		t.Fatal(err)
	}

	r, _ := ApiRequest("POST", "/post", bytes.NewBuffer(jsonBody))

	var response post.Post
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
}

var TestCreatePostFailTitleMissing = func(t *testing.T) {
	jsonBody, err := json.Marshal(map[string]string{
		"body": "Test Body",
	})
	if err != nil {
		t.Fatal(err)
	}

	r, _ := ApiRequest("POST", "/post", bytes.NewBuffer(jsonBody))

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
}

var TestCreatePostTitleTooShort = func(t *testing.T) {
	jsonBody, err := json.Marshal(map[string]string{
		"title": "Te",
		"body":  "Test Body",
	})
	if err != nil {
		t.Fatal(err)
	}

	r, _ := ApiRequest("POST", "/post", bytes.NewBuffer(jsonBody))

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
}

var TestGetPostSuccess = func(t *testing.T) {
	var createdPost post.Post
	var postToCreate = post.Post{
		Title: "Test Title",
		Body:  "Test Body",
	}
	db := GetDb()
	db.Create(&postToCreate).Scan(&createdPost)

	r, _ := ApiRequest("GET", "/post/"+strconv.Itoa(int(createdPost.ID)), nil)

	var response post.Post
	bd, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bd, &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, createdPost.ID, response.ID)
	assert.Equal(t, createdPost.Title, response.Title)
	assert.Equal(t, createdPost.Body, response.Body)
}

var TestGetAllPostsSuccess = func(t *testing.T) {
	var createdPost post.Post
	var postToCreate = post.Post{
		Title: "Test Title",
		Body:  "Test Body",
	}
	db := GetDb()
	db.Create(&postToCreate).Scan(&createdPost)

	r, _ := ApiRequest("GET", "/post", nil)

	var response []post.Post
	bd, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bd, &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, createdPost.Title, response[0].Title)
	assert.Equal(t, createdPost.Body, response[0].Body)
}

var TestPostCreate = func(t *testing.T) {
	t.Run("Should return an post with the given title and body", TestCreatePostSuccess)
	t.Run("Should return an error if the title is missing", TestCreatePostFailTitleMissing)
	t.Run("Should return an error if the title is too short", TestCreatePostTitleTooShort)
}

var TestPostGet = func(t *testing.T) {
	t.Run("Should return an post with the given id", TestGetPostSuccess)
}

var TestPostGetAll = func(t *testing.T) {
	t.Run("Should return all posts", TestGetAllPostsSuccess)
}

func TestPost(t *testing.T) {
	StartServerIfNotAlready()
	WaitForServer()

	t.Run("Create", TestPostCreate)
	t.Run("Get", TestPostGet)
	t.Run("GetAll", TestPostGetAll)
}
