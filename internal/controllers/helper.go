package controllers

import (
	"html/template"
	"net/http"
)

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

func Use(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		handler(w, r)
	}
}
