package controllers

import (
	"html/template"
	"net/http"
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
