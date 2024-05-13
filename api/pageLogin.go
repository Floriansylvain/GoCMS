package api

import (
	"net/http"
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
		bs, err := Container.GetPageUseCase.GetPage("login", map[string]interface{}{
			"PageError": loginPage.PageError,
			"Username":  loginPage.Username,
			"Head":      headTmpl,
		})
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(bs)
	}
}

func PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials := LoginCredentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	err := validate.Struct(credentials)
	if err != nil {
		r.Method = http.MethodGet
		GetLoginPageHandler(&LoginPage{
			PageError: NewPageError("Invalid form data format."),
			Username:  r.FormValue("username"),
		})(w, r)
		return
	}

	dbUser, err := getUserFromCredentials(credentials)
	if err != nil {
		r.Method = http.MethodGet
		GetLoginPageHandler(&LoginPage{
			PageError: NewPageError("Invalid username or password."),
			Username:  r.FormValue("username"),
		})(w, r)
		return
	}

	_ = SetJwtCookie(&w, dbUser.ID)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
