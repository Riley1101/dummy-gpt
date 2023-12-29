package endpoints

import (
	"database/sql"
	"dummygpt/database"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type EndpointHandler struct {
	EndpointDB *database.EndpointDB
}

var owner = 1

func (e *EndpointHandler) Endpoints(r chi.Router) {
	templates := []string{"templates/admin/endpoint.html", "templates/base.html"}
	ts, err := template.ParseFiles(templates...)

	if err != nil {
		panic(err)
	}

	r.Get("/endpoints", func(w http.ResponseWriter, r *http.Request) {
		endpoints, err := e.EndpointDB.GetEndpointsByOwner(owner)
		fmt.Println(endpoints)
		if err != nil {
			panic(err)
		}
		ts.ExecuteTemplate(w, "base", endpoints)

	})

	r.Post("/endpoints", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post endpoints"))
	})
}

func InitEndpoints(r chi.Router, db *sql.DB) {
	endPoint := database.EndpointDB{DB: db}
	endpointHandler := &EndpointHandler{
		EndpointDB: &endPoint,
	}
	endpointHandler.EndpointDB.CreateEndpointTable()
	endpointHandler.Endpoints(r)
}
