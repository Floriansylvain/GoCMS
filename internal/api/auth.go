package api

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Floriansylvain/GohCMS/internal/database"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var UsersLocation = database.Location{Database: "gohcms", Collection: "users"}

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:          "GohCMS",
	Key:            []byte(os.Getenv("APP_JWT_SECRET")),
	SendCookie:     true,
	CookieHTTPOnly: true,
	CookieSameSite: http.SameSiteStrictMode,
	Timeout:        time.Hour,
	MaxRefresh:     time.Hour,
	LoginResponse:  JWTLoginResponse,
	Authenticator:  JWTAuthenticator,
})

func JWTLoginResponse(c *gin.Context, code int, message string, expire time.Time) {
	if code == http.StatusOK {
		c.JSON(code, gin.H{"code": code, "message": "Successfully logged in!", "expire": expire.Format(time.RFC3339)})
	} else {
		c.JSON(code, gin.H{"code": code, "message": "Something wrong has happened."})
	}
}

func JWTAuthenticator(c *gin.Context) (interface{}, error) {
	var user = User{}
	err := c.BindJSON(&user)
	if err != nil {
		return nil, errors.New("wrong credentials json format.")
	}

	_, err = database.GetUniqueDocument(UsersLocation, bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		return nil, errors.New("wrong email or password.")
	}

	return gin.H{"email": user.Email}, nil
}
