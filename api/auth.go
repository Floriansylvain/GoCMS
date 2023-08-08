package api

import (
	"GohCMS2/useCases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserRegister struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

var TokenAuth *jwtauth.JWTAuth

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
		Secure:   false, // TODO false in dev, true in prod
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

func isAllowedToCreateUser() bool {
	users := container.ListUsersUseCase.ListUsers()
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

func login(w http.ResponseWriter, r *http.Request) {
	var user UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, bodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser, err := container.GetUserByUsernameUseCase.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(w, logsErrorMessage, http.StatusForbidden)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, logsErrorMessage, http.StatusForbidden)
		return
	}

	_ = SetJwtCookie(&w, dbUser.ID)

	message, _ := json.Marshal(map[string]interface{}{"message": "User logged in! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func register(w http.ResponseWriter, r *http.Request) {
	if !IsLoggedIn(r) && !isAllowedToCreateUser() {
		http.Error(w, "You are not allowed to create a user. Log in or reset database.", http.StatusForbidden)
		return
	}

	var user UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, bodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := container.CreateUserUseCase.CreateUser(useCases.CreateUserCommand{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = SetJwtCookie(&w, createdUser.ID)

	message, _ := json.Marshal(map[string]interface{}{"message": "User registered! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func logout(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		Secure:   false, // TODO false in dev, true in prod
		HttpOnly: true,
		Path:     "/v1/",
	})

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