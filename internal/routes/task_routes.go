package routes

import (
	"github.com/Skapar/task-management/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router, taskHandler *handler.TaskHandler) {
    r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
    r.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
    r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
}
