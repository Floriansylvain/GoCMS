package api

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed static
var staticFolder embed.FS

var contentTypes = map[string]string{
	".css":  "text/css",
	".js":   "application/javascript",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".webp": "image/webp",
	".svg":  "image/svg+xml",
	".ico":  "image/x-icon",
}

type PageError struct {
	Message string `json:"message"`
	IsError bool   `json:"isError"`
}

func NewPageError(message string) *PageError {
	return &PageError{
		Message: message,
		IsError: strings.Compare(message, "") != 0,
	}
}

func IsLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsLoggedIn(r) {
			http.Redirect(w, r, LoginRoute, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, LoginRoute, http.StatusPermanentRedirect)
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
	RemoveJwtCookie(w)
	http.Redirect(w, r, LoginRoute, http.StatusSeeOther)
}

func StaticFileServerWithContentType(fsys http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if ext := filepath.Ext(path); ext != "" {
			if ct, ok := contentTypes[ext]; ok {
				w.Header().Set("Content-Type", ct)
			}
		}
		http.FileServer(fsys).ServeHTTP(w, r)
	})
}

func NewPageRouter() http.Handler {
	r := chi.NewRouter()

	contentStatic := fs.FS(staticFolder)

	r.Handle("/static/*", StaticFileServerWithContentType(http.FS(contentStatic)))
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		_, err := http.FS(staticFolder).Open("favicon.ico")
		if err != nil {
			return
		}
	})

	r.Get("/", GetLogin)
	r.Get(LoginRoute, GetLoginPageHandler(EmptyLoginPage))
	r.Post(LoginRoute, GetLoginPageHandler(EmptyLoginPage))
	r.Get("/register", GetRegisterPageHandler(EmptyRegisterPage))
	r.Post("/register", GetRegisterPageHandler(EmptyRegisterPage))
	r.Get("/logout", GetLogout)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Get("/register-confirm", GetRegisterConfirmPage)
		r.Get("/home", GetHomePage)
	})

	return r
}
