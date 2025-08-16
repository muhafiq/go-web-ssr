package controllers

import (
	"go-web-ssr/internal/middleware"
	"go-web-ssr/internal/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	session := middleware.GetSession(r)

	success, errors := parseFlashes(session)
	session.Save(r, w)

	renderTemplate(w, "login", map[string]interface{}{
		"Errors":  errors,
		"Success": success,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	session := middleware.GetSession(r)

	u := &models.User{}
	user, err := u.GetByEmail(email)
	if err != nil {
		session.AddFlash("Error fetch data user.", "error")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if user == nil {
		session.AddFlash("Invalid Credentials. Go away!", "error")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		session.AddFlash("Invalid Credentials. Go away!", "error")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session.Values["userId"] = user.ID
	session.Values["email"] = user.Email
	session.AddFlash("Login to dashboard successfully.", "success")
	session.Save(r, w)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	session := middleware.GetSession(r)

	session.Values["userId"] = nil
	session.Values["email"] = nil
	session.AddFlash("Logout successfully.", "success")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func DashboardView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "dashboard", nil)
}
