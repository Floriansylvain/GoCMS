package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

const LoginRoute = "/login"

type LoginPage struct {
	PageError *PageError `json:"error"`
	Username  string     `json:"username"`
}

var EmptyLoginPage = &LoginPage{
	PageError: NewPageError(""),
	Username:  "",
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
		if IsUserTableEmpty() {
			http.Redirect(w, r, "/register", http.StatusSeeOther)
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
		GetLoginPageHandler(&LoginPage{
			PageError: NewPageError("Invalid form data format."),
			Username:  r.FormValue("username"),
		})(w, r)
		return
	}

	response, err := http.Post(
		"http://localhost:"+os.Getenv("PORT")+"/v1/auth/login",
		"application/json",
		bytes.NewBuffer(credentials))

	if err != nil || response.StatusCode != http.StatusOK {
		r.Method = http.MethodGet
		GetLoginPageHandler(&LoginPage{
			PageError: NewPageError("Invalid username or password."),
			Username:  r.FormValue("username"),
		})(w, r)
		return
	}

	w.Header().Set("Set-Cookie", response.Header.Get("Set-Cookie"))

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
