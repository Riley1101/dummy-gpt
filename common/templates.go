package common

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templates []string, data interface{}) {
	ts, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
