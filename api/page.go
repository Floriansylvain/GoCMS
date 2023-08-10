package api

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var contentTypes = map[string]string{
	".css":  "text/css",
	".js":   "application/javascript",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".webp": "image/webp",
	".svg":  "image/svg+xml",
	".ico":  "image/x-icon",
}

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
	tmpl, err := template.ParseFiles("./web/templates/" + page + ".html")
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

	// TODO replace url with som env variable
	response, err := http.Post("http://localhost:8080/v1/auth/login", "application/json", bytes.NewBuffer(credentials))
	// TODO handle possible errors in separate file (api package)
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
	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
}

func getLogout(w http.ResponseWriter, r *http.Request) {
	removeJwtCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func IsLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsLoggedIn(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func staticFileServerWithContentType(dir http.Dir) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedFilePath := r.URL.Path
		fileExtension := filepath.Ext(requestedFilePath)

		if contentType, ok := contentTypes[fileExtension]; ok {
			w.Header().Set("Content-Type", contentType)
		}

		http.FileServer(dir).ServeHTTP(w, r)
	})
}

func NewPageRouter() http.Handler {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", staticFileServerWithContentType("./web/static")))
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/static/favicon.ico")
	})

	r.Get("/", getLogin)
	r.Get("/login", getLoginPageHandler(NewLoginPage("", "")))
	r.Post("/login", getLoginPageHandler(NewLoginPage("", "")))
	r.Get("/logout", getLogout)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Get("/home", getHomePage)
	})

	return r
}
