package main

import (
	"database/sql"
	"dummygpt/common"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
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
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	//	user, password, host, port, dbname)
	// db, _ := sql.Open("postgres", connStr)

	// public
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome!"))
		})
	})
	// admin
	r.Group(func(r chi.Router) {
		auth := common.Auth{}
		common.InitAuth(&auth)
		r.Use(jwtauth.Verifier(auth.Token))
		r.Use(jwtauth.Authenticator(auth.Token))
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})
	http.ListenAndServe(":3000", r)
}

//table := common.Table{
//	Name: "users",
//	Fields: []common.Field{
//		{
//			Name:     "id",
//			Datatype: "integer",
//		},
//		{
//			Name:     "name",
//			Datatype: "varchar(255)",
//		},
//		{
//			Name:     "email",
//			Datatype: "varchar(255)",
//		},
//	},
//}

//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//	const userid = "Arkar"
//	formBody := common.SchemaForm{
//		Schema: r.FormValue("schema"),
//	}
//	dbSchema := common.DbSchema{
//		Name: "users",
//	}
//	dbSchema.ParseSchema(formBody.Schema)
//	dbSchema.DescribeSchema()
//	dbSchema.GenerateQuery()
//	files := []string{
//		"templates/index.tmpl",
//		"templates/base.tmpl",
//	}
//	ts, err := template.ParseFiles(files...)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	err = ts.ExecuteTemplate(w, "base", table)
//})
