package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kerimcanbalkan/url-shortener/cmd/api"
	"github.com/kerimcanbalkan/url-shortener/config"
	"github.com/kerimcanbalkan/url-shortener/db"
)

func main() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.Host,
		config.Envs.Port,
		config.Envs.DBName,
	)
	// connStr := "postgres://postgres:010203@localhost:5432/postgres?sslmode=disable"

	// Initialize database
	database, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize API
	urlAPI := api.NewAPI(database)

	// Start the server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", urlAPI.Router))
}
