package api

import (
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
)

func PostRegisterPendingPage(w http.ResponseWriter, r *http.Request) {
	token, _ := jwtauth.VerifyRequest(
		TokenAuth, r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	userId := token.PrivateClaims()["user_id"].(float64)
	err := Container.DeleteUserUseCase.DeleteUser(uint32(userId))
	if err != nil {
		log.Println(err)
		r.Method = http.MethodGet
		http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
	}

	RemoveJwtCookie(w)

	http.Redirect(w, r, "/register", http.StatusSeeOther)
}

func GetRegisterPendingPage(w http.ResponseWriter, _ *http.Request) {
	registerPendingTmpl, _ := Container.GetPageUseCase.GetPage("registerPending", map[string]interface{}{
		"Head": headTmpl,
	})
	_, _ = w.Write(registerPendingTmpl)
}
