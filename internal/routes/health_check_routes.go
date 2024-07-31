package routes

import (
	"github.com/Skapar/task-management/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterHealthCheckRoutes(r *mux.Router, healthHandler *handler.HealthHandler) {
    r.HandleFunc("/health", healthHandler.HealthCheck).Methods("GET")
}
