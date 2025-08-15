package controllers

import "net/http"

func LoginView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

}
