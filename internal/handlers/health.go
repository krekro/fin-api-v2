package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

// HealthHandler handles health check endpoints
type HealthHandler struct {
	// Add any state you want to maintain
	requestCount int
	mu           sync.Mutex // For thread-safe access to requestCount
}

var (
	instance *HealthHandler
	once     sync.Once
)

// NewHealthHandler creates a new health handler (singleton)
func NewHealthHandler() *HealthHandler {
	once.Do(func() {
		instance = &HealthHandler{
			requestCount: 0,
		}
	})
	return instance
}

// Check handles the health check endpoint
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	// Thread-safe increment of request count
	h.mu.Lock()
	h.requestCount++
	count := h.requestCount
	h.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"status":       "healthy",
		"message":      "Service is running",
		"requestCount": count,
	}

	//log.Printf("Testing w : %v", w)
	json.NewEncoder(w).Encode(response)
}
