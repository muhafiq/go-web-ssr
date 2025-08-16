package controllers

import (
	"go-web-ssr/internal/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	u := &models.User{}
	user, err := u.GetByEmail(email)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// TODO: create new user session and store to http only cookie.

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

}

func DashboardView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "dashboard", nil)
}
