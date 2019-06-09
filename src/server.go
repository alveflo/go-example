package main

import (
	"log"
	"net/http"

	"./httphandlers"
)

func main() {
	http.HandleFunc("/view/", httphandlers.MakeHandler(httphandlers.ViewPageHandler))
	http.HandleFunc("/edit/", httphandlers.MakeHandler(httphandlers.EditPageHandler))
	http.HandleFunc("/save/", httphandlers.MakeHandler(httphandlers.SavePageHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
