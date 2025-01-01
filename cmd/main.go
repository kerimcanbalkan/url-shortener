package main

import (
	"log"
	"net/http"

	"github.com/kerimcanbalkan/url-shortener/cmd/api"
	"github.com/kerimcanbalkan/url-shortener/config"
	"github.com/kerimcanbalkan/url-shortener/db"
)

func main() {
	// Database connection string
	connStr := "user=" + config.Envs.DBUser + " password=" + config.Envs.DBPassword + " dbname=" + config.Envs.DBName + "sslmode=disable"

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
