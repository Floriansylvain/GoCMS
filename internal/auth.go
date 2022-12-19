package internal

import (
	"errors"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var USERS_LOCATION = Location{Database: "gohcms", Collection: "users"}

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:         "GohCMS",
	Key:           []byte(os.Getenv("JWT_SECRET")),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour,
	Authenticator: JWTAuthenticator,
})

func JWTAuthenticator(c *gin.Context) (interface{}, error) {
	var user = User{}
	c.BindJSON(&user)
	if user.Email == "dddeschamps2022" && user.Password == "1234" {
		return gin.H{"email": "dddeschamps2022"}, nil
	}
	return nil, errors.New("can't verify credentials.")
}
