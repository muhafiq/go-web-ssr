package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

type contextKey string

const sessionKey contextKey = "session"

var (
	key   = []byte("session-secret")
	store = sessions.NewCookieStore(key)
)

func Session(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "sid")

		ctx := context.WithValue(r.Context(), sessionKey, session)
		next(w, r.WithContext(ctx))
	}
}

func GetSession(r *http.Request) *sessions.Session {
	sess := r.Context().Value(sessionKey)
	if sess == nil {
		return nil
	}
	return sess.(*sessions.Session)
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := GetSession(r)

		if strings.HasPrefix(r.URL.Path, "/dashboard") {
			if session.Values["userId"] == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		} else if strings.HasPrefix(r.URL.Path, "/login") {
			if session.Values["userId"] != nil {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			}
		}

		next(w, r)
	}
}
