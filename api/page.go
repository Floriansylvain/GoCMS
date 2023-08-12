package api

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed web/templates/*
var templateFiles embed.FS

//go:embed web/static
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

const loginRoute = "/login"

type LoginPage struct {
	IsError  bool   `json:"isError"`
	Error    string `json:"error"`
	Username string `json:"username"`
}

func NewLoginPage(error string, username string) *LoginPage {
	return &LoginPage{
		IsError:  strings.Compare(error, "") != 0,
		Error:    error,
		Username: username,
	}
}

func getPage(page string, data interface{}) ([]byte, error) {
	var processedHTML bytes.Buffer
	tmpl, err := template.ParseFS(templateFiles, "web/templates/"+page+".html")
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&processedHTML, data)
	if err != nil {
		return nil, err
	}
	return processedHTML.Bytes(), nil
}

func getLoginPageHandler(loginPage *LoginPage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsLoggedIn(r) {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		if r.Method == http.MethodPost {
			postLoginPage(w, r)
			return
		}
		bs, err := getPage("login", &loginPage)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(bs)
	}
}

func postLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials, err := json.Marshal(&UserLogin{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		r.Method = http.MethodGet
		getLoginPageHandler(NewLoginPage("Invalid username or password.", r.FormValue("username")))(w, r)
		return
	}

	response, err := http.Post(
		"http://localhost:"+os.Getenv("PORT")+"/v1/auth/login",
		"application/json",
		bytes.NewBuffer(credentials))

	if err != nil || response.StatusCode != http.StatusOK {
		r.Method = http.MethodGet
		getLoginPageHandler(NewLoginPage("Invalid username or password.", r.FormValue("username")))(w, r)
		return
	}

	w.Header().Set("Set-Cookie", response.Header.Get("Set-Cookie"))

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func getHomePage(w http.ResponseWriter, _ *http.Request) {
	navbarTmpl, _ := getPage("componentNavbar", nil)
	homeTmpl, _ := getPage("home", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
	})
	_, _ = w.Write(homeTmpl)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, loginRoute, http.StatusPermanentRedirect)
}

func getLogout(w http.ResponseWriter, r *http.Request) {
	removeJwtCookie(w)
	http.Redirect(w, r, loginRoute, http.StatusSeeOther)
}

func IsLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsLoggedIn(r) {
			http.Redirect(w, r, loginRoute, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func staticFileServerWithContentType(fsys http.FileSystem) http.Handler {
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

	contentStatic, _ := fs.Sub(fs.FS(staticFolder), "web/static")

	r.Handle("/static/*", staticFileServerWithContentType(http.FS(contentStatic)))
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		_, err := http.FS(staticFolder).Open("favicon.ico")
		if err != nil {
			return
		}
	})

	r.Get("/", getLogin)
	r.Get(loginRoute, getLoginPageHandler(NewLoginPage("", "")))
	r.Post(loginRoute, getLoginPageHandler(NewLoginPage("", "")))
	r.Get("/logout", getLogout)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Get("/home", getHomePage)
	})

	return r
}
