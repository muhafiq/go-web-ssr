package main

import (
	"go-web-ssr/internal/config"
	"go-web-ssr/internal/controllers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

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
	http.HandleFunc("/", controllers.Use(map[string]http.HandlerFunc{
		"GET": controllers.HomeView,
	}))
	http.HandleFunc("/login", controllers.Use(map[string]http.HandlerFunc{
		"GET":  controllers.LoginView,
		"POST": controllers.LoginUser,
	}))
	http.HandleFunc("/dashboard", controllers.Use(map[string]http.HandlerFunc{
		"GET": controllers.DashboardView,
	}))

	log.Println("HTTP Server start on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
