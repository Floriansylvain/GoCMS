package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type RegisterPage struct {
	PageError *PageError `json:"error"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
}

var EmptyRegisterPage = &RegisterPage{
	PageError: NewPageError(""),
	Username:  "",
	Email:     "",
}

func GetRegisterPageHandler(registerPage *RegisterPage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsLoggedIn(r) || !IsUserTableEmpty() {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		if r.Method == http.MethodPost {
			PostRegisterPage(w, r)
			return
		}
		bs, err := Container.GetPageUseCase.GetPage("setup1", &registerPage)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(bs)
	}
}

func PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials, err := json.Marshal(&UserRegister{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
	})
	if err != nil {
		r.Method = http.MethodGet
		GetRegisterPageHandler(&RegisterPage{
			PageError: NewPageError("Invalid register form data format."),
			Username:  r.FormValue("username"),
			Email:     r.FormValue("email"),
		})(w, r)
		return
	}

	response, err := http.Post(
		"http://localhost:"+os.Getenv("PORT")+"/v1/auth/register",
		"application/json",
		bytes.NewBuffer(credentials))

	if err != nil || response.StatusCode != http.StatusOK {
		r.Method = http.MethodGet
		GetRegisterPageHandler(&RegisterPage{
			PageError: NewPageError("Username should be between 3 and 20 characters long, password should be between 8 and 20 characters long, and email should be a valid email address."),
			Username:  r.FormValue("username"),
			Email:     r.FormValue("email"),
		})(w, r)
		return
	}

	w.Header().Set("Set-Cookie", response.Header.Get("Set-Cookie"))

	http.Redirect(w, r, "/register-confirm", http.StatusSeeOther)
}
