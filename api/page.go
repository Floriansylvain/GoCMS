package api

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed static
var staticFolder embed.FS

var headTmpl template.HTML

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

func IsVerifiedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsVerified(r) {
			http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsNotVerifiedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsVerified(r) {
			http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
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
				w.Header().Set("Cache-Control", "public, max-age=31536000")
				w.Header().Set("Content-Type", ct)
			}
		}
		http.FileServer(fsys).ServeHTTP(w, r)
	})
}

func InitHeadTmpl() {
	headTmplHtml, _ := Container.GetPageUseCase.GetPage("utilsHead", nil)
	headTmpl = template.HTML(headTmplHtml)
}

func NewPageRouter() http.Handler {
	r := chi.NewRouter()

	contentStatic := fs.FS(staticFolder)

	InitHeadTmpl()

	r.Handle("/static/*", StaticFileServerWithContentType(http.FS(contentStatic)))

	r.Get("/", GetLogin)
	r.Get(LoginRoute, GetLoginPageHandler(EmptyLoginPage))
	r.Post(LoginRoute, GetLoginPageHandler(EmptyLoginPage))
	r.Get("/register", GetRegisterPageHandler(EmptyRegisterPage))
	r.Post("/register", GetRegisterPageHandler(EmptyRegisterPage))
	r.Get("/logout", GetLogout)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Use(IsNotVerifiedMiddleware)
		r.Get("/register/pending", GetRegisterPendingPage)
		r.Post("/register/pending", PostRegisterPendingPage)
		r.Get("/register/validate", GetRegisterValidatePage)
	})

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Use(IsVerifiedMiddleware)
		r.Get("/home", GetHomePage)
		r.Get("/post", GetPostsPage)
		r.Get("/post/{id}/edit", GetPostEditPage)
		r.Post("/post/{id}/edit", PostPostEditPage)
		r.Get("/post/{id}/delete", GetPostDeletePage)
		r.Get("/post/create", GetPostCreatePage)
		r.Post("/post/create", PostPostCreatePage)
	})

	return r
}
