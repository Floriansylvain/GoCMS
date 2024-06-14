package api

import (
	"github.com/google/uuid"
	"net/http"
	"os"
)

type PasswordResetRequest struct {
	Email string `json:"email"`
}

var successMessage = "An email has been sent to the provided email address if it exists."

func GetPasswordResetRequest(w http.ResponseWriter, r *http.Request) {
	success := r.URL.Query().Get("success")
	email := r.URL.Query().Get("email")

	bs, _ := Container.GetPageUseCase.GetPage("passwordResetRequest", map[string]interface{}{
		"Head":    headTmpl,
		"Email":   email,
		"Success": success,
	})
	_, _ = w.Write(bs)
}

func PostPasswordResetRequest(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	passResetReq := PasswordResetRequest{
		Email: r.FormValue("email"),
	}
	var getRedirectUrl = "/register/reset/request?success=" + successMessage + "&email=" + passResetReq.Email

	err := validate.Struct(passResetReq)
	if err != nil {
		http.Redirect(w, r, getRedirectUrl, http.StatusSeeOther)
		return
	}

	fetchedUser, err := Container.GetUserUseCase.GetUserByEmail(passResetReq.Email)
	if err != nil {
		http.Redirect(w, r, getRedirectUrl, http.StatusSeeOther)
		return
	}

	verificationCode := uuid.NewString()
	updatedUser, _ := Container.UpdateUserUseCase.UpdatePasswordResetCode(fetchedUser.ID, verificationCode)

	_ = Container.SendMailUseCase.SendMail(updatedUser.Email, "passwordReset", map[string]string{
		"Host":             os.Getenv("HOST"),
		"VerificationCode": verificationCode,
		"Email":            updatedUser.Email,
	})

	http.Redirect(w, r, getRedirectUrl, http.StatusSeeOther)
}
