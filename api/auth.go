package api

import (
	"GohCMS2/domain/user"
	"GohCMS2/useCases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type RegisterCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

var shouldCookieBeSecure = os.Getenv("ENVIRONMENT") == "production"
var TokenAuth *jwtauth.JWTAuth

// TODO Move into its own file or package that handles api errors
const logsErrorMessage = "Access to the requested resource is forbidden due to incorrect password and/or username."
const bodyErrorMessage = "The request cannot be processed due to a mismatch in the format of the body."

func SetJwtCookie(w *http.ResponseWriter, userId uint32) error {
	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}
	http.SetCookie(*w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   shouldCookieBeSecure,
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

func IsUserTableEmpty() bool {
	users := Container.ListUsersUseCase.ListUsers()
	return len(users) == 0
}

func IsLoggedIn(r *http.Request) bool {
	token, err := jwtauth.VerifyRequest(
		TokenAuth,
		r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	return token != nil && err == nil
}

func getUserFromCredentials(credentials LoginCredentials) (user.User, error) {
	dbUser, err := Container.GetUserUseCase.GetUserByUsername(credentials.Username)
	if err != nil {
		return user.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(credentials.Password))
	if err != nil {
		return user.User{}, err
	}

	return dbUser, nil
}

func login(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, bodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser, err := getUserFromCredentials(credentials)
	if err != nil {
		http.Error(w, logsErrorMessage, http.StatusForbidden)
	}

	_ = SetJwtCookie(&w, dbUser.ID)

	message, _ := json.Marshal(map[string]interface{}{"message": "User logged in! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func register(w http.ResponseWriter, r *http.Request) {
	if !IsLoggedIn(r) && !IsUserTableEmpty() {
		http.Error(w, "You are not allowed to create a user. Log in or reset database.", http.StatusForbidden)
		return
	}

	var credentials RegisterCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, bodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := Container.CreateUserUseCase.CreateUser(useCases.CreateUserCommand{
		Username: credentials.Username,
		Password: credentials.Password,
		Email:    credentials.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = SetJwtCookie(&w, createdUser.ID)

	message, _ := json.Marshal(map[string]interface{}{"message": "User registered! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func RemoveJwtCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		Secure:   shouldCookieBeSecure,
		HttpOnly: true,
		Path:     "/",
	})
}

func logout(w http.ResponseWriter, _ *http.Request) {
	RemoveJwtCookie(w)
	message, _ := json.Marshal(map[string]interface{}{"message": "User logged out! HTTPonly jwt cookie deleted"})
	_, _ = w.Write(message)
}

func NewAuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/register", register)
	r.Post("/login", login)
	r.Post("/logout", logout)
	return r
}
