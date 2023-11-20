package main

import (
	"net/http"

	"github.com/jba/muxpatterns"
)

func (app *application) routes() *muxpatterns.ServeMux {
	mux := muxpatterns.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /contacts", app.contactsViewAll)
	mux.HandleFunc("GET /contacts/{id}", app.contactView)
	mux.HandleFunc("POST /contacts", app.contactCreate)

	return mux
}
