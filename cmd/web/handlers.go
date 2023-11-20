package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/jba/muxpatterns"
	"github.com/matthewrankin/hypermedia-contact/ui"
)

// Define a home handler function which writes a byte slice containing "Hello
// from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// Create a slice containing the filepath patterns for the templates we
		// want to parse.
		patterns := []string{
			"html/base.tmpl.html", "html/partials/*.tmpl.html", page,
		}

		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem.
		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			app.serverError(w, r, err)
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.serverError(w, r, err)
		}
	}
}

// contactsViewAll provides the handler function to view all contacts.
func (app *application) contactsViewAll(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Display all contacts..."))
	if err != nil {
		app.logger.Error(err.Error())
	}
}

// contactView provides the handler function to view a single contact.
func (app *application) contactView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(muxpatterns.PathValue(r, "id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	_, err = fmt.Fprintf(w, "Display a specific contact with ID %d...", id)
	if err != nil {
		log.Fatal(err)
	}
}

// contactCreate provides the handler function to create a new contact.
func (app *application) contactCreate(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Create a new contact..."))
	if err != nil {
		log.Fatal(err)
	}
}
