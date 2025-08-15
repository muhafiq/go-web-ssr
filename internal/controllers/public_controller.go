package controllers

import (
	"net/http"
)

func HomeView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", map[string]interface{}{
		"Title": "Home",
	})
}
