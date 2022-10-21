package internal

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var USERS_LOCATION = Location{Database: "gohcms", Collection: "users"}

func getUserHashedPassword(user User) string {
	password := sha256.New()
	password.Write([]byte(user.Password))
	return fmt.Sprintf("%x", password.Sum(nil))
}

func isUserLoggedIn(user User) bool {
	for i := 0; i < len(SESSIONS); i++ {
		hash1 := fmt.Sprint(SESSIONS[i].Token.Sum(nil))
		hash2 := fmt.Sprint(generateSessionToken(user).Sum(nil))
		if hash1 != hash2 {
			continue
		}
		sessionExpired := isSessionExpired(SESSIONS[i])
		if sessionExpired {
			removeSession(user)
		}
		return !sessionExpired
	}
	return false
}

func isUserReal(user User) bool {
	docUsers, _ := getDocuments(USERS_LOCATION, user)
	return len(docUsers) == 1
}

func parseUserFromContext(c *gin.Context) (User, error) {
	var user User
	if c.BindJSON(&user) != nil {
		return User{}, errors.New("could not correctly parse user crendentials")
	}
	user.Password = getUserHashedPassword(user)
	return user, nil
}

func LoginUser(c *gin.Context) {
	user, err := parseUserFromContext(c)
	if err != nil {
		SendErrorMessageToClient(c, err.Error())
		return
	}
	if isUserLoggedIn(user) {
		SendErrorMessageToClient(c, "User is already logged in!")
		return
	}
	if !isUserReal(user) {
		SendErrorMessageToClient(c, "Unknown email or wrong password.")
		return
	}
	addSession(user)
	SendOkMessageToClient(c, "User successfully logged in.")
}

func LogoutUser(c *gin.Context) {
	user, err := parseUserFromContext(c)
	if err != nil {
		SendErrorMessageToClient(c, err.Error())
		return
	}
	if !isUserLoggedIn(user) {
		SendErrorMessageToClient(c, "User is not logged in!")
		return
	}
	removeSession(user)
	SendOkMessageToClient(c, "User successfully logged out.")
}

func AuthCheck(c *gin.Context) {
	var user User
	username, password, isOk := c.Request.BasicAuth()
	if !isOk {
		SendErrorMessageToClient(c, "Incorrect or missing user credentials.")
		c.Abort()
		return
	}

	user.Email = username
	user.Password = password
	user.Password = getUserHashedPassword(user)

	if !isUserLoggedIn(user) {
		SendErrorMessageToClient(c, "Authentification failed, credentials could be wrong, user may not be logged in, session may have expired.")
		c.Abort()
		return
	}
}
