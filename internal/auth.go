package internal

import (
	"errors"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var UsersLocation = Location{Database: "gohcms", Collection: "users"}

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:         "GohCMS",
	Key:           []byte(os.Getenv("APP_JWT_SECRET")),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour,
	Authenticator: JWTAuthenticator,
})

func JWTAuthenticator(c *gin.Context) (interface{}, error) {
	var user = User{}
	err := c.BindJSON(&user)
	if err != nil {
		return nil, errors.New("wrong credentials json format.")
	}

	_, err = getUniqueDocument(UsersLocation, bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		return nil, errors.New("wrong email or password.")
	}

	return gin.H{"email": user.Email}, nil
}
