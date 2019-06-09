package templates

import (
	"html/template"
	"net/http"

	"../entities"
)

var templates = template.Must(template.ParseFiles(
	"templates/edit.html",
	"templates/view.html",
))

// RenderTemplate renders provided template from cache
func RenderTemplate(w http.ResponseWriter, tmpl string, page *entities.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
