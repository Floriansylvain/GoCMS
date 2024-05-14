package api

import "net/http"

func GetRegisterPendingPage(w http.ResponseWriter, _ *http.Request) {
	registerPendingTmpl, _ := Container.GetPageUseCase.GetPage("registerPending", map[string]interface{}{
		"Head": headTmpl,
	})
	_, _ = w.Write(registerPendingTmpl)
}
