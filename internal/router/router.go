package router

import (
	"net/http"

	"fin-api-go/internal/handlers"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()

	mux.HandleFunc("GET /health", healthHandler.Check)
	mux.HandleFunc("GET /transactions", handlers.GetTransactions)
	mux.HandleFunc("GET /expenses", handlers.GetExpenseSummary)

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", mux))

	return api
}
