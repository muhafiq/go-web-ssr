package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

/*
This helper function use to render html template. But this function have a problem.
It's slow because every request, it will read from file, not from memory.
So this need to fix to cache the template to memory. But i'll do it later.
*/
func renderTemplate(w http.ResponseWriter, views string, data map[string]interface{}) {
	tmpl, err := template.ParseFiles(
		"views/layout.html",
		"views/pages/"+views+".html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
`Use` function use to set http method directly when we register new route.
*/
func Use(methodHandlers map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := methodHandlers[r.Method]; ok {
			handler(w, r)
			return
		}
		http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
	}
}

/* parsing flash messages */
func parseFlashes(session *sessions.Session) ([]string, []string) {
	errs := session.Flashes("error")
	errors := make([]string, len(errs))
	for i, f := range errs {
		if msg, ok := f.(string); ok {
			errors[i] = msg
		}
	}

	succs := session.Flashes("success")
	success := make([]string, len(succs))
	for i, f := range succs {
		if msg, ok := f.(string); ok {
			success[i] = msg
		}
	}

	return success, errors
}
