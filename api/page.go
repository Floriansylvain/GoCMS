package api

import (
	"GohCMS2/adapters/secondary"
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type LoginPage struct {
	IsError bool   `json:"isError"`
	Error   string `json:"error"`
}

func getPage(page string, data interface{}) ([]byte, error) {
	template := secondary.GetTemplate(page)
	processed, err := secondary.ProcessTemplate(template, data)
	if err != nil {
		return nil, err
	}
	return processed, nil
}

func getLoginPageHandler(errMsg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := getPage("login", &LoginPage{
			IsError: strings.Compare(errMsg, "") != 0,
			Error:   errMsg,
		})
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(bs)
	}
}

func postLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials, err := json.Marshal(&UserLogin{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		getLoginPageHandler("Missing username or password.")(w, r)
		return
	}

	// TODO replace url with som env variable
	response, err := http.Post("http://localhost:8080/v1/auth/login", "application/json", bytes.NewBuffer(credentials))
	// TODO handle possible errors in separate file (api package)
	if err != nil || response.StatusCode != http.StatusOK {
		getLoginPageHandler("Invalid username or password.")(w, r)
		return
	}

	w.Header().Set("Set-Cookie", response.Header.Get("Set-Cookie"))

	http.Redirect(w, r, "/home", http.StatusFound)
}

func getHomePage(w http.ResponseWriter, _ *http.Request) {
	bs, err := getPage("home", nil)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
	_, _ = w.Write(bs)
}

func IsLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsLoggedIn(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewPageRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/login", getLoginPageHandler(""))
	r.Post("/login", postLoginPage)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Get("/home", getHomePage)
	})

	return r
}
