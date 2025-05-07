package main

import (
	"log"
	"net/http"

	"fin-api-go/internal/middleware"
	"fin-api-go/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// Setup routes
	mux := router.SetupRoutes()

	err := godotenv.Load("cmd/api/.env")
	if err != nil {
		log.Fatalf("Error loading .env file : %s", err)
	}

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", middleware.Logger(mux)); err != nil {
		log.Fatalf("Server failed to start: %+v", err)
	}
}
