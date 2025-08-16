package middleware

import "net/http"

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
