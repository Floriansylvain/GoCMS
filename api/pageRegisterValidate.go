package api

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"time"
)

func GetRegisterValidatePage(w http.ResponseWriter, r *http.Request) {
	queryVerificationCode := r.URL.Query().Get("c")
	if queryVerificationCode == "" {
		http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
		return
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	userId, _ := claims["user_id"].(uint32)

	user, _ := Container.GetUserUseCase.GetUser(userId)
	errorMessage := ""

	if user.VerificationCode != queryVerificationCode || user.VerificationExpiration.Before(time.Now()) {
		errorMessage = "Verification link is incorrect or has expired."
	} else {
		// TODO New usecase "UpdateUserUseCase" to update its verification status
	}

	registerValidateTmpl, _ := Container.GetPageUseCase.GetPage("registerValidate", map[string]interface{}{
		"Head":      headTmpl,
		"PageError": NewPageError(errorMessage),
	})
	_, _ = w.Write(registerValidateTmpl)
}
