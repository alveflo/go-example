package httphandlers

import (
	"net/http"

	"../entities"
)

// SavePageHandler saves or updates a provided page
func SavePageHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	p := &entities.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
