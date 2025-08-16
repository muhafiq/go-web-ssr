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
		middleware.Use(map[string]http.HandlerFunc{"GET": controllers.HomeView}),
		middleware.Logger,
	))
	http.HandleFunc("/login", Chain(
		middleware.Use(map[string]http.HandlerFunc{
			"GET":  controllers.LoginView,
			"POST": controllers.LoginUser,
		}),
		middleware.Logger,
	))
	http.HandleFunc("/dashboard", Chain(
		middleware.Use(map[string]http.HandlerFunc{"GET": controllers.DashboardView}),
		middleware.Logger,
	))

	log.Println("HTTP Server start on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
