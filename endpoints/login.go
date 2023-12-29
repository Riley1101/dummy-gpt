package endpoints

import (
	"database/sql"
	c "dummygpt/common"
	"dummygpt/database"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type AuthHandler struct {
	UserDb *database.UserDb
}

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h *AuthHandler) Register(r chi.Router) {
	templates := []string{
		"templates/register.html",
		"templates/base.html",
	}
	ts, err := template.ParseFiles(templates...)

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		register := RegisterForm{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}
		_, err := c.ValidateStruct(register)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := c.HashPassword(register.Password)

		userExists, err := h.UserDb.CheckUserExists(register.Username)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if userExists {
			err = ts.ExecuteTemplate(w, "base", AuthResponse{
				Success: true,
				Message: "User created successfully",
			})
			return
		}

		dbUser := database.User{
			Username: register.Username,
			Password: hashedPassword,
			Email:    register.Email,
		}

		err = h.UserDb.CreateUser(&dbUser)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = ts.ExecuteTemplate(w, "base", AuthResponse{
			Success: true,
			Message: "User created successfully",
		})
	})
}

func (h *AuthHandler) Login(r chi.Router) {
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"templates/login.html",
			"templates/base.html",
		}
		ts, err := template.ParseFiles(templates...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
		fmt.Println(err)
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
	UserDb := database.UserDb{DB: db}
	authHandler := AuthHandler{
		UserDb: &UserDb,
	}
	authHandler.UserDb.CreateUserTable()
	authHandler.Register(r)
	authHandler.Login(r)
}
