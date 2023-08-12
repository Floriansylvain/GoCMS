package api

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"io/fs"
	"net/http"
	"path/filepath"
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
	r.Get(LoginRoute, GetLoginPageHandler(NewLoginPage("", "")))
	r.Post(LoginRoute, GetLoginPageHandler(NewLoginPage("", "")))
	r.Get("/logout", GetLogout)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Get("/home", GetHomePage)
	})

	return r
}
