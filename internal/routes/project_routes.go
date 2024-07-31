package routes

import (
	"github.com/Skapar/task-management/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(r *mux.Router, projectHandler *handler.ProjectHandler) {
    r.HandleFunc("/projects", projectHandler.GetAllProjects).Methods("GET")
    r.HandleFunc("/projects/{id}", projectHandler.GetProjectByID).Methods("GET")
    r.HandleFunc("/projects", projectHandler.CreateProject).Methods("POST")
    r.HandleFunc("/projects/{id}", projectHandler.UpdateProject).Methods("PUT")
    r.HandleFunc("/projects/{id}", projectHandler.DeleteProject).Methods("DELETE")
}
