package api

import (
	"net/http"
)

func GetRegisterConfirmPage(w http.ResponseWriter, _ *http.Request) {
	registerConfirmTmpl, _ := Container.GetPageUseCase.GetPage("setup2", nil)
	_, _ = w.Write(registerConfirmTmpl)
}
