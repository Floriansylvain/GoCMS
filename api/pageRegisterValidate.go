package api

import (
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetRegisterValidatePage(w http.ResponseWriter, r *http.Request) {
	queryVerificationCode := r.URL.Query().Get("c")
	if queryVerificationCode == "" {
		http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
		return
	}

	token, _ := jwtauth.VerifyRequest(
		TokenAuth,
		r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)

	userId := token.PrivateClaims()["user_id"].(float64)
	user, _ := Container.GetUserUseCase.GetUser(uint32(userId))
	errorMessage := ""

	err := bcrypt.CompareHashAndPassword([]byte(user.VerificationCode), []byte(queryVerificationCode))
	if err != nil || user.VerificationExpiration.Before(time.Now()) {
		errorMessage = "Verification link is incorrect or has expired."
	} else {
		_, err := Container.UpdateUserUseCase.UpdateVerificationStatus(user.ID, true)
		if err != nil {
			errorMessage = "Something went wrong server-side. User account may not exist."
		}
	}

	registerValidateTmpl, _ := Container.GetPageUseCase.GetPage("registerValidate", map[string]interface{}{
		"Head":      headTmpl,
		"PageError": NewPageError(errorMessage),
	})
	_, _ = w.Write(registerValidateTmpl)
}
