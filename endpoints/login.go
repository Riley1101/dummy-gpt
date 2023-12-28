package endpoints

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitAuthEndpoint(r chi.Router) {
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
		return
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		success := true
		//		formValues := LoginForm{
		//			Username: r.FormValue("username"),
		//			Password: r.FormValue("password"),
		//		}
		if success {
			// set cookie
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
