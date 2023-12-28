package main

import (
	"database/sql"
	"dummygpt/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
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

	table := common.Table{
		Name: "users",
		Fields: []common.Field{
			{
				Name:     "id",
				Datatype: "integer",
			},
			{
				Name:     "name",
				Datatype: "varchar(255)",
			},
			{
				Name:     "email",
				Datatype: "varchar(255)",
			},
		},
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		formBody := common.SchemaForm{
			Schema: r.FormValue("schema"),
		}
		dbSchema := common.DbSchema{
			Name: "users",
		}
		dbSchema.ParseSchema(formBody.Schema)
		dbSchema.DescribeSchema()
		dbSchema.GenerateQuery()
		tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))
		tmpl.Execute(w, table)
	})
	http.ListenAndServe(":3000", r)
}
