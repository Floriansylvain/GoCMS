package test

import (
	"GoCMS/adapters/secondary/gateways/models"
	"GoCMS/api"
	"GoCMS/main/server"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"time"
)

const testDbFile = "test.db"

var ApiUrl string
var AuthorizationCookie *http.Cookie
var HttpClient = http.Client{}

func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(testDbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&models.Post{}, &models.User{})
	return db
}

func StartServerIfNotAlready() {
	_, err := http.Get(ApiUrl)
	if err == nil {
		return
	}
	_ = os.Remove(testDbFile)
	_ = os.Setenv("DB_FILE", testDbFile)
	go func(url *string) {
		router := server.InitServer()
		*url = "http://localhost:" + os.Getenv("PORT") + "/v1"
		err := server.StartServer(router)
		if err != nil {
			panic(err)
		}
	}(&ApiUrl)
}

func getAuthorizationCookie(userId uint32) *http.Cookie {
	_, tokenString, err := api.TokenAuth.Encode(map[string]interface{}{"user_id": userId})
	if err != nil {
		panic(err)
	}
	return &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}
}

func SetAuthorizationCookieIfNotAlready(r *http.Request) {
	if AuthorizationCookie != nil {
		r.AddCookie(AuthorizationCookie)
		return
	}
	db := GetDb()
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "testemail@a.com",
	}
	var createdUser models.User
	db.Create(&user).Scan(&createdUser)
	AuthorizationCookie = getAuthorizationCookie(createdUser.ID)
	r.AddCookie(AuthorizationCookie)
}

func WaitForServer() {
	for {
		time.Sleep(100 * time.Millisecond)
		if ApiUrl == "" {
			continue
		}
		_, err := http.Get(ApiUrl)
		if err == nil {
			break
		}
	}
}

func ApiRequest(method string, route string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, ApiUrl+route, body)
	if err != nil {
		return nil, err
	}
	SetAuthorizationCookieIfNotAlready(request)

	response, err := HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
