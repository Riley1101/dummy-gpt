package endpoints

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Register(r chi.Router) {
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"templates/register.tmpl",
			"templates/base.tmpl",
		}

		ts, err := template.ParseFiles(templates...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
	})
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		register := LoginForm{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		w.Write([]byte(register.Username))
	})
}

func (h *AuthHandler) Login(r chi.Router) {
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"templates/login.tmpl",
			"templates/base.tmpl",
		}
		ts, err := template.ParseFiles(templates...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		success := true
		//		formValues := LoginForm{
		//			Username: r.FormValue("username"),
		//			Password: r.FormValue("password"),
		//		}
		if success {
			cookie := http.Cookie{
				Name:     "token",
				Value:    "Hello world!",
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	})
	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("logout"))
	})

}

func InitAuthEndpoint(r chi.Router, db *sql.DB) {
	authHandler := AuthHandler{DB: db}
	authHandler.Register(r)
	authHandler.Login(r)
}
