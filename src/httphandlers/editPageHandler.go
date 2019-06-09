package httphandlers

import (
	"net/http"

	"../entities"
	"../templates"
)

// EditPageHandler renders edit.html template
func EditPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := entities.LoadPage(title)
	if err != nil {
		page = &entities.Page{Title: title}
	}

	templates.RenderTemplate(w, "edit", page)
}
