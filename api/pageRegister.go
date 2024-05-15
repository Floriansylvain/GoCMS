package api

import (
	"github.com/google/uuid"
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
		bs, err := Container.GetPageUseCase.GetPage("register", map[string]interface{}{
			"PageError": registerPage.PageError,
			"Username":  registerPage.Username,
			"Email":     registerPage.Email,
			"Head":      headTmpl,
		})
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(bs)
	}
}

func PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials := RegisterCredentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
	}
	err := validate.Struct(credentials)
	if err != nil {
		r.Method = http.MethodGet
		GetRegisterPageHandler(&RegisterPage{
			PageError: NewPageError("Invalid register form data format."),
			Username:  r.FormValue("username"),
			Email:     r.FormValue("email"),
		})(w, r)
		return
	}

	verificationCode := uuid.NewString()
	createdUser, err := getNewUser(credentials, verificationCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Container.SendMailUseCase.SendMail(createdUser.Email, "mailValidation", map[string]string{
		"Host":             os.Getenv("HOST"),
		"VerificationCode": verificationCode,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_ = SetJwtCookie(&w, createdUser.ID)

	http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
}
