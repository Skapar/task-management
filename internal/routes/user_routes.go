package routes

import (
	"github.com/Skapar/task-management/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, userHandler *handler.UserHandler) {
    r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
    r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}
