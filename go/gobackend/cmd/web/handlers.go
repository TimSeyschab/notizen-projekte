package main

import (
	"errors"
	"fmt"
	"gobackend/internal/models"
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Test-Id", "TimTest")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, r, err)
		return
	}
	for _, snippet := range snippets {
		app.logger.Info(fmt.Sprintf("%+v\n", snippet))
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message, use
	// the http.Error() function to send an Internal Server Error response to the
	// user, and then return from the handler so no subsequent code is executed.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, r, err)
		return
	}
	// Then we use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serveError(w, r, err)
	}
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		http.NotFound(w, r)
		return
	}

	// return a 404 Not Found response.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoSnippet) {
			http.NotFound(w, r)
		} else {
			app.serveError(w, r, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", snippet)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ShowCreate"))
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	id, err := app.snippets.Insert("snail", "oompaloompa", 7)
	if err != nil {
		app.serveError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/view/%s", id), http.StatusSeeOther)
}
