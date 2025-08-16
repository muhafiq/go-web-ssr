package main

import (
	"go-web-ssr/internal/config"
	"go-web-ssr/internal/controllers"
	"go-web-ssr/internal/middleware"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

/* helper function to register route with chained middlewares */
func Chain(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func main() {
	// init
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config.ConnectDB()
	defer config.DB.Close()
	dir := http.Dir("./public")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(dir))
	http.Handle("/public/", staticFileHandler)

	// routes
	http.HandleFunc("/", Chain(
		controllers.Use(map[string]http.HandlerFunc{"GET": controllers.HomeView}),
		middleware.Logger,
		middleware.Session,
	))
	http.HandleFunc("/login", Chain(
		controllers.Use(map[string]http.HandlerFunc{
			"GET":  controllers.LoginView,
			"POST": controllers.LoginUser,
		}),
		middleware.Logger,
		middleware.Session,
		middleware.Auth,
	))
	http.HandleFunc("/dashboard", Chain(
		controllers.Use(map[string]http.HandlerFunc{"GET": controllers.DashboardView}),
		middleware.Logger,
		middleware.Session,
		middleware.Auth,
	))
	http.HandleFunc("/logout", Chain(
		controllers.Use(map[string]http.HandlerFunc{"POST": controllers.LogoutUser}),
		middleware.Logger,
		middleware.Session,
		middleware.Auth,
	))

	log.Println("HTTP Server start on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
