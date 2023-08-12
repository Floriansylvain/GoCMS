package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

const LoginRoute = "/login"

type LoginPage struct {
	IsError  bool   `json:"isError"`
	Error    string `json:"error"`
	Username string `json:"username"`
}

func NewLoginPage(error string, username string) *LoginPage {
	return &LoginPage{
		IsError:  strings.Compare(error, "") != 0,
		Error:    error,
		Username: username,
	}
}

func GetLoginPageHandler(loginPage *LoginPage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsLoggedIn(r) {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		if r.Method == http.MethodPost {
			PostLoginPage(w, r)
			return
		}
		bs, err := Container.GetPageUseCase.GetPage("login", &loginPage)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(bs)
	}
}

func PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials, err := json.Marshal(&UserLogin{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		r.Method = http.MethodGet
		GetLoginPageHandler(NewLoginPage("Invalid username or password.", r.FormValue("username")))(w, r)
		return
	}

	response, err := http.Post(
		"http://localhost:"+os.Getenv("PORT")+"/v1/auth/login",
		"application/json",
		bytes.NewBuffer(credentials))

	if err != nil || response.StatusCode != http.StatusOK {
		r.Method = http.MethodGet
		GetLoginPageHandler(NewLoginPage("Invalid username or password.", r.FormValue("username")))(w, r)
		return
	}

	w.Header().Set("Set-Cookie", response.Header.Get("Set-Cookie"))

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, LoginRoute, http.StatusPermanentRedirect)
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
	RemoveJwtCookie(w)
	http.Redirect(w, r, LoginRoute, http.StatusSeeOther)
}

func IsLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsLoggedIn(r) {
			http.Redirect(w, r, LoginRoute, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
