package main

import (
	"database/sql"
	"dummygpt/common"
	"dummygpt/endpoints"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "dummy-gpt"
)

var db *sql.DB

func main() {
	godotenv.Load(".env")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname)
	db, _ := sql.Open("postgres", connStr)

	// public
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome!"))
		})
		endpoints.InitAuthEndpoint(r, db)
	})

	// admin
	r.Group(func(r chi.Router) {
		r.Use(common.SessionAuthMiddleware)
		r.Route("/admin", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Main admin"))
			})
			endpoints.InitEndpoints(r, db)

		})
	})

	http.ListenAndServe(":3000", r)
}
