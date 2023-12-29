package endpoints

import (
	"database/sql"
	c "dummygpt/common"
	"dummygpt/database"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TemplateData struct {
	Data interface{}
	Err  error
	Sql  interface{}
}

type EndpointForm struct {
	EndpointName string `json:"endpoint_name" validate:"required"`
	Schema       string `json:"schema" validate:"required"`
}

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
		if err != nil {
			panic(err)
		}
		ts.ExecuteTemplate(w, "base", TemplateData{
			Data: endpoints,
			Err:  err,
			Sql:  nil,
		})
	})

	r.Post("/endpoints", func(w http.ResponseWriter, r *http.Request) {
		formData := EndpointForm{
			EndpointName: r.FormValue("endpointName"),
			Schema:       r.FormValue("schema"),
		}
		c.ValidateStruct(formData)
		dbSchema := c.DbSchema{
			Name: formData.EndpointName,
		}
		dbSchema.ParseSchema(formData.Schema)
		dbSchema.DescribeSchema()
		query := dbSchema.GenerateQuery()
		log.Println(query)
		ts.ExecuteTemplate(w, "base", TemplateData{
			Data: nil,
			Err:  err,
			Sql:  query,
		})
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
