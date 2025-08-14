package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func ConnectDB() {
	url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
		os.Exit(1)
	}

	DB = db
}
